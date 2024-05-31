//   __________
//  /\____;;___\
// | /         /
// `. ())oo() .    Stash
//  |\(%()*^^()^\
//  | |-%-------|  A hoarder's dream come true.
//  \ | %  ))   |
//   \|%________|

package monitor

import (
	"github.com/fsnotify/fsnotify"
	"github.com/zerotohero-dev/stash/internal/env"
	"log"
)

func Watch(process func(string), closer <-chan struct{}) {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatalf("failed to create monitor: %v", err)
	}

	links := env.StashInbox()

	log.Printf("Links folder to watch: '%s'\n", links)

	err = watcher.Add(links)
	if err != nil {
		log.Fatalf("failed to add links to monitor: %v\n", err)
		return
	}

	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					log.Println("Watcher events channel closed")
					return
				}
				if event.Op&fsnotify.Create == fsnotify.Create {
					log.Printf("New file detected: %s\n", event.Name)
					process(event.Name)
				}
			case err, ok := <-watcher.Errors:
				if ok && err != nil {
					log.Printf("Watcher error: %v\n", err)
				}
			case <-closer:
				err := watcher.Close()
				if err != nil {
					log.Printf("Error while closing monitor: %v\n", err)
				}
			}
		}
	}()
}
