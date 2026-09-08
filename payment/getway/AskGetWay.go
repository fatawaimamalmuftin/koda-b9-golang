package getway

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/fatawaimamalmuftin/koda-b9-golang/internal/model"
)

func AskCheckout() bool{
	reader := bufio.NewReader(os.Stdin)

	fmt.Print(" Pilih metode pembayaran -> ")
	input,_ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	switch input {
	case "1":
		model.Clear()

	case "2":
		model.Clear()

	case "0":
		model.Clear()
		return true

	default:
		model.Clear()
		fmt.Println("--------------------------------------------") 
		fmt.Println(" Menu tidak tersedia!") 
		fmt.Println(" Silakan pilih menu 1 - 2 dan 0 untuk ke menu utama.") 
		fmt.Println("--------------------------------------------") 
	}
	return false
}