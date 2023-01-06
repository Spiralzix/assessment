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

func TestGetExpenseHandler(t *testing.T) {

	t.Run("Pass-condition", func(t *testing.T) {

		//Data preparation
		mockRows := sqlmock.NewRows([]string{"id", "title", "amount", "note", "tags"}).
			AddRow(1, "NMD ADIDAS", 5600, "Student discount", pq.Array([]string{"shoes", "adidas"}))
		expected := `{"id":1,"title":"NMD ADIDAS","amount":5600,"note":"Student discount","tags":["shoes","adidas"]}`

		// Arrange
		e := echo.New()
		// var id int = 1
		req := httptest.NewRequest(http.MethodGet, "/expense/1", strings.NewReader(""))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		db, mock, err := sqlmock.New()
		mock.ExpectQuery("SELECT (.+) FROM expenses").WithArgs().WillReturnRows(mockRows)
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
		}
		h := handler{db}
		c := e.NewContext(req, rec)

		// Act
		err = h.QueryExpenseHandler(c)

		// Assertions
		if assert.NoError(t, err) {
			assert.Equal(t, http.StatusOK, rec.Code)
			assert.Equal(t, expected, strings.TrimSpace(rec.Body.String()))
		}
	})
}
