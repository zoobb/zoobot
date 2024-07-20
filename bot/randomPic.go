package bot

import (
	"math/rand"
	"os"
	"path/filepath"
	"time"
)

func RandomPic(path string) (string, error) {
	files, err := os.ReadDir(path)
	if err != nil {
		return "", err
	}
	rand.Seed(time.Now().UnixNano())

	randomIndex := rand.Intn(len(files))
	randomFileName := files[randomIndex].Name()

	fullPath := filepath.Join(path, randomFileName)

	return fullPath, nil
}
