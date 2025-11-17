package view

import (
	"RaspberryWeather/settings"
	"context"
	"fmt"
	"os"
	"os/exec"
	"time"
)

func TakePicture() bool {

	if _, err := exec.LookPath("fswebcam"); err != nil {
		fmt.Println("fswebcam not found. Call sudo apt install fswebcam")
		return false
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cmd := exec.CommandContext(ctx, settings.Config.Command, settings.Config.CommandArguments...)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		// timeout?
		if ctx.Err() == context.DeadlineExceeded {
			fmt.Println("fswebcam: Timeout")
		}
		fmt.Printf("fswebcam failed: %v\n", err)
		return false
	}

	return true
}
