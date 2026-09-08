package choice

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)


type Person struct{
	Name string
	Addres string
	Phone string
}

// Constraction
func NewPerson(Name string, Addres string, Phone string) *Person{
	return &Person{
		Name: Name,
		Addres: Addres,
		Phone: Phone,
	}
}

// Getter Show all person
func (p *Person) GetPerson() string{
	return fmt.Sprintf("Nama : %s\nAlamat : %s\nPhone : %s", p.Name, p.Addres, p.Phone)
}

// Setter with greet
func (p *Person) SetPerson(newName string){
	p.Name = newName
	p.Greet()
}

// greet
func (greet *Person) Greet(){
	fmt.Printf("\n\n hallo.. %s\n", greet.Name)
}

func Minitask7(){
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Nama -> ")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)

	fmt.Print("Alamat -> ")
	addr, _ := reader.ReadString('\n')
	addr = strings.TrimSpace(addr)

	fmt.Print("No.HP -> ")
	notel, _ := reader.ReadString('\n')
	notel = strings.TrimSpace(notel)

	person := NewPerson(name,addr,notel)
	person.Greet()

	fmt.Print("Masukan Nama baru : ")
	newName,_ := reader.ReadString('\n')
	newName = strings.TrimSpace(newName)

	person.SetPerson(newName)
}