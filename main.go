package main

import (
	"fmt"
	"github.com/TehreemJaved/quiz2/statuspackage"
)

// Keep count of functions visited
var pk, sk, fr, pm bool
func main() {
	var run bool
	for !(run) {
		fmt.Println("1-Print Covid19 cases Pakistan")
		fmt.Println("2-Print Covid19 cases SouthKorea")
		fmt.Println("3-Print Covid19 cases France")
		fmt.Println("4-Print my personalized message to coronavirus")
		fmt.Println("0-Exit")
		var input int
		fmt.Scanln(&input)
		switch input {
		case 0:
			if !(pk && sk && fr && pm) {
				fmt.Println("Cant exit without going through all the options.")
			} else {
				run=true
			}
		case 1:
			pk=statusprint.Pak()
		case 2:
			sk=statusprint.SK()
		case 3:
			fr=statusprint.France()
		case 4:
			pm=statusprint.PerM()
		default:
			fmt.Println("Invalid number entered. Choose again!")
		}
	}
}
