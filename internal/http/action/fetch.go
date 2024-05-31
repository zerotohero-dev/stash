//   __________
//  /\____;;___\
// | /         /
// `. ())oo() .    Stash
//  |\(%()*^^()^\
//  | |-%-------|  A hoarder's dream come true.
//  \ | %  ))   |
//   \|%________|

package action

import (
	"bufio"
	"fmt"
	"github.com/zerotohero-dev/stash/internal/env"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"
)

// TODO: this is repeated.
const versionPrefix = "version="

func FetchRemoteLinks() (int, []string, error) {
	url := env.SeedUrl()

	resp, err := http.Get(url)
	if err != nil {
		return -1, nil, err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Printf("Error while closing response body: %v\n", err)
		}
	}(resp.Body)

	// TODO: we might need to follow 302 redirects.
	if resp.StatusCode != http.StatusOK {
		log.Printf("Failed to fetch remote links: %v %v\n", resp.Status, resp.StatusCode)

		return -1, nil, nil
	}

	var links []string

	scanner := bufio.NewScanner(resp.Body)

	if !scanner.Scan() {
		log.Printf("Failed to read remote links: %v\n", scanner.Err())
		return -1, nil, fmt.Errorf("failed to read remote links: %v", scanner.Err())
	}

	versionLine := scanner.Text()
	version, err := strconv.Atoi(strings.TrimPrefix(versionLine, versionPrefix))
	if err != nil {
		log.Printf("Failed to parse version: %v\n", err)
		return -1, nil, fmt.Errorf("failed to parse version: %v", err)
	}

	for scanner.Scan() {
		link := scanner.Text()
		link = strings.TrimSpace(link)
		if link == "" {
			continue
		}
		links = append(links, link)
	}

	if err := scanner.Err(); err != nil {
		log.Printf("Error reading remote links: %v\n", err)
		return -1, nil, fmt.Errorf("error reading remote links: %v", err)
	}

	return version, links, nil
}
