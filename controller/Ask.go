package controller

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	choice "github.com/fatawaimamalmuftin/koda-b9-golang/controller/choice"
)


func Ask() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("---------------------")
	fmt.Print("pilih -> ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	switch input {
	case "1":
		choice.Build()
	}
}