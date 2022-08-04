package controllers

import (
	"github.com/dakasakti/postingan/deliveries/helpers"
	"github.com/dakasakti/postingan/deliveries/middlewares"
	"github.com/dakasakti/postingan/entities"
	ts "github.com/dakasakti/postingan/services/posttype"
	"github.com/dakasakti/postingan/services/validation"

	"github.com/labstack/echo/v4"
)

type postTypeController struct {
	Ts ts.PostTypeService
	Vs validation.Validation
}

func NewPostTypeController(ts ts.PostTypeService, vs validation.Validation) *postTypeController {
	return &postTypeController{Ts: ts, Vs: vs}
}

func (tc *postTypeController) Register(ctx echo.Context) error {
	var data entities.PostTypeRequest
	user_id := uint(middlewares.ExtractTokenUserId(ctx))

	err := ctx.Bind(&data)
	if err != nil {
		return ctx.JSON(400, helpers.ResponseJSON{
			Status:  400,
			Message: err.Error(),
			Data:    nil,
		})
	}

	err = tc.Vs.Validate(data)
	if err != nil {
		return ctx.JSON(400, helpers.ResponseJSON{
			Status:  400,
			Message: "input data is not valid",
			Data:    validation.MessageValidate(err),
		})
	}

	err = tc.Ts.Register(user_id, data)
	if err != nil {
		return ctx.JSON(400, helpers.ResponseJSON{
			Status:  400,
			Message: err.Error(),
			Data:    nil,
		})
	}

	return ctx.JSON(201, helpers.ResponseJSON{
		Status:  201,
		Message: "successfully created",
		Data:    nil,
	})
}

func (tc *postTypeController) GetAll(ctx echo.Context) error {
	data, err := tc.Ts.GetAll()
	if err != nil {
		return ctx.JSON(400, helpers.ResponseJSON{
			Status:  400,
			Message: err.Error(),
			Data:    nil,
		})
	}

	if data == nil {
		return ctx.JSON(404, helpers.ResponseJSON{
			Status:  404,
			Message: "record not found",
			Data:    nil,
		})
	}

	return ctx.JSON(200, helpers.ResponseJSON{
		Status:  200,
		Message: "successfully retrieved",
		Data:    data,
	})
}

func (tc *postTypeController) GetById(ctx echo.Context) error {
	param := ctx.Param("id")

	id, err := tc.Ts.CheckParamId(param)
	if err != nil {
		return ctx.JSON(400, helpers.ResponseJSON{
			Status:  400,
			Message: "param only use id (number)",
			Data:    nil,
		})
	}

	data, err := tc.Ts.GetById(id)
	if err != nil {
		return ctx.JSON(404, helpers.ResponseJSON{
			Status:  404,
			Message: err.Error(),
			Data:    nil,
		})
	}

	return ctx.JSON(200, helpers.ResponseJSON{
		Status:  200,
		Message: "successfully retrieved",
		Data:    data,
	})
}

func (tc *postTypeController) UpdateById(ctx echo.Context) error {
	param := ctx.Param("id")
	var data entities.PostTypeUpdateRequest

	id, err := tc.Ts.CheckParamId(param)
	if err != nil {
		return ctx.JSON(400, helpers.ResponseJSON{
			Status:  400,
			Message: "param only use id (number)",
			Data:    nil,
		})
	}

	if err := ctx.Bind(&data); err != nil {
		return ctx.JSON(400, helpers.ResponseJSON{
			Status:  400,
			Message: err.Error(),
			Data:    nil,
		})
	}

	err = tc.Vs.Validate(data)
	if err != nil {
		return ctx.JSON(400, helpers.ResponseJSON{
			Status:  400,
			Message: "input data is not valid",
			Data:    validation.MessageValidate(err),
		})
	}

	err = tc.Ts.UpdateById(id, data)
	if err != nil {
		return ctx.JSON(400, helpers.ResponseJSON{
			Status:  400,
			Message: err.Error(),
			Data:    nil,
		})
	}

	return ctx.JSON(200, helpers.ResponseJSON{
		Status:  200,
		Message: "successfully updated",
		Data:    nil,
	})
}

func (tc *postTypeController) DeleteById(ctx echo.Context) error {
	param := ctx.Param("id")

	id, err := tc.Ts.CheckParamId(param)
	if err != nil {
		return ctx.JSON(400, helpers.ResponseJSON{
			Status:  400,
			Message: "param only use id (number)",
			Data:    nil,
		})
	}

	err = tc.Ts.DeleteById(id)
	if err != nil {
		return ctx.JSON(400, helpers.ResponseJSON{
			Status:  400,
			Message: err.Error(),
			Data:    nil,
		})
	}

	return ctx.JSON(200, helpers.ResponseJSON{
		Status:  200,
		Message: "successfully deleted",
		Data:    nil,
	})
}
