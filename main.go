package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"
	"os"
//	"strconv"

	"github.com/ecc1/cc2500"
)

func main() {
	r := cc2500.Open()
	hw := r.Hardware()
	log.Printf("connected to %s radio on %s", hw.Name(), hw.Device())
//	seconds := os.Args[1]
//	s := strconv.Itoa(420)
	hours := time.Tick(420 * time.Second)
	readings := r.ReceiveReadings()
	numReadings := 0
	for {
		if r.Error() != nil {
			log.Fatal(r.Error())
		}
		select {
		case <-hours:
			fmt.Printf("%d readings in previous hour\n", numReadings)
			numReadings = 0
			os.Exit(0)
		case reading := <-readings:
			if reading != nil {
				print(reading)
				numReadings++
				os.Exit(0)
			}
		}
	}
}

func print(r *cc2500.Packet) {
//	b, err := json.MarshalIndent(*r, "", "  ")
	b, err := json.Marshal(*r)
	if err != nil {
		fmt.Println(err)
		fmt.Printf("%+v", *r)
	} else {
		fmt.Println(string(b))
	}
}
