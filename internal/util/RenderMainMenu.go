package util

import (
	"fmt"

	"github.com/fatawaimamalmuftin/koda-b9-golang/controller"
)

func RenderMainMenu() {
	for {
		fmt.Println()
		fmt.Println("========================================")
		fmt.Println(" PROGRAM PERSEGI PANJANG")
		fmt.Println("========================================")
		fmt.Println(" 1. Build Persegi Panjang")
		fmt.Println(" 2. Hitung Keliling")
		fmt.Println(" 3. Hitung Luas")
		fmt.Println(" 4. Hitung Keliling & Luas")
		fmt.Println(" 5. Build Window")
		fmt.Println(" 6. Insert Array")
		fmt.Println(" 0. Keluar")
		fmt.Println("========================================")

		controller.Ask()
	}
}