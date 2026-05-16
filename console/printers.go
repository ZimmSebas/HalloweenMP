package console

import (
	"fmt"
	"math/rand"
)

func glitchLine(s string, intensity float64) string {
	runes := []rune(s)
	for i := range runes {
		if rand.Float64() < intensity {
			runes[i] = rune(33 + rand.Intn(94)) // random printable char
		}
	}
	return string(runes)
}

func GlitchPrintf(intensity float64, format string, a ...interface{}) {
	text := fmt.Sprintf(format, a...)
	glitched := glitchLine(text, intensity)
	fmt.Print(glitched)
}

func ErrorPrintf(text string) {

}

// func (c *Console) PrintColor (ctx *app.Context, text string, color string) {
// 	switch color {
// 		case "red":
// 	}
//
// }
