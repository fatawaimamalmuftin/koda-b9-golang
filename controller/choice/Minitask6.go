package choice

import (
	"bufio"
	"fmt"
	"os"
)

func Minitask6(path string) {
	defer func ()  {
		if err := recover(); err != nil {
			fmt.Println("Recovered panic . . .")			
		}	
	}()

	read, err := os.Open(path)

	if err != nil {
		fmt.Println("error os.open", err)
		return
	}
	defer func(){
		read.Close()
		fmt.Println("berhasil tutup . . . ")
	}()
	
	scan := bufio.NewScanner(read)

	for scan.Scan(){
		fmt.Println(scan.Text())
	}

	if err := scan.Err(); err != nil {
		panic(fmt.Sprintf("error scan line 25 -> %s\n", err))
	}
}