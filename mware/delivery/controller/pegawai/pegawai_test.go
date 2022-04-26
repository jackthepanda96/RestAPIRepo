package pegawai

import (
	"apiex/mware/entity"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

var token string

func TestLogin(t *testing.T) {
	e := echo.New()
	requestBody, _ := json.Marshal(map[string]interface{}{
		"hp":       12345,
		"password": "jerry123",
	})
	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(string(requestBody)))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	res := httptest.NewRecorder()
	context := e.NewContext(req, res)
	context.SetPath("/login")

	controller := New(&mockRepoPegawai{}, validator.New())
	controller.Login(context)

	type ResponseStructure struct {
		Code    int
		Message string
		Status  bool
		Data    interface{}
	}

	var response ResponseStructure

	json.Unmarshal([]byte(res.Body.Bytes()), &response)
	log.Warn(response.Data)
	assert.True(t, response.Status)
	assert.NotNil(t, response.Data)
	data := response.Data.(map[string]interface{})
	log.Warn(data)
	token = data["Token"].(string)
}

func TestInsert(t *testing.T) {

}

func TestGetAllPegawai(t *testing.T) {
	t.Run("Get All Pegawai", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodPost, "/", nil)
		req.Header.Set(echo.HeaderAuthorization, "Bearer "+token)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/coba")

		pegawaiController := New(&mockRepoPegawai{}, validator.New())
		middleware.JWTWithConfig(middleware.JWTConfig{SigningMethod: "HS256", SigningKey: []byte("RH$SI4")})(pegawaiController.GetAllPegawai())(context)

		type response struct {
			Code    int
			Message string
			Status  bool
			Data    []entity.Pegawai
		}

		var resp response

		json.Unmarshal([]byte(res.Body.Bytes()), &resp)
		assert.Equal(t, 200, resp.Code)
		assert.Equal(t, resp.Data[0].Nama, "Jerry")

	})
}

type mockRepoPegawai struct{}

func (mrp *mockRepoPegawai) Insert(newPegawai entity.Pegawai) (entity.Pegawai, error) {
	return newPegawai, nil
}
func (mrp *mockRepoPegawai) GetAll() ([]entity.Pegawai, error) {
	return []entity.Pegawai{{Nama: "Jerry", HP: 12345}}, nil
}
func (mrp *mockRepoPegawai) Login(hp int, password string) (entity.Pegawai, error) {
	return entity.Pegawai{Model: gorm.Model{ID: uint(1)}, Nama: "Jerry", HP: 12345}, nil
}
