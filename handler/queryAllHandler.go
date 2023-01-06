package handler

import (
	"net/http"

	"github.com/Spiralzix/assessment/entity"
	"github.com/labstack/echo"
	"github.com/lib/pq"
)

func (h *handler) QueryAllExpenseHandler(c echo.Context) error {
	rows, err := h.DB.Query("SELECT * FROM expenses")
	if err != nil {
		return err
	}
	defer rows.Close()

	expenses := []entity.Expense{}
	for rows.Next() {
		u := entity.Expense{}
		err := rows.Scan(&u.ID, &u.Title, &u.Amount, &u.Note, pq.Array(&u.Tags))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, entity.Err{Message: "can't scan user:" + err.Error()})
		}
		expenses = append(expenses, u)
	}
	return c.JSON(http.StatusOK, expenses)
}
