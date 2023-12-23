package main

import (
	"github.com/kzumi/timer/timer"
)

func main() {
	timetoint, result := timer.GetInfo()
	if result == "start" {
		timer.Run(timetoint)
	}
}
