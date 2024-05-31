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
	"github.com/zerotohero-dev/stash/internal/http/action"
	"github.com/zerotohero-dev/stash/internal/http/version"
	"github.com/zerotohero-dev/stash/internal/io"
	"log"
	"time"
)

func WatchUpstream() {
	ticker := time.NewTicker(5 * time.Second)
	go func() {
		for {
			select {
			case <-ticker.C:
				remoteVersion, links, err := action.FetchRemoteLinks()
				if err != nil {
					log.Printf("Failed to download links: %v\n", err)
					continue
				}

				currentVersion := version.Current()
				if remoteVersion <= currentVersion {
					continue
				}

				if err := io.Save(links); err != nil {
					log.Printf("Failed to save links file: %v\n", err)
					continue
				}

				if err := version.Update(remoteVersion); err != nil {
					log.Printf("Failed to update version: %v\n", err)
					continue
				}
			}
		}
	}()
}
