/*
표준 입출력을 받기위한 예제입니다.
*/

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var reader *bufio.Reader = bufio.NewReader(os.Stdin)

	// data 를 받아서 저장할 slice
	slice := make([]int, 0)
	for i := 0; i < 2; i++ {
		var a, b int
		fmt.Fscanln(reader, &a, &b)
		slice = append(slice, a, b)
	}

	fmt.Println(slice)

	// txt 에 데이터 쓰기

	fd, err := os.Create("write.txt")
	if err != nil {
		panic("create error")
		return
	}
	defer fd.Close()
	var writer *bufio.Writer = bufio.NewWriter(fd)

	_, err = writer.WriteString("문자열 쓰기 예제입니다")
	if err != nil {
		panic("문자열 쓰기 에러\n")
		return
	}

	err = writer.Flush()
	if err != nil {
		panic("플러시 에러")
		return
	}

}
