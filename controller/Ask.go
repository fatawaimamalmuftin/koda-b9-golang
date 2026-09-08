package controller

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	choice "github.com/fatawaimamalmuftin/koda-b9-golang/controller/choice"
	"github.com/fatawaimamalmuftin/koda-b9-golang/internal/model"
	"github.com/fatawaimamalmuftin/koda-b9-golang/payment"
)

func Ask() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print(" Pilih menu -> ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	switch input {
	case "1":
		model.Clear()
		fmt.Println("---------- BIODATA DIRI ----------")
		choice.Build()
	
	case "2":
		model.Clear()
		fmt.Println("------------- HITUNG KELILING -------------")
		fmt.Print("Masukkan Panjang -> ")
		p, _ := reader.ReadString('\n')
		p = strings.TrimSpace(p)
		panjang, _ := strconv.Atoi(p)
		fmt.Print("Masukkan Lebar -> ")
		L, _ := reader.ReadString('\n')
		L = strings.TrimSpace(L)
		lebar, _ := strconv.Atoi(L)
		hasilKeliling := choice.Keliling(uint8(panjang), uint8(lebar))
		fmt.Println("--------------------------------------------")
		fmt.Printf("Keliling Persegi Panjang : %d\n", hasilKeliling) 
		
	case "3": 
		model.Clear() 
		fmt.Println("--------------- HITUNG LUAS ----------------") 
		fmt.Print("Masukkan Panjang -> ") 
		p, _ := reader.ReadString('\n') 
		p = strings.TrimSpace(p) 
		panjang, _ := strconv.Atoi(p) 
		fmt.Print("Masukkan Lebar -> ") 
		L, _ := reader.ReadString('\n') 
		L = strings.TrimSpace(L) 
		lebar, _ := strconv.Atoi(L) 
		hasilLuas := choice.Luas(uint8(panjang), uint8(lebar)) 
		fmt.Println("--------------------------------------------") 
		fmt.Printf("Luas Persegi Panjang : %d\n", hasilLuas)
	
	case "4": 
		model.Clear() 
		fmt.Println("---------- HITUNG KELILING & LUAS ----------") 
		fmt.Print("Masukkan Panjang -> ")
		p, _ := reader.ReadString('\n')
		p = strings.TrimSpace(p)
		panjang, _ := strconv.Atoi(p) 
		fmt.Print("Masukkan Lebar -> ")
		L, _ := reader.ReadString('\n')
		L = strings.TrimSpace(L)
		lebar, _ := strconv.Atoi(L)
		hasilKeliling, hasilLuas := choice.LnK( uint8(panjang), uint8(lebar), ) 
		fmt.Println("--------------------------------------------") 
		fmt.Printf("Keliling : %d\n", hasilKeliling) 
		fmt.Printf("Luas : %d\n", hasilLuas)
	
	case "5": 
		model.Clear() 
		fmt.Println("--------------- BUILD WINDOW ---------------") 
		fmt.Print("Masukkan Luas Window -> ") 
		L, _ := reader.ReadString('\n') 
		L = strings.TrimSpace(L) 
		lebar, _ := strconv.Atoi(L) 
		window := choice.BuildWindow(uint8(lebar)) 
		if window != nil { 
			fmt.Printf("Error : %s\n", window) 
		return } 
		fmt.Println("Window berhasil dibuat!") 
		
	case "6": 
		model.Clear() 
		fmt.Println("--------------- INSERT ARRAY ---------------") 
		result := choice.InsertArry( []int{50, 75, 66, 20, 32, 90}, 88, ) 
		fmt.Printf("Hasil Array : %v\n", result) 

	case "7":
		fmt.Print("File path -> ")
		filePath,_ := reader.ReadString('\n')
		filePath = strings.TrimSpace(filePath)
		choice.Minitask6(filePath)

	case "8":
		model.Clear()
		choice.Minitask7()

	case "9":
		model.Clear()
		payment.MenuPayment()
		
	case "0": 
		model.Clear() 
		fmt.Println("========================================") 
		fmt.Println(" Terima kasih sudah mencoba! ") 
		fmt.Println("========================================") 
		os.Exit(0) 
		
	default: 
		model.Clear() 
		fmt.Println("--------------------------------------------") 
		fmt.Println(" Menu tidak tersedia!") 
		fmt.Println(" Silakan pilih menu 1 - 9 dan 0 untuk exit.") 
		fmt.Println("--------------------------------------------") }
}