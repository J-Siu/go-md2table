package lib

import (
	"strings"

	"github.com/J-Siu/go-helper/v2/array"
	"github.com/J-Siu/go-md2table/global"
)

func ToTable(lines *[]string) *array.Array[string] {
	var (
		hasLink        bool
		headerLever    int
		headerLevelMax int
		table          array.Array[string]
	)
	// Get max level
	for _, line := range *lines {
		level := tableLevel(&line)
		if level > headerLevelMax {
			headerLevelMax = level
		}
	}
	// Add title
	if headerLevelMax > 0 {
		table.Add(strings.Repeat("|", headerLevelMax) + "||" + global.LineBreak)
		table.Add(strings.Repeat("|--", headerLevelMax) + "|--|")
	}
	// Rows
	for n, l := range *lines {
		if len(l) == 0 {
			continue
		}
		line := strings.TrimSpace(l)
		headerLever = tableLevel(&line)
		if headerLever > 0 {
			// line end
			if n > 0 {
				if hasLink {
					table.Add("|")
				} else {
					table.Add(strings.Repeat("|", headerLevelMax-headerLever+1))
				}
			}
			// line start
			table.Add(global.LineBreak + strings.Repeat("|", headerLever))
			// header
			table.Add(strings.TrimSpace(strings.TrimLeft(line, "#")) + strings.Repeat("|", headerLevelMax-headerLever+1))
			hasLink = false
		} else {
			table.Add(" " + line)
			hasLink = true
		}
	}
	// Ending table fence of last line
	if hasLink {
		table.Add("|")
	}
	return &table
}

func tableLevel(line *string) (n int) {
	for _, c := range *line {
		if c == '#' {
			n++
		} else {
			break
		}
	}
	return n
}
