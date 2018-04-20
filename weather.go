package main

import (
	"fmt"
	"math/rand"
)

var Pogoda map[string][]string

var Wind []string = []string{
	"Brak",
	"Leciutki 🌬",
	"Mocny  💨",
	"Wichura 🌪",
}

type TempInfo struct {
	min int
	max int
}

func (t TempInfo) Random() int {
	return rand.Intn(t.max-t.min) + t.min
}

var Temperature map[string]TempInfo = map[string]TempInfo{
	"lato":   TempInfo{10, 35},
	"wiosna": TempInfo{-5, 10},
	"jesień": TempInfo{0, 20},
	"zima":   TempInfo{-25, 0},
}

func GetTemperature(season string) string {
	return fmt.Sprintf("%d℃", Temperature[season].Random())
}

func CreateWeather() {
	Pogoda = map[string][]string{
		"zima": {
			"Słońce, bezchmurne niebo ☀️",
			"Słonecznie, małe zachmurzenie 🌤",
			"Słonecznie, umiarkowane zachmurzenie 🌥",
			"Duże zachmurzenie  ☁️",
			"Drobny deszcz  🌧",
			"Ulewa ⛈",
			"Drobny śnieg 🌨",
			"Śnieżyca ❄️",
		},
		"wiosna": {
			"Słońce, bezchmurne niebo ☀️",
			"Słonecznie, małe zachmurzenie 🌤",
			"Słonecznie, umiarkowane zachmurzenie 🌥",
			"Duże zachmurzenie  ☁️",
			"Drobny deszcz  🌧",
			"Ulewa  ⛈",
			"Łagodna burza 🌩",
			"Drobny śnieg 🌨",
		},
		"lato": {
			"Słońce, bezchmurne niebo ☀️",
			"Słonecznie, małe zachmurzenie 🌤",
			"Słonecznie, umiarkowane zachmurzenie 🌥",
			"Duże zachmurzenie  ☁️",
			"Drobny deszcz  🌧",
			"Ulewa  ⛈",
			"Łagodna burza 🌩",
			"Gwałtowna burza ⚡️",
		},
		"jesień": {
			"Słońce, bezchmurne niebo ☀️",
			"Słonecznie, małe zachmurzenie 🌤",
			"Słonecznie, umiarkowane zachmurzenie 🌥",
			"Duże zachmurzenie  ☁️",
			"Drobny deszcz  🌧",
			"Ulewa  ⛈",
			"Burza 🌩",
			"Drobny śnieg 🌨",
		},
	}

}

func GetWeatherReport(season string) string {
	return fmt.Sprintf("%s\nWiatr: %s\nTemperatura: %s",
		GetWeather(season),
		GetWind(),
		GetTemperature(season))
}

func GetWeather(season string) string {
	var s []string = Pogoda[season]
	l := len(s)
	return s[rand.Intn(l)]
}

func GetWind() string {
	return Wind[rand.Intn(len(Wind))]
}
