//go:build unit
// +build unit

package handler

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/labstack/echo"
	"github.com/lib/pq"
	"github.com/stretchr/testify/assert"
)

func TestUpdateExpenseHandler(t *testing.T) {

	t.Run("Pass-condition", func(t *testing.T) {

		//Data preparation
		mockRow := sqlmock.NewRows([]string{"id", "title", "amount", "Note", "tags"}).
			AddRow(1, "NMD ADIDAS", 5600, "Student discount", pq.Array([]string{"shoes", "adidas"}))
		// mockUpdateRow := sqlmock.NewRows([]string{"id", "title", "amount", "Note", "tags"}).
		// 	AddRow(1, "ULTRA BOOST ADIDAS", 6400, "Student discount", pq.Array([]string{"shoes", "adidas"}))
		jsonPayload := `{"title":"ULTRA BOOST ADIDAS","amount":6400,"note":"Student discount","tags":["shoes","adidas"]}`
		expected := `{"id":0,"title":"ULTRA BOOST ADIDAS","amount":6400,"note":"Student discount","tags":["shoes","adidas"]}`

		// Arrange
		e := echo.New()
		req := httptest.NewRequest(http.MethodPut, "/expenses/1", strings.NewReader(jsonPayload))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		db, mock, err := sqlmock.New()
		mock.ExpectPrepare("SELECT (.+) expenses").ExpectQuery().WithArgs().WillReturnRows(mockRow)
		mock.ExpectPrepare("UPDATE expenses").ExpectExec().WithArgs(1, "ULTRA BOOST ADIDAS", 6400, "Student discount", pq.Array([]string{"shoes", "adidas"})).WillReturnResult(sqlmock.NewResult(1, 1))

		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
		}
		h := handler{db}
		c := e.NewContext(req, rec)

		// Act
		err = h.UpdateExpenseHandler(c)

		// Assertions
		if assert.NoError(t, err) {
			assert.Equal(t, http.StatusOK, rec.Code)
			assert.Equal(t, expected, strings.TrimSpace(rec.Body.String()))
		}
	})
}
