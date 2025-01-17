package main

import (
	"flag"
	"log"
	"strconv"
	"time"
)

var art = flag.String("a", "assets/images/art.png", "Enter the path of the art file. (51x7)")
var date = flag.String("d", "2025-01-05", "You must enter the date of the Sunday of the first week other than the half week.")

func init() {
	flag.Parse()
}

func run() error {
	img, err := getArt(*art)
	if err != nil {
		return err
	}

	firstDate, err := time.Parse("2006-01-02", *date)
	if err != nil {
		return err
	}

	log.Println("starting")

	size := img.Bounds().Size()
	var currentDay time.Duration
	pixel := 1
	for x := 0; x < size.X; x++ {
		for y := 0; y < size.Y; y++ {
			if _, _, _, a := img.At(x, y).RGBA(); a != 0 {
				if err := commandExec("git", "commit", "--allow-empty", "--date", firstDate.Add(time.Hour*24*currentDay).Format("2006-01-02")+" "+randClock(), "-m", "pixel: "+strconv.Itoa(pixel)); err != nil {
					return err
				}
				pixel++
				log.Println("a commit")
			}
			currentDay++
		}
	}

	log.Println("push")
	return commandExec("git", "push")
}

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}
