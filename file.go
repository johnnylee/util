package util

import (
	"os"
	"os/user"
	"path/filepath"
)

// fileExists: Return true if the path exists. Arguments are joined with
// filepath.Join to construct the full path.
func FileExists(elem ...string) bool {
	path := filepath.Join(elem...)
	_, err := os.Stat(path)
	return err == nil
}

// ExpandPath: Expand the path to its full path.
func ExpandPath(path string) (string, error) {
	var err error

	if len(path) == 0 {
		return path, nil
	}

	if path[0] == '~' {
		usr, err := user.Current()
		if err != nil {
			return "", err
		}
		path = filepath.Join(usr.HomeDir, path[1:])
	}

	if path, err = filepath.Abs(path); err != nil {
		return "", err
	}

	return path, nil
}
