package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
)

func main() {
	engine := html.New("./", ".html")

	app := fiber.New(fiber.Config{Views: engine})

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("index", nil)
	})

	app.Post("/upload", func(c *fiber.Ctx) error {
		var input struct {
			nama_gambar string
		}

		if err := c.BodyParser(&input); err != nil {
			return c.Status(500).JSON(fiber.Map{
				"message": err.Error(),
			})
		}

		gambar, err := c.FormFile("gambar")
		if err != nil {
			return c.Status(500).JSON(fiber.Map{
				"message": err.Error(),
			})
		}
		fmt.Printf("Nama file : %s \n", gambar.Filename)
		fmt.Printf("Ukuran file (bytes): %d \n", gambar.Size)

		folderUpload := filepath.Join(".", "uploads")
		if err := os.MkdirAll(folderUpload, 0770); err != nil {
			return c.Status(500).JSON(fiber.Map{
				"message": err.Error(),
			})
		}

		if err := c.SaveFile(gambar, "./uploads/"+namaFileBaru); err != nil {
			return c.Status(500).JSON(fiber.Map{
				"message": err.Error(),
			})
		}

		return c.JSON(fiber.Map{
			"title":       input.nama_gambar,
			"nama_gambar": namaFileBaru,
			"message":     "gambar berhasil diupload",
		})
	})

	app.Listen(":8080")
}
