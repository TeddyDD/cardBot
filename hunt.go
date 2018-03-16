package main

import (
	"math/rand"
)

var Animals []string = []string{
	"🐀",
	"🦀",
	"🦌",
	"🐇",
	"🐗",
	"🐣",
	"🐝",
	"🕷",
	"🐟",
	"🐻",
	"🐍",
	"🦆",
	"🦊",
	"Ryś",
	"Łoś",
	"Kuropatwa",
}

func GetAnimal() string {
	return Animals[rand.Intn(len(Animals))]
}
