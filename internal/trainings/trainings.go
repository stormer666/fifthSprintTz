package trainings

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

type Training struct {
	// TODO: добавить поля
	Steps        int
	TrainingType string
	Duration     time.Duration
	personaldata.Personal
}

func (t *Training) Parse(datastring string) (err error) {
	// TODO: реализовать функцию
	str := strings.Split(datastring, ", ")

	if len(str) != 3 {
		err := errors.New("invalid string length")
		log.Println(err)
		return err
	}

	steps, err := strconv.Atoi(str[0])
	if err != nil {
		err := errors.New("step conversion error")
		log.Println(err)
		return err
	}
	if steps <= 0 {
		err := errors.New("incorrect number of steps")
		log.Println(err)
		return err
	}
	t.Steps = steps

	t.TrainingType = str[1]

	timeDuration, err := time.ParseDuration(str[2])
	if err != nil {
		err := errors.New("time conversion error")
		log.Println(err)
		return err
	}
	if timeDuration <= 0 {
		err := errors.New("incorrect amount of time")
		log.Println(err)
		return err
	}
	t.Duration = timeDuration

	return nil
}

func (t Training) ActionInfo() (string, error) {
	// TODO: реализовать функцию
	distance := spentenergy.Distance(t.Steps, t.Height)

	var ccal float64

	averageSpeed := spentenergy.MeanSpeed(t.Steps, t.Height, t.Duration)

	switch t.TrainingType {
	case "Ходьба":
		var err error
		ccal, err = spentenergy.WalkingSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
		if err != nil {
			err := errors.New("the error of counting calories for Ходьба")
			log.Println(err)
			return "", err
		}
	case "Бег":
		var err error
		ccal, err = spentenergy.RunningSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
		if err != nil {
			err := errors.New("the error of counting calories for Бег")
			log.Println(err)
			return "", err
		}
	default:
		err := errors.New("unknown type of training")
		log.Println(err)
		return "", err
	}

	return fmt.Sprintf("Тип тренировки:%s\nДлительность:%.2f\n,Дистанция:%.2f\nСкорость:%.2f\nСожгли калорий:%.2f\n", t.TrainingType, float64(t.Duration)/float64(time.Hour), distance, averageSpeed, ccal), nil
}
