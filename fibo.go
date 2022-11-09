package main

import (
	"fmt"
	//"bufio"
	//"os"
	//"strconv"
)

func fibo(n) int{
	if n <= 2{
		return 1
	}

    else{
        return fibo(n - 1) + fibo(n - 2)
	}

}
func main() {
	fmt.Printf(fibo(100))

}