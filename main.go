package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
	"unicode/utf8"
)

func tabsToSpaces(lines []string) []string {
	var ret []string
	for _, l := range lines {
		l = strings.Replace(l, "\t", " ", -1) // source: https://pkg.go.dev/strings#Replace
		ret = append(ret, l)
	}
	return ret
}

func calculateMaxWidth(lines []string) int {
	w := 0
	for _, l := range lines {
		len := utf8.RuneCountInString(l)
		if len > w {
			w = len
		}
	}
	return w
}

func normalizeStringsLength(lines []string, maxwidth int) []string {
	var ret []string
	for _, l := range lines {
		s := l + strings.Repeat(" ", maxwidth-utf8.RuneCountInString(l))
		ret = append(ret, s)
	}
	return ret
}

func buildBalloon(lines []string, maxwidth int) string {
	var borders []string
	count := len(lines)
	var ret []string

	borders = []string{"/", "\\", "\\", "/", "|", "<", ">"}

	top := " " + strings.Repeat("_", maxwidth+2)
	bottom := " " + strings.Repeat("-", maxwidth+2)

	ret = append(ret, top)
	if count == 1 {
		s := fmt.Sprintf("%s %s %s", borders[5], lines[0], borders[6])
		ret = append(ret, s)
	} else {
		s := fmt.Sprintf("%s %s %s", borders[0], lines[0], borders[1])
		ret = append(ret, s)
		i := 1
		for ; i < count-1; i++ {
			s = fmt.Sprintf("%s %s %s", borders[4], lines[i], borders[4])
			ret = append(ret, s)
		}
		s = fmt.Sprintf("%s %s %s", borders[2], lines[i], borders[3])
		ret = append(ret, s)
	}

	ret = append(ret, bottom)
	return strings.Join(ret, "\n")
}

func printFigure(name string) {
	var cow = `         \  ^__^
          \ (oo)\_______
	    (__)\       )\/\
	        ||----w |
	        ||     ||
		`

	var stegosaurus = `         \                      .       .
		\                    / ` + "`" + `.   .' "
		 \           .---.  <    > <    >  .---.
		  \          |    \  \ - ~ ~ - /  /    |
		_____           ..-~             ~-..-~
	   |     |   \~~~\\.'                    ` + "`" + `./~~~/
	  ---------   \__/                         \__/
	 .'  O    \     /               /       \  "
	(_____,    ` + "`" + `._.'               |         }  \/~~~/
	 ` + "`" + `----.          /       }     |        /    \__/
		   ` + "`" + `-.      |       /      |       /      ` + "`" + `. ,~~|
			   ~-.__|      /_ - ~ ^|      /- _      ` + "`" + `..-'
					|     /        |     /     ~-.     ` + "`" + `-. _  _  _
					|_____|        |_____|         ~ - . _ _ _ _ _>

  `
	switch name {
	case "cow":
		fmt.Println(cow)
	case "stegosaurus":
		fmt.Println(stegosaurus)
	default:
		fmt.Println(cow)
	}
}

func main() {
	// Get file information about standard input
	info, _ := os.Stdin.Stat()

	// Check if input is from a terminal (character device) using a bitwise operation
	if info.Mode()&os.ModeCharDevice != 0 {
		fmt.Println("The command  is intended to work with pipes.")
		fmt.Println("Usage: fortune | gocowsay")
		return
	}

	// Create new buffered reader
	reader := bufio.NewReader(os.Stdin)

	// Select figure
	var figure string
	flag.StringVar(&figure, "f", "cow", "the figure name. Valid values are `cow` and `stegosaurus`")
	flag.Parse()

	// Initialise empty string slice
	var lines []string

	// Infinite loop
	for {
		input, _, err := reader.ReadLine()
		if err != nil && err == io.EOF {
			break
		}
		lines = append(lines, string(input))
	}

	lines = tabsToSpaces(lines)
	maxwidth := calculateMaxWidth(lines)
	messages := normalizeStringsLength(lines, maxwidth)
	balloon := buildBalloon(messages, maxwidth)
	fmt.Println(balloon)
	printFigure(figure)
	fmt.Println()
}
