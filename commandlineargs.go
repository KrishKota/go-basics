
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// To create dynamic array
	arr := make([]string, 0)
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Enter Text: ")
		// Scans a line from Stdin(Console)
		scanner.Scan()
		// Holds the string that scanned
		text := scanner.Text()
		if len(text) != 0 {
			fmt.Println("given text:", text)
			arr = append(arr, text)
		} else {
			break
		}

	}
	// Use collected inputs
	fmt.Println("names printed:", arr)
}
