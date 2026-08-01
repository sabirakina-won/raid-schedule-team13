package main 

import "fmt"

func printMessage(rows []string, busy string, important string) {
	errorMessage := validate(rows)

	if errorMessage != "" {
		fmtPrintln(errorMessage)
		return
	}

columns := len(rows[0])

for _, row := range rows {
	printBorder (columns)
	fmt.Print("|")

	for i := 0; i < columns; i++ {
		cell := string(row[i])

		if cell == "0" {
			fmr.Print(makeCell("", ""))
		} else if cell == "1" {
			fmt.Print(maleCell(busy, "red"))
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
	printSchedule([]string{"0102", "01100", "2001"}, "###", "!!!")

	fmt.Println()
    
	printSchedule([]string{"01020", "011002", "00210"}, "*", "$")

	fmt.println()

	printSchedule([]string{"0102", "1X00"}, "###", "!!!")
}
