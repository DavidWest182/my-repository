package input

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Inpfromkey() (float64, error) {

	inputFromUser := bufio.NewReader(os.Stdin)
	inputStr, err := inputFromUser.ReadString('n')
	if err != nil {
		return 0, fmt.Errorf("invalid input or system error")
	}
	inputTrimed := strings.TrimSpace(inputStr)
	inputParsed, err := strconv.ParseFloat(inputTrimed, 64)
	if err != nil {
		return 0, fmt.Errorf("invalid input or system error")
	}
	return inputParsed, nil

}
