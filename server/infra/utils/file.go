package utils

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

func FileSize(path string) int64 {
	fi, err := os.Stat(path)
	if err != nil {
		return 0
	}
	return fi.Size()
}

func FileMtime(path string) int64 {
	fi, err := os.Stat(path)
	if err != nil {
		return 0
	}
	return fi.ModTime().Unix()
}

func FileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func EnsureDirExists(filename string) {
	_, err := os.Stat(filename)
	if os.IsNotExist(err) {
		os.MkdirAll(filename, 0755)
	}
}

func ReadJSONFile(fpath string, obj any) {
	d, err := os.ReadFile(fpath)
	if err == nil {
		json.Unmarshal(d, obj)
	}
}

func WriteJSONFile(fpath string, obj any) error {
	var err error
	var d []byte

	if d, err = json.MarshalIndent(obj, "", "  "); err != nil {
		return err
	}
	return os.WriteFile(fpath, d, 0644)
}

func FileChecksum(fpath string) []byte {
	f, err := os.Open(fpath)
	if err != nil {
		return []byte("")
	}

	hash := md5.New()
	if _, err := io.Copy(hash, f); err != nil {
		log.Fatal(err)
	}
	return hash.Sum(nil)
}

func FileChecksumHex(fpath string) string {
	return hex.EncodeToString(FileChecksum(fpath))
}

// file needs to be closed and removed later
func DownloadToTempFile(url string, destDir string) (string, error) {
	EnsureDirExists(destDir)
	tmpfile, err := os.CreateTemp(destDir, ".temp-")
	if err != nil {
		return "", fmt.Errorf("create temp file: %w", err)
	}

	// defer func() {
	// 	tmpfile.Close()
	// 	os.Remove(tmpPath)
	// }()

	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	_, err = io.Copy(tmpfile, resp.Body)
	if err != nil {
		return "", err
	}
	return tmpfile.Name(), nil
}

func DownloadToFile(url string, dest string) error {
	EnsureDirExists(filepath.Dir(dest))
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	out, err := os.Create(dest)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	return err
}

func RemoveDir(path string) error {
	if strings.Count(path, "/") < 2 {
		return fmt.Errorf("refusing to remove %s", path)
	}

	// Get the file info for the path
	info, err := os.Stat(path)
	if err != nil {
		return err
	}

	// Check if the path is a directory
	if !info.IsDir() {
		return fmt.Errorf("%s is not a directory", path)
	}

	// Get all files and subdirectories in the directory
	files, err := filepath.Glob(filepath.Join(path, "*"))
	if err != nil {
		return err
	}

	// Remove all files and subdirectories
	for _, file := range files {
		err = os.RemoveAll(file)
		if err != nil {
			return err
		}
	}

	// Remove the directory itself
	err = os.Remove(path)
	if err != nil {
		return err
	}

	return nil
}
