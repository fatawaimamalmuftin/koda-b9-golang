package choice

func InsertArry(arryNum []int, num int) []int{
	return append(arryNum[:3], append([]int{num}, arryNum[3:]...)...)
}