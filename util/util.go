package util

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"

	jww "github.com/spf13/jwalterweatherman"
)

// Exists checks if a file or directory exists.
func Exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

// FindInDirs looks for a filename in configured directories,
// and returns the first matching file path.
func FindInDirs(fname string, dirs []string) (string, error) {
	jww.INFO.Printf("searching for %s in %s", fname, dirs)
	for _, d := range dirs {
		fpath := filepath.Join(d, fname)
		if b, _ := Exists(fpath); b {
			return fpath, nil
		}
	}

	jww.INFO.Printf("unable to find %s in %s", fname, dirs)
	return "", fmt.Errorf("unable to find %s in %s", fname, dirs)
}

// ForceWrite forcibly writes a string to a given filepath.
func ForceWrite(path string, contents string) error {
	if err := os.MkdirAll(filepath.Dir(path), 0755); err != nil {
		return err
	}
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()
	_, err = f.WriteString(contents)
	if err != nil {
		return err
	}
	return nil
}

// GetCA returns the certificate authority pem bytes.
func GetCA(dirs []string) ([]byte, error) {
	caPath, err := FindInDirs("ca.pem", dirs)
	if err != nil {
		return nil, err
	}
	ca, err := ioutil.ReadFile(caPath)
	if err != nil {
		return nil, err
	}
	return ca, nil
}

// Timeout starts a go routine which writes true to the given channel
// after the given time.
func Timeout(d time.Duration) <-chan bool {
	ch := make(chan bool, 1)
	go func() {
		time.Sleep(d)
		ch <- true
	}()
	return ch
}
