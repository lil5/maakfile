package classes

import (
	"fmt"
	"io/ioutil"

	"github.com/BurntSushi/toml"
	"github.com/lil5/maakfile/global"
)

type Config struct {
	Sequential *map[string]interface{} `toml:"sequential"`
	Parallel   *map[string][]string    `toml:"parallel"`
	Watch      *map[string]string      `toml:"watch"`
}

func GetConfig() (*Config, error) {
	file, err := ioutil.ReadFile(global.Fpath)
	if err != nil {
		return nil, err
	}
	var c Config

	if _, err = toml.Decode(string(file), &c); err != nil {
		return nil, err
	}

	if c.Sequential == nil && c.Parallel == nil && c.Watch == nil {
		return nil, global.ErrEmptyConfig
	}

	return &c, nil
}

func (c *Config) FindCommandFromName(name string) (*Command, error) {
	// scripts
	if c.Sequential != nil {
		for key, value := range *(c.Sequential) {
			if key != name {
				continue
			}

			r, err := fromStringOrStrings(value)
			if err != nil {
				return nil, global.ErrBadCommandType(name)
			}

			return &Command{
				Type: TypeSequential,
				Run:  r,
			}, nil
		}
	}

	// parallel
	if c.Parallel != nil {
		for key, value := range *(c.Parallel) {
			if key != name {
				continue
			}

			v := value
			return &Command{
				Type: TypeParallel,
				Run:  v,
			}, nil
		}
	}

	// watch
	if c.Watch != nil {
		for key, value := range *(c.Watch) {
			watchKey := fmt.Sprintf("watch-%s", key)
			if watchKey != name {
				continue
			}

			return &Command{
				Type: TypeWatch,
				Run:  []string{key, value},
			}, nil
		}
	}

	return nil, global.ErrEmptyConfig
}

func fromStringOrStrings(value interface{}) ([]string, error) {
	r := []string{}
	switch v := value.(type) {
	case string:
		r = []string{v}
	case []interface{}:
		for i := range v {
			if s, ok := v[i].(string); ok {
				r = append(r, s)
			} else {
				return nil, fmt.Errorf("invalid value")

			}
		}
	}
	return r, nil
}
