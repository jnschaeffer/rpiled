package main

import (
	"log"
	"time"

	"github.com/kidoman/embd"

	_ "github.com/kidoman/embd/host/rpi"
)

func main() {
	errGPIO := embd.InitGPIO()
	if errGPIO != nil {
		log.Fatal(errGPIO)
	}

	defer embd.CloseGPIO()

	pin, errPin := embd.NewDigitalPin("P1_7")
	if errPin != nil {
		log.Fatal(errPin)
	}

	defer pin.Close()
	
	errDir := pin.SetDirection(embd.Out)
	if errDir != nil {
		log.Fatal(errDir)
	}

	for {
		errWrite := pin.Write(embd.Low)
		if errWrite != nil {
			log.Fatal(errWrite)
		}

		time.Sleep(1 * time.Second)

		errWrite = pin.Write(embd.High)
		if errWrite != nil {
			log.Fatal(errWrite)
		}

		time.Sleep(1 * time.Second)
	}
}
