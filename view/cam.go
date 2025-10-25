package view

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"time"
)

type CamConfig struct {
	Command          string   `json:"Command"`
	CommandArguments []string `json:"CommandArguments"`
}

func TakePicture() bool {
	var cfg = loadConfig()

	if _, err := exec.LookPath("fswebcam"); err != nil {
		fmt.Println("fswebcam not found. Call sudo apt install fswebcam")
		return false
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cmd := exec.CommandContext(ctx, cfg.Command, cfg.CommandArguments...)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		// War es ein Timeout?
		if ctx.Err() == context.DeadlineExceeded {
			fmt.Println("fswebcam: Timeout")
		}
		fmt.Printf("fswebcam failed: %v\n", err)
		return false
	}

	return true
}

func loadConfig() CamConfig {
	f, err := os.Open("CamConfig.json")
	if err != nil {
		panic(err)
	}
	defer func(f *os.File) {
		_ = f.Close()
	}(f)

	var cfg CamConfig
	if err := json.NewDecoder(f).Decode(&cfg); err != nil {
		panic(err)
	}

	return cfg
}
