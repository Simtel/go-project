package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	// Инициализация пользователя
	user := User{name: "Simtel", lastname: "Simuls", email: "email@example.com"}

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

	// Вывод информации о локациях
	fmt.Println("Локация:", location.getName())
	fmt.Println("Родительская локация:", location.parent.getName())

	// Вывод текущей даты (вместо устаревшего time.DateOnly, так как такого метода нет в стандартной библиотеке)
	fmt.Println("Текущая дата:", time.Now().Format("2006-01-02"))
}

func addRegion(l *Location) {
	l.name = "Region:" + l.name
}
