package formatting

import "text/template"

func FuncMap() template.FuncMap {
	return template.FuncMap{
		"include":   Include,
		"dump":      DumpArgs,
		"blue":      Blue,
		"cyan":      Cyan,
		"green":     Green,
		"magenta":   Magenta,
		"red":       Red,
		"yellow":    Yellow,
		"white":     White,
		"hiblue":    HighBlue,
		"hicyan":    HighCyan,
		"higreen":   HighGreen,
		"himagenta": HighMagenta,
		"hired":     HighRed,
		"hiyellow":  HighYellow,
		"hiwhite":   HighWhite,
	}
}
