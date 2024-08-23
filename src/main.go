package main

import (
	"flag"
	"fmt"
	"strings"
)

func usage() {
	fmt.Println(`Usage: hr4go [OPTIONS] [TITLE] COMMAND [ARGS]...

Print horizontal rules.

Arguments:
  [TITLE]  An optional title.

Options:
  -l, --length INTEGER  Minimum character length.  [default: 50]
  -b, --border TEXT     Character(s) to use for outer borders.  [default: //]
  -f, --filler TEXT     Character to use for inner fillers.  [default: -]
  -p, --paragraph       Prepend 'BEGIN' and 'END' before title.
  -u, --upper           Convert title to uppercase.
  -n, --no-newline      Do not print a new-line character at the end.
  -v, --version         Show version and exit.
  -h, --help            Show this message and exit.`)
}

func main() {
	var length int
	flag.IntVar(&length, "l", 50, "Minimum character length.")
	flag.IntVar(&length, "length", 50, "Minimum character length.")

	var border string
	flag.StringVar(&border, "b", "//", "Character(s) to use for outer borders.")
	flag.StringVar(&border, "border", "//", "Character to use for inner fillers.")

	var filler string
	flag.StringVar(&filler, "f", "-", "Character(s) to use for outer borders.")
	flag.StringVar(&filler, "filler", "-", "Character to use for inner fillers.")

	var asParagraph bool
	flag.BoolVar(&asParagraph, "p", false, "Prepend 'BEGIN' and 'END' before title.")
	flag.BoolVar(&asParagraph, "paragraph", false, "Prepend 'BEGIN' and 'END' before title.")

	var upper bool
	flag.BoolVar(&upper, "u", false, "Convert title to uppercase.")
	flag.BoolVar(&upper, "upper", false, "Convert title to uppercase.")

	var noNewline bool
	flag.BoolVar(&noNewline, "n", false, "Do not print a new-line character at the end.")
	flag.BoolVar(&noNewline, "no-newline", false, "Do not print a new-line character at the end.")

	var showVersion bool
	flag.BoolVar(&showVersion, "v", false, "Show version and exit.")
	flag.BoolVar(&showVersion, "version", false, "Show version and exit.")

	flag.Usage = usage
	flag.Parse()

	var title = flag.Arg(0)

	if showVersion {
		fmt.Println("Not defined.")
		return
	}

	if upper {
		title = strings.ToUpper(title)
	}

	var titles []string
	if asParagraph {
		if title == "" {
			titles = append(titles, "BEGIN")
			titles = append(titles, "END")
		} else {
			titles = append(titles, fmt.Sprintf("BEGIN %s", title))
			titles = append(titles, fmt.Sprintf("END %s", title))
		}
	} else {
		titles = append(titles, title)
	}

	for _, title := range titles {
		horizontalRule := GetHorizontalRule(title, length, border, string(filler[0]))
		if noNewline {
			fmt.Print(horizontalRule)
		} else {
			fmt.Println(horizontalRule)
		}
	}
}
