package choice

import (
	"fmt"
	"sync"
	"time"
)

func mandi(){
	fmt.Println("mulai mandi")
	fmt.Println("sedang mandi . . . ")
	time.Sleep(30 * time.Second)
	fmt.Println("mandi selesai")
}

func buatKopi(){
	fmt.Println("mulai bikin kopi")
	fmt.Println("proses membuat kopi . . . ")
	time.Sleep(5 * time.Second)
	fmt.Println("kopi selesai")
}

func buatSarapan(){
	fmt.Println("mulai bikin sarapan")
	fmt.Println("proses membuat sarapan . . . ")
	time.Sleep(30 * time.Second)
	fmt.Println("selesai sarapan")
}

func rapikanKamar(){
	fmt.Println("mulai merapikan")
	fmt.Println("proses merapikan . . . ")
	time.Sleep(5 * time.Second)
	fmt.Println("rapikan kamar selesai")
}

func Minitask9(){
	var wg sync.WaitGroup
	fmt.Println("mulai kegiatan pagi")

	wg.Go(mandi)
	wg.Go(buatKopi)
	wg.Go(buatSarapan)
	wg.Go(rapikanKamar)

	wg.Wait()

	fmt.Println("berangkat kerja")
}