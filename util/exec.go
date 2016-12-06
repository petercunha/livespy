package util

import (
	"os"
	"fmt"
	"os/exec"
	"time"
)

func Execute(cmd string) (string, error) {
	fileName := string(time.Now().Unix())

	makeTemp(cmd, fileName)
	output, err := runCmd(fileName)
	cleanup(fileName)

	return output, err
}

func cleanup(fileName string) {
	os.Remove(fileName)
}

func runCmd(fileName string) (string, error) {
	out, err := exec.Command("sh", fileName).Output()
	if err != nil {
		fmt.Println(err)
	}
	return string(out), err
}

func makeTemp(cmd string, fileName string) {
	f, err := os.Create(fileName)
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