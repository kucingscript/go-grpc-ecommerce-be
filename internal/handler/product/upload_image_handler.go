package product

import (
	"fmt"
	"net/http"
	"path/filepath"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
)

const MAX_FILE_SIZE = 2 * 1024 * 1024

func UploadProductImageHandler(c *fiber.Ctx) error {
	file, err := c.FormFile("image")
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Image file not found",
		})
	}

	if file.Size > MAX_FILE_SIZE {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "File size exceeds the limit of 2 MB",
		})
	}

	ext := strings.ToLower(filepath.Ext(file.Filename))
	allowedExts := map[string]bool{
		".jpg":  true,
		".jpeg": true,
		".png":  true,
		".webp": true,
	}

	if !allowedExts[ext] {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "extension not allowed (jpg, jpeg, png, webp)",
		})
	}

	contentType := file.Header.Get("Content-Type")
	allowedContentTypes := map[string]bool{
		"image/jpeg": true,
		"image/png":  true,
		"image/webp": true,
	}

	if !allowedContentTypes[contentType] {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "content type not allowed (jpg, jpeg, png, webp)",
		})
	}

	timestamp := time.Now().UnixNano()
	fileName := fmt.Sprintf("product_%d%s", timestamp, filepath.Ext(file.Filename))
	uploadPath := "./storage/product/" + fileName

	err = c.SaveFile(file, uploadPath)
	if err != nil {
		fmt.Println(err)
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "Failed to save file",
		})
	}

	return c.JSON(fiber.Map{
		"success":   true,
		"message":   "Image uploaded successfully",
		"file_name": fileName,
	})
}
