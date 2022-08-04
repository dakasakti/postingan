package controllers

import "github.com/labstack/echo/v4"

type PostTypeController interface {
	Register(ctx echo.Context) error
	GetAll(ctx echo.Context) error
	GetById(ctx echo.Context) error
	UpdateById(ctx echo.Context) error
	DeleteById(ctx echo.Context) error
}
