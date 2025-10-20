package view

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"time"
)

const PictureName = "shot.jpg"

const Command = "fswebcam"

var CommandArgs = []string{"-d", "/dev/video0", "-r", "1280x720", "--set", "white_balance_automatic=0", "--set", "white_balance_temperature=6200", "--set", "gain=0", "--set", "backlight_compensation=0", "--skip", "5", "--jpeg", "95", "--no-banner", PictureName}

func TakePicture() bool {
	if _, err := exec.LookPath("fswebcam"); err != nil {
		fmt.Println("fswebcam nicht gefunden. Installiere es z.B. mit: sudo apt install fswebcam")
		return false
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cmd := exec.CommandContext(ctx, Command, CommandArgs...)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		// War es ein Timeout?
		if ctx.Err() == context.DeadlineExceeded {
			fmt.Println("fswebcam: Timeout (Kamera busy oder kein Signal?)")
		}
		fmt.Printf("fswebcam fehlgeschlagen: %v\n", err)
		return false
	}

	fmt.Println("Bild gespeichert: shot.jpg")
	return true
}
