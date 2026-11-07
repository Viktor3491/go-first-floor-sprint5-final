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

// Данные о дневной активности.
type DaySteps struct {
	Steps    int
	Duration time.Duration
	personaldata.Personal
}

// Парсит строку и заполняет поля.
func (ds *DaySteps) Parse(datastring string) (err error) {
	parts := strings.Split(datastring, ",")

	// Проверяем, что элементов ровно 2
	if len(parts) != 2 {
		return errors.New("неверный формат данных прогулки")
	}

	// Парсим шаги
	steps, err := strconv.Atoi(parts[0])
	if err != nil {
		return err
	}
	// Проверка на ноль и отрицательные значения
	if steps <= 0 {
		return errors.New("количество шагов должно быть больше нуля")
	}
	ds.Steps = steps

	// Парсим длительность
	duration, err := time.ParseDuration(parts[1])
	if err != nil {
		return err
	}
	// Проверка на ноль и отрицательные значения
	if duration <= 0 {
		return errors.New("продолжительность должна быть больше нуля")
	}
	ds.Duration = duration

	return nil
}

// Формирует строку с информацией о прогулке.
func (ds DaySteps) ActionInfo() (string, error) {
	// Вычисляем дистанцию
	distance := spentenergy.Distance(ds.Steps, ds.Height)

	// Вычисляем калории
	calories, err := spentenergy.WalkingSpentCalories(ds.Steps, ds.Weight, ds.Height, ds.Duration)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf(
		"Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n",
		ds.Steps,
		distance,
		calories,
	), nil
}
