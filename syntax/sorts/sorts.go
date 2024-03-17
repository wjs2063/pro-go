/*
이 모듈은 정렬에 관한 연습예제입니다.
*/

package main

import (
	"fmt"
	"sort"
	"reflect"
)


func main(){

	slice := make([]int,0,10)
	arr := [11]int{}

	// 모두 역순으로 넣어줍니다
	for i := 10;i>=0;i--{
		slice = append(slice,i)
		arr[10-i] = i
	}

	// 정렬한 결과를 보여줍니다

	sort.Slice(slice,func(i,j int) bool {
		return slice[i] < slice[j]
	})

	
	sort.Ints(arr[:])

	fmt.Println("slice : ",slice,reflect.TypeOf(slice))
	fmt.Println("array : ",arr ,reflect.TypeOf(arr))


}
