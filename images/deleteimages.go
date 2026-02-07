package images

import (
	"fmt"
	"os"
)

// search for dwdc images ("dwdc-1.png", ...) in user's desktop area and delete them
func deleteDwdcImages() {
	dwdcImages := getDwdcImagesPath()

	for _, imagePath := range dwdcImages {
		err := os.Remove(imagePath)
		if err != nil {
			fmt.Println("failed to delete file:", err)
		}
	}
}
