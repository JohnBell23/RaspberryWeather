package main

import (
	"RaspberryWeather/temperature"
	"RaspberryWeather/view"
	"fmt"
	"time"

	"github.com/eiannone/keyboard"
)

var done bool

func main() {

	fmt.Printf("Start Weather App\n")
	done = false

	go getKey()

	counter := 0
	interval := 20

	var sensor temperature.Bmp280
	sensor.Initialize(temperature.DefaultIc2TempSensorAddr)
	for !done {
		if counter%interval == 0 {
			counter = 0
			fmt.Println("----------------------")
			fmt.Println(time.Now().Format(time.RFC3339))
			sensor.UpdateSensorData()
			fmt.Printf("Temperatur: %.2f °C\n", sensor.Temperature)
			fmt.Printf("Druck: %.2f hPa\n", sensor.Pressure)
			fmt.Println("")

			view.TakePicture()
		}
		time.Sleep(1 * time.Second)
		counter++
	}

	sensor.Uninitialize()
	fmt.Printf("goodbye\n")
}

func getKey() {
	if err := keyboard.Open(); err != nil {
		panic(err)
	}
	defer keyboard.Close()

	_, _, err := keyboard.GetKey()
	if err != nil {
	}
	fmt.Println("got key")
	done = true
}
