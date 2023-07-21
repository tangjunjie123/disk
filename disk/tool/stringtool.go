package tool

import (
	"crypto/md5"
	"fmt"
	"github.com/google/uuid"
	"math/rand"
)

func Md5(s string) string {
	sum := md5.Sum([]byte(s))
	sprintf := fmt.Sprintf("%x", sum)
	return sprintf
}

func RandInt() uint32 {
	return rand.Uint32()%1000 + 1000
}

func UUID() string {
	newUUID, err := uuid.NewUUID()
	if err != nil {
		fmt.Println(err)
	}
	return newUUID.String()
}
