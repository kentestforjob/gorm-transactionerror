package httpDelivery

import (
	"net/http"
	"test/gormtransactionerr/app/usecases"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

type DummyHandler struct {
	dummy_usecase usecases.InterfaceDummyUsecase
}

// initialize the resources endpoint
func NewDummyHandler(route_engine *gin.Engine,
	dummy_usecase usecases.InterfaceDummyUsecase,

) *DummyHandler {

	handler := &DummyHandler{
		dummy_usecase: dummy_usecase,
	}
	return handler
}

func (dh *DummyHandler) PostUpdate(ctx *gin.Context) {

	resp_obj := NewCustomResponseJson(ctx.Request.Context())

	var request_body struct {
		UserId uint32 `json:"user_id" validate:"required"`
		Email  string `json:"email" validate:"required"`
	}

	err := ctx.ShouldBindJSON(&request_body)
	if err != nil {
		ctx.JSON(http.StatusOK, resp_obj.Fail(err))
		return
	}
	validate := validator.New()
	err = validate.Struct(&request_body)
	if err != nil {
		ctx.JSON(http.StatusOK, resp_obj.Fail(err))
		return
	}

	err = dh.dummy_usecase.UpdateDummy(request_body.UserId, request_body.Email)
	if err != nil {
		ctx.JSON(http.StatusOK, resp_obj.Fail(err))
		return
	}

	result := resp_obj.Success()
	ctx.JSON(http.StatusOK, result)
}

func (dh *DummyHandler) GetList(ctx *gin.Context) {

	resp_obj := NewCustomResponseJson(ctx.Request.Context())

	list, err := dh.dummy_usecase.FindAll()
	if err != nil {
		ctx.JSON(http.StatusOK, resp_obj.Fail(err))
		return
	}

	result := resp_obj.Success()
	result["list"] = list
	ctx.JSON(http.StatusOK, result)
}
