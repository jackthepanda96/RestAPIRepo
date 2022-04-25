package controller

import (
	"apiex/testing/entity"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"github.com/stretchr/testify/assert"
)

func TestGetAllPegawai(t *testing.T) {
	t.Run("Success Get All", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/users")

		pegawaiController := New(&mockUserRepository{}, validator.New())
		pegawaiController.GetAllPegawai(context)

		type response struct {
			Code    int
			Message string
			Status  bool
			Data    []entity.Pegawai
		}

		var resp response

		json.Unmarshal([]byte(res.Body.Bytes()), &resp)
		assert.Equal(t, resp.Data[0].Nama, "Jerry")
	})
	t.Run("Error Get All", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/users")

		pegawaiController := New(&erorrMockUserRepository{}, validator.New())
		pegawaiController.GetAllPegawai(context)

		type response struct {
			Code    int
			Message string
			Status  bool
			Data    []entity.Pegawai
		}

		var resp response

		json.Unmarshal([]byte(res.Body.Bytes()), &resp)
		assert.Nil(t, resp.Data)
		assert.False(t, resp.Status)
		assert.Equal(t, 500, resp.Code)
	})
}

func TestInsert(t *testing.T) {
	t.Run("Success Insert", func(t *testing.T) {
		e := echo.New()
		requestBody, _ := json.Marshal(map[string]interface{}{
			"nama": "Galang",
			"hp":   1234,
			"gaji": 10000,
		})
		req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(string(requestBody)))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON) // Set Content to JSON
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/pegawai")

		pegawaiController := New(&mockUserRepository{}, validator.New())
		pegawaiController.Insert(context)

		type response struct {
			Code    int
			Message string
			Status  bool
			Data    interface{}
		}

		var resp response

		json.Unmarshal([]byte(res.Body.Bytes()), &resp)
		assert.Equal(t, "Galang", resp.Data.(map[string]interface{})["Nama"])
		assert.True(t, resp.Status)
		assert.Equal(t, 201, resp.Code)
	})
	t.Run("Error Validasi", func(t *testing.T) {
		e := echo.New()
		requestBody, _ := json.Marshal(map[string]interface{}{
			"hp": 12345,
		})

		req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(string(requestBody)))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/pegawai")

		pegawaiController := New(&erorrMockUserRepository{}, validator.New())
		pegawaiController.Insert(context)

		type response struct {
			Code    int
			Message string
			Status  bool
			Data    interface{}
		}

		var resp response

		json.Unmarshal([]byte(res.Body.Bytes()), &resp)
		log.Warn(resp)
		assert.False(t, resp.Status)
		assert.Nil(t, resp.Data)
		assert.Equal(t, 400, resp.Code)
	})
	t.Run("Error Bind", func(t *testing.T) {
		e := echo.New()
		requestBody, _ := json.Marshal(map[string]interface{}{
			"hp": 12345,
		})

		req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(string(requestBody)))
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/pegawai")

		pegawaiController := New(&erorrMockUserRepository{}, validator.New())
		pegawaiController.Insert(context)

		type response struct {
			Code    int
			Message string
			Status  bool
			Data    interface{}
		}

		var resp response

		json.Unmarshal([]byte(res.Body.Bytes()), &resp)
		log.Warn(resp)
		assert.False(t, resp.Status)
		assert.Nil(t, resp.Data)
		assert.Equal(t, 400, resp.Code)
	})
	t.Run("Error Insert DB", func(t *testing.T) {
		e := echo.New()
		requestBody, _ := json.Marshal(map[string]interface{}{
			"nama": "Galang",
			"hp":   12345,
		})

		req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(string(requestBody)))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/pegawai")

		pegawaiController := New(&erorrMockUserRepository{}, validator.New())
		pegawaiController.Insert(context)

		type response struct {
			Code    int
			Message string
			Status  bool
			Data    interface{}
		}

		var resp response

		json.Unmarshal([]byte(res.Body.Bytes()), &resp)
		assert.False(t, resp.Status)
		assert.Nil(t, resp.Data)
		assert.Equal(t, 500, resp.Code)
	})
}

// Dummy Data

type mockUserRepository struct{}

func (mur *mockUserRepository) Insert(newPegawai entity.Pegawai) (entity.Pegawai, error) {
	return newPegawai, nil
}
func (mur *mockUserRepository) GetAll() ([]entity.Pegawai, error) {
	return []entity.Pegawai{{Nama: "Jerry", HP: 12345}}, nil
}

type erorrMockUserRepository struct{}

func (emur *erorrMockUserRepository) Insert(newPegawai entity.Pegawai) (entity.Pegawai, error) {
	return entity.Pegawai{}, errors.New("tidak bisa insert data")
}
func (emur *erorrMockUserRepository) GetAll() ([]entity.Pegawai, error) {
	return nil, errors.New("tidak bisa select data")
}
