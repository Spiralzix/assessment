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

func TestCreateExpenseHandler(t *testing.T) {

	t.Run("Success-condition, 201Created", func(t *testing.T) {
		//Data preparation
		jsonPayload := `{"title":"NMD ADIDAS","amount":5600,"note":"Student discount","tags":["shoes","adidas"]}`
		expected := `{"id":1,"title":"NMD ADIDAS","amount":5600,"note":"Student discount","tags":["shoes","adidas"]}`

		// Arrange
		e := echo.New()
		req := httptest.NewRequest(http.MethodPost, "/expenses", strings.NewReader(jsonPayload))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		db, mock, err := sqlmock.New()
		mockRows := sqlmock.NewRows([]string{"id"}).AddRow(1)
		mock.ExpectQuery(("INSERT INTO expenses")).WithArgs("NMD ADIDAS", 5600, "Student discount", pq.Array([]string{"shoes", "adidas"})).WillReturnRows(mockRows)
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
		}
		h := handler{db}
		c := e.NewContext(req, rec)

		// Act
		err = h.CreateExpenseHandler(c)

		// Assertions
		if assert.NoError(t, err) {
			assert.Equal(t, http.StatusCreated, rec.Code)
			assert.Equal(t, expected, strings.TrimSpace(rec.Body.String()))
		}
	})

	t.Run("Fail-condition, 400BadRequest", func(t *testing.T) {
		// Arrange
		e := echo.New()
		req := httptest.NewRequest(http.MethodPost, "/expenses", strings.NewReader(""))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		db, mock, err := sqlmock.New()
		mockRows := sqlmock.NewRows([]string{"id"}).AddRow(1)
		mock.ExpectQuery(("INSERT INTO expenses")).WithArgs("NMD ADIDAS", 5600, "Student discount", pq.Array([]string{"shoes", "adidas"})).WillReturnRows(mockRows)
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
		}
		h := handler{db}
		c := e.NewContext(req, rec)

		// Act
		err = h.CreateExpenseHandler(c)

		// Assertions
		if assert.NoError(t, err) {
			assert.Equal(t, http.StatusBadRequest, rec.Code)
		}
	})
}
