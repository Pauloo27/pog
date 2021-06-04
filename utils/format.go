package utils

import "fmt"

const (
	ColorBold   = "\033[1m"
	ColorReset  = "\033[0m"
	ColorRed    = "\033[31m"
	ColorGreen  = "\033[32m"
	ColorYellow = "\033[33m"
	ColorBlue   = "\033[34m"
	ColorWhite  = "\033[39m"
)

func HideCursor() {
	fmt.Printf("\033[?25l")
}

func ShowCursor() {
	fmt.Printf("\033[?25h")
}

func ClearScreen() {
	MoveCursorTo(1, 1)
	ClearAfterCursor()
}

func ClearLine() {
	fmt.Print("\033[K")
}

func ClearAfterCursor() {
	fmt.Print("\033[J")
}

func MoveCursorTo(line, column int) {
	fmt.Printf("\033[%d;%df", line, column)
}

func MoveCursorUp(lineCount int) {
	fmt.Printf("\033[%dF", lineCount)
}

func EditLastLine() {
	MoveCursorUp(1)
}
