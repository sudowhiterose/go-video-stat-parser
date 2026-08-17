package main

import (
	"fmt"
	"path/filepath"

	mgo "github.com/mowshon/moviego/v2"
)

func main() {
	//init video
	files, err := filepath.Glob("storage/*")
	//checking on error
	if err != nil {
		fmt.Println("File search pattern error:", err)
	}

	if len(files) == 0 {
		fmt.Println("Files not found. Please, create a 'storage' folder and upload videos there")
		return
	}
	//checking all video in folder
	for i, file := range files {
		info, err := mgo.Probe(file)
		if err != nil {
			fmt.Printf("File parsing error %s: %v\n", file, err)
			continue
		}
		//parsing result
		fmt.Printf("\n--- VIDEO #%d ---\n", i+1)
		fmt.Printf("Format: %s\n", info.Format)
		fmt.Printf("Duration: %v\n", info.Duration)
		fmt.Printf("File size: %d\n", info.Size)
		fmt.Printf("Rate: %d\n", info.Rate)
		fmt.Printf("Variable fps: %t\n", info.VariableFPS)
		fmt.Printf("codec: %v\n", info.Codec)
		//checking alvailable audio in video
		if info.Audio != nil {
			fmt.Printf("AUDIO STAT:\n")
			fmt.Printf("Audio codec: %s\n", info.Audio.Codec)
			fmt.Printf("Audio channels: %d\n", info.Audio.Channels)
			fmt.Printf("Sample rate: %d Hz\n", info.Audio.SampleRate)
		}
	}
}
