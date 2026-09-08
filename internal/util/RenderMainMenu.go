package util

import (
	"fmt"

	"github.com/fatawaimamalmuftin/koda-b9-golang/controller"
	"github.com/fatawaimamalmuftin/koda-b9-golang/internal/model"
)

func RenderMainMenu() {
	model.Clear()
	for {
		fmt.Println()
		fmt.Println("========================================")
		fmt.Println(" MENU UTAMA ")
		fmt.Println("========================================")
		fmt.Println(" 1. Build Persegi Panjang")
		fmt.Println(" 2. Hitung Keliling")
		fmt.Println(" 3. Hitung Luas")
		fmt.Println(" 4. Hitung Keliling & Luas")
		fmt.Println(" 5. Build Window")
		fmt.Println(" 6. Insert Array")
		fmt.Println(" 7. Minitask 6")
		fmt.Println(" 8. Minitask 7")
		fmt.Println(" 0. Keluar")
		fmt.Println("========================================")

		controller.Ask()
	}
}