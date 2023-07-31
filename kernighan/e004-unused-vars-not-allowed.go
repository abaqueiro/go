/* Listing command line arguments received */
package main

import (
	"fmt"
	"os"
)

func main(){
	fmt.Println("VALUE")
	// will not compile, unused index variable is not allowed
	for index, value := range os.Args {
		fmt.Printf("%v\n", value)
	}
}
