package main

import "fmt"
func main() {
	luas := luas(5,10)
	fmt.Println(luas)

	keliling := keliling(5,10)
	fmt.Println(keliling)

	K,L := LnK(5,10)
	fmt.Println(K,L)
}

// rumus luas : panjang x lebar
func luas(panjang uint8, lebar uint8) (luas uint8){
	luas = panjang * lebar
	return luas
}

//rumus keliling : 2 x (panjang x lebar)
func keliling(panjang uint8, lebar uint8) (keliling uint8){
	keliling = 2 * (panjang * lebar)
	return keliling	
}

// keliling dan lebar
func LnK(panjang uint8, lebar uint8)(K uint8, L uint8){
	K = keliling(5,10)
	L = luas(5,10)
	return K, L
}