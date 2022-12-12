package main

import (
	"errors"
	"fmt"
	"io/fs"
	"log"
	"os"
	"strconv"
	"strings"
	"text/template"

	"github.com/gammban/numtow"
	"github.com/gammban/numtow/lang"
	"tailscale.com/util/must"
)

type AdventDay struct {
	Year    int
	Day     int
	DayWord string
}

func ToWord(i int) string {
	return strings.Title(numtow.MustInt64(int64(i), lang.EN))
}

// Generates
func main() {
	if len(os.Args) != 3 {
		log.Fatalf("expected 2 args got %d", len(os.Args)-1)
	}
	year := must.Get(strconv.Atoi(os.Args[1]))
	day := must.Get(strconv.Atoi(os.Args[2]))

	basePath := fmt.Sprintf("./year%d/day%02d", year, day)
	goPath := fmt.Sprintf("%s.go", basePath)
	testPath := fmt.Sprintf("%s_test.go", basePath)

	ctx := &AdventDay{
		Year:    year,
		Day:     day,
		DayWord: ToWord(day),
	}

	err := writeTemplate(goPath, goTemplate, ctx)
	if err != nil {
		log.Fatalf("error writing go template: %s", err)
	}
	log.Printf("write go file to %s", goPath)

	err = writeTemplate(testPath, testTemplate, ctx)
	if err != nil {
		log.Fatalf("error writing test template: %s", err)
	}
	log.Printf("write test file to %s", testPath)
}

func writeTemplate(path string, tmpl string, ctx *AdventDay) error {
	_, err := os.Stat(path)
	if !errors.Is(err, fs.ErrNotExist) {
		return fmt.Errorf("file already exists: %s", path)
	}

	t := template.Must(template.New("file").Parse(tmpl))
	f, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("file create err: %s", err)
	}
	defer f.Close()
	return t.Execute(f, ctx)
}

const goTemplate = `
package year{{ .Year }}

import (
	"bufio"
	"log"
	"strings"
)

type Day{{ .DayWord }}Input struct {
}
type Day{{ .DayWord }}Output struct {
	PartOneAnswer int
}

func parseDay{{ .DayWord }}(rawInput string) (*Day{{ .DayWord }}Input, error) {
	scanner := bufio.NewScanner(strings.NewReader(rawInput))
	in := &Day{{ .DayWord }}Input{}

	for scanner.Scan() {
		line := scanner.Text()
		log.Printf("line: %s", line)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return in, nil
}

func solveDay{{ .DayWord }}(in *Day{{ .DayWord }}Input) (*Day{{ .DayWord }}Output, error) {
	out := &Day{{ .DayWord }}Output{}

	// part one

	// part two

	return out, nil
}

func Day{{ .DayWord }}(rawInput string) (*Day{{ .DayWord }}Output, error) {
	in, err := parseDay{{ .DayWord }}(rawInput)
	if err != nil {
		return nil, err
	}

	out, err := solveDay{{ .DayWord }}(in)
	if err != nil {
		return nil, err
	}

	return out, nil
}
`

const testTemplate = `
package year{{ .Year }}

import (
	"testing"

	"github.com/davecgh/go-spew/spew"
	"github.com/phinze/adventofcode/aoc"
	"github.com/stretchr/testify/require"
)

func TestDay{{ .DayWord }}_simple(t *testing.T) {
	in := ` + "`" + `
` + "`" + `
	out, err := Day{{ .DayWord }}(in)
	if err != nil {
		t.Fatalf("err: %s", err)
	}
	expected := 0
	require.Equal(t, expected, out.PartOneAnswer)
}

func TestDay{{ .DayWord }}(t *testing.T) {
	in, err := aoc.FetchInput({{ .Year }}, {{ .Day }})
	if err != nil {
		t.Fatalf("err: %s", err)
	}
	out, err := Day{{ .DayWord }}(in)
	if err != nil {
		t.Fatalf("err: %s", err)
	}
	t.Logf("out: %s", spew.Sdump(out))
}
`
