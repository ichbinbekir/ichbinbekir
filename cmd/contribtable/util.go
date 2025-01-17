package main

import (
	"fmt"
	"image"
	"image/png"
	"math/rand"
	"os"
	"os/exec"
)

func getArt(name string) (image.Image, error) {
	file, err := os.Open(name)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	return png.Decode(file)
}

func commandExec(name string, args ...string) error {
	cmd := exec.Command(name, args...)
	return cmd.Run()
}

func randClock() string {
	hour := rand.Intn(24)
	minute := rand.Intn(60)
	second := rand.Intn(60)

	return fmt.Sprintf("%02d:%02d:%02d", hour, minute, second)
}
