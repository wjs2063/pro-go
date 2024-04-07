/*
defer 예제입니다.
*/
package main

import "os"

func main() {
	f, err := os.Open("1.txt")
	if err != nil {
		panic(err)
	}
	// stack 에 f.Close() 추가됨
	defer f.Close()
	bytes := make([]byte, 1024)
	f.Read(bytes)
	println(len(bytes))
}
