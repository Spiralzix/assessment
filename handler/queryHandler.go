package handler

import (
	"database/sql"
	"net/http"

	"github.com/Spiralzix/assessment/entity"
	"github.com/labstack/echo"
	"github.com/lib/pq"
)

func (h *handler) QueryExpenseHandler(c echo.Context) error {
	id := c.Param("id")
	e := entity.Expense{}
	row := h.DB.QueryRow("SELECT id, title, amount, note, tags FROM expenses WHERE id = $1", id)
	err := row.Scan(&e.ID, &e.Title, &e.Amount, &e.Note, pq.Array(&e.Tags))
	switch err {
	case sql.ErrNoRows:
		return c.JSON(http.StatusNotFound, entity.Err{Message: "expense not found"})
	case nil:
		return c.JSON(http.StatusOK, e)
	default:
		return c.JSON(http.StatusInternalServerError, entity.Err{Message: "can't scan expense:"})
	}
}
