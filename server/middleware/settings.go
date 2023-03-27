package middleware

import (
	"backend/twitch-bot/database"
	"backend/twitch-bot/models"

	"github.com/gofiber/fiber/v2"
)

// Middleware that gets the current Song Queue Settings
func GetSongQueueSettings(c * fiber.Ctx) error {
	// Creates a Query Struct for the query parameters the GET request will take in
	type Query struct {
		Channel string `query:"channel"`
	}
	query := new(Query)

	if err := c.QueryParser(query); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"error": err,
		})
	}
	// Checks if query is empty if it is then return back to the request that the query is missing
	if query.Channel == "" {
		clientData := models.ClientData{
			Status:  "fail",
			Message: "missing channel",
			Data:    nil,
		}
		return c.Status(fiber.StatusBadRequest).JSON(map[string]interface{}{
			"error": clientData.Message,
		})
	}

	db, err := database.InitializeSettingsDBConnection()
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(map[string]interface{}{
			"error": err.Error(),
		})
	}

	settings, db, err  := database.GetSongQueueSettings(query.Channel, db)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(map[string]interface{}{
			"error": err.Error(),
		})
	}
	db.Close()
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"settings": settings,
	})

}
