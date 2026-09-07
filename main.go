package main

import "fmt"
func main() {
	// luas(5,10)
	// keliling(5,10)
	K,L := LnK(5,10)
	fmt.Printf("hasil keliling persegi: %d\nhasil luas persegi: %d\n", K, L)
}

// rumus luas : panjang x lebar
func luas(panjang uint8, lebar uint8) uint8 {
	luas := panjang * lebar
	return luas
}

//rumus keliling : 2 x (panjang x lebar)
func keliling(panjang uint8, lebar uint8) uint8 {
	keliling := 2 * (panjang * lebar)
	return keliling
}

// keliling dan lebar
func LnK(panjang uint8, lebar uint8) (hasilKeliling uint8, hasilLuas uint8){
	hasilKeliling = keliling(panjang, lebar)
	// keliling := 2 * (panjang * lebar)
	hasilLuas = luas(panjang, lebar)
	// luas := panjang * lebar	
	// fmt.Printf("hasil keliling persegi: %d\nhasil luas persegi: %d\n", K, L)
	return hasilKeliling, hasilLuas
}