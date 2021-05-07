package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Print("Input your score:")        // 입력 유도 프롬프트
	reader := bufio.NewReader(os.Stdin)   // 키보드 입력 감지
	input, err := reader.ReadString('\n') // 개행문자까지 입력한 내용 입력
	// 에러 발생시 대응
	if err != nil {
		log.Fatal(err)
	}

	//공백 제거
	input = strings.TrimSpace(input)
	// 타입 변환
	// grade 는 선언, err는 할당  => 매번 새로 선언하는 번거로움을 줄이기 위해 이와같은 방식 지원
	grade, err := strconv.ParseFloat(input, 64)

	if err != nil {
		log.Fatal(err)
	}

	var status string // 스코프를 생각한 변수 선언
	if grade >= 60 {
		status = "passing"
	} else {
		status = "failed"
	}

	fmt.Println("You are", status)
}
