package main

import (
	"RaspberryWeather/temperature"
	"fmt"
	"time"
)

func main() {

	var sensor temperature.Bmp280

	sensor.Initialize(temperature.DefaultIc2TempSensorAddr)

	for i := 0; i < 10; i++ {
		sensor.UpdateSensorData()

		fmt.Printf("Temperatur: %.2f °C\n", sensor.Temperature)
		fmt.Printf("Druck: %.2f hPa\n", sensor.Pressure)
		time.Sleep(2 * time.Second)
	}

	sensor.Uninitialize()
}
