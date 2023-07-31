/* implementation of echo unix command */
package main

import "fmt"
import "os"

func main(){
	var arg_n int = len(os.Args)
	if arg_n > 1 {
		fmt.Print( os.Args[1] )
	}
	for i:=2; i < arg_n; i++ {
		fmt.Print( " " )
		fmt.Print( os.Args[i] )
	}
	fmt.Println("")
}
