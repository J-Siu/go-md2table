package lib

import (
	"strings"

	"github.com/J-Siu/go-helper/v2/array"
	"github.com/J-Siu/go-md2table/global"
)

func ToTable(lines *[]string) *array.Array[string] {
	var (
		hasEntry       bool
		headerLevel    int
		headerLevelMax int
		table          array.Array[string]
	)
	// Get max level
	for _, line := range *lines {
		level := getHeaderLevel(&line)
		if level > headerLevelMax {
			headerLevelMax = level
		}
	}
	// header row
	if headerLevelMax > 0 {
		table.Add(strings.Repeat("|", headerLevelMax) + "||" + global.LineBreak)
		table.Add(strings.Repeat("|--", headerLevelMax) + "|--|")
	}
	// rows
	for n, l := range *lines {
		if len(l) == 0 {
			continue
		}
		line := strings.TrimSpace(l)
		headerLevel = getHeaderLevel(&line)
		if headerLevel > 0 {
			if n > 0 { // not first line
				if hasEntry {
					table.Add("|")
				} else {
					table.Add(strings.Repeat("|", headerLevelMax-headerLevel+1))
				}
			}
			// line start
			table.Add(global.LineBreak + strings.Repeat("|", headerLevel))
			// header to column
			table.Add(strings.TrimSpace(strings.TrimLeft(line, "#")) + strings.Repeat("|", headerLevelMax-headerLevel+1))
			hasEntry = false
		} else {
			table.Add(" " + strings.TrimSpace(strings.TrimLeft(line, "- ")))
			hasEntry = true
		}
	}
	// table ending fence
	if hasEntry {
		table.Add("|")
	}
	return &table
}

func getHeaderLevel(line *string) (n int) {
	for _, c := range *line {
		if c == '#' {
			n++
		} else {
			break
		}
	}
	return n
}
