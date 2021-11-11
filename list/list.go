package list

import (
	"fmt"
	"strings"

	"github.com/lil5/maakfile/global"
	"github.com/lil5/maakfile/global/classes"
)

type list struct {
	Sequential []string
	Parallel   []string
	Watch      []string
}

func PrintList(c *classes.Config) {
	l := list{}
	// collect data to l
	if c.Sequential != nil {
		for key := range *(c.Sequential) {
			l.Sequential = append(l.Sequential, key)
		}
	}
	if c.Parallel != nil {
		for key := range *(c.Parallel) {
			l.Parallel = append(l.Parallel, key)
		}
	}
	if c.Watch != nil {
		for key := range *(c.Watch) {
			l.Watch = append(l.Watch, fmt.Sprintf("watch-%s", key))
		}
	}

	// print l to stdout
	print("Sequential", l.Sequential)
	print("Parallel", l.Parallel)
	print("Watch", l.Watch)

}

func print(name string, values []string) {
	if len(values) != 0 {
		fmt.Printf(
			"%s%s:%s\n\t%s\n", global.ColorGreen, name, global.ColorReset, strings.Join(values, ", "))
	}
}
