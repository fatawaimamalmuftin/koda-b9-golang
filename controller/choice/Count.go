package choice

// rumus Luas : panjang x lebar
func Luas(panjang uint8, lebar uint8) uint8 {
	Luas := panjang * lebar
	return Luas
}

// rumus Keliling : 2 x (panjang x lebar)
func Keliling(panjang uint8, lebar uint8) uint8 {
	Keliling := 2 * (panjang * lebar)
	return Keliling
}

// Keliling dan lebar
func LnK(panjang uint8, lebar uint8) (hasilKeliling uint8, hasilLuas uint8){
	hasilKeliling = Keliling(panjang, lebar)
	hasilLuas = Luas(panjang, lebar)
	return hasilKeliling, hasilLuas
}