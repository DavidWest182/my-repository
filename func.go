// calculator of paint consumption for a wall
package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// function take input from keyboard and return float value
func input() (float64, error) {

	inp := bufio.NewReader(os.Stdin)
	strValue, err := inp.ReadString('\n') //string value
	procesErr(err)
	trimStrValue := strings.TrimSpace(strValue)        //trimed string value
	value, err := strconv.ParseFloat(trimStrValue, 64) //float value
	procesErr(err)
	// input function error return
	if value <= 0 {
		err := errors.New("error")
		return 0, err
	}
	return value, nil
}

// function of program math
func consCalc(w, h, c float64) float64 {
	return w * h / c
}

// processing errors
func procesErr(err error) {

	if err != nil {
		fmt.Errorf("вы ввели не допустимое значение")
		main()
	}
}

func main() {

	// program body
	fmt.Print("Введите ширину стены:")
	width, err := input()
	procesErr(err)
	fmt.Print("Введите высоту стены:")
	height, err := input()
	procesErr(err)
	fmt.Print("Введите расход краски:")
	cons, err := input() // consumption of paint
	procesErr(err)
	fmt.Printf("вам понадобится %.2f литров краски.\n", consCalc(width, height, cons))
	fmt.Print("Для завершения нажмите \"Enter\".")
	//close program with enter
	close := bufio.NewReader(os.Stdin)
	close.ReadString('\n')
}
