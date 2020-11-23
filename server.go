package main

import (
	"fmt"
	"net/http"
	// "os/exec"
)

// func build() {

// 	cmd := exec.Command("./src/build.sh")

// 	err := cmd.Run()

// 	if err != nil {
// 		panic("Command can not be executed")
// 	} else {
// 		fmt.Println("Build successful")
// 	}
// }

func main() {

	// build()

	port := ":5000"
	fmt.Println("Runnig Server at port:", port)

	http.Handle("/", http.FileServer(http.Dir(".")))

	http.ListenAndServe(port, nil)
}
