package osx

import (
	"os"
	"os/user"
	"path/filepath"
)

func FindFiles(path string) (files []string, err error) {
	var entries []os.DirEntry

	if entries, err = os.ReadDir(path); err != nil {
		return nil, err
	}

	for _, entry := range entries {
		var info os.FileInfo

		if info, err = entry.Info(); err != nil {
			return nil, err
		}

		if !info.IsDir() {
			files = append(files, info.Name())
		}
	}

	return files, nil
}

func GetHomeDirectory() (home string, err error) {
	var currentUser *user.User

	if currentUser, err = user.Current(); err != nil {
		return "", err
	}

	return currentUser.HomeDir, nil
}

func IsHomeDirectory(path string) (isHome bool, err error) {
	var home string

	if home, err = GetHomeDirectory(); err != nil {
		return false, err
	}

	if path, err = filepath.EvalSymlinks(path); err != nil {
		return false, err
	}

	return home == path, nil
}

type WalkDirectoryUpFunc func(path string) error

// WalkDirectoryUp Walks up the directory structure until the home directory.
// A function is called for every directory.
// Returns an error if it failed.
func WalkDirectoryUp(path string, fn WalkDirectoryUpFunc) (err error) {
	var ok bool

	for {
		err = fn(path)

		if ok, _ = IsHomeDirectory(path); ok {
			break
		}

		path = filepath.Dir(path)
	}

	return err
}
