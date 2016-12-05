package util

import (
	"os"
	"fmt"
	"os/exec"
)
func Execute(cmd string) (string, error) {
	makeTemp(cmd)
	output, err := runCmd()
	cleanup()

	return output, err
}

func cleanup() {
	os.Remove("temp.sh")
}

func runCmd() (string, error) {
	out, err := exec.Command("sh", "temp.sh").Output()
	if err != nil {
		fmt.Println(err)
	}
	return string(out), err
}

func makeTemp(cmd string) {
	f, err := os.Create("temp.sh")
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()

	_, err = f.WriteString(cmd)
	if err != nil {
		fmt.Println(err)
	}
	f.Sync()
}