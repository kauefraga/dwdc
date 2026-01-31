package images

import (
	"fmt"
	"os"
)

func DeleteDwdcImages() {
	dwdcImages := getDwdcImagesPath()

	for _, imagePath := range dwdcImages {
		err := os.Remove(imagePath)
		if err != nil {
			fmt.Println("failed to delete file:", err)
		}
	}
}
