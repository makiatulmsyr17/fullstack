package controllers

import (
	"makiatul/database"
	"makiatul/models"

	"github.com/gofiber/fiber/v2"
)

func Reads(c *fiber.Ctx) error {
	var users []models.User

	database.DB.Find(&users)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Data tersedia",
		"Data":    users,
	})
}
func Read(c *fiber.Ctx) error {
	var user models.User

	id := c.Params("id")

	if id == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "id tidak boleh kosong",
		})
	}
	if err := database.DB.Where("id = ?", id).First(&user).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "gagal memuat data user",
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Data Ditemukan",
		"data":    user,
	})
}

func Create(c *fiber.Ctx) error {
	var UserReq models.UserReq

	if err := c.BodyParser(&UserReq); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"Messege": "Gagal parsing JSON",
		})
	}
	var user models.User
	user.Nama = UserReq.Nama
	user.Kelas = UserReq.Kelas
	user.Prodi = UserReq.Prodi

	if err := database.DB.Create(&user).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Gagal Menambahkan data User",
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Data user berhasil di buat",
		"user":    user,
	})
}
func Update(c *fiber.Ctx) error {
	userUpdate := new(models.UserReq)
	if err := c.BodyParser(&userUpdate); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"Messege": "Gagal parsing JSON",
		})
	}
	id := c.Params("id")

	if id == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "id tidak boleh kosong",
		})

	}
	user := models.User{}

	if err := database.DB.First(&user, "id =?", id).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "id tidak boleh kosong",
		})
	}
	user.Nama = userUpdate.Nama
	user.Kelas = userUpdate.Kelas
	user.Prodi = userUpdate.Prodi

	if err := database.DB.Model(&user).Error; err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{})
	}
	result := database.DB.Where("id = ?", id).Updates(&user)
	if result.Error != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": " eror model",
		})
	}
	if result.RowsAffected == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": " eror RowsAFFected",
		})

	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"Message": "Data berhasil di update",
		"data":    user,
	})
}

func Delete(c *fiber.Ctx) error {
	c.Set("Acces-Control-Allow_origin", "hhtp:/localhost")
	user := models.User{}
	id := c.Params("id")

	if err := database.DB.First(&user, "id = ?", id).Delete(&user).Error; err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"messege": "Data gagal",
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"massage": "Data berhasil di hapus",
	})
}
