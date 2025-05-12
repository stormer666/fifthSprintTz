package trainings

import (
	"errors"
	"fmt"
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
	data := strings.Split(datastring, ",")
	if len(data) != 3 {
		return errors.New("invalid string length")
	}

	steps, err := strconv.Atoi(data[0])
	if err != nil {
		return fmt.Errorf("invalid number of steps")
	}
	if steps <= 0 {
		return errors.New("error steps count")
	}
	timeDuration, err := time.ParseDuration(data[2])
	if err != nil || timeDuration <= 0 {
		return errors.New("invalid duration format or non-positive duration")
	}

	t.Steps = steps
	t.TrainingType = data[1]
	t.Duration = timeDuration

	return nil
}

func (t Training) ActionInfo() (string, error) {
	// TODO: реализовать функцию
	distance := spentenergy.Distance(t.Steps, t.Height)
	speed := spentenergy.MeanSpeed(t.Steps, t.Height, t.Duration)
	var err error
	var ccal float64

	switch t.TrainingType {
	case "Бег":
		ccal, err = spentenergy.RunningSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
	case "Ходьба":
		ccal, err = spentenergy.WalkingSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
	default:
		return "", fmt.Errorf("unknown training type: %s", t.TrainingType)
	}
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n",
		t.TrainingType, t.Duration.Hours(), distance, speed, ccal), nil
}
