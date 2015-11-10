package scraper

import (
	"fmt"
	"os/exec"
	"strings"
)

func GetCommandOutput(cmd string) []byte {
	parts := strings.Fields(cmd)
	head := parts[0]
	parts = parts[1:len(parts)]

	out, err := exec.Command(head, parts...).Output()
	if err != nil {
		fmt.Printf("%s", err)
	}

	return out
}
