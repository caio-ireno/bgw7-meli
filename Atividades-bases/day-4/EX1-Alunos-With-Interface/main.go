package main

import "fmt"

type StudentManager interface {
	Details() string
	SetName(Name string)
	SetLastName(LastName string)
	SetDNI(DNI int)
	SetDate(Date string)
}

type Student struct {
	Name     string
	LastName string
	DNI      int
	Date     string
}

func (s *Student) SetName(Name string) {
	s.Name = Name
}

func (s *Student) SetLastName(LastName string) {
	s.LastName = LastName
}

func (s *Student) SetDNI(DNI int) {
	s.DNI = DNI
}

func (s *Student) SetDate(Date string) {
	s.Date = Date
}

func (s Student) Details() string {
	dados := "Name: " + s.Name + " Last Name: " + s.LastName + " DNI: " + fmt.Sprint(s.DNI) + " Date: " + s.Date
	return dados
}

func main() {

	var s1 StudentManager = &Student{
		Name:     "João",
		LastName: "Silva",
		DNI:      123456789,
		Date:     "01/01/2000",
	}

	s1.SetName("Maria")
	s1.SetLastName("Santos")
	fmt.Println(s1.Details())

}
