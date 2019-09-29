package main

import (
	"fmt"
	"time"
)

func main() {
	Now := time.Now()
	Year := 365
	YearDay := Now.YearDay()
	Percentage := (float64(YearDay) / float64(Year)) * 100
	echo(47, 30, fmt.Sprintf("\n %d年已经过去了 ", Now.Year()))
	echo(41, 37, fmt.Sprintf(" %.0f%% ", Percentage))
	echo(46, 37, fmt.Sprintf(" %d/%d \n", YearDay, Year))
	for i := 0; i < Year; i++ {
		if i != 0 && i%32 == 0 {
			fmt.Println()
		}
		if i < YearDay {
			echo(43, 30, "|")
		} else {
			echo(47, 30, "|")
		}
	}
	echo(41, 37, fmt.Sprintf("%d\n\n", Now.Year()+1))
}

func echo(b, t int, s string) {
	bg := b   // 背景色
	text := t // 前景色
	fmt.Printf("%c[%d;%d;%dm%s%c[0m", 0x1B, 0, bg, text, s, 0x1B)
}
