package handlers

import (
	"context"

	"github.com/gofiber/fiber/v3"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/teesh3rt/wizardapi/internal/database"
)

type WizardHandler struct {
	Queries *database.Queries
}

func (h *WizardHandler) GetWizard(ctx fiber.Ctx) error {
	idStr := ctx.Params("id")
	var id pgtype.UUID
	id.Scan(idStr)
	wizard, err := h.Queries.GetWizard(context.Background(), id)
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"success": false,
			"reason":  "Wizard not found",
		})
	}
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"id":      wizard.ID,
		"name":    wizard.Name,
		"level":   wizard.Level,
		"bio":     wizard.Bio,
	})
}

func (h *WizardHandler) GetAllWizards(ctx fiber.Ctx) error {
	wizards, err := h.Queries.GetAllWizards(context.Background())
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"reason":  "Failed to fetch wizards",
		})
	}
	// Convert to non-Go case (snake_case) for each wizard
	var result []fiber.Map
	for _, wizard := range wizards {
		result = append(result, fiber.Map{
			"id":    wizard.ID,
			"name":  wizard.Name,
			"level": wizard.Level,
			"bio":   wizard.Bio,
		})
	}
	return ctx.Status(fiber.StatusOK).JSON(result)
}

func (h *WizardHandler) DeleteWizard(ctx fiber.Ctx) error {
	idStr := ctx.Params("id")
	var id pgtype.UUID
	if err := id.Scan(idStr); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"reason":  "ID is not a valid UUID",
		})
	}
	err := h.Queries.DeleteWizard(context.Background(), id)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"reason":  "Failed to delete wizard",
		})
	}
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"success": true})
}

func (h *WizardHandler) CreateWizard(ctx fiber.Ctx) error {
	type reqBody struct {
		Name string `json:"name"`
		Bio  string `json:"bio"`
	}
	var body reqBody
	if err := ctx.Bind().Body(&body); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"reason":  "Invalid request body",
		})
	}
	// Optionally, you may want to return the created wizard. For now, just return success.
	err := h.Queries.CreateWizard(context.Background(), database.CreateWizardParams{
		Name: body.Name,
		Bio:  body.Bio,
	})
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"reason":  "Failed to create wizard",
		})
	}
	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{"success": true})
}
