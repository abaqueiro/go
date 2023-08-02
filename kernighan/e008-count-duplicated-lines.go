/* Count duplicated lines from standard input */
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main(){
	// create a map that works as an associative array or dictionary
	countH := make( map[string]int )
	// bufio contains methods for buffered I/O
	// a Scanner has methods to read a line and the return the text read
	// os.Stdin is a file handle to standard input stream
	scanner := bufio.NewScanner( os.Stdin )
	// scanner.Scan returns true on a line loaded or false on end of stream
	for scanner.Scan() {
		// access the line read
		line := scanner.Text()
		countH[ line ]++
	}

	msg := "------------------------------------------------------------"
	fmt.Println(msg)
	fmt.Println("TIMES\tLINE")
	fmt.Println(msg)
	for line, count := range(countH) {
		if count > 1 {
			fmt.Print(count)
			fmt.Print("\t")
			fmt.Print(line)
			fmt.Println("")
		}
	}
}
