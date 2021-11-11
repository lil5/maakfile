package main

import (
	"log"

	"github.com/rjeczalik/notify"
)

func main() {
	// Make the channel buffered to ensure no event is dropped. Notify will drop
	// an event if the receiver is not able to keep up the sending pace.
	c := make(chan notify.EventInfo, 1)

	// Set up a watchpoint listening for inotify-specific events within a
	// current working directory. Dispatch each InCloseWrite and InMovedTo
	// events separately to c.
	if err := notify.Watch(".", c, notify.InCloseWrite, notify.InMovedTo); err != nil {
		log.Fatal(err)
	}
	defer notify.Stop(c)

	// Block until an event is received.
	switch ei := <-c; ei.Event() {
	case notify.InCloseWrite:
		log.Println("Editing of", ei.Path(), "file is done.")
	case notify.InMovedTo:
		log.Println("File", ei.Path(), "was swapped/moved into the watched directory.")
	}
}
