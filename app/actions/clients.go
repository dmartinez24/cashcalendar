package actions

import (
	"cashcalendar/app/models"
	"fmt"
	"net/http"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop/v6"
	"github.com/gofrs/uuid"
)

type ClientsResource struct {
	buffalo.Resource
}

func (e ClientsResource) List(c buffalo.Context) error {
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return fmt.Errorf("no transaction found")
	}

	clients := models.Clients{}
	err := tx.Where("user_id = ?", uuid.FromStringOrNil("7221e463-c869-40ce-8de6-b789d8bff5e2")).All(&clients)
	if err != nil {
		return err
	}

	c.Set("clients", clients)
	c.Set("page", "clients")
	return c.Render(http.StatusOK, r.HTML("/home/index.plush.html"))
}
