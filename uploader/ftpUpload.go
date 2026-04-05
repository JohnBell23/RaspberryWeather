package uploader

import (
	"RaspberryWeather/settings"
	"fmt"
	"os"
	"path"
	"time"

	"github.com/jlaffaye/ftp"
)

func UploadFiles() {
	fmt.Println("upload files start")

	files := []string{
		settings.Config.TemperatureFileName,
		settings.Config.PictureName,
	}

	for _, file := range files {
		if err := upload(file); err != nil {
			fmt.Println(err)
			continue
		}

		fmt.Printf("upload %s done\n", file)
	}

	fmt.Println("upload files done")
}

func upload(fileName string) error {
	c, err := ftp.Dial(settings.Config.FtpHost, ftp.DialWithTimeout(30*time.Second))
	if err != nil {
		return err
	}
	defer c.Quit()

	if err := c.Login(settings.Config.FtpUser, settings.Config.FtpPassword); err != nil {
		return err
	}

	f, err := os.Open(fileName)
	if err != nil {
		return err
	}
	defer f.Close()

	return c.Stor(path.Join(settings.Config.FtpTargetPath, path.Base(fileName)), f)
}
