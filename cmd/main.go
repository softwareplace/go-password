package main

import (
	"errors"
	"github.com/softwareplace/go-password/pkg/str"
	"io"
	"log"
	"os/exec"
)

func copyToClipboard(text string) error {
	// Try xclip first, fall back to xsel
	var cmd *exec.Cmd
	if _, err := exec.LookPath("xclip"); err == nil {
		cmd = exec.Command("xclip", "-selection", "clipboard")
	} else if _, err := exec.LookPath("xsel"); err == nil {
		cmd = exec.Command("xsel", "--clipboard", "--input")
	} else {
		return errors.New("neither xclip nor xsel are installed")
	}

	stdin, err := cmd.StdinPipe()
	if err != nil {
		return err
	}

	go func() {
		defer func(stdin io.WriteCloser) {
			_ = stdin.Close()
		}(stdin)

		_, _ = io.WriteString(stdin, text)
	}()

	return cmd.Run()
}

func main() {
	generate := str.Default().
		Generate()
	err := copyToClipboard(generate)

	if err == nil {
		log.Printf("Generated password copied to clipboard\n\n")
	}

	println(generate)
}
