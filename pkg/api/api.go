package api

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type BaseApi struct {
	Ctx echo.Context
	Orm *gorm.DB
}

func (a *BaseApi) GetOrm() {

}
