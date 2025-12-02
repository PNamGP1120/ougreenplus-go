package auth

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"

	"github.com/PNamGP1120/ougreenplus-go/internal/config"
	"github.com/PNamGP1120/ougreenplus-go/internal/database"
	"github.com/PNamGP1120/ougreenplus-go/internal/user"
)

type Handler struct {
	cfg *config.Config
}

func NewHandler(cfg *config.Config) *Handler {
	return &Handler{cfg: cfg}
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (h *Handler) Login(c *fiber.Ctx) error {
	var req LoginRequest
	if err := c.BodyParser(&req); err != nil {
		return fiber.ErrBadRequest
	}

	var u user.User
	if err := database.DB.Where("email = ?", req.Email).First(&u).Error; err != nil {
		return fiber.ErrUnauthorized
	}

	if err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(req.Password)); err != nil {
		return fiber.ErrUnauthorized
	}

	access, refresh, err := generateTokens(&u, h.cfg)
	if err != nil {
		return fiber.ErrInternalServerError
	}

	return c.JSON(fiber.Map{
		"access_token":  access,
		"refresh_token": refresh,
	})
}

type RefreshRequest struct {
	RefreshToken string `json:"refresh_token"`
}

func (h *Handler) Refresh(c *fiber.Ctx) error {
	var req RefreshRequest
	if err := c.BodyParser(&req); err != nil {
		return fiber.ErrBadRequest
	}

	token, err := jwt.Parse(req.RefreshToken, func(t *jwt.Token) (interface{}, error) {
		return []byte(h.cfg.JWTSecret), nil
	})
	if err != nil || !token.Valid {
		return fiber.ErrUnauthorized
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return fiber.ErrUnauthorized
	}

	uidFloat, ok := claims["sub"].(float64)
	if !ok {
		return fiber.ErrUnauthorized
	}
	uid := uint(uidFloat)

	var u user.User
	if err := database.DB.First(&u, uid).Error; err != nil {
		return fiber.ErrUnauthorized
	}

	access, refresh, err := generateTokens(&u, h.cfg)
	if err != nil {
		return fiber.ErrInternalServerError
	}

	return c.JSON(fiber.Map{
		"access_token":  access,
		"refresh_token": refresh,
	})
}

func (h *Handler) Me(c *fiber.Ctx) error {
	uidVal := c.Locals("user_id")
	uid, ok := uidVal.(uint)
	if !ok {
		return fiber.ErrUnauthorized
	}

	var u user.User
	if err := database.DB.First(&u, uid).Error; err != nil {
		return fiber.ErrNotFound
	}

	// Không trả password
	u.Password = ""
	return c.JSON(u)
}

func (h *Handler) Logout(c *fiber.Ctx) error {
	// MVP: frontend tự xoá token
	return c.JSON(fiber.Map{"success": true})
}
