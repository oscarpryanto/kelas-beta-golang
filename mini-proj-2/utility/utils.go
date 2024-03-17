package utility

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

func runCmd(name string, arg ...string) {
	cmd := exec.Command(name, arg...)
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func ClearTerminal() {
	switch runtime.GOOS {
	case "darwin":
		runCmd("clear")
	case "linux":
		runCmd("clear")
	case "windows":
		runCmd("cmd", "/c", "cls")
	default:
		runCmd("clear")
	}
}

func InputMultiKata() string {
	var inputanMultiKata = bufio.NewReader(os.Stdin)
	input, error := inputanMultiKata.ReadString('\n')
	if error != nil {
		fmt.Print("Terjadi kesalahan pada saat input data buku : ", error)
		os.Exit(0)
	}
	input = strings.TrimSpace(input)
	return input
}
