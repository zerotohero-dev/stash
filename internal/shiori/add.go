//   __________
//  /\____;;___\
// | /         /
// `. ())oo() .    Stash
//  |\(%()*^^()^\
//  | |-%-------|  A hoarder's dream come true.
//  \ | %  ))   |
//   \|%________|

package shiori

import (
	"fmt"
	"github.com/zerotohero-dev/stash/internal/env"
	"log"
	"os/exec"
	"os/user"
	"strings"
)

const versionPrefix = "version="

func Add(link string) {
	if strings.HasPrefix(link, versionPrefix) {
		return
	}

	cmd := exec.Command(env.StashShiori(), "add", link)

	usr, err := user.Current()
	if err != nil {
		fmt.Printf("failed to get current user: %v", err)
		return
	}

	cmd.Dir = usr.HomeDir

	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Printf("Failed to execute shiori: %v, Output: %s\n", err, output)
		return
	}

	fmt.Printf("link:\n%s\nout:\n%s\n", link, output)
}
