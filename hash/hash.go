package hash

import (
	"hash"
	"io"
	"os"
	"path/filepath"
)

func hashText(hash hash.Hash, text string) ([]byte, error) {
	_, err := hash.Write([]byte(text))
	if err != nil {
		return nil, err
	}
	return hash.Sum(nil), nil
}

func hashFile(hash hash.Hash, path string) ([]byte, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	if _, err := io.Copy(hash, file); err != nil {
		return nil, err
	}
	return hash.Sum(nil), nil
}

func hashPath(hash hash.Hash, path string) ([]byte, error) {
	info, err := os.Stat(path)
	if err != nil {
		return nil, err
	}
	switch mode := info.Mode(); {
	case mode.IsDir():
		return hashDir(hash, path)
	case mode.IsRegular():
		return hashFile(hash, path)
	}
	return nil, nil // TODO: Handle Error
}

func hashDir(hash hash.Hash, path string) ([]byte, error) {
	err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return nil
		}
		io.WriteString(hash, path)
		return nil
	})
	if err != nil {
		return nil, nil
	}
	return hash.Sum(nil), nil
}
