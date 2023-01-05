package handler

import (
	"net/http"

	"github.com/Spiralzix/assessment/entity"
	"github.com/labstack/echo"
	"github.com/lib/pq"
)

func (h *handler) CreateExpenseHandler(c echo.Context) error {
	e := entity.Expense{}
	err := c.Bind(&e)
	if err != nil {
		return c.JSON(http.StatusBadRequest, entity.Err{Message: err.Error()})
	}
	row := h.DB.QueryRow("INSERT INTO expenses (title, amount, note, tags) values ($1, $2, $3, $4)  RETURNING id", e.Title, e.Amount, e.Note, pq.Array(&e.Tags))
	err = row.Scan(&e.ID)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, entity.Err{Message: err.Error()})
	}
	return c.JSON(http.StatusCreated, e)
}
