package img

import (
	"errors"
	"fmt"
	"io"
	"os/exec"
	"strings"
	"time"
)

func escapeQuotes(str string) string {
	return strings.ReplaceAll(str, `"`, `\"`)
}

var stdin io.WriteCloser
var ready = false

func StartDaemon() error {
	cmd := exec.Command("ueberzug", "layer", "--parser", "json")
	var err error
	stdin, err = cmd.StdinPipe()
	if err != nil {
		return err
	}
	ready = true
	return cmd.Run()
}

func SendCommand(command string) error {
	if stdin == nil {
		return errors.New("Stdin not opened")
	}
	_, err := io.WriteString(stdin, command+"\n")
	return err
}

func ShowImageWithSize(path, id string, x, y, w, h int) error {
	return SendCommand(
		fmt.Sprintf(
			`{"action": "add", "x": %d, "y": %d, "width": %d, "height": %d, "path": "%s", "identifier": "%s"}`,
			x, y, w, h,
			escapeQuotes(path),
			escapeQuotes(id),
		),
	)
}

func ShowImage(path, id string, x, y int) error {
	return SendCommand(
		fmt.Sprintf(
			`{"action": "add", "x": %d, "y": %d, "path": "%s", "identifier": "%s"}`,
			x, y,
			escapeQuotes(path),
			escapeQuotes(id),
		),
	)
}

func RemoveImage(id string) error {
	return SendCommand(fmt.Sprintf(`{"action": "remove", "identifier": "%s"}`, escapeQuotes(id)))
}

func WaitForDaemon() {
	for {
		if ready {
			break
		}
		time.Sleep(10 * time.Millisecond)
	}
}
