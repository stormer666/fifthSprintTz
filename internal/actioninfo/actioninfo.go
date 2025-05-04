package actioninfo

import (
	"errors"
	"log"
)

type DataParser interface {
	// TODO: добавить методы
	Parse(string) error
	ActionInfo() (string, error)
}

func Info(dataset []string, dp DataParser) {
	// TODO: реализовать функцию
	for v := range dataset {
		g := dataset[v]
		dp.Parse(dataset[v])
		if len(g) <= 0 {
			err := errors.New("string parsing error")
			log.Println(err)
			continue
		}
		dp.ActionInfo()
	}

}
