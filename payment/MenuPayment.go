package payment

import (
	"fmt"
)

func Payment(){
	for{
		fmt.Println()
		fmt.Println("========================================")
		fmt.Println(" PAYMENT ")
		fmt.Println("========================================")
		fmt.Println(" 1. Paymen with Bank")
		fmt.Println(" 2. Paymen with Online")
		fmt.Println(" 0. Kembali ke menu utama")
		fmt.Println("========================================")

		AskPayment()
	}
}