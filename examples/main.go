package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	mezonsdk "github.com/nccasia/mezon-go-sdk"
)

func main() {
	player, err := mezonsdk.NewAudioPlayer("123456", "123456", "123456", "komu", "token")
	if err != nil {
		fmt.Println("error", err)
		return
	}

	// Cleanup function (called when program stops)
	cleanup := func() {
		fmt.Println("Cleaning up...")
		player.Close("123456")
	}

	// Catch the program stop signal
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	go func() {
		if err := player.Play("https://raw.githubusercontent.com/mezonai/mezon-go-bot/refs/heads/main/audio/ncc8.ogg"); err != nil {
			fmt.Println("Error playing audio:", err)
			// End if music playback fails
			stop <- syscall.SIGTERM
		}
	}()

	// Wait for stop signal
	<-stop
	fmt.Println("Stopping program...")

	// Call cleanup function before exiting
	cleanup()

	return
}
