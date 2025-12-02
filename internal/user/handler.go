package user

import (
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"

	"github.com/PNamGP1120/ougreenplus-go/internal/common"
	"github.com/PNamGP1120/ougreenplus-go/internal/database"
)

func List(c *fiber.Ctx) error {
	var users []User
	database.DB.Order("id ASC").Find(&users)

	return c.JSON(common.Success(users))
}

func Get(c *fiber.Ctx) error {
	id, err := common.ParseUint(c.Params("id"))
	if err != nil {
		return fiber.ErrBadRequest
	}

	var u User
	if err := database.DB.First(&u, id).Error; err != nil {
		return fiber.ErrNotFound
	}

	u.Password = ""
	return c.JSON(common.Success(u))
}

func Create(c *fiber.Ctx) error {
	var dto CreateUserDTO
	if err := c.BodyParser(&dto); err != nil {
		return fiber.ErrBadRequest
	}

	hash, _ := bcrypt.GenerateFromPassword([]byte(dto.Password), 12)

	u := &User{
		Email:    dto.Email,
		Password: string(hash),
		Role:     dto.Role,
	}

	if err := database.DB.Create(u).Error; err != nil {
		return err
	}

	u.Password = ""
	return c.JSON(common.Success(u))
}

func Update(c *fiber.Ctx) error {
	id, err := common.ParseUint(c.Params("id"))
	if err != nil {
		return fiber.ErrBadRequest
	}

	var dto UpdateUserDTO
	if err := c.BodyParser(&dto); err != nil {
		return fiber.ErrBadRequest
	}

	var u User
	if err := database.DB.First(&u, id).Error; err != nil {
		return fiber.ErrNotFound
	}

	u.Email = dto.Email
	u.Role = dto.Role

	if dto.Password != "" {
		hash, _ := bcrypt.GenerateFromPassword([]byte(dto.Password), 12)
		u.Password = string(hash)
	}

	if err := database.DB.Save(&u).Error; err != nil {
		return err
	}

	u.Password = ""
	return c.JSON(common.Success(u))
}

func Delete(c *fiber.Ctx) error {
	id, err := common.ParseUint(c.Params("id"))
	if err != nil {
		return fiber.ErrBadRequest
	}

	if err := database.DB.Delete(&User{}, id).Error; err != nil {
		return err
	}

	return c.JSON(common.Success(true))
}
