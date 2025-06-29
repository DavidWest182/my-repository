// The game of guessing number
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

	//generator of random values
	seconds := time.Now().Unix()
	rand.Seed(seconds)
	target := rand.Intn(100) + 1

	//program greeting massege
	fmt.Println("Привет,я загадал случайное число от 0 до 100.")
	fmt.Println("Попробуй отгадать.))")

	var cikle int

	//cikle, it takes 10 attempts to user
	for cikle = 0; cikle < 10; cikle++ {
		fmt.Println("У тебя осталось", 10-cikle, "попыток!")

		//take user input value
		fmt.Print("Введи число:")
		keybInp := bufio.NewReader(os.Stdin)
		userInput, err := keybInp.ReadString('\n')
		if err != nil {

			log.Fatal(err)
		}

		//take user input and transformate to integer number
		userValue := strings.TrimSpace(userInput)
		userNumber, err := strconv.Atoi(userValue)
		if err != nil {

			log.Fatal(err)
		}

		if target > userNumber {

			fmt.Println("Ваше число меньше загадоного!")

		} else if target < userNumber {

			fmt.Println("Ваше число больше загадоного!")

		} else {
			fmt.Println("Поздравляю вы отгадали загадоное число:", target)
			fmt.Print("Нажмите \"Enter\" для завершения.")
			keybInp := bufio.NewReader(os.Stdin)
			keybInp.ReadString('\n')
			break
		}
	}

	if cikle == 10 {

		fmt.Println("К сожалению вы не отгадали.")
		fmt.Println("Загаданое число:", target)
		fmt.Print("Нажмите \"Enter\" для завершения.")
		keybInp := bufio.NewReader(os.Stdin)
		keybInp.ReadString('\n')
	}
}
