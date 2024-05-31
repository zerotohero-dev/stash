//   __________
//  /\____;;___\
// | /         /
// `. ())oo() .    Stash
//  |\(%()*^^()^\
//  | |-%-------|  A hoarder's dream come true.
//  \ | %  ))   |
//   \|%________|

package version

import (
	"github.com/zerotohero-dev/stash/internal/env"
	"log"
	"os"
	"strconv"
)

func Current() int {
	versionFileName := env.VersionFileName()
	data, err := os.ReadFile(versionFileName)
	if err != nil {
		if os.IsNotExist(err) {
			log.Println("version file does not exist. Assuming version 0")
			return 0
		}

		log.Println("failed to read version file:", err.Error())
		return 0
	}

	version, err := strconv.Atoi(string(data))
	if err != nil {
		log.Println("failed to parse version:", err.Error())
		return 0
	}

	return version
}

func Update(newVersion int) error {
	versionFileName := env.VersionFileName()

	return os.WriteFile(versionFileName, []byte(strconv.Itoa(newVersion)), 0644)
}
