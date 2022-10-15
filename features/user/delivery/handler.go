package delivery

import (
	"net/http"
	"strconv"

	"github.com/ALTA-BE12-KhalidRianda/Deployment/features/user/domain"
	"github.com/labstack/echo/v4"
)

type userHandler struct {
	srv domain.Service
}

func New(e *echo.Echo, srv domain.Service) {
	handler := userHandler{srv: srv}
	e.GET("/users", handler.ShowAllUser())
	e.POST("/users", handler.AddUser())
	e.POST("/users/update", handler.UpdateUser())
	e.GET("/users/id", handler.ShowUserByID())
	e.GET("/id/:id", handler.ShowUserByID())
	e.DELETE("/users", handler.DeleteByID())
	e.DELETE("/users/:id", handler.DeleteByID())

	// e.POST("/users/login", handler.LoginUser())
}

func (us *userHandler) AddUser() echo.HandlerFunc {
	return func(c echo.Context) error {
		var input RegisterFormat
		if err := c.Bind(&input); err != nil {
			return c.JSON(http.StatusBadRequest, FailResponse("cannot bind input"))
		}
		cnv := ToDomain(input)
		res, err := us.srv.AddUser(cnv)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, FailResponse(err.Error()))
		}

		return c.JSON(http.StatusCreated, SuccessResponse("berhasil register", ToResponse(res, "reg")))
	}

}

func (us *userHandler) UpdateUser() echo.HandlerFunc {
	return func(c echo.Context) error {
		var input UpdateFormat
		if err := c.Bind(&input); err != nil {
			return c.JSON(http.StatusBadRequest, FailResponse("cannot bind input"))
		}
		cnv := ToDomain(input)
		res, err := us.srv.UpdateProfile(cnv)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, FailResponse(err.Error()))
		}

		return c.JSON(http.StatusCreated, SuccessResponse("berhasil update", ToResponse(res, "reg")))
	}

}

func (us *userHandler) ShowAllUser() echo.HandlerFunc {
	return func(c echo.Context) error {
		res, err := us.srv.ShowAllUser()
		if err != nil {
			return c.JSON(http.StatusInternalServerError, FailResponse(err.Error()))
		}

		return c.JSON(http.StatusOK, SuccessResponse("success get all user", ToResponse(res, "all")))
	}
}

func (us *userHandler) ShowUserByID() echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.QueryParam("id")
		u64, err := strconv.ParseUint(id, 10, 32)
		if err != nil {
			return c.JSON(http.StatusBadRequest, FailResponse(err.Error()))
		}
		wd := uint(u64)

		res, err := us.srv.Profile(uint(wd))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, FailResponse(err.Error()))
		}

		return c.JSON(http.StatusOK, SuccessResponse("success get user", ToResponse(res, "reg")))
	}
}

func (us *userHandler) DeleteByID() echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.QueryParam("id")
		u64, err := strconv.ParseUint(id, 10, 32)
		if err != nil {
			return c.JSON(http.StatusBadRequest, FailResponse(err.Error()))
		}
		wd := uint(u64)

		err = us.srv.Delete(uint(wd))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, FailResponse(err.Error()))
		}

		return c.JSON(http.StatusOK, SuccessDelete("success delete user"))
	}
}

// func (us *userHandler) LoginUser() echo.HandlerFunc {
// 	//autentikasi user login
// 	return func(c echo.Context) error {
// 		var resQry LoginFormat
// 		if err := c.Bind(&resQry); err != nil {
// 			return c.JSON(http.StatusBadRequest, FailResponse("cannot bind input"))
// 		}

// 		cnv := ToDomain(resQry)
// 		res, err := us.srv.LoginUser(cnv)
// 		if err != nil {
// 			return c.JSON(http.StatusInternalServerError, FailResponse(err.Error()))
// 		}
// 		token := us.srv.GenerateToken(res.ID)
// 		return c.JSON(http.StatusCreated, SuccessLogin("berhasil register", token, ToResponse(res, "reg")))
// 	}
// }
