package main

import (
	"fmt"
	"os"
)

func main() {
	defer foo()


	_, err := os.Open("no-file.txt")

	if err != nil {

		//log.Fatalln(err)
		//log.Panicln(err)
		panic(err)
	}

}

func foo(){

	fmt.Println("when os.Exit() is called, deferred function don't run")

}


