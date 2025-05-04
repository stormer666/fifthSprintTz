package personaldata

import (
	"errors"
	"fmt"
	"log"
)

type Personal struct {
	// TODO: добавить поля
	Name   string
	Weight float64
	Height float64
}

func (p Personal) Print() {
	// TODO: реализовать функцию
	if p.Name == "" {
		err := errors.New("the name should not be empty")
		log.Println(err)
		return
	}
	if p.Weight <= 0.0 {
		err := errors.New("zero weight output value")
		log.Println(err)
		return
	}
	if p.Height <= 0.0 {
		err := errors.New("zero height output value")
		log.Println(err)
		return
	}
	fmt.Printf("Имя: %s\nВес: %.2f кг.\nРост: %.2f м.\n\n", p.Name, p.Weight, p.Height)
}
