package product

import (
	"log"
	"mime"
	"net/http"
	"os"
	"path"

	"github.com/gofiber/fiber/v2"
)

func GetProductImageHandler(c *fiber.Ctx) error {
	fileNameParam := c.Params("filename")
	filePath := path.Join("storage", "product", fileNameParam)

	if _, err := os.Stat(filePath); err != nil {
		if os.IsNotExist(err) {
			return c.Status(http.StatusNotFound).JSON(fiber.Map{
				"success": false,
				"message": "Image file not found",
			})
		}

		log.Println(err)
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "Failed to get image file",
		})
	}

	file, err := os.Open(filePath)
	if err != nil {
		log.Println(err)
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "Failed to open image file",
		})
	}

	ext := path.Ext(filePath)
	mimeType := mime.TypeByExtension(ext)

	c.Set("Content-Type", mimeType)
	return c.SendStream(file)
}
