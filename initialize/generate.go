package initialize

import (
	"os"

	"github.com/lil5/maakfile/global"
)

const configContents = "[sequential]\n" +
	"test = \"echo 'this is a test'\"\n" +
	"[parallel]\n" +
	"test:parallel = [\"echo 'test1'\", \"echo 'test2'\"]\n" +
	"[watch]\n" +
	"test = \"*\"\n"

func GenerateConfig() error {
	if _, err := os.Stat(global.Fpath); !os.IsNotExist(err) {
		return global.ErrFileAlreadyExists
	}

	if err := os.WriteFile(global.Fpath, []byte(configContents), 0644); err != nil {
		// fmt.Fprintf(os.Stderr, "%v\n", err)
		return global.ErrWriteFile
	}

	return nil
}
