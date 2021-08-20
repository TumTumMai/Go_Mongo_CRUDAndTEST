package controller

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testgo/repository"
	"testgo/service"
	"testing"

	"github.com/labstack/echo/v4"

	"github.com/stretchr/testify/assert"
)

// var userbody = `{"name": "LnW","city": "MIng","age": 34}`
var getallres = `[{"Id":"1","name":"HadesGod","city":"Hell","age":22},{"Id":"2","name":"TitonGod","city":"Heaven","age":30}]`

func TestGetAll(t *testing.T) {
	//Setup Request
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	// req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	///Setup Dbmok
	userRepositoryMock := repository.NewUserRepositoryMock()
	userService := service.NewuserService(userRepositoryMock)
	h := NewuserHandler(userService)
	_ = h.GetAllUser(c)
	fmt.Println(rec.Body)
	res := rec.Body.String()
	res = strings.Replace(res, "\n", "", -1)

	fmt.Printf("res: %v\n", res)

	fmt.Println(res == getallres)
	// Assertions
	if assert.NoError(t, h.GetAllUser(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, getallres, res)

	}

}
