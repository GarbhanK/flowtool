package utils

import (
	"fmt"

	"golang.design/x/clipboard"
)

func ExportToClipboard(templatedStr string) {
	err := clipboard.Init()
	if err != nil {
		fmt.Printf("Unable to init clipboard, %s", err.Error())
	}

	byteSql := []byte(templatedStr)
	clipboard.Write(clipboard.FmtText, byteSql)
}

func ClipboardToString() string {
	err := clipboard.Init()
	if err != nil {
		fmt.Printf("Unable to init clipboard, %s", err.Error())
	}

	curr_clipboard := clipboard.Read(clipboard.FmtText)

	return string(curr_clipboard)
}
