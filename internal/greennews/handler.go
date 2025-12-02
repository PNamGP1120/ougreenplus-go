package greennews

import (
	"github.com/PNamGP1120/ougreenplus-go/internal/common"
	"github.com/gofiber/fiber/v2"
)

type Handler struct {
	repo Repository
}

func NewHandler() *Handler {
	return &Handler{repo: NewRepository()}
}

func (h *Handler) List(c *fiber.Ctx) error {
	month := c.QueryInt("month")
	year := c.QueryInt("year")

	items, err := h.repo.List(month, year)
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

	g, err := h.repo.GetByID(id)
	if err != nil {
		return fiber.ErrNotFound
	}

	return c.JSON(common.Success(g))
}

func (h *Handler) Create(c *fiber.Ctx) error {
	var dto CreateUpdateGreennewsDTO
	if err := c.BodyParser(&dto); err != nil {
		return fiber.ErrBadRequest
	}

	g := &Greennews{
		Number: dto.Number,
		Month:  dto.Month,
		Year:   dto.Year,
	}

	if err := h.repo.Create(g); err != nil {
		return err
	}

	return c.JSON(common.Success(g))
}

func (h *Handler) Update(c *fiber.Ctx) error {
	id, err := common.ParseUint(c.Params("id"))
	if err != nil {
		return fiber.ErrBadRequest
	}

	g, err := h.repo.GetByID(id)
	if err != nil {
		return fiber.ErrNotFound
	}

	var dto CreateUpdateGreennewsDTO
	if err := c.BodyParser(&dto); err != nil {
		return fiber.ErrBadRequest
	}

	g.Number = dto.Number
	g.Month = dto.Month
	g.Year = dto.Year

	if err := h.repo.Update(g); err != nil {
		return err
	}

	return c.JSON(common.Success(g))
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
