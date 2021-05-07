// 난수 생성 후 사용자로부터 입력을받아 생성된 난수를 맞추는 프로그램
package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	// 난수 생성 시드를 위한 현재시간 기반 정수값
	seconds := time.Now().Unix()
	// 난수 생성 시드 -> 해당 과정이 없는경우 계속 동일한 값으로 지정됨
	rand.Seed(seconds)
	// 난수 생성
	target := rand.Intn(100) + 1

	fmt.Println("I've chosen a random number between 1 to 100")
	fmt.Println("Let's guess my number")
	success := false

	for i := 0; i < 10; i++ {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Guess number: ")
		input, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		input = strings.TrimSpace(input)
		guess, err := strconv.Atoi(input)
		if err != nil {
			log.Fatal(err)
		}

		if guess < target {
			fmt.Println("Target number is BIGGER than you guess")
		} else if guess > target {
			fmt.Println("Target number is SMALLER than you guess")
		} else {
			fmt.Println("That's right. Congratulation!!")
			success = true
			break
		}
	}

	if !success {
		fmt.Println("Failed. The number is ", target)
	}

}
