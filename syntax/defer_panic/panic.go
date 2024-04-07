/*
panic 예제를 위한 공간입니다.
*/
package main

import (
	"fmt"
	"os"
)

func main() {
	openFile("invalid.txt")
	fmt.Println("Done")

}

func openFile(fn string) {
	f, err := os.Open(fn)
	// err 이 나오면 panic 발동후 모든 defer함수를 실행
	if err != nil {
		panic(err)
	}

	defer f.Close()
}
