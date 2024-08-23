package main

import (
	"flag"
	"fmt"
	"strings"

	"github.com/own-neufeldm/hr4go/lib"
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
	for _, name := range []string{"l", "length"} {
		flag.IntVar(&length, name, 50, "Minimum character length.")
	}

	var border string
	for _, name := range []string{"b", "border"} {
		flag.StringVar(&border, name, "//", "Character(s) to use for outer borders.")
	}

	var filler string
	for _, name := range []string{"f", "filler"} {
		flag.StringVar(&filler, name, "-", "Character(s) to use for outer borders.")
	}

	var asParagraph bool
	for _, name := range []string{"p", "paragraph"} {
		flag.BoolVar(&asParagraph, name, false, "Prepend 'BEGIN' and 'END' before title.")
	}

	var upper bool
	for _, name := range []string{"u", "upper"} {
		flag.BoolVar(&upper, name, false, "Convert title to uppercase.")
	}

	var noNewline bool
	for _, name := range []string{"n", "no-newline"} {
		flag.BoolVar(&noNewline, name, false, "Do not print a new-line character at the end.")
	}

	var showVersion bool
	for _, name := range []string{"v", "version"} {
		flag.BoolVar(&showVersion, name, false, "Show version and exit.")
	}

	flag.Usage = usage
	flag.Parse()

	if showVersion {
		fmt.Println("v1.0.0")
		return
	}

	var title = flag.Arg(0)
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
		horizontalRule := lib.GetHorizontalRule(title, length, border, string(filler[0]))
		if noNewline {
			fmt.Print(horizontalRule)
		} else {
			fmt.Println(horizontalRule)
		}
	}
}
