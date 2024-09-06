package views

import (
	"app/repository"
	"app/service"
	"app/templates/components"
	"fmt"

	"github.com/labstack/echo/v4"
)

func Routes(e *echo.Echo) {
	e.GET("/", func(c echo.Context) error {
		components.Layout(Index()).Render(c.Request().Context(), c.Response())
		return nil
	})
	e.POST("/ping", Ping)

	e.GET("/search", func(c echo.Context) error {
		components.Layout(CardSearch()).Render(c.Request().Context(), c.Response())
		return nil
	})

	e.POST("/search", ShowSearchResults)
}

func Ping(c echo.Context) error {
	components.Ping().Render(c.Request().Context(), c.Response())
	return nil
}

func ShowSearchResults(c echo.Context) error {
	cardSearchService := service.CardSearchService{
		CardRepo: repository.NewCardRepository(),
	}

	cardName := c.FormValue("card-name")
	fmt.Println(cardName)

	cards := cardSearchService.GetCardByName(cardName)
	fmt.Println(len(cards))

	components.SearchResults(cards).Render(c.Request().Context(), c.Response())

	return nil
}
