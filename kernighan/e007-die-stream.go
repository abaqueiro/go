/* Genereate a stream of n random numbers between 1 and 6
	n is read from os.Args[1]
	if os.Args[1] is not provided by default n = 60
*/
package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
)

func main() {
	var n int

	if len( os.Args ) > 1 {
		var msg string
		t, err := strconv.ParseInt( os.Args[1], 10, 32 )
		if err != nil {
			msg = fmt.Sprintf( "ERROR: %v\n", err )
			fmt.Fprintf( os.Stderr, msg )
			os.Exit(1)
		}
		n = int(t)
	} else {
		n = 60
	}

	for i:=0; i<n; i++ {
		fmt.Println( "dice came up", 1 + rand.Intn(6) )
	}
}
