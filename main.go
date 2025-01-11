package main

import (
	"image/png"
	"log"
	"os"
	"time"
)

func run() error {
	f, err := os.Open("4x7-ichbinbekir.png")
	if err != nil {
		return err
	}
	defer f.Close()

	img, err := png.Decode(f)
	if err != nil {
		return err
	}

	startDay := time.Date(2025, time.January, 5, 0, 0, 0, 0, time.UTC)

	day := int(time.Now().UTC().Sub(startDay).Hours() / 24)
	week := day / 7

	log.Println("day:", day)
	log.Println("week:", week)
	log.Println("day mod with seven:", day%7)
	log.Println("pixel:", img.At(week, day%7))

	_, _, _, alpha := img.At(week, day%7).RGBA()
	if alpha != 0 {
		log.Println("that day today")
		return nil
	}

	log.Println("you are not in your day")
	return nil
}

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
	time.Sleep(time.Hour)
}
