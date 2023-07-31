/* Listing command line arguments received */
package main

import (
	"fmt"
	"os"
)

func main(){
	// using for as iterator and range to take arguments one at a time
	fmt.Println("POSITION: VALUE")
	for index, value := range os.Args {
		fmt.Printf("%v: %v\n", index, value)
	}
}
