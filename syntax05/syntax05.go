package main

import (
	"fmt"
)

// shadowing
func main() {
	// 자료타입을 변수명으로 사용
	var test1 float64 = 9.1
	var test2 float64 = 7.9
	var univ string = "inha"

	var f1 string = "functions"
	var f2 = append([]string{}, "함수")

	fmt.Println(test1 + test2)
	fmt.Println(univ)
	fmt.Println(f1)
	fmt.Println(f2)

	//패키지를 변수명으로 사용
	//var fmt string = "inha"
	//fmt.Println()

	//함수를 변수명으로 사용
	//var append string = "functions"
	//var f = append([]string{}, "함수")
}

/*package main

import (
	"fmt"
	"log"
	"os"
)

// 입력(0과 1처리))된 수의 소수 판정 프로그램 v0.9
func main() {
	var number int

	fmt.Print("정수 입력 : ")
	_, err := fmt.Scanln(&number)

	//fmt.Println(n)
	if err != nil {
		log.Fatal(err)
	}
	for number < 2 {
		fmt.Println(number, "는(은) 소수가 아닙니다!")
		os.Exit(0)
	}

	isPrime := true

	for i := 2; i < number; i++ {
		if number%i == 0 {
			isPrime = false
			break
		}
	}

	if isPrime {
		fmt.Println(number, "는(은) 소수입니다!")
	} else {
		fmt.Println(number, "는(은) 소수가 아닙니다!")
	}

}*/

/*package main

import (
	"fmt"
	"log"
)

// 입력(fmt 패키지의 Scanln())된 수의 소수 판정 프로그램 v0.7
func main() {
	var number int
	fmt.Print("정수 입력 : ")
	_, err := fmt.Scanln(&number)

	//fmt.Println(n)

	if err != nil {
		log.Fatal(err)
	}

	isPrime := true

	for i := 2; i < number; i++ {
		if number%i == 0 {
			isPrime = false
			break
		}
	}

	if isPrime {
		fmt.Println(number, "는(은) 소수입니다!")
	} else {
		fmt.Println(number, "는(은) 소수가 아닙니다!")
	}

}*/

/*package main

import (
	"fmt"
	"strconv"
	"strings"

	//"math/rand"
	"bufio"
	"log"
	"os"
	//"time"
)

// 입력된 수의 소수 판정 프로그램 v0.7
func main() {
	fmt.Print("정수 입력 : ")
	rd := bufio.NewReader(os.Stdin)
	in, err := rd.ReadString('\n')

	if err != nil {
		log.Fatal(err)
	}

	in = strings.TrimSpace(in)
	number, err := strconv.Atoi(in)
	if err != nil {
		log.Fatal(err)
	}

	isPrime := true

	for i := 2; i < number; i++ {
		if number%i == 0 {
			isPrime = false
			break
		}
	}

	if isPrime {
		fmt.Println(number, "는(은) 소수입니다!")
	} else {
		fmt.Println(number, "는(은) 소수가 아닙니다!")
	}

}
*/
