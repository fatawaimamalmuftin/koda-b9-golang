package util

import (
	"fmt"

	"github.com/fatawaimamalmuftin/koda-b9-golang/controller"
)

func RenderMainMenu(){
	for{
		fmt.Println("===== Main Menu =====")
		fmt.Println("1. Biodata Diri")
		fmt.Println("2. Hitung Keliling Persegi Panjang")
		fmt.Println("3. Hitung Luas Persegi Panjang")
		fmt.Println("4. Hitung Keliling dan Luas Persegi Panjang")
		fmt.Println("0. Exit")

		controller.Ask()
	}
}