package main

import (
	"RaspberryWeather/settings"
	"RaspberryWeather/temperature"
	"RaspberryWeather/uploader"
	"RaspberryWeather/view"

	"fmt"
	"os/exec"
	"time"
)

var done bool

func main() {

	fmt.Printf("Start Weather App\n")
	done = false

	counter := 0
	interval := 1

	var sensor temperature.Bmp280
	sensor.Initialize(temperature.DefaultIc2TempSensorAddr)

	for !done {

		settings.LoadConfig()

		if counter%interval == 0 {

			interval = settings.Config.IntervalSec

			counter = 0
			fmt.Println("")
			fmt.Println(time.Now().Format(time.RFC3339))
			sensor.UpdateSensorData()
			fmt.Printf("temperature: %.2f °C\n", sensor.Temperature)
			fmt.Printf("pressure: %.2f hPa\n", sensor.Pressure)
			fmt.Println("")

			view.TakePicture()

			fmt.Println("")

			saveWeatherData(sensor.Temperature, sensor.Pressure)

			fmt.Println("")

			uploader.UploadFiles()

			fmt.Println("")

			fmt.Printf("wait time %d\n", interval)

		}
		time.Sleep(1 * time.Second)
		counter++
	}

	sensor.Uninitialize()
	fmt.Printf("goodbye\n")
}

func saveWeatherData(temp float64, pressure float64) {
	fmt.Printf("write data to %s\n", settings.Config.TemperatureFileName)
	arg := fmt.Sprintf("echo %.1f°C %.2f > %s", temp, pressure, settings.Config.TemperatureFileName)
	cmd := exec.Command("bash", "-c", arg)
	err := cmd.Run()
	if err != nil {
		panic(err)
	}
}
