package main

import "fmt"

type ISmartphone interface {
	setModel(model string)
	setOS(os string)
	getModel() string
	getOS() string
}

type Smartphone struct {
	model string
	os    string
}

func (s *Smartphone) setModel(model string) {
	s.model = model
}

func (s *Smartphone) getModel() string {
	return s.model
}

func (s *Smartphone) setOS(os string) {
	s.os = os
}

func (s *Smartphone) getOS() string {
	return s.os
}

type Samsung_Galaxy_S21 struct {
	Smartphone
}

func newSamsung_Galaxy_S21() ISmartphone {
	return &Samsung_Galaxy_S21{
		Smartphone: Smartphone{
			model: "Samsung Galaxy S-21",
			os:    "Android",
		},
	}
}

type iPhone14 struct {
	Smartphone
}

func newiPhone14() ISmartphone {
	return &iPhone14{
		Smartphone: Smartphone{
			model: "iPhone 14",
			os:    "iOS",
		},
	}
}

func getSmartphone(phoneType string) (ISmartphone, error) {
	if phoneType == "Samsung_Galaxy_S21" {
		return newSamsung_Galaxy_S21(), nil
	}
	if phoneType == "iPhone14" {
		return newiPhone14(), nil
	}
	return nil, fmt.Errorf("Wrong phone type passed")
}

func main() {
	Samsung_Galaxy_S21, _ := getSmartphone("Samsung_Galaxy_S21")
	iPhone14, _ := getSmartphone("iPhone14")

	printDetails(Samsung_Galaxy_S21)
	printDetails(iPhone14)
}

func printDetails(s ISmartphone) {
	fmt.Printf("Smartphone: %s", s.getModel())
	fmt.Println()
	fmt.Printf("OS: %s", s.getOS())
	fmt.Println()
}
