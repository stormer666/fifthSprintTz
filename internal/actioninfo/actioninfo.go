package actioninfo

import (
	"errors"
	"fmt"
	"log"
)

type DataParser interface {
	// TODO: добавить методы
	Parse(datastring string) error
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

		result, err := dp.ActionInfo()
		if err != nil {
			log.Printf("error in ActionInfo(): %v", err)
			continue
		}
		fmt.Println(result)
	}

}
