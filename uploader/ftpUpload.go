package uploader

import (
	"RaspberryWeather/settings"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/jlaffaye/ftp"
)

func UploadFiles() {
	fmt.Println("upload files")

	upload(settings.Config.TemperatureFileName)
	upload(settings.Config.PictureName)
}

func upload(fileName string) {
	// connect
	c, err := ftp.Dial(settings.Config.FtpHost, ftp.DialWithTimeout(5*time.Second))
	if err != nil {
		log.Fatal(err)
	}
	defer func(c *ftp.ServerConn) {
		err := c.Quit()
		if err != nil {
			// eat it
		}
	}(c)

	// login
	if err := c.Login(settings.Config.FtpUser, settings.Config.FtpPassword); err != nil {
		fmt.Println("login error, ", err)
	}

	// local file
	f, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// upload
	if err := c.Stor(
		fmt.Sprintf("%s/%s", settings.Config.FtpTargetPath, fileName), f); err != nil {
		fmt.Println("upload error, ", err)
	}

	fmt.Println("upload done")
}
