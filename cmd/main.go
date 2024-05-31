//   __________
//  /\____;;___\
// | /         /
// `. ())oo() .    Stash
//  |\(%()*^^()^\
//  | |-%-------|  A hoarder's dream come true.
//  \ | %  ))   |
//   \|%________|

package main

import (
	"fmt"
	"github.com/zerotohero-dev/stash/internal/io"
	"github.com/zerotohero-dev/stash/internal/monitor"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	closer := make(chan struct{})

	monitor.Watch(io.Parse, closer)
	monitor.WatchUpstream()

	// Block the process from exiting, but also be graceful and honor the
	// termination signals that may come from the orchestrator.
	s := make(chan os.Signal, 1)
	signal.Notify(s, syscall.SIGINT, syscall.SIGTERM)
	select {
	case e := <-s:
		closer <- struct{}{}
		fmt.Println(e)
		panic("bye cruel world!")
	}

}
