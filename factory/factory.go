package factory

import "fmt"

type Lunch interface {
	Cook()
}

type School struct {
}

func (s School) Cook() {
	fmt.Println("School Cook ...")
}

type Family struct {
}

func (f Family) Cook() {
	fmt.Println("Family Cook ......")
}

type Factory interface {
	GetSchoolFood() Lunch
	GetFamilyFood() Lunch
}

type SimpleFactory struct {
}

func NewSimpleFactory() Factory {
	return &SimpleFactory{}
}

func (s SimpleFactory) GetSchoolFood() Lunch {
	return &School{}
}

func (s SimpleFactory) GetFamilyFood() Lunch {
	return &Family{}
}
