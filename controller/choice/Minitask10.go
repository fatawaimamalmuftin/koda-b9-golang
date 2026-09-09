package choice

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/fatawaimamalmuftin/koda-b9-golang/internal/model"
)

type note struct{
	writer string
	notelog string
}

func (n note) addNote(whiteBoard *[]note){
	*whiteBoard = append([]note{n}, *whiteBoard...)
}

func writer(name string, noteText string, whiteBoard *[]note, wbCH chan string){
	newNote := note{
		writer : name,
		notelog : noteText,
	}

	newNote.addNote(whiteBoard)

	wbCH <- "berhasil di tambahkan"
}
var whiteBoardDummy = []note{{
		writer: "muftin",
		notelog: "apa aja bisa",
	},{
		writer: "fajar",
		notelog: "apa aja bisa",
	},{
		writer: "cupan",
		notelog: "apa aja bisa",
	},{
		writer: "cana",
		notelog: "apa aja bisa",
	},{
		writer: "pijo",
		notelog: "apa aja bisa",
	},
}

func WhiteBoard(){
	model.Clear()
	for{
		for i,v := range whiteBoardDummy{
			fmt.Printf("%d . %s : %s \n",i+1, v.writer, v.notelog)
		}
	
		reader := bufio.NewReader(os.Stdin)
		wbCH := make(chan string)
		
		fmt.Print("Tekan Y untuk nambah pesan Y/N : ")
		YorN,_ := reader.ReadString('\n')
		YorN = strings.ToLower(strings.TrimSpace(YorN))
		
		if YorN == "y"{
			fmt.Print("Masukan nama anda -> ")
			name,_ := reader.ReadString('\n')
			name = strings.TrimSpace(name)
	
			fmt.Print("Masukan note anda -> ")
			note,_ := reader.ReadString('\n')
			note = strings.TrimSpace(note)
	
			go writer(name, note, &whiteBoardDummy, wbCH)
			
			massage := <- wbCH
			fmt.Println(massage)
		}
		
		if YorN == "n" {
			return
		}
	}
}
