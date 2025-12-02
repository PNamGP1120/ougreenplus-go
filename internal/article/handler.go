package article

import (
	"strconv"
	"time"

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

	var catID uint
	if catStr := c.Query("category"); catStr != "" {
		if v, err := strconv.ParseUint(catStr, 10, 64); err == nil {
			catID = uint(v)
		}
	}

	var status Status
	if s := c.Query("status"); s != "" {
		status = Status(s)
	}

	items, total, err := h.repo.List(page, size, catID, status)
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

	a, err := h.repo.GetByID(id)
	if err != nil {
		return fiber.ErrNotFound
	}
	return c.JSON(a)
}

func (h *Handler) Create(c *fiber.Ctx) error {
	var dto CreateUpdateArticleDTO
	if err := c.BodyParser(&dto); err != nil {
		return fiber.ErrBadRequest
	}

	if dto.Title == "" || dto.CategoryID == 0 {
		return fiber.NewError(fiber.StatusBadRequest, "title and category_id are required")
	}

	a := &Article{
		Title:      dto.Title,
		Summary:    dto.Summary,
		Content:    dto.Content,
		Thumbnail:  dto.Thumbnail,
		CategoryID: dto.CategoryID,
		Type:       dto.Type,
		Status:     dto.Status,
	}

	if dto.Status == StatusPub {
		now := time.Now()
		a.PublishedAt = &now
	}

	if err := h.repo.Create(a); err != nil {
		return err
	}

	return c.Status(fiber.StatusCreated).JSON(a)
}

func (h *Handler) Update(c *fiber.Ctx) error {
	id, err := parseUintParam(c, "id")
	if err != nil {
		return fiber.ErrBadRequest
	}

	a, err := h.repo.GetByID(id)
	if err != nil {
		return fiber.ErrNotFound
	}

	var dto CreateUpdateArticleDTO
	if err := c.BodyParser(&dto); err != nil {
		return fiber.ErrBadRequest
	}

	a.Title = dto.Title
	a.Summary = dto.Summary
	a.Content = dto.Content
	a.Thumbnail = dto.Thumbnail
	a.CategoryID = dto.CategoryID
	a.Type = dto.Type
	a.Status = dto.Status

	if dto.Status == StatusPub && a.PublishedAt == nil {
		now := time.Now()
		a.PublishedAt = &now
	}

	if err := h.repo.Update(a); err != nil {
		return err
	}
	return c.JSON(a)
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

func (h *Handler) Related(c *fiber.Ctx) error {
	id, err := parseUintParam(c, "id")
	if err != nil {
		return fiber.ErrBadRequest
	}

	a, err := h.repo.GetByID(id)
	if err != nil {
		return fiber.ErrNotFound
	}

	items, err := h.repo.ListRelated(a.ID, a.CategoryID, 5)
	if err != nil {
		return err
	}

	return c.JSON(items)
}
