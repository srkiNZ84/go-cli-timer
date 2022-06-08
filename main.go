package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/faiface/beep"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
)

func main() {
	// Below code is derived from: https://github.com/faiface/beep/wiki/Hello,-Beep!
	fmt.Println("CLI Timer")

	// Track is from https://freemusicarchive.org/home
	f, err := os.Open("Scott_Holmes_Music_-_Strategy.mp3")
	if err != nil {
		log.Fatalf("Error opening MP3: %v", err)
	}

	streamer, format, err := mp3.Decode(f)
	if err != nil {
		log.Fatalf("Error decoding MP3: %v", err)
	}
	defer streamer.Close()

	speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))

	done := make(chan bool)

	fmt.Println("Timer set and started for 25mins!")
	time.Sleep(25 * time.Minute)

	fmt.Println("Time for a break!")
	speaker.Play(beep.Seq(streamer, beep.Callback(func() {
		done <- true
	})))

	<-done
}
