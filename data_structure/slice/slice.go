/*
golang 의 slice 연습용 예제입니다.
*/
package main

import "fmt"

func main() {

	// slice 는 기본적으로 배열기반이지만 현재 slice 크기까지만 배열로 이어져있고
	// 다음에 크기를 늘려야할때면 메모리주소값이 연속적이지않다.

	// 8byte 로 선언해보자
	// []int64 형 slice, len,cap, 현재길이0이고 용량 1024인 slice 생성
	slice := make([]int64, 0, 1024)

	var idx int64

	// 주소가 기본적으로 8byte 차이가있다는것을 알수있다. 하지만 간간히 2byte 차이가나는 구간도 존재한다.
	// 초기를 0으로하면 2 4 8 16 일때 동떨어진 주소값을 사용한다는 사실도 알수있다.
	for idx = 0; idx < 10; idx++ {
		slice = append(slice, idx)
		n := len(slice)
		fmt.Println(&slice[n-1])
	}

	for k, v := range slice {
		fmt.Println(k, v)
	}
}
