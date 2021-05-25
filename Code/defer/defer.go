package main

import (
	"io"
	"os"
)

func main() {
	CopyFile("/Users/oliverwerlen/Documents/6th/PCP/PCP_GOLANG/code/defer/testCopy.txt", "/Users/oliverwerlen/Documents/6th/PCP/PCP_GOLANG/code/defer/test.txt")
}

func CopyFile(dstName, srcName string) (written int64, err error) {
	src, err := os.Open(srcName)
	if err != nil {
		return
	}
	defer src.Close()

	dst, err := os.Create(dstName)
	if err != nil {
		return
	}
	defer dst.Close()

	return io.Copy(dst, src)
}
