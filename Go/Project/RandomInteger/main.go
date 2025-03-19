package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

///////////////////////////////////////////////////////////
// Random int game
///////////////////////////////////////////////////////////
// Object
// 1. 0 ~ 99 사이 랜덤 숫자 하나를 설정
// 2. 사용자 입력 저장
// 3. 입력값과 랜덤값을 비교하여 비교 연산자를 활용해 결과 출력
// 4. 서로 값이 맞을 경우 Return의 i++값을 증가시키고 출력
// 5. 프로그램 종료
///////////////////////////////////////////////////////////

var stdin = bufio.NewReader(os.Stdin)

func InputIntValue() (int, error) {
	var n int
	_, err := fmt.Scanln(&n)

	if err != nil {
		stdin.ReadString('\n')
	}

	return n, err
}

func main() {
	setTime := time.Now().UnixNano()

	rand.Seed(setTime) // Set random time [UnixNano]

	r := rand.Intn(100)
	cnt := 1

	fmt.Println(r)
	for {
		fmt.Printf("Input value :")
		n, err := InputIntValue()

		if err != nil {
			fmt.Println("Only input to int")
		} else {
			if n > r {
				fmt.Println("Input value > Random vlaue")
			} else if n < r {
				fmt.Println("Input value < Random value")
			} else {
				fmt.Println("Try input : ", cnt)
				break
			}

			cnt++
		}
	}
}
