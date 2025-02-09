package business

import (
	"go-web/dto"
	"strings"
)

var answerMap map[string]string

func init() {
	answerMap = make(map[string]string)
	answerMap["你好"] = "你好！如果你有任何问题或需要帮助，请随时告诉我。我会尽力为你提供支持和信息。"
	answerMap["1+1等于几"] = "1+1=2"
	answerMap["李白是什么朝代的"] = "李白是唐朝的诗人"
	answerMap["default"] = "非常抱歉，您的问题超出了我现有的知识库范围，我无法直接为您提供准确的答案。"
}

type HelloBusiness struct {
}

func NewHelloBusiness() *HelloBusiness {
	return &HelloBusiness{}
}

func (h *HelloBusiness) GetResponse(request *dto.ModelQueryRequest) *dto.ModelQueryResponse {
	response := &dto.ModelQueryResponse{}
	for k, v := range answerMap {
		if strings.Contains(request.Question, k) {
			response.Response = v
			break
		}
	}
	if len(response.Response) == 0 {
		response.Response = answerMap["default"]
	}
	return response
}
