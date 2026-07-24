package lib

import (
	"strings"

	"github.com/J-Siu/go-md2table/global"
)

func ToTable(lines *[]string) *[]string {
	var (
		hasEntry       bool
		headerLevel    int
		headerLevelMax int
		table          []string
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
		table = append(table, strings.Repeat("|", headerLevelMax)+"||"+global.LineBreak)
		table = append(table, strings.Repeat("|--", headerLevelMax)+"|--|")
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
					table = append(table, "|")
				} else {
					table = append(table, strings.Repeat("|", headerLevelMax-headerLevel+1))
				}
			}
			// line start
			table = append(table, global.LineBreak+strings.Repeat("|", headerLevel))
			// header to column
			table = append(table, strings.TrimSpace(strings.TrimLeft(line, "#"))+strings.Repeat("|", headerLevelMax-headerLevel+1))
			hasEntry = false
		} else {
			table = append(table, " "+strings.TrimSpace(strings.TrimLeft(line, "- ")))
			hasEntry = true
		}
	}
	// table ending fence
	if hasEntry {
		table = append(table, "|")
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
