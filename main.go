package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"sync"
)

func recordScreen(outputFile string) {
	// Define screen dimensions and frame rate
	screenWidth := 1920
	screenHeight := 1080
	framerate := 30

	// Construct the FFmpeg command to capture the screen
	cmdArgs := []string{
		"-f", "avfoundation",
		"-framerate", fmt.Sprintf("%d", framerate),
		"-video_size", fmt.Sprintf("%dx%d", screenWidth, screenHeight),
		"-i", "1",
		"-c:v", "libx264",
		"-preset", "ultrafast",
		outputFile,
	}

	// Create the command
	cmd := exec.Command("ffmpeg", cmdArgs...)

	// Redirect FFmpeg output to stdout and stderr
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	// Run the FFmpeg command to record the screen
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	// Number of screen recording copies
	numCopies := 3

	// Create a WaitGroup for goroutine synchronization
	var wg sync.WaitGroup

	// Start screen recording and copies in separate goroutines
	for i := 0; i <= numCopies; i++ {
		outputFile := fmt.Sprintf("output%d.mp4", i)
		wg.Add(1)
		go func(outputFile string) {
			defer wg.Done()
			recordScreen(outputFile)
		}(outputFile)
	}

	// Wait for all goroutines to finish
	wg.Wait()
}
