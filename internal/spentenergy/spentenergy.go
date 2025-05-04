package spentenergy

import (
	"errors"
	"log"
	"time"
)

// Основные константы, необходимые для расчетов.
const (
	mInKm                      = 1000 // количество метров в километре.
	minInH                     = 60   // количество минут в часе.
	stepLengthCoefficient      = 0.45 // коэффициент для расчета длины шага на основе роста.
	walkingCaloriesCoefficient = 0.5  // коэффициент для расчета калорий при ходьбе.
)

func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// TODO: реализовать функцию
	if steps <= 0 {
		err := errors.New("zero or negative number of steps")
		log.Println(err)
		return 0, err
	}
	if weight <= 0 {
		err := errors.New("incorrect weight value")
		log.Println(err)
		return 0, err
	}
	if height <= 0 {
		err := errors.New("incorrect height value")
		log.Println(err)
		return 0, err
	}
	if duration <= 0 {
		err := errors.New("incorrect time value")
		log.Println(err)
		return 0, err
	}
	averageSpeed := MeanSpeed(steps, height, duration)

	minutes := duration.Minutes()

	gap := weight * averageSpeed * minutes

	gap1 := gap / minInH

	result := gap1 * walkingCaloriesCoefficient

	return result, nil

}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// TODO: реализовать функцию
	if steps <= 0 {
		err := errors.New("zero or negative number of steps")
		log.Println(err)
		return 0, err
	}
	if weight <= 0 {
		err := errors.New("incorrect weight value")
		log.Println(err)
		return 0, err
	}
	if height <= 0 {
		err := errors.New("incorrect height value")
		log.Println(err)
		return 0, err
	}
	if duration <= 0 {
		err := errors.New("incorrect time value")
		log.Println(err)
		return 0, err
	}

	averageSpeed := MeanSpeed(steps, height, duration)

	minutes := duration.Minutes()

	gap := weight * averageSpeed * minutes

	result := gap / minInH

	return result, nil
}

func MeanSpeed(steps int, height float64, duration time.Duration) float64 {
	// TODO: реализовать функцию

	if steps <= 0 {
		err := errors.New("the number of steps must not be less than or equal to zero")
		log.Println(err)
		return 0
	}

	if duration <= 0 {
		err := errors.New("the time must not be less than or equal to zero")
		log.Println(err)
		return 0
	}

	distant := Distance(steps, height)

	hours := duration.Hours()

	averageSpeed := distant / hours

	return averageSpeed

}

func Distance(steps int, height float64) float64 {
	// TODO: реализовать функцию

	stepLength := height * stepLengthCoefficient
	gap := float64(steps) * stepLength
	resultDistance := gap / mInKm

	return resultDistance

}
