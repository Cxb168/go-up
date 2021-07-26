package protocol

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"testing"
)

func TestProtocol_Encode(t *testing.T) {
	// bytes.Buffer非并发安全
	buf := new(bytes.Buffer)
	proto := NewProtocol(buf)
	// 模拟客户端发包
	//go func() {
	for i := 0; i < 20; i++ {
		bodyStr := fmt.Sprintf("[%d]Msg...", i)
		bodyByte := proto.Encode([]byte(bodyStr))
		buf.Write(bodyByte)
	}
	//}()

	proto.Run()
}

func TestPutUnit32(t *testing.T) {
	rtn := make([]byte, 32)
	binary.BigEndian.PutUint32(rtn, 21)
	binary.BigEndian.PutUint16(rtn[4:6], 12)
	binary.BigEndian.PutUint16(rtn[6:8], 1)
	binary.BigEndian.PutUint16(rtn[8:10], 2)
	binary.BigEndian.PutUint16(rtn[10:12], 3)
	fmt.Printf("%v\n", rtn)
}
