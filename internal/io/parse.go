//   __________
//  /\____;;___\
// | /         /
// `. ())oo() .    Stash
//  |\(%()*^^()^\
//  | |-%-------|  A hoarder's dream come true.
//  \ | %  ))   |
//   \|%________|

package io

import (
	"bufio"
	"github.com/zerotohero-dev/stash/internal/shiori"
	"log"
	"os"
	"strings"
)

const linkFileSuffix = "links.txt"

func Parse(fileName string) {
	if !strings.HasSuffix(fileName, linkFileSuffix) {
		return
	}

	file, err := os.Open(fileName)
	if err != nil {
		log.Printf("Failed to open file: %v\n", err)
		return
	}

	scanner := bufio.NewScanner(file)
	log.Printf("Scanning: %s\n", file.Name())
	for scanner.Scan() {
		link := scanner.Text()
		log.Println("link:", link)
		link = strings.TrimSpace(link)
		if link == "" {
			continue
		}
		shiori.Add(link)
	}

	if err := scanner.Err(); err != nil {
		log.Printf("Error reading file: %v\n", err)
	}

	// Since we are deleting the file later, close it explicitly
	// instead of a a `defer file.Close()`.
	err = file.Close()
	if err != nil {
		log.Printf("Error closing file: %v\n", err)
		return
	}

	// Done with the file; remove it.
	// TODO: maybe rename it with a timestamp instead of removing it.
	// or move it to a "recycle bin" folder, just in case.
	if err := os.Remove(fileName); err != nil {
		log.Printf("Failed to delete file: %v\n", err)
	}
}
