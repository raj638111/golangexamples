// `main` is a special package name which denotes main entry point of the application
package main

import "fmt"

func main() {
	fmt.Print("Hello world!")
	fmt.Println("Hellow world!!")
	var str = ""
	var str2 string = ""
	fmt.Printf("This is a string %v\n", str)
	fmt.Print(str2)
	ebt, profit, ratio := calculate(1, 2, 3)
}

/**
	1. If all the parameters are of same type, one type annotation in the end should be enough
	2. Return types can also be given name if required
*/
func calculate(revenue, expenses, taxRate float64) (float64, float64, float64) {
	ebt := revenue - expenses
	profit := ebt * (1 - taxRate / 100)
	ratio := ebt / profit
	return ebt, profit, ratio

}