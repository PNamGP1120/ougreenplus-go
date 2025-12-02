package newsletter

import (
	"github.com/gofiber/fiber/v2"

	"github.com/PNamGP1120/ougreenplus-go/internal/common"
)

type Handler struct {
	repo Repository
}

func NewHandler() *Handler {
	return &Handler{repo: NewRepository()}
}

func (h *Handler) Subscribe(c *fiber.Ctx) error {
	var dto SubscribeDTO
	if err := c.BodyParser(&dto); err != nil {
		return fiber.ErrBadRequest
	}

	s := &Subscriber{
		Email:    dto.Email,
		IsActive: true,
	}

	if err := h.repo.Subscribe(s); err != nil {
		return err
	}

	return c.JSON(common.Success(s))
}

func (h *Handler) Unsubscribe(c *fiber.Ctx) error {
	var dto SubscribeDTO
	if err := c.BodyParser(&dto); err != nil {
		return fiber.ErrBadRequest
	}

	if err := h.repo.Unsubscribe(dto.Email); err != nil {
		return err
	}

	return c.JSON(common.Success(true))
}

func (h *Handler) List(c *fiber.Ctx) error {
	list, err := h.repo.All()
	if err != nil {
		return err
	}
	return c.JSON(common.Success(list))
}

func (h *Handler) Send(c *fiber.Ctx) error {
	var dto SendMailDTO
	if err := c.BodyParser(&dto); err != nil {
		return fiber.ErrBadRequest
	}

	// TODO: integrate with SendGrid or Resend

	return c.JSON(common.Success("Email sent (mock)"))
}
