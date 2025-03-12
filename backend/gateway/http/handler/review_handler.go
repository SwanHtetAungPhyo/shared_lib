package handlers

import "github.com/gofiber/fiber/v3"

type UnImplementedReviewHandler interface {
	CreateReview(c *fiber.Ctx) error
	ReadReview(c *fiber.Ctx) error
	UpdateReview(c *fiber.Ctx) error
}

type ReviewHandler struct {
	UnImplementedReviewHandler
}

func NewReviewHandler() *ReviewHandler {
	return &ReviewHandler{}
}
func (R *ReviewHandler) CreateReview(c *fiber.Ctx) error {
	return nil
}

func (R *ReviewHandler) ReadReview(c *fiber.Ctx) error {
	return nil
}
func (R *ReviewHandler) UpdateReview(c *fiber.Ctx) error {
	return nil
}
