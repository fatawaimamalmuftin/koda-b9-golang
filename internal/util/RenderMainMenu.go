package util

import (
	"fmt"

	"github.com/fatawaimamalmuftin/koda-b9-golang/controller"
)

func RenderMainMenu(){
	for{
		fmt.Println("===== Main Menu =====")
		fmt.Println("1. Biodata Diri")
		controller.Ask()
	}
}