package main

import "fmt"
const red ="\033[31m"
const yellow = "\033[33m"
const reset = "\033[0m"
func printBorder(columns int) {
	fmt.Print ("+")
	for i:=0; i< columns; i++ {
		fmt.Print("---------+")

	}
	fmt.Println()
}
func makeSpaces(count int) string {
	result := ""
	for i:=0; i<count; i++ {
		result +=" "
	}
	return result
}
func makeCell(text string, color string) string {
 	spaces := 9 - len(text)
 	leftSpaces := spaces / 2
 	rightSpaces := spaces - leftSpaces

 	cell := makeSpaces(leftSpaces) + text + makeSpaces(rightSpaces)

 	if color == "red" {
  		return red + cell + reset
 	}

 	if color == "yellow" {
  		return yellow + cell + reset
 	}

 return cell
}
