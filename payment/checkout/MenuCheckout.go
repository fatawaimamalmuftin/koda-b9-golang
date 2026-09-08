package checkout

import (
	"fmt"

	"github.com/fatawaimamalmuftin/koda-b9-golang/internal/model"
)

func MenuCheckout(){
	model.Clear()
	for{
		fmt.Println()
		fmt.Println("========================================")
		fmt.Println(" Metode Pembayaran Checkout ")
		fmt.Println("========================================")
		fmt.Println(" 1. Bank")
		fmt.Println(" 2. Online")
		fmt.Println(" 0. Kembali ke menu utama")
		fmt.Println("========================================")

		if AskCheckout() {
			return
		}
		
	}
}