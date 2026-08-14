package main

import (
	"fmt"

	mgo "github.com/mowshon/moviego/v2"
)

func main() {
	info, err := mgo.Probe("video.mp4")
	if err != nil {
		panic(err)
	}

	fmt.Printf("VIDEO STAT:\n")
	fmt.Printf("Format: %s\n", info.Format)
	fmt.Printf("Duration: %v\n", info.Duration) // Returns time.Duration, not float64!
	fmt.Printf("File size: %d bytes\n", info.Size)
	fmt.Printf("AUDIO STAT:\n")
	fmt.Printf("Audio codec: %s\n", info.Audio.Codec)
	fmt.Printf("Audio channels: %d\n", info.Audio.Channels)
	fmt.Printf("Sample rate: %d Hz\n", info.Audio.SampleRate)
}
