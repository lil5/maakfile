package global

import "fmt"

var ErrEmptyConfig = fmt.Errorf("no scripts or parallel or watch configuration found")
var ErrBadCommandType = func(name string) error {
	return fmt.Errorf("command %s is does not match string type or []string type", name)
}

var ErrWriteFile = fmt.Errorf("unable to write config file")
var ErrFileAlreadyExists = fmt.Errorf("config file already exists")
