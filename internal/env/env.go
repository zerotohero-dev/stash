//   __________
//  /\____;;___\
// | /         /
// `. ())oo() .    Stash
//  |\(%()*^^()^\
//  | |-%-------|  A hoarder's dream come true.
//  \ | %  ))   |
//   \|%________|

package env

import "os"

func SeedUrl() string {
	return os.Getenv("STASH_SEED_URL")
}

func StashInbox() string {
	return os.Getenv("STASH_INBOX")
}

func StashShiori() string {
	return os.Getenv("STASH_SHIORI")
}

func VersionFileName() string {
	return os.Getenv("STASH_VERSION_FILE_NAME")
}
