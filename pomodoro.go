package main

import (
	"github.com/faiface/beep"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
	"github.com/schollz/progressbar/v3"
	"log"
	"os"
	"strconv"
	"time"
)

func main() {
	var countTo int = 0

	if len(os.Args) > 1 {
		userInput, err := strconv.Atoi(os.Args[1])

		if err != nil {
			log.Fatal(err)
		}
		countTo = userInput
	} else {
		countTo = 600 // default 10 minutes value
	}

	soundFile, err := os.Open("bell.mp3")
	if err != nil {
		log.Fatal(err)
	}
	streamer, format, err := mp3.Decode(soundFile)
	if err != nil {
		log.Fatal(err)
	}
	defer streamer.Close()
	speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))

	bar := progressbar.Default(int64(countTo))

	for i := 0; i < countTo; i++ {
		bar.Add(1)
		time.Sleep(1 * time.Second)
	}

	done := make(chan bool)
	speaker.Play(beep.Seq(streamer, beep.Callback(func() {
		done <- true
	})))
	<-done
}
