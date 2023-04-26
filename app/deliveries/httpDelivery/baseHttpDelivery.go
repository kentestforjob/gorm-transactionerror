package httpDelivery

import (
	"context"
	"net/http"
	"test/gormtransactionerr/app/domains"
	"time"
)

type CustomResponseJson struct {
	output map[string]interface{}
	sctx   context.Context
}

func NewCustomResponseJson(ctx context.Context) CustomResponseJson {
	// output := make(map[string]interface{})
	return CustomResponseJson{
		output: make(map[string]interface{}),
		sctx:   ctx,
	}
}

func (resp *CustomResponseJson) Fail(err error) map[string]interface{} {
	resp.output = make(map[string]interface{})
	resp.output["success"] = 0
	resp.output["message"] = err.Error()
	resp.output["error_code"] = getStatusCode(err)
	resp.output["system_time"] = time.Now().UTC().Unix()

	return resp.output
}

func (resp *CustomResponseJson) Success() map[string]interface{} {
	resp.output = make(map[string]interface{})

	resp.output["success"] = 1
	resp.output["message"] = ""
	resp.output["system_time"] = time.Now().UTC().Unix()

	return resp.output
}

func GetStatusCode(err error) int {
	return getStatusCode(err)
}

func getStatusCode(err error) int {
	if err == nil {
		return http.StatusOK
	}

	switch err {

	case domains.ErrBadRequest:
		return http.StatusBadRequest
	case domains.ErrInternalServerError:
		return http.StatusInternalServerError
	case domains.ErrNotFound:
		return http.StatusNotFound
	case domains.ErrConflict:
		return http.StatusConflict
	default:
		return http.StatusInternalServerError
	}
}
