package cmd

import (
	"fmt"
	"strings"

	"github.com/rodaine/table"
)

func returnTable() table.Table {
	table.DefaultHeaderFormatter = func(format string, args ...interface{}) string {
		return strings.ToUpper(fmt.Sprintf(format, args...))
	}

	return table.New("ID", "Title", "Created").WithHeaderSeparatorRow('-')
}
