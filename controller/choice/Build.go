package choice

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/fatawaimamalmuftin/koda-b9-golang/internal/model"
)

func Build() {
	reader := bufio.NewReader(os.Stdin)
	var Biodata model.Bio
	var Education model.Education

	fmt.Print("Nama -> ")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)
	// ----------------
	Biodata.Name = name

	fmt.Print("Profile -> ")
	url, _ := reader.ReadString('\n')
	url = strings.TrimSpace(url)
	// ----------------
	Biodata.UrlProfile = url

	fmt.Print("Email -> ")
	email, _ := reader.ReadString('\n')
	email = strings.TrimSpace(email)
	// ----------------
	Biodata.Email = email

	fmt.Print("Umur -> ")
	inputAge, _ := reader.ReadString('\n')
	inputAge = strings.TrimSpace(inputAge)
	age, _ := strconv.Atoi(inputAge)
	// ----------------
	Biodata.Age = age

	fmt.Print("No.Hp -> ")
	phon, _ := reader.ReadString('\n')
	phon = strings.TrimSpace(phon)
	// ----------------
	Biodata.PhonNum = phon

	fmt.Print("Sudah nikah Y/N -> ")
	wed, _ := reader.ReadString('\n')
	wed = strings.TrimSpace(strings.ToLower(wed))
	var isWed bool
	switch wed {
	case "y":
		isWed = true
	case "n":
		isWed = false
	}
	// ----------------
	Biodata.IsWedding = isWed

	fmt.Print("Pendidikan -> ")
	edu, _ := reader.ReadString('\n')
	edu = strings.TrimSpace(edu)
	// ----------------
	Education.NameEdu = edu

	fmt.Print("Jurusan -> ")
	maj, _ := reader.ReadString('\n')
	maj = strings.TrimSpace(maj)
	// ----------------
	Education.NameMaj = maj

	fmt.Println("\n==============================")
	fmt.Println("        BIODATA DIRI")
	fmt.Println("==============================")
	fmt.Println("Nama        :", Biodata.Name)
	fmt.Println("Profile     :", Biodata.UrlProfile)
	fmt.Println("Email       :", Biodata.Email)
	fmt.Println("Umur        :", Biodata.Age)
	fmt.Println("No. HP      :", Biodata.PhonNum)

	if Biodata.IsWedding {
		fmt.Println("Status      : Sudah menikah")
	} else {
		fmt.Println("Status      : Belum menikah")
	}

	fmt.Println("Pendidikan  :", Education.NameEdu)
	fmt.Println("Jurusan     :", Education.NameMaj)
	fmt.Println("==============================")
}
