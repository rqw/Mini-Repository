package util

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"errors"
	"fmt"
	"os"
	"path"
)

func CreateParentIfNotExist(file string) error {
	dirPath := path.Dir(file)

	if stat, err := os.Stat(dirPath); err != nil && os.IsNotExist(err) {
		if err = os.MkdirAll(dirPath, 0755); err != nil {
			return err
		}
	} else if !stat.IsDir() {
		return errors.New(fmt.Sprintf("%s is not a dir\n", dirPath))
	}
	return nil
}

func CreateFileIfNotExist(file string) error {
	if _, err := os.Stat(file); err != nil && os.IsNotExist(err) {
		if err = CreateParentIfNotExist(file); err != nil {
			return err
		}
		if _, err := os.Create(file); err != nil {
			return err
		}
	}
	return nil
}

func CheckFileExist(file string) (bool, error) {
	if _, err := os.Stat(file); err != nil && os.IsNotExist(err) {
		return false, nil
	} else if err == nil {
		return true, nil
	} else {
		return false, err
	}
}

func GetHash(file []byte, hash string) []byte {
	switch hash {
	case "md5":
		return []byte(fmt.Sprintf("%x", md5.Sum(file)))
	case "sha1":
		return []byte(fmt.Sprintf("%x", sha1.Sum(file)))
	case "sha256":
		return []byte(fmt.Sprintf("%x", sha256.Sum256(file)))
	case "sha512":
		return []byte(fmt.Sprintf("%x", sha512.Sum512(file)))
	default:
		return nil
	}
}
