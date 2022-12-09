package year2022

import (
	"bufio"
	"fmt"
	"path/filepath"
	"sort"
	"strconv"
	"strings"

	"tailscale.com/util/must"
)

type WalkFunc func(*File) error

type File struct {
	name     string
	size     int
	children []*File
	parent   *File
}

func (f *File) AddFile(name string, size int) {
	f.children = append(f.children, &File{name: name, size: size, parent: f})
}

func (f *File) AddDir(name string) {
	f.children = append(f.children, &File{name: name, parent: f, children: []*File{}})
}

func (f *File) Chdir(name string) *File {
	switch name {
	case "/":
		return f.Root()
	case "..":
		return f.Parent()
	default:
		for _, c := range f.children {
			if c.name == name {
				return c
			}
		}

		return nil
	}
}

func (f *File) IsRoot() bool {
	return f.name == "/"
}

func (f *File) IsDir() bool {
	return f.size == 0
}

func (f *File) Parent() *File {
	if f.IsRoot() {
		return f
	} else {
		return f.parent
	}
}

func (f *File) Root() *File {
	if f.Parent() == f {
		return f
	} else {
		return f.Parent().Root()
	}
}

func (f *File) Path() string {
	if f.IsRoot() {
		return "/"
	} else {
		return filepath.Join(f.Parent().Path(), f.name)
	}
}

func (f *File) TotalSize() int {
	if f.IsDir() {
		sum := 0
		for _, c := range f.children {
			sum += c.TotalSize()
		}
		return sum
	} else {
		return f.size
	}
}

func (f *File) Walk(fn WalkFunc) error {
	if !f.IsDir() {
		return fn(f)
	}
	if err := fn(f); err != nil {
		return err
	}
	for _, c := range f.children {
		if err := c.Walk(fn); err != nil {
			return err
		}
	}
	return nil
}

type DaySevenInput struct {
	Root *File
}
type DaySevenOutput struct {
	SmallDirsTotal int
	DirToDelete    int
}

func parseDaySeven(rawInput string) (*DaySevenInput, error) {
	scanner := bufio.NewScanner(strings.NewReader(rawInput))
	in := &DaySevenInput{
		Root: &File{name: "/", children: []*File{}},
	}

	cwdStr := ""
	cwd := in.Root

	for scanner.Scan() {
		line := scanner.Text()
		tokens := strings.Fields(line)
		if tokens[0] == "$" {
			// command
			switch tokens[1] {
			case "cd":
				cwdStr = filepath.Join(cwdStr, tokens[2])
				cwd = cwd.Chdir(tokens[2])
			case "ls":
				// don't need to do anything, file listing follows
			default:
				return nil, fmt.Errorf("unknown command: %s", tokens[1])
			}
		} else {
			// output
			if tokens[0] == "dir" {
				cwd.AddDir(tokens[1])
			} else {
				size := must.Get(strconv.Atoi(tokens[0]))
				cwd.AddFile(tokens[1], size)
			}
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return in, nil
}

func solveDaySeven(in *DaySevenInput) (*DaySevenOutput, error) {
	out := &DaySevenOutput{}

	// part one
	smallDirSize := 100000
	in.Root.Walk(func(f *File) error {
		if f.IsDir() && !f.IsRoot() {
			ts := f.TotalSize()
			if ts <= smallDirSize {
				out.SmallDirsTotal += ts
			}
		}
		return nil
	})

	// part two
	diskSize := 70000000
	freeNeeded := 30000000
	free := diskSize - in.Root.TotalSize()

	candidates := []*File{}
	in.Root.Walk(func(f *File) error {
		if f.IsDir() {
			freeIfDeleted := free + f.TotalSize()
			if freeIfDeleted >= freeNeeded {
				candidates = append(candidates, f)
			}
		}
		return nil
	})

	sort.Slice(candidates, func(i, j int) bool {
		return candidates[i].TotalSize() < candidates[j].TotalSize()
	})
	out.DirToDelete = candidates[0].TotalSize()

	return out, nil
}

func DaySeven(rawInput string) (*DaySevenOutput, error) {
	in, err := parseDaySeven(rawInput)
	if err != nil {
		return nil, err
	}

	out, err := solveDaySeven(in)
	if err != nil {
		return nil, err
	}

	return out, nil
}
