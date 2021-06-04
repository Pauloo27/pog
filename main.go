package main

import (
	"fmt"
	"os"
	"time"

	"github.com/Pauloo27/pog/img"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: pog <image path>")
		os.Exit(-1)
	}
	path := os.Args[1]
	go func() {
		err := img.StartDaemon()
		if err != nil {
			panic(err)
		}
	}()
	img.WaitForDaemon()
	err := img.ShowImageWithSize(path, "img", 1, 1, 50, 0)
	if err != nil {
		panic(err)
	}
	time.Sleep(5 * time.Second)
}
