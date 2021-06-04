package main

import (
	"fmt"
	"os"
	"time"

	"github.com/Pauloo27/pog/img"
	"github.com/Pauloo27/pog/utils"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: pog <image path>")
		os.Exit(-1)
	}

	// TODO: check if file exists
	// TODO: check if file is a image
	// TODO: handle http(s) url
	path := os.Args[1]

	go func() {
		err := img.StartDaemon()
		if err != nil {
			panic(err)
		}
	}()
	img.WaitForDaemon()

	// TODO: listen for terminal resize

	utils.MoveCursorTo(1, 1)
	utils.ClearAfterCursor()
	fmt.Printf("%sViewing file %s%s\n", utils.ColorGreen, utils.ColorWhite, path)
	utils.HideCursor()

	// TODO: find a better size than just 50, 0
	err := img.ShowImageWithSize(path, "img", 2, 1, 50, 0)
	if err != nil {
		panic(err)
	}
	// TODO: wait for user interaction
	time.Sleep(5 * time.Second)
	utils.ShowCursor()
}
