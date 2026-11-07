package actioninfo

import (
	"fmt"
	"log"
)

// Интерфейс.
type DataParser interface {
	Parse(datastring string) (err error)
	ActionInfo() (string, error)
}

// Принимает слайс строк и экземпляр структуры,
// парсит каждую строку и выводит информацию.
func Info(dataset []string, dp DataParser) {
	for _, data := range dataset {
		// Пытаемся распарсить строку
		err := dp.Parse(data)
		if err != nil {
			// Если ошибка парсинга, то логируем и переходим к следующей строке
			log.Println(err)
			continue
		}

		// Формируем строку с информацией
		info, err := dp.ActionInfo()
		if err != nil {
			// Если ошибка при формировании, то логируем
			log.Println(err)
			continue
		}

		// Вывод результата на экран
		fmt.Println(info)
	}
}
