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

func TestGetAllExpenseHandler(t *testing.T) {

	t.Run("Pass-condition", func(t *testing.T) {

		//Data preparation
		mockRows := sqlmock.NewRows([]string{"id", "title", "amount", "note", "tags"}).
			AddRow(1, "NMD ADIDAS", 5600, "Student discount", pq.Array([]string{"shoes", "adidas"})).
			AddRow(2, "strawberry smoothie", 79, "night market promotion discount 10 bath", pq.Array([]string{"food", "beverage"}))
		expected := `[{"id":1,"title":"NMD ADIDAS","amount":5600,"note":"Student discount","tags":["shoes","adidas"]},{"id":2,"title":"strawberry smoothie","amount":79,"note":"night market promotion discount 10 bath","tags":["food","beverage"]}]`

		// Arrange
		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/expenses", strings.NewReader(""))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		db, mock, err := sqlmock.New()
		mock.ExpectQuery("SELECT (.+) FROM expenses").WillReturnRows(mockRows)
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
		}
		h := handler{db}
		c := e.NewContext(req, rec)

		// Act
		err = h.QueryAllExpenseHandler(c)

		// Assertions
		if assert.NoError(t, err) {
			assert.Equal(t, http.StatusOK, rec.Code)
			assert.Equal(t, expected, strings.TrimSpace(rec.Body.String()))
		}
	})
}
