/*
이 모듈은 자료구조 Queue library 인 contanier/list 예제연습용 입니다.
*/

package main

import (
	"container/list"
	"fmt"
)

func main() {

	// list 생성
	q := list.New()

	// 0부터 9 까지 q에 넣자
	for i := 0; i < 10; i++ {
		q.PushFront(i)
	}

	// Front 부터 순회해봅니다.

	for e := q.Front(); e != nil; e = e.Next() {
		fmt.Printf("%d -> ", e.Value)
	}

	fmt.Println()
	// Back 부터 순회해봅니다

	for e := q.Back(); e != nil; e = e.Prev() {
		fmt.Printf("%d -> ", e.Value)
	}

	fmt.Println()
	// 아이템을 삭제해봅니다.
	fmt.Println(q.Len())
	for q.Len() != 0 {
		data := q.Front()
		q.Remove(data)

	}

	fmt.Println(q.Front(), q.Back())

}
