package handler

import (
	"net/http"
	"strconv"

	"github.com/Spiralzix/assessment/entity"
	"github.com/labstack/echo"
	"github.com/lib/pq"
)

func (h *handler) UpdateExpenseHandler(c echo.Context) error {
	id := c.Param("id")
	expenseId, _ := strconv.Atoi(id)
	entityPayload := entity.Expense{ID: expenseId}
	err := c.Bind(&entityPayload)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, entity.Err{Message: "can't bind data" + err.Error()})
	}

	stmt, err := h.DB.Prepare("SELECT id, title, amount, note, tags FROM expenses WHERE id = $1")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, entity.Err{Message: "can't prepare query expense history:" + err.Error()})
	}
	row := stmt.QueryRow(id)
	dbExpense := entity.Expense{}
	err = row.Scan(&dbExpense.ID, &dbExpense.Title, &dbExpense.Amount, &dbExpense.Note, pq.Array(&dbExpense.Tags))
	if err != nil {
		return c.JSON(http.StatusNotFound, entity.Err{Message: "expense not found"})
	}

	stmt, err = h.DB.Prepare("UPDATE expenses SET title=$2, amount=$3, note=$4, tags=$5 WHERE id=$1;")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, entity.Err{Message: "can't prepare update expense history:" + err.Error()})
	}
	if _, err := stmt.Exec(dbExpense.ID, entityPayload.Title, entityPayload.Amount, entityPayload.Note, pq.Array(&entityPayload.Tags)); err != nil {
		return c.JSON(http.StatusInternalServerError, entity.Err{Message: err.Error()})
	}
	return c.JSON(http.StatusOK, entityPayload)
}
