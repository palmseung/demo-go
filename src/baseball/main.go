package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	//Seed에 항상 다른 값이 들어가도록 설정해야 실행할 때마다 랜덤값이 나온다. 그렇지 않으면 같은값만 계속 나옴
	rand.Seed(time.Now().UnixNano())
	count := 0

	for {
		count++
		// 무작위 숫자 3개를 만든다
		numbers := MakeNumbers()

		// 사용자의 입력을 받는다
		inputNumbers := InputNumbers()

		// 결과를 비교한다
		result := CompareNumbers(numbers, inputNumbers)

		PrintResult(result)

		// 3 strikes 인가?
		if IsGameEnd(result) {
			//게임끝
			break
		}
	}

	//몇 번만에 맞혔는지 출력
	fmt.Printf("%d 번만에 맞혔습니다.\n", count)
}

func MakeNumbers() [3]int {
	// 0 ~ 9 사이에 겹치지 않는 숫자 3개
	var rst [3]int

	for i := 0; i < 3; i++ {
		for {
			duplicated := false
			//Intn(n)은 [0, n) -> 0<= randomNumber < n
			n := rand.Intn(10)
			for j := 0; j < i; j++ {
				if rst[j] == n {
					//겹친다. 다시 뽑는다.
					duplicated = true
					break
				}
			}
			if !duplicated {
				rst[i] = n
				break
			}
		}

	}
	fmt.Println(rst)
	return rst
}

func InputNumbers() [3]int {
	//키보드로부터 겹치지 않는 0 ~ 9 사이의 숫자 3개를 입력받아 반환
	var rst [3]int
	return rst
}

func CompareNumbers(numbers, inputNumbers [3]int) bool {
	//두 개의 숫자 3개를 비교해서 결과를 반환한다
	return true
}

func PrintResult(result bool) {
	fmt.Println(result)
}

func IsGameEnd(result bool) bool {
	// 비교결과가 3 스트라이크인지 확인
	return result
}
