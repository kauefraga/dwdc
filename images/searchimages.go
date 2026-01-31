package images

import (
	"log"
	"os"
	"os/user"
	"path/filepath"
	"strings"
)

func getUserDesktopDirectoryPath() string {
	user, err := user.Current()
	if err != nil {
		log.Fatal("user.current failed to resolve")
	}

	return filepath.Join(user.HomeDir, "desktop")
}

func getDwdcImagesPath() []string {
	desktopPath := getUserDesktopDirectoryPath()

	entries, err := os.ReadDir(desktopPath)
	if err != nil {
		log.Fatal("failed to read desktop directory")
	}

	var dwdcImages []string

	for _, e := range entries {
		if e.IsDir() {
			continue
		}

		isDwdcImage := strings.HasPrefix(e.Name(), "dwdc") && strings.HasSuffix(e.Name(), ".png")

		if isDwdcImage {
			imagePath := filepath.Join(desktopPath, e.Name())
			dwdcImages = append(dwdcImages, imagePath)
		}
	}

	return dwdcImages
}
