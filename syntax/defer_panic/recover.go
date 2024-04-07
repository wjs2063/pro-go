/*
recover 함수예제를 위한 공간입니다.

[참고]
https://go.dev/blog/defer-panic-and-recover
*/
package main

import (
	"fmt"
	"os"
)

func main() {
	// 잘못된 파일명을 넣음
	OpenFile("Invalid.txt")

	// recover에 의해
	// 이 문장 실행됨
	println("Done")
}

func OpenFile(fn string) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("파일 열기 에러")
		}
	}()

	f, err := os.Open(fn)
	if err != nil {
		panic(err)
	}
	defer f.Close()

}
