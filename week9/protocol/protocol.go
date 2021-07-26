package protocol

import (
	"encoding/binary"
	"fmt"
	"io"
)

const (
	PackageLength = 4
	HeaderLength  = 2
	ProtoVersion  = 2
	Operation     = 2
	SequenceId    = 2
	HeaderLenVal  = PackageLength + HeaderLength + ProtoVersion + Operation + SequenceId
)

type PackageInfo struct {
	PackageLen   uint32
	HeaderLen    uint16
	ProtoVersion uint16
	Operation    uint16
	SequenceId   uint16
	Body         []byte
}

// Protocol 协议解析类，对TCP长连接进行解析，通过chan返回对应协议内容
type Protocol struct {
	conn io.ReadWriter
	ch   chan *PackageInfo
}

func NewProtocol(conn io.ReadWriter) *Protocol {
	return &Protocol{
		conn: conn,
		ch:   make(chan *PackageInfo),
	}
}

func (p *Protocol) Run() {
	for {
		packageInfo := new(PackageInfo)
		pLenByte := make([]byte, PackageLength)
		_, err := p.conn.Read(pLenByte)

		if err != nil {
			panic(err)
		}
		pLen := binary.BigEndian.Uint32(pLenByte)
		if pLen < HeaderLenVal {
			fmt.Printf("pLenByte: %v\n", pLenByte)
			panic(fmt.Errorf("pLen长度错误[%d]\n", pLen))
		}
		remainByte := make([]byte, pLen-PackageLength)
		// TODO 是否判断读取字节数
		_, err = p.conn.Read(remainByte)
		if err != nil {
			panic(err)
		}
		packageInfo.PackageLen = pLen
		packageInfo.HeaderLen = binary.BigEndian.Uint16(remainByte[:HeaderLength])
		packageInfo.ProtoVersion = binary.BigEndian.Uint16(remainByte[HeaderLength : HeaderLength+ProtoVersion])
		packageInfo.Operation = binary.BigEndian.Uint16(remainByte[HeaderLength+ProtoVersion : HeaderLength+ProtoVersion+Operation])
		packageInfo.SequenceId = binary.BigEndian.Uint16(remainByte[HeaderLength+ProtoVersion+Operation : HeaderLength+ProtoVersion+Operation+SequenceId])
		packageInfo.Body = remainByte[HeaderLength+ProtoVersion+Operation+SequenceId:]

		//fmt.Printf("PackageInfo: %+v\n", packageInfo)
		p.ch <- packageInfo
	}
}

func (p *Protocol) GetChan() chan *PackageInfo {
	return p.ch
}

func (p *Protocol) Encode(body []byte) []byte {
	bodyLen := len(body)
	packageLen := uint32(bodyLen) + HeaderLenVal
	packageInfo := new(PackageInfo)
	packageInfo.PackageLen = packageLen
	packageInfo.HeaderLen = HeaderLenVal
	packageInfo.ProtoVersion = 1
	packageInfo.Operation = 2
	packageInfo.SequenceId = 3
	packageInfo.Body = body

	rtn := make([]byte, packageLen)
	binary.BigEndian.PutUint32(rtn[:PackageLength], packageInfo.PackageLen)
	binary.BigEndian.PutUint16(rtn[PackageLength:PackageLength+HeaderLength], packageInfo.HeaderLen)
	binary.BigEndian.PutUint16(rtn[PackageLength+HeaderLength:PackageLength+HeaderLength+ProtoVersion], packageInfo.ProtoVersion)
	binary.BigEndian.PutUint16(rtn[PackageLength+HeaderLength+ProtoVersion:PackageLength+HeaderLength+ProtoVersion+Operation], packageInfo.Operation)
	binary.BigEndian.PutUint16(rtn[PackageLength+HeaderLength+ProtoVersion+Operation:PackageLength+HeaderLength+ProtoVersion+Operation+SequenceId], packageInfo.SequenceId)
	copy(rtn[HeaderLenVal:], body)
	//fmt.Printf("%v\n", rtn)
	return rtn
}
