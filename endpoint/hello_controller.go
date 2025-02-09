//endpoint.hello_controller.go
package endpoint

import (
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"go-web/business"
	"go-web/dto"
	"log"
	"net/http"
)

type HelloController struct {
}

func NewHelloController() *HelloController {
	return &HelloController{}
}

func (h *HelloController) Handle(ctx *gin.Context) {
	var request dto.ModelQueryRequest
	//转换请求体
	if err := ctx.ShouldBindJSON(&request); err != nil {
		failResponse := dto.NewFailResponse(500, err.Error())
		ctx.JSON(http.StatusInternalServerError, failResponse)
		return
	}
	bytes, _ := json.Marshal(request)
	log.Printf("接收到请求：traceId = {%s}, request = %s\n", request.Base.TraceId, string(bytes))
	if err := validRequest(request); err != nil {
		failResponse := dto.NewFailResponse(500, err.Error())
		ctx.JSON(http.StatusUnauthorized, failResponse)
		return
	}
	response := business.NewHelloBusiness().GetResponse(&request)
	responseBytes, _ := json.Marshal(response)
	log.Printf("接口返回: traceId = {%s}, response = %s\n", request.Base.TraceId, string(responseBytes))
	ctx.JSON(http.StatusOK, dto.NewSuccessResponse(response))
}

func validRequest(request dto.ModelQueryRequest) error {
	if request.Base == nil {
		return errors.New("base不能为空")
	}
	if len(request.Base.TraceId) == 0 || request.Base.TraceId == "" {
		return errors.New("traceId不能为空")
	}
	return nil
}
