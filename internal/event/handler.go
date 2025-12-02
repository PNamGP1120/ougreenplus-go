package event

import (
	"strconv"

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
	status := Status(c.Query("status"))

	items, err := h.repo.List(status)
	if err != nil {
		return err
	}
	return c.JSON(common.Success(items))
}

func (h *Handler) Get(c *fiber.Ctx) error {
	id, err := common.ParseUint(c.Params("id"))
	if err != nil {
		return fiber.ErrBadRequest
	}

	e, err := h.repo.GetByID(id)
	if err != nil {
		return fiber.ErrNotFound
	}
	return c.JSON(common.Success(e))
}

func (h *Handler) Create(c *fiber.Ctx) error {
	var dto CreateUpdateEventDTO
	if err := c.BodyParser(&dto); err != nil {
		return fiber.ErrBadRequest
	}

	e := &Event{
		Title:       dto.Title,
		Description: dto.Description,
		StartDate:   dto.StartDate,
		EndDate:     dto.EndDate,
		Location:    dto.Location,
		PosterURL:   dto.PosterURL,
		Status:      dto.Status,
	}

	if err := h.repo.Create(e); err != nil {
		return err
	}

	return c.JSON(common.Success(e))
}

func (h *Handler) Update(c *fiber.Ctx) error {
	id, err := common.ParseUint(c.Params("id"))
	if err != nil {
		return fiber.ErrBadRequest
	}

	event, err := h.repo.GetByID(id)
	if err != nil {
		return fiber.ErrNotFound
	}

	var dto CreateUpdateEventDTO
	if err := c.BodyParser(&dto); err != nil {
		return fiber.ErrBadRequest
	}

	event.Title = dto.Title
	event.Description = dto.Description
	event.StartDate = dto.StartDate
	event.EndDate = dto.EndDate
	event.Location = dto.Location
	event.PosterURL = dto.PosterURL
	event.Status = dto.Status

	if err := h.repo.Update(event); err != nil {
		return err
	}

	return c.JSON(common.Success(event))
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

func (h *Handler) Register(c *fiber.Ctx) error {
	idStr := c.Params("id")
	eid, _ := strconv.ParseUint(idStr, 10, 64)

	var dto RegisterDTO
	if err := c.BodyParser(&dto); err != nil {
		return fiber.ErrBadRequest
	}

	reg := &Registration{
		EventID:   uint(eid),
		Name:      dto.Name,
		Email:     dto.Email,
		Phone:     dto.Phone,
		StudentID: dto.StudentID,
	}

	if err := h.repo.Register(reg); err != nil {
		return err
	}

	return c.JSON(common.Success(reg))
}

func (h *Handler) ListRegistrations(c *fiber.Ctx) error {
	id, err := common.ParseUint(c.Params("id"))
	if err != nil {
		return fiber.ErrBadRequest
	}

	list, err := h.repo.ListRegistrations(id)
	if err != nil {
		return err
	}

	return c.JSON(common.Success(list))
}
