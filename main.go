package main

import (
	"fmt"
	"os"
	"time"

	"github.com/IvanMenshykov/MoonPhase"
	"github.com/bglebrun/WeatherGo"
)

func main() {
	city, state := os.Args[1], os.Args[2]
	if city == "" || state == "" {
		fmt.Print("USAGE: weather [city] [state]")
		return
	}
	w := WeatherGo.MakeQuery(WeatherGo.BuildURL(WeatherGo.BuildLocation(city, state)))
	if w == nil {
		fmt.Printf("Program Error")
	} else {
		fmt.Printf("Temperature: %s %s, %s, Humidity: %s", w.Temp, w.Tp, w.Weth, w.Humidity)
	}

	phase := getPhaseString()

	fmt.Printf("\nMoon Phase: %s \n", phase)

}

func getPhaseString() string {
	m := MoonPhase.New(time.Now())
	return m.PhaseName()
}
