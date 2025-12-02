package tag

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

func (h *Handler) List(c *fiber.Ctx) error {
	items, err := h.repo.List()
	if err != nil {
		return err
	}
	return c.JSON(common.Success(items))
}

func (h *Handler) Create(c *fiber.Ctx) error {
	var dto CreateUpdateTagDTO
	if err := c.BodyParser(&dto); err != nil {
		return fiber.ErrBadRequest
	}

	t := &Tag{
		Name: dto.Name,
	}

	if err := h.repo.Create(t); err != nil {
		return err
	}

	return c.JSON(common.Success(t))
}

func (h *Handler) Update(c *fiber.Ctx) error {
	id, err := common.ParseUint(c.Params("id"))
	if err != nil {
		return fiber.ErrBadRequest
	}

	var dto CreateUpdateTagDTO
	c.BodyParser(&dto)

	t := &Tag{
		ID:   id,
		Name: dto.Name,
	}

	if err := h.repo.Update(t); err != nil {
		return err
	}

	return c.JSON(common.Success(t))
}

func (h *Handler) Delete(c *fiber.Ctx) error {
	id, err := common.ParseUint(c.Params("id"))
	if err != nil {
		return fiber.ErrBadRequest
	}

	if err := h.repo.Delete(id); err != nil {
		return err
	}

	return c.JSON(common.Success(true))
}
