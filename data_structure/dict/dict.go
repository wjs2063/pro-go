/*
golang 의 map 에 대한 예제입니다.
*/
package main

import (
	"fmt"
	"strconv"
)

func main() {

	// dict 선언
	m_dict := make(map[string]int)
	dict := map[string]int{}

	for i := 0; i < 10; i++ {
		dict[strconv.Itoa(i)] = i
		m_dict[strconv.Itoa(i)] = i
	}
	// 모든 key,val 을 순회해봅니다
	for key, val := range dict {
		fmt.Printf("key : %s, val : %d\n", key, val)
	}

	// 50이있는지 체크
	ok, exist := dict["50"]
	if !exist {
		fmt.Println("50은없어요")
	} else {
		fmt.Println(ok)
	}

	// 특정 key를 삭제해봅니다
	delete(dict, "1")

	// 그리고나서 체크해보자
	// 50이있는지 체크
	ok, exist = dict["1"]
	if !exist {
		fmt.Println("1은없어요")
	} else {
		fmt.Println(ok)
	}

	// 굿

}
