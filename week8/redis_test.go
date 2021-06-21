package week8

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"math/rand"
	"testing"
)

const (
	kvNum     = 50000
	keySize   = 10
	valueSize = 10
)

var (
	ctx       = context.Background()
	charsPool = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
)

func genRandomStr(size int) string {
	rtn := make([]byte, size)
	for i := 0; i < size; i++ {
		rtn[i] = charsPool[rand.Intn(len(charsPool))]
		//rtn = append(rtn, charsPool[rand.Intn(len(charsPool))])
	}
	return string(rtn)
}

func TestWrite2Redis(t *testing.T) {
	client := redis.NewClient(&redis.Options{
		Addr: "127.0.0.1:6379",
		DB:   0,
	})

	for i := 0; i < kvNum; i++ {
		k := genRandomStr(keySize)
		v := genRandomStr(valueSize)
		client.Set(ctx, k, v, 0)
		if i%10000 == 0 {
			fmt.Println("------", i)
		}
	}

}
