package spentenergy

import (
	"errors"
	"time"
)

// Основные константы, необходимые для расчетов.
const (
	mInKm                      = 1000 // количество метров в километре.
	minInH                     = 60   // количество минут в часе.
	stepLengthCoefficient      = 0.45 // коэффициент для расчета длины шага на основе роста.
	walkingCaloriesCoefficient = 0.5  // коэффициент для расчета калорий при ходьбе.
)

// Рассчитывает калории при ходьбе.
func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// 1. Проверяем входные данные на корректность
	if steps <= 0 || weight <= 0 || height <= 0 || duration <= 0 {
		return 0, errors.New("некорректные входные параметры")
	}

	// 2. Рассчитываем среднюю скорость
	meanSpeed := MeanSpeed(steps, height, duration)

	// 3. Считаем калории по формуле
	durationInMinutes := duration.Minutes()
	calories := (weight * meanSpeed * durationInMinutes) / minInH

	// 4. Умножаем на коэффициент ходьбы и возвращаем
	return calories * walkingCaloriesCoefficient, nil
}

// Рассчитывает калории при беге.
func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// 1. Проверяем входные данные
	if steps <= 0 || weight <= 0 || height <= 0 || duration <= 0 {
		return 0, errors.New("некорректные входные параметры")
	}

	// 2. Рассчитываем среднюю скорость
	meanSpeed := MeanSpeed(steps, height, duration)

	// 3. Считаем калории
	durationInMinutes := duration.Minutes()
	calories := (weight * meanSpeed * durationInMinutes) / minInH

	return calories, nil
}

// Рассчитывает среднюю скорость.
func MeanSpeed(steps int, height float64, duration time.Duration) float64 {
	// Проверяем время и шаги
	if duration <= 0 || steps < 0 {
		return 0
	}

	// Считаем дистанцию и делим на часы
	distance := Distance(steps, height)
	return distance / duration.Hours()
}

// Рассчитывает пройденную дистанцию в километрах.
func Distance(steps int, height float64) float64 {
	stepLength := height * stepLengthCoefficient
	distance := float64(steps) * stepLength
	return distance / mInKm
}
