package main

import (
	"fmt"
	"go-project/users"
	"os"
	"time"
)

func main() {
	// Инициализация пользователя
	user := User{name: "Simtel", lastname: "Simuls", email: "email@example.com"}
	contact := users.Contact{Phone: "4343434343", Address: "address home"}
	user.contact = contact

	// Вывод информации о пользователе
	fmt.Println("Имя пользователя:", user.name)
	fmt.Println("Полное имя:", user.getFullName())

	// Аргументы командной строки
	args := os.Args
	fmt.Println("Аргументы командной строки (количество):", len(args))

	// Инициализация локаций
	location := &Location{id: 1, name: "Ulyanovsk"}
	parentLocation := &Location{id: 2, name: "Russia"}
	location.setParent(parentLocation)

	// Добавление региона к локации
	addRegion(location)

	scores := []int{1, 2, 3, 4, 5}
	scores = append(scores, 6)

	fmt.Println("Кол-во элементов в массиве", len(arrays()))
	fmt.Println("Кол-во элементов в срезе", len(scores))
	fmt.Println("Кол-во элементов в карте", len(makeMap()))
	// Вывод информации о локациях
	fmt.Println("Локация:", location.getName())
	fmt.Println("Родительская локация:", location.parent.getName())

	// Вывод текущей даты (вместо устаревшего time.DateOnly, так как такого метода нет в стандартной библиотеке)
	fmt.Println("Текущая дата:", time.Now().Format("2006-01-02"))
}

func addRegion(l *Location) {
	l.name = "Region:" + l.name
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
