package personaldata

import (
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
	if p.Weight <= 0 {
		log.Printf("incorrect weight value")
	}
	if p.Name <= "" {
		log.Printf("incorrect name value")
	}
	if p.Height <= 0 {
		log.Printf("incorrect height value")
	}
	fmt.Printf("Имя: %s\nВес: %.2f кг.\nРост: %.2f м.\n\n", p.Name, p.Weight, p.Height)
}
