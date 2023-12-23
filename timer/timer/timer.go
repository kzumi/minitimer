package timer

import (
	"fmt"
	"os"
	"time"

	"github.com/fatih/color"
)

var (
	Hour     int
	Minute   int
	Second   int
	ho       int
	min      int
	se       int
	Startmsg string
)

func Run(timing int) {
	LogTime(timing)
	SetTimer(timing)
}
func SetTimer(timing int) {
	ho = timing / 3600
	min = (timing % 3600) / 60
	se = ((timing % 3600) % 60)
	var duration int = timing

	for duration >= 0 {
		fmt.Printf("\r [ ** ] time to left :  %v : %v : %v", ho, min, se)
		se = se - 1
		time.Sleep(time.Second)
		duration--
		if se == 0 {
			if min > 0 && ho == 0 {
				swapse := 59
				se = swapse
				swapse--
				min = min - 1
			}
			if ho > 0 {
				swapmin := 59
				min = swapmin
				swapmin--
				ho--
				swapse := 60
				se = swapse
				swapse--
				min = min - 1

			}
		}

	}

	fmt.Println(color.BlueString(" [**] "), "timer is end ! ")
}

// func EndTimer(timing int) {
//todo
// }

func LogTime(timing int) {
	fmt.Println(color.GreenString("[start] timer is started !  "))
	fmt.Println(color.RedString("[end] timer end at :"))

}

func GetInfo() (int, string) {

	fmt.Print("[STAT] enter your time you want : [format(with space ) : hour minute second ]")
	fmt.Scanf("%d %d %d ", &Hour, &Minute, &Second)
	fmt.Print("type start to begin from now : ")
	fmt.Print("console > ")
	fmt.Scanf("%s", &Startmsg)
	var time_result int = (3600 * Hour) + (60 * Minute) + Second
	if Startmsg == "start" {
		return time_result, Startmsg

	} else {
		fmt.Println("DOES NOT MATCH , woould you try again : [y/n]")
		var yes string
		fmt.Scanf("%s", &yes)
		switch yes {
		case "y":
			GetInfo()
		case "n":
			os.Exit(0)
		}
	}
	return 0, ""
}
