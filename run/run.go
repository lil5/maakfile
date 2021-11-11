package run

import "github.com/lil5/maakfile/global/classes"

func Run(c *classes.Config, command *classes.Command) error {
	switch command.Type {
	case classes.TypeSequential:
		execSequential(command.Run)
	case classes.TypeParallel:
		execParallel(command.Run)
	case classes.TypeWatch:
		err := runWatch(c, command.Name, command.Run)
		if err != nil {
			return err
		}
	}

	return nil
}

func runWatch(c *classes.Config, name string, extensions []string) error {
	command, err := c.FindCommandFromName(name)
	if err != nil {
		return err
	}

	run := func() error {
		execSequential(command.Run)
		return nil
	}

	return watch(extensions, run)
}
