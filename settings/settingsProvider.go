package settings

import (
	"encoding/json"
	"os"
)

type Settings struct {
	TemperatureFileName string   `json:"TemperatureFileName"`
	FtpHost             string   `json:"FtpHost"`
	FtpUser             string   `json:"FtpUser"`
	FtpPassword         string   `json:"FtpPassword"`
	FtpTargetPath       string   `json:"FtpTargetPath"`
	PictureName         string   `json:"PictureName"`
	IntervalSec         int      `json:"IntervalSec"`
	Command             string   `json:"Command"`
	CommandArguments    []string `json:"CommandArguments"`
}

var Config Settings

func LoadConfig() {
	f, err := os.Open("Settings.json")
	if err != nil {
		panic(err)
	}
	defer func(f *os.File) {
		_ = f.Close()
	}(f)

	if err := json.NewDecoder(f).Decode(&Config); err != nil {
		panic(err)
	}
}
