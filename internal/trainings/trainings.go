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

	trainingType := strings.TrimSpace(str[1])
	if trainingType == "" {
		err := errors.New("invalid training type; cannot be empty")
		log.Println(err)
		return err
	}
	t.TrainingType = trainingType

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
	var err error
	switch t.TrainingType {
	case "Ходьба":
		ccal, err = spentenergy.WalkingSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
		if err != nil {
			log.Printf("Error while calculating calories for Ходьба: %v", err)
			return "", fmt.Errorf("error calculating calories for Ходьба: %w", err)
		}
	case "Бег":
		ccal, err = spentenergy.RunningSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
		if err != nil {
			log.Printf("Error while calculating calories for Бег: %v", err)
			return "", fmt.Errorf("error calculating calories for Бег: %w", err)
		}
	default:
		err := fmt.Errorf("unknown training type: %s", t.TrainingType)
		log.Println(err)
		return "", err
	}
	return fmt.Sprintf("Тип тренировки:%s\nДлительность:%.2f ч.\n,Дистанция:%.2f км.\nСкорость:%.2f км/ч\nСожгли калорий:%.2f\n", t.TrainingType, float64(t.Duration)/float64(time.Hour), distance, averageSpeed, ccal), nil
}
