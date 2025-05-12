package daysteps

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
)

type DaySteps struct {
	// TODO: добавить поля
	Steps    int
	Duration time.Duration
	personaldata.Personal
}

func (ds *DaySteps) Parse(datastring string) (err error) {
	// TODO: реализовать функцию
	str := strings.Split(datastring, ", ")

	if len(str) != 2 {
		//err := errors.New("invalid string length !=2")
		//log.Println(err)
		return errors.New("invalid string length !=2")
	}

	steps, err := strconv.Atoi(str[0])
	if err != nil {
		//err := errors.New("step conversion error")
		//log.Println(err)
		return err
	}
	if steps <= 0 {
		//err := errors.New("incorrect number of steps")
		//log.Println(err)
		return errors.New("incorrect number of steps")

	}
	ds.Steps = steps

	timeDuration, err := time.ParseDuration(str[1])
	if err != nil {
		//err := errors.New("time conversion error")
		//log.Println(err)
		return err
	}
	if timeDuration <= 0 {
		//err := errors.New("incorrect amount of time")
		//log.Println(err)
		return errors.New("incorrect amount of time")
	}

	ds.Duration = timeDuration

	return nil
}

func (ds DaySteps) ActionInfo() (string, error) {
	distance := spentenergy.Distance(ds.Steps, ds.Height)
	if distance <= 0 {
		err := errors.New("distance value error")
		log.Println(err)
		return "", err
	}

	burnedCCals, err := spentenergy.WalkingSpentCalories(ds.Steps, ds.Weight, ds.Height, ds.Duration)
	if err != nil {
		err := errors.New("calorie counting error")
		log.Println(err)
		return "", err
	}
	return fmt.Sprintf("Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n", ds.Steps, distance, burnedCCals), nil
}
