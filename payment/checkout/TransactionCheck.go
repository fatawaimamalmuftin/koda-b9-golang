package checkout

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type interTrans interface{
	Pay() string
}

type Bank struct {}
func (b Bank) Pay() string{
	return "Bank"
}

type Online struct {}
func (o Online) Pay() string{
	return "Online"
}

func TypePaymen(interTrans interTrans) string{
	return interTrans.Pay()
}

type Tabel struct{
	nominal string
	typePay string
}

var tabel []Tabel

func TransactionCheck(paymentType interTrans){
	typePaymen := TypePaymen(paymentType)
	
	reader := bufio.NewReader(os.Stdin)

	fmt.Println()
	fmt.Println("========================================")
	fmt.Println("           TAMBAH PEMBAYARAN")
	fmt.Println("========================================")
	
	fmt.Printf("Tipe Pembayaran : %s\n", typePaymen)
	fmt.Print("Masukan nominal : ")
	
	price,_ := reader.ReadString('\n')
	price = strings.TrimSpace(price)
	
	transaction := Tabel{
		nominal: price,
		typePay: typePaymen,
	}
	
	tabel = append(tabel, transaction)
	
	fmt.Println("========================================")
}

func ShowTransaction() {
	fmt.Println()
	fmt.Println("========================================")
	fmt.Println("          RIWAYAT PEMBAYARAN")
	fmt.Println("========================================")
	fmt.Printf(" %5s %15s %15s\n", "No", "Nominal", "Metode")
	fmt.Println("------------------------------------------------------")
	
	for i, transaction := range tabel {
		fmt.Printf(" %5d Rp%13s %15s\n",i+1,transaction.nominal, transaction.typePay)
	}
	fmt.Println("========================================")
}