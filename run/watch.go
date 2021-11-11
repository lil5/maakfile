package run

import (
	"fmt"
	"log"
	"path/filepath"
	"strings"
	"time"

	"github.com/rjeczalik/notify"
)

// const batchTime = time.Millisecond * 200
// const eventBuffer = 3

// CommonExcludes is a list of commonly excluded files suitable for passing in
// the excludes parameter to Watch - includes repo directories, temporary
// files, and so forth.
var CommonExcludes = []string{
	// VCS
	"**/.git/**",
	"**/.hg/**",
	"**/.svn/**",
	"**/.bzr/**",

	// OSX
	"**/.DS_Store/**",

	// Temporary files
	"**.tmp",
	"**~",
	"**#",
	"**.bak",
	"**.swp",
	"**.___jb_old___",
	"**.___jb_bak___",
	"**mage_output_file.go",

	// Python
	"**.py[cod]",

	// Node
	"**/node_modules/**",
}

func watch(extensions []string, run func() error) error {
	fsEvents := make(chan notify.EventInfo, 1)

	err := notify.Watch(`./...`, fsEvents, notify.All)
	if err != nil {
		log.Fatal(err)
	}

	err = run()
	if err != nil {
		return err
	}

	for {
		event := <-fsEvents
		if !matchExtOrDir(extensions, event.Path()) {
			fmt.Printf("ignored %s\n", event.Path())
			continue
		}

		fmt.Printf("%s\n", event.Path())
		err = run()
		if err != nil {
			return err
		}

		time.Sleep(200 * time.Millisecond)
		select {
		case <-fsEvents:
		default:
		}
	}
}

func matchExtOrDir(extensions []string, path string) bool {
	ok := false
	base := filepath.Base(path)

	// only checks if the name has no extension
	isDir := filepath.Ext(base) == ""

	if isDir {
		ok = true
	} else {
		for _, e := range extensions {
			if e == "*" {
				ok = true
				break
			}

			if strings.HasSuffix(base, e) {
				ok = true
				break
			}
		}
	}
	//fmt.Printf("ok: %t\text: %v\tpath: %s\n", ok, extensions, path)

	return ok
}
