package temperature

import (
	"log"
	"time"

	"periph.io/x/conn/v3/i2c"
	"periph.io/x/conn/v3/i2c/i2creg"
	"periph.io/x/conn/v3/physic"
	"periph.io/x/devices/v3/bmxx80"
	"periph.io/x/host/v3"
)

type TemperatureSensor interface {
	Initialize(addr int)
	UpdateSensorData()
	Uninitialize()
}

const DefaultIc2TempSensorAddr = 0x76

type Bmp280 struct {
	bus         i2c.BusCloser
	addr        int
	dev         *bmxx80.Dev
	Temperature float64
	Pressure    float64
}

func (bmp *Bmp280) Uninitialize() {
	_ = bmp.bus.Close()
	_ = bmp.dev.Halt()
}

func (bmp *Bmp280) Initialize(addr int) {

	bmp.addr = addr

	if _, err := host.Init(); err != nil {
		log.Fatalf("host.Init: %v", err)
	}

	bus, err := i2creg.Open("")
	if err != nil {
		log.Fatalf("i2creg.Open: %v", err)
	}
	bmp.bus = bus

	dev, err2 := bmxx80.NewI2C(bus, uint16(bmp.addr), &bmxx80.DefaultOpts)
	if err2 != nil {
		log.Fatalf("bmxx80.NewI2C: %v", err)
	}
	bmp.dev = dev

	time.Sleep(300 * time.Millisecond)
}

func (bmp *Bmp280) UpdateSensorData() {
	var env physic.Env
	if err := bmp.dev.Sense(&env); err != nil {
		log.Fatalf("Sense: %v", err)
	}

	bmp.Temperature = env.Temperature.Celsius()
	bmp.Pressure = float64(env.Pressure) / float64(physic.Pascal) / 100.0
}
