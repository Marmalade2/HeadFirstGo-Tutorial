package main

import "fmt"

func main() {
	// 섀도잉 -> 동일한 스코프에 변수, 타입, 메서드 등의 이름이 겹치면 이미 존재하는 대상을 shadow한다.

	/*
		var int int = 12
		var append string = "minutes of bonus footage"
		var fmt string = "DVD"
		var count int // 위 선언된 int 참조 (no type int)
		var languages = append([]string{}, "Espanol") // 위 선언된 append 참조 (no method append)
		fmt.Println(int, append, "on", fmt, languages) // 위 선언된 fmt 참조 (no package fmt)
	*/

	// 따라서 같은 스코프 내에서 미리 사용된 이름을 사용하지 않는것이 좋다.
	var count int = 12
	var suffix string = "minutes of bonus footage"
	var format string = "DVD"
	var languages = append([]string{}, "Espanol")
	fmt.Println(count, suffix, "on", format, languages)
}
