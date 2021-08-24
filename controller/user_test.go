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

var userinserttest = `{"Id":"3","name":"aaa","city":"Hell","age":22}`

var userinputtest = `{"Id":"1","name":"Obelisk","city":"valhalla","age":1000}`

var userid = "1"
var deleteres = "Delsusecfull :" + userid

func TestDel(t *testing.T) {
	//Setup Request
	e := echo.New()
	req := httptest.NewRequest(http.MethodDelete, "/", strings.NewReader(userid))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	c.SetPath("/:id")
	c.SetParamNames("id")
	c.SetParamValues("1")
	///Setup Dbmok
	userRepositoryMock := repository.NewUserRepositoryMock()
	userService := service.NewuserService(userRepositoryMock)
	h := NewuserHandler(userService)
	_ = h.DeleteById(c)
	fmt.Println(rec.Body)
	res := rec.Body.String()
	res = strings.Replace(res, "\n", "", -1)

	// fmt.Printf("res: %v\n", res)

	// fmt.Println(res == usertest)
	// Assertions
	if assert.NoError(t, h.DeleteById(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, deleteres, res)

	}

}

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

func TestInsert(t *testing.T) {
	//Setup Request
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(userinserttest))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	///Setup Dbmok
	userRepositoryMock := repository.NewUserRepositoryMock()
	userService := service.NewuserService(userRepositoryMock)
	h := NewuserHandler(userService)
	_ = h.Insert(c)
	fmt.Println(rec.Body)
	res := rec.Body.String()
	res = strings.Replace(res, "\n", "", -1)

	// fmt.Printf("res: %v\n", res)

	// fmt.Println(res == usertest)
	// Assertions
	if assert.NoError(t, h.Insert(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, userinserttest, res)

	}

}

func TestInput(t *testing.T) {
	//Setup Request
	e := echo.New()
	req := httptest.NewRequest(http.MethodPut, "/", strings.NewReader(userinputtest))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	c.SetPath("/:id")
	c.SetParamNames("id")
	c.SetParamValues("1")
	///Setup Dbmok
	userRepositoryMock := repository.NewUserRepositoryMock()
	userService := service.NewuserService(userRepositoryMock)
	h := NewuserHandler(userService)
	_ = h.UpdateOne(c)
	fmt.Println(rec.Body)
	res := rec.Body.String()
	res = strings.Replace(res, "\n", "", -1)

	// fmt.Printf("res: %v\n", res)

	// fmt.Println(res == usertest)
	// Assertions
	if assert.NoError(t, h.UpdateOne(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, userinputtest, res)

	}

}
