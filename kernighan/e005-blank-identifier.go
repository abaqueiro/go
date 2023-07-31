/* Listing command line arguments received */
package main

import (
	"fmt"
	"os"
)

func main(){
	fmt.Println("POSITION: VALUE")
	i := 1
	// _ is a blank identifier, used to ignore a variable
	for _, value := range os.Args {
		fmt.Printf("%d %v\n", i, value)
		i++
	}
}
