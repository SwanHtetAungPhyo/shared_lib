package handlers

import "github.com/gofiber/fiber/v3"

type ServiceBlockHandler interface {
	CreateServiceOnPlatform(c fiber.Ctx) error
	UpdateServiceOnPlatform(c fiber.Ctx) error
	DeleteServiceOnPlatform(c fiber.Ctx) error
	ListServiceBlocks(c fiber.Ctx) error
}

// BlocksHandler TODO: For digital asset
type BlocksHandler struct {
	Service ServiceBlockHandler
}

func NewBlocksHandler() *BlocksHandler {
	return &BlocksHandler{}
}
func (h *BlocksHandler) CreateServiceOnPlatform(c fiber.Ctx) error {
	return nil
}

func (h *BlocksHandler) UpdateServiceOnPlatform(c fiber.Ctx) error {
	return nil
}

func (h *BlocksHandler) DeleteServiceOnPlatform(c fiber.Ctx) error {
	return nil
}

func (h *BlocksHandler) ListServiceBlocks(c fiber.Ctx) error {
	return nil
}
