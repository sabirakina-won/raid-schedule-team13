package main

import "fmt"

func printSchedule(rows []string, busy string, important string) {
	errorMessage := validate(rows)

	if errorMessage != "" {
		fmt.Println(errorMessage)
		return
	}

	columns := len(rows[0])

	for _, row := range rows {
		printBorder(columns)
		fmt.Print("|")

		for i := 0; i < columns; i++ {
			cell := string(row[i])

			if cell == "0" {
				fmt.Print(makeCell("", ""))
			} else if cell == "1" {
				fmt.Print(makeCell(busy, "red"))
			} else {
				fmt.Print(makeCell(important, "yellow"))
			}

			fmt.Print("|")
		}

		fmt.Println()
	}

	printBorder(columns)
}

func main() {
	printSchedule([]string{"0102", "0110", "2001"}, "###", "!!!")

	fmt.Println()

	printSchedule([]string{"01020", "01102", "00210"}, "*", "$")

	fmt.Println()

	printSchedule([]string{"0102", "1X00"}, "###", "!!!")
}
