package main

import "fmt"

func main() {
	// % -> formatting verb. 형식 지정
	// ex)
	fmt.Printf("%%s 는 %s와 같은 string을 프린트합니다.\n", "abcde")
	fmt.Printf("%%d 는 %d와 같은 decimal을 프린트합니다.\n", 1024)
	fmt.Printf("%%f 는 %f와 같은 float을 프린트합니다.\n", 0.125)
	fmt.Printf("%%t 는 %t와 같은 boolean을 프린트합니다.\n", true, false)
	fmt.Printf("%%v 는 %v, %v, %v와 같이 각 변수에 맞는 형식화한 값을 프린트 합니다.\n", 1.2, "\t", false)
	fmt.Printf("%%#v 는 %#v, %#v, %#v와 같이 프로그래밍 단계에서 보이는 값을 그대로 프린트 합니다.\n", 2, "\t", false)
	fmt.Printf("%%T 는 %T, %T, %T와 같이 각 값의 타입을 프린트 합니다.\n", 2, "\t", true)
	fmt.Printf("%%%% 는 %%의 이스케이프 문자입니다.\n")
	fmt.Printf("만약  %10s(%%10s) or %5d(%%5d)와 같이 %% 와 형식동사 사이에 정수형 숫자를 넣으면, 각 값은 최소 너비를 가집니다.\n", "String", 123)
	fmt.Printf("최소 넓이와 함께 float형 변수를 프린트 할 경우 %7.3f (%%7.3f) 와 같이 소수점자릿수를 지정하여 해당 자릿수까지 반올림한 수를 프린트할 수 있습니다.\n", 1.82377)

}
