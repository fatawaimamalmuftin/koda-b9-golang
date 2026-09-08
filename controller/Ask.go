package controller

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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

	case "2":
		fmt.Print("Masukan Panjang -> ")
		p, _ := reader.ReadString('\n')
		p = strings.TrimSpace(p)
		panjang, _ := strconv.Atoi(p)

		fmt.Print("Masukan Luas -> ")
		L, _ := reader.ReadString('\n')
		L = strings.TrimSpace(L)
		lebar, _ := strconv.Atoi(L)

		hasilKeliling := choice.Keliling(uint8(panjang), uint8(lebar))

		fmt.Printf("Keliling Persegi panjang ada lah : %d\n", hasilKeliling)

	case "3":
		fmt.Print("Masukan Panjang -> ")
		p, _ := reader.ReadString('\n')
		p = strings.TrimSpace(p)
		panjang, _ := strconv.Atoi(p)

		fmt.Print("Masukan Luas -> ")
		L, _ := reader.ReadString('\n')
		L = strings.TrimSpace(L)
		lebar, _ := strconv.Atoi(L)

		hasilLuas := choice.Luas(uint8(panjang), uint8(lebar))

		fmt.Printf("Luas Persegi panjang ada lah : %d\n", hasilLuas)

	case "4":
		fmt.Print("Masukan Panjang -> ")
		p, _ := reader.ReadString('\n')
		p = strings.TrimSpace(p)
		panjang, _ := strconv.Atoi(p)

		fmt.Print("Masukan Luas -> ")
		L, _ := reader.ReadString('\n')
		L = strings.TrimSpace(L)
		lebar, _ := strconv.Atoi(L)

		hasilKeliling, hasilLuas := choice.LnK(uint8(panjang), uint8(lebar))

		fmt.Printf("Keliling dan Luas Persegi panjang ada lah : Keliling=%d, Luas=%d\n", hasilKeliling, hasilLuas)

	case "5":
		fmt.Print("Masukan Luas Window -> ")
		L, _ := reader.ReadString('\n')
		L = strings.TrimSpace(L)

		lebar, _ := strconv.Atoi(L)

		window := choice.BuildWindow(uint8(lebar))
		if window != nil{
			fmt.Printf("error : %s\n", window)
			return
		}

		fmt.Println(window)

	case "6":
		result := choice.InsertArry([]int{50, 75, 66, 20, 32, 90}, 88)
		fmt.Println(result)

	case "0":
		fmt.Println("Program selesai...")
		os.Exit(0)

	default:
		fmt.Println("Menu tidak tersedia")
	}
}