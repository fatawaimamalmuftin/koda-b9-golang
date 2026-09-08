package getway

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"

	"github.com/fatawaimamalmuftin/koda-b9-golang/internal/model"
)

type InterTrans interface{
	Pay() string
}

type Bank struct {}
func (b Bank) Pay() string{
	return "Bank"
}

type Online struct {}
func (O Online) Pay() string{
	return "Online"
}

func TypePaymen(InterTrans InterTrans) string{
	return InterTrans.Pay()
}



func Transaction(paymentType InterTrans) {
	price := []int{12000, 13000, 14000}
	random := price[rand.Intn(3)]

	typePaymen := TypePaymen(paymentType)

	reader := bufio.NewReader(os.Stdin)

	fmt.Println()
	fmt.Println("========================================")
	fmt.Println("          DETAIL PEMBAYARAN")
	fmt.Println("========================================")
	fmt.Printf(" Metode Pembayaran : %s\n", typePaymen)
	fmt.Printf(" Total Pembayaran   : Rp%d\n", random)
	fmt.Println("========================================")

	fmt.Print(" Masukan Nominal yang sama : ")
	payment, _ := reader.ReadString('\n')
	payment = strings.TrimSpace(payment)

	if payment == strconv.Itoa(random) {
		model.Clear()
		fmt.Println()
		fmt.Println("========================================")
		fmt.Println("          PEMBAYARAN BERHASIL")
		fmt.Println("========================================")
		fmt.Printf(" Metode : %s\n", typePaymen)
		fmt.Printf(" Nominal : Rp%s\n", payment)
		fmt.Println(" Status  : SUCCESS")
		fmt.Println("========================================")
		fmt.Println()
	}
}