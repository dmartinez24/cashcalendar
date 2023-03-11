package actions

import (
	"cashcalendar/app/models"
	"fmt"
	"net/http"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop/v5"
)

func Create(c buffalo.Context) error {
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return fmt.Errorf("no transaction found")
	}

	event := models.Event{}
	if err := c.Bind(&event); err != nil {
		return err
	}

	verrs := event.Validate()
	if verrs.HasAny() {
		c.Set("errors", verrs)
		c.Set("page", "calendar")
		return c.Render(http.StatusUnprocessableEntity, r.HTML("home/index.plush.html"))
	}

	if err := tx.Create(&event); err != nil {
		return err
	}

	c.Set("page", "calendar")
	return c.Render(http.StatusUnprocessableEntity, r.HTML("home/index.plush.html"))
}
