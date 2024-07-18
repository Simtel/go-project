package main

import (
	"errors"
	"fmt"
	"go-project/contracts"
	"go-project/models"
	"go-project/users"
	"os"
	"time"
)

func main() {
	// Инициализация пользователя
	user := &models.User{Name: "Simtel", Lastname: "Simuls", Email: "email@example.com"}
	contact := users.Contact{Phone: "4343434343", Address: "address home"}
	user.Contact = contact

	// Вывод информации о пользователе
	fmt.Println("Имя пользователя:", user.Name)
	fmt.Println("Полное имя:", user.GetFullName())

	// Аргументы командной строки
	args := os.Args
	fmt.Println("Аргументы командной строки (количество):", len(args))

	// Инициализация локаций
	location := &models.Location{Id: 1, Name: "Ulyanovsk"}
	parentLocation := &models.Location{Id: 2, Name: "Russia"}
	location.SetParent(parentLocation)

	// Добавление региона к локации
	addRegion(location)

	scores := []int{1, 2, 3, 4, 5}
	scores = append(scores, 6)

	sum, errorSum := Sum(2, 1)
	fmt.Println("Сумма", sum)
	if errorSum != nil {
		fmt.Println("Ошибка", errorSum)
	}
	sum2, errorSum2 := Sum(1, 2)
	if errorSum2 == nil {
		fmt.Println("Сумма 2:", sum2)
	}
	fmt.Println("Ошибка2:", errorSum2)

	fmt.Println("Кол-во элементов в массиве", len(arrays()))
	fmt.Println("Кол-во элементов в срезе", len(scores))
	fmt.Println("Кол-во элементов в карте", len(makeMap()))
	// Вывод информации о локациях
	fmt.Println("Локация:", location.GetName())
	fmt.Println("Родительская локация:", location.Parent.GetName())

	// Вывод текущей даты (вместо устаревшего time.DateOnly, так как такого метода нет в стандартной библиотеке)
	fmt.Println("Текущая дата:", time.Now().Format("2006-01-02"))

	myModels := []contracts.Models{user, location}
	for _, model := range myModels {
		fmt.Println(model.GetName())
	}

	ShowDomains()
}

func addRegion(l *models.Location) {
	l.Name = "Region:" + l.Name
}

func arrays() [2]int {
	var ar [2]int
	var resultAr [2]int
	ar[0] = 1
	ar[1] = 2
	for index, element := range ar {
		resultAr[index] = element + 10
	}
	return resultAr
}

func makeMap() map[string]int {
	source := make(map[string]int)

	source["first"] = 1
	source["second"] = 2

	delete(source, "first")

	return source
}

func Sum(a int, b int) (int, error) {
	if b > a {
		return 0, errors.New("b must be less than a")
	}
	return a + b, nil
}
