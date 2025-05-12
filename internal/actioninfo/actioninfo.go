package actioninfo

import (
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
	for _, v := range dataset {
		err := dp.Parse(v)
		if err != nil {
			log.Printf("error Parse() for '%s': %v", v, err)
			continue
		}

		actionInfo, err := dp.ActionInfo()
		if err != nil {
			log.Printf("error ActionInfo() '%s': %v", v, err)
			continue
		}
		fmt.Println(actionInfo)
	}
}
