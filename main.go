package main

import (
	"fmt"
	"strings"
)

func main() {
	printSchedule([]string{"0102", "1100", "0010", "2001"}, "###", "!!!")
	printSchedule([]string{"01020", "11002", "00210"}, "*", "$")
}

func printSchedule(rows []string, busy string, important string) {
	// 1. ПРОВЕРКА ВВОДА
	// 1.1 проверка на пустое расписание
	if len(rows) == 0 || len(rows[0]) == 0 {
		fmt.Println("error : empty schedule")
		return
	}
	cols := len(rows[0])

	for _, row := range rows {

		// 1.2 проверка на ряды разной длины
		if len(row) != cols {
			fmt.Println("error: rows have different lengths")
			return
		}

		// 1.3 проверка на символы 0/1/2 (только они допускаются, все остальное ошибка)
		for _, num := range row {
			if num != '0' && num != '1' && num != '2' {
				fmt.Println("error: invalid symbol (only 0/1/2 allowed)")
				return
			}
		}
	}

	// 2. РИСУЕМ ТАБЛИЦУ
	// 2.1 Рисуем пограничную линию (вот +---------+---------+---------+---------+)
	border := "+"
	for i := 0; i < cols; i++{
		border += "---------+"
	}
	fmt.Println(border)

	// 2.2 Рисуем ячейку
	for _, row := range rows {
		fmt.Print("|")
		for _, num := range row {
			switch num {
			case '0':
				fmt.Printf("%s|", center("", 9))
			case '1':
				cell := center(busy, 9)
				fmt.Printf("%s%s%s|", red, cell, reset)
			case '2':
				cell := center(important, 9)
				fmt.Printf("%s%s%s|", yellow, cell, reset)
			}
		}
		fmt.Println()
		fmt.Println(border)
	}
}


// 2.3 Раскрашиваем таблицу
const (
	red = "\033[31m"
	yellow = "\033[33m"
	reset = "\033[0m"
)

// 2.4 Отцентровка ячеек
func center(text string, width int) string {
	left := (width - len(text)) / 2
	right := width - len(text) - left
	return strings.Repeat(" ", left) + text + strings.Repeat(" ", right)
}
