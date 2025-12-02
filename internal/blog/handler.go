package blog

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

func parseUintParam(c *fiber.Ctx, name string) (uint, error) {
	idStr := c.Params(name)
	id64, err := strconv.ParseUint(idStr, 10, 64)
	return uint(id64), err
}

func (h *Handler) List(c *fiber.Ctx) error {
	page, _ := strconv.Atoi(c.Query("page", "1"))
	size, _ := strconv.Atoi(c.Query("size", "10"))

	items, total, err := h.repo.List(page, size)
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

func (h *Handler) Get(c *fiber.Ctx) error {
	id, err := parseUintParam(c, "id")
	if err != nil {
		return fiber.ErrBadRequest
	}

	b, err := h.repo.GetByID(id)
	if err != nil {
		return fiber.ErrNotFound
	}
	return c.JSON(b)
}

func (h *Handler) Create(c *fiber.Ctx) error {
	var dto CreateUpdateBlogDTO
	if err := c.BodyParser(&dto); err != nil {
		return fiber.ErrBadRequest
	}

	if dto.Title == "" {
		return fiber.NewError(fiber.StatusBadRequest, "title is required")
	}

	b := &Blog{
		Title:     dto.Title,
		Summary:   dto.Summary,
		Content:   dto.Content,
		Thumbnail: dto.Thumbnail,
	}

	if err := h.repo.Create(b); err != nil {
		return err
	}

	return c.Status(fiber.StatusCreated).JSON(b)
}

func (h *Handler) Update(c *fiber.Ctx) error {
	id, err := parseUintParam(c, "id")
	if err != nil {
		return fiber.ErrBadRequest
	}

	b, err := h.repo.GetByID(id)
	if err != nil {
		return fiber.ErrNotFound
	}

	var dto CreateUpdateBlogDTO
	if err := c.BodyParser(&dto); err != nil {
		return fiber.ErrBadRequest
	}

	b.Title = dto.Title
	b.Summary = dto.Summary
	b.Content = dto.Content
	b.Thumbnail = dto.Thumbnail

	if err := h.repo.Update(b); err != nil {
		return err
	}

	return c.JSON(b)
}

func (h *Handler) Delete(c *fiber.Ctx) error {
	id, err := parseUintParam(c, "id")
	if err != nil {
		return fiber.ErrBadRequest
	}

	if err := h.repo.Delete(id); err != nil {
		return err
	}
	return c.JSON(fiber.Map{"success": true})
}
