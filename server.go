package main

import (
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"time"
)

func build() {

	os.Setenv("GOOS", "js")
	os.Setenv("GOARCH", "wasm")

	cmd := exec.Command("go", "build", "-o", "../dist/main.wasm")

	cmd.Dir = "src"

	err := cmd.Run()

	if err != nil {
		panic("Command can not be executed")
	} else {
		fmt.Println("Build successful")

		time.Sleep(500 * time.Millisecond)

		build()
	}
}

func main() {

	go build()

	port := ":5000"
	fmt.Println("Runnig Server at port:", port)

	http.Handle("/", http.FileServer(http.Dir("./dist")))

	http.ListenAndServe(port, nil)
}
