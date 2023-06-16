package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

const size = 100 * 1024 * 1024

func main() {
	stat, err := os.Stat("test/test.mp4")
	if err != nil {
		fmt.Println(err)
	}
	Num := int(math.Ceil(float64(stat.Size()) / size))
	file, _ := os.OpenFile("test/test.mp4", os.O_RDONLY, 0666)
	for i := 0; i < Num; i++ {
		b := make([]byte, size)
		file.Seek(int64(i*size), 0)
		if size > stat.Size()-int64(i*size) {
			b = make([]byte, stat.Size()-int64(i*size))
		}
		file.Read(b)
		openFile, _ := os.OpenFile("./"+strconv.Itoa(i)+".chunk", os.O_CREATE, os.ModePerm)
		openFile.Write(b)
		openFile.Close()
		fmt.Println(111)
	}
	file.Close()

}
