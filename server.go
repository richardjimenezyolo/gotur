package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"time"
)

var file = ""

var reload = true

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

	port := flag.String("p", ":5000", "Listening Port")

	flag.Parse()

	// port := ":5000"

	http.Handle("/", http.FileServer(http.Dir("./dist")))

	fmt.Println("Runnig Server at port:", *port)

	http.ListenAndServe(*port, nil)

}
