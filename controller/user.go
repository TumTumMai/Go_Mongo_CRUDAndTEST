package controller

import (
	"fmt"
	"net/http"
	"testgo/service"

	"github.com/labstack/echo/v4"
)

type userHandler struct {
	userService service.UserService
}

func NewuserHandler(cusService service.UserService) userHandler {
	return userHandler{
		userService: cusService,
	}
}

func (h userHandler) GetAllUser(c echo.Context) error {
	users, err := h.userService.GetUsers()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, users)
}

func (h userHandler) GetUserById(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		return c.String(http.StatusBadRequest, "Invalid Id")
	}
	user, err := h.userService.GetUser(id)
	if err != nil {
		if err.Error() == "not found" {
			return c.JSON(http.StatusNotFound, map[string]string{"message": fmt.Sprintf("Not Found id %s", id)})
		}
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, user)

}

func (h userHandler) Insert(c echo.Context) error {
	// users, err := h.userService.Insert(users)
	testuser := service.UserResponse{}
	if err := c.Bind(&testuser); err != nil {
		return c.NoContent(http.StatusBadRequest)
	}
	users, err := h.userService.Insert(testuser)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, users)
}
/////////////////////////////////
