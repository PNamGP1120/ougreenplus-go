package category

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type Handler struct {
	repo Repository
}

func NewHandler() *Handler {
	return &Handler{
		repo: NewRepository(),
	}
}

func (h *Handler) List(c *fiber.Ctx) error {
	items, err := h.repo.List()
	if err != nil {
		return err
	}
	return c.JSON(items)
}

func (h *Handler) Create(c *fiber.Ctx) error {
	var dto CreateUpdateCategoryDTO
	if err := c.BodyParser(&dto); err != nil {
		return fiber.ErrBadRequest
	}

	if dto.Name == "" {
		return fiber.NewError(fiber.StatusBadRequest, "name is required")
	}

	cat := &Category{
		Name:        dto.Name,
		Description: dto.Description,
	}

	if err := h.repo.Create(cat); err != nil {
		return err
	}
	return c.Status(fiber.StatusCreated).JSON(cat)
}

func (h *Handler) Update(c *fiber.Ctx) error {
	idStr := c.Params("id")
	id64, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		return fiber.ErrBadRequest
	}
	id := uint(id64)

	var dto CreateUpdateCategoryDTO
	if err := c.BodyParser(&dto); err != nil {
		return fiber.ErrBadRequest
	}

	cat := &Category{
		ID:          id,
		Name:        dto.Name,
		Description: dto.Description,
	}

	if err := h.repo.Update(cat); err != nil {
		return err
	}
	return c.JSON(cat)
}

func (h *Handler) Delete(c *fiber.Ctx) error {
	idStr := c.Params("id")
	id64, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		return fiber.ErrBadRequest
	}
	id := uint(id64)

	if err := h.repo.Delete(id); err != nil {
		return err
	}
	return c.JSON(fiber.Map{"success": true})
}

func (h *Handler) ListArticles(c *fiber.Ctx) error {
	idStr := c.Params("id")
	id64, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		return fiber.ErrBadRequest
	}
	categoryID := uint(id64)

	page, _ := strconv.Atoi(c.Query("page", "1"))
	size, _ := strconv.Atoi(c.Query("size", "10"))

	items, total, err := h.repo.ListArticles(categoryID, page, size)
	if err != nil {
		return err
	}

	return c.JSON(fiber.Map{
		"data":  items,
		"page":  page,
		"size":  size,
		"total": total,
	})
}
