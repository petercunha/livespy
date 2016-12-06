package util

import (
	"fmt"
)

func CaptureScreen() error {
	// Currently only works on MacOS and OSX
	_, err := Execute("screencapture -x site/latest.jpg")
	if err != nil {
		fmt.Println(err)
	}
	return err
}