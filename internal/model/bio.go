package model

type Bio struct {
	Name string
	UrlProfile string
	Email string
	Age int
	PhonNum string
	IsWedding bool
	Education []Education
}

type Education struct{
	NameEdu string
	NameMaj string
}