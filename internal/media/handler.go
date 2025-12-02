package media

import (
	"github.com/gofiber/fiber/v2"

	"github.com/PNamGP1120/ougreenplus-go/internal/common"
	"github.com/PNamGP1120/ougreenplus-go/internal/config"
)

type Handler struct {
	repo Repository
	cfg  *config.Config
}

func NewHandler(cfg *config.Config) *Handler {
	return &Handler{
		repo: NewRepository(),
		cfg:  cfg,
	}
}

func (h *Handler) Upload(c *fiber.Ctx) error {
	file, err := c.FormFile("file")
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "file is required")
	}

	url, err := UploadToS3(h.cfg, file)
	if err != nil {
		return fiber.ErrInternalServerError
	}

	userID := c.Locals("user_id").(uint)

	media := &Media{
		FileName:   file.Filename,
		FileSize:   file.Size,
		FileType:   file.Header.Get("Content-Type"),
		URL:        url,
		UploadedBy: userID,
	}

	if err := h.repo.Create(media); err != nil {
		return err
	}

	return c.JSON(common.Success(media))
}

func (h *Handler) List(c *fiber.Ctx) error {
	items, err := h.repo.List()
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
	m, err := h.repo.Get(id)
	if err != nil {
		return fiber.ErrNotFound
	}
	return c.JSON(common.Success(m))
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
