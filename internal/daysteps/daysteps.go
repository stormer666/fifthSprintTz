package daysteps

import (
	"errors"
	"fmt"
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
	str := strings.Split(datastring, ",")
	if len(str) != 2 {
		return errors.New("invalid string length, expected 2 components")
	}

	steps, err := strconv.Atoi(str[0])
	if err != nil {
		return fmt.Errorf("failed to convert steps to integer: %w", err)
	}

	if steps <= 0 {
		return errors.New("incorrect number of steps, must be positive")
	}

	ds.Steps = steps

	timeDuration, err := time.ParseDuration(str[1])
	if err != nil {
		return fmt.Errorf("duration format error: %w", err)
	}

	if timeDuration <= 0 {
		return errors.New("mistake duration count")
	}
	ds.Duration = timeDuration

	return nil
}

func (ds DaySteps) ActionInfo() (string, error) {
	if ds.Steps <= 0 || ds.Duration <= 0 || ds.Weight <= 0 || ds.Height <= 0 {
		return "", errors.New("steps, duration, weight, and height must be positive")
	}
	// TODO: реализовать функцию
	distance := spentenergy.Distance(ds.Steps, ds.Height)

	ccal, err := spentenergy.WalkingSpentCalories(ds.Steps, ds.Weight, ds.Height, ds.Duration)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n",
		ds.Steps, distance, ccal), nil
}
