package routes

import (
	"memnixrest/handlers"

	"github.com/gofiber/fiber/v2"
)

// New
func New() *fiber.App {
	// Create new app
	app := fiber.New()

	// Api group
	api := app.Group("/api")

	// v1 group "/api/v1"
	v1 := api.Group("/v1", func(c *fiber.Ctx) error {
		return c.Next()
	})

	api.Get("/", func(c *fiber.Ctx) error {
		return fiber.NewError(fiber.StatusForbidden, "This is not a valid route") // Custom error
	})

	v1.Get("/", func(c *fiber.Ctx) error {
		return fiber.NewError(fiber.StatusForbidden, "This is not a valid route") // Custom error
	})

	// Users
	v1.Get("/users", handlers.GetAllUsers)          // Get all users
	v1.Get("/user/id/:id", handlers.GetUserByID)    // Get user by id
	v1.Post("/user/new", handlers.CreateNewUser)    // Create a new user
	v1.Put("/user/id/:id", handlers.UpdateUserByID) // Update an user using his id

	// Identifiers
	v1.Get("/identifiers", handlers.GetAllIdentifiers)                            // Get all identifiers
	v1.Get("/identifier/id/:id", handlers.GetIdentifierByID)                      // Get identifier by id
	v1.Get("/identifier/userid/:userID", handlers.GetIdentifierByUserID)          // Get identifier by user_id
	v1.Get("/identifier/discordid/:discordID", handlers.GetIdentifierByDiscordID) // Get identifier by discord_id
	v1.Post("/identifier/new", handlers.CreateNewIdentifier)                      // Create a new identifier

	// Decks
	v1.Get("/decks", handlers.GetAllDecks)
	v1.Get("/deck/id/:id", handlers.GetDeckByID)
	v1.Post("/deck/new", handlers.CreateNewDeck)

	// Cards
	v1.Get("/cards", handlers.GetAllCards)
	v1.Get("/card/id/:id", handlers.GetCardByID)
	v1.Get("/card/deck/:deckID", handlers.GetCardsFromDeck)
	v1.Post("/card/new", handlers.CreateNewCard)

	// Revision
	v1.Get("/revisions", handlers.GetAllRevisions)
	v1.Get("/revision/id/:id", handlers.GetRevisionByID)
	v1.Get("/revision/userid/:userID", handlers.GetRevisionByUserID)
	v1.Get("/revision/cardid/cardID", handlers.GetRevisionByCardID)
	v1.Post("/revision/new", handlers.CreateNewRevision)

	// Access
	// TODO

	// History
	// TODO

	return app

}
