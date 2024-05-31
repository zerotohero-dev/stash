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
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"github.com/zerotohero-dev/stash/internal/env"
	"log"
	"os"
	"path/filepath"
)

func randomFileName() string {
	bytes := make([]byte, 16)
	if _, err := rand.Read(bytes); err != nil {
		log.Fatalf("Failed to generate random file name: %v\n", err)
	}
	return hex.EncodeToString(bytes) + ".links.txt"
}

func Save(links []string) error {
	// TODO: ensure these variables exist at app startup.
	inbox := env.StashInbox()
	fileName := filepath.Join(inbox, randomFileName())

	file, err := os.Create(fileName)
	if err != nil {
		return fmt.Errorf("failed to create file: %v", err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Printf("Error while closing file: %v\n", err)
		}
	}(file)

	for _, link := range links {
		if _, err := file.WriteString(link + "\n"); err != nil {
			return fmt.Errorf("failed to write link: %v", err)
		}
	}

	return nil
}
