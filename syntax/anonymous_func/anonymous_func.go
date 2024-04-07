/*
익명함수 및 일급함수 예제입니다.
*/
package main

import "fmt"

//type add func(int, int) int

func main() {
	// 가변인자
	sum := func(n ...int) int {
		s := 0
		for _, i := range n {
			s += i
		}
		return s
	}
	// 1 ~ 10 까지 모두 합
	fmt.Println(sum(1, 2, 3, 4, 5, 6, 7, 8, 9, 10))

	// 일급 함수
	/*
		함수 자체가 다른 함수의 파라미터로 사용될수있다
	*/
	add := func(a int, b int) int {
		return a + b
	}
	// add func 정의
	r1 := func(add func(int, int) int, a int, b int) int {
		var s int = add(a, b)
		return s
	}

	fmt.Println(r1(add, 10, 20))

}
