package diff

import (
	"strings"
)

type DiffType int

const (
	DiffAdded DiffType = iota
	DiffRemoved
	DiffContext
)

type Line struct {
	Type    DiffType
	Content string
	OldNum  int
	NewNum  int
}

type Hunk struct {
	OldStart int
	OldCount int
	NewStart int
	NewCount int
	Lines    []Line
}

type FileDiff struct {
	OldName string
	NewName string
	Hunks   []Hunk
}

func ParseDiff(content string, filename string) *FileDiff {
	return &FileDiff{OldName: filename, NewName: filename}
}

func (d *FileDiff) Stats() (added, removed int) {
	for _, h := range d.Hunks {
		for _, l := range h.Lines {
			switch l.Type {
			case DiffAdded:
				added++
			case DiffRemoved:
				removed++
			}
		}
	}
	return
}

func JoinPath(base, file string) string {
	if strings.HasPrefix(file, "/") {
		return file
	}
	return base + "/" + file
}
