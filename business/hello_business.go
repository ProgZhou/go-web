package business

import (
	"go-web/dto"
)

var xinghuoAnswerMap map[string]string
var doubaoAnswerMap map[string]string

var xinghuoModel *dto.XingHuoModel
var doubaoModel *dto.DoubaoModel

func init() {
	xinghuoAnswerMap = make(map[string]string)
	xinghuoAnswerMap["你好"] = "你好！如果你有任何问题或需要帮助，请随时告诉我。我会尽力为你提供支持和信息。"
	xinghuoAnswerMap["1+1等于几"] = "1+1=2"
	xinghuoAnswerMap["李白是什么朝代的"] = "李白是唐朝的诗人"
	xinghuoAnswerMap["你是谁，你来自哪里"] = "我是讯飞星火认知大模型，我来自中国。创造我的公司是科大讯飞，成立于1999年，是亚太地区知名的智能语音和人工智能上市企业，总部位于安徽合肥。"
	doubaoAnswerMap = make(map[string]string)
	doubaoAnswerMap["你好"] = "你好呀，有什么可以帮助你的吗？无论是有趣的话题、问题咨询还是其他任何事，都可以随时和我说。"
	doubaoAnswerMap["1+1等于几"] = "哈哈，这可真是个经典的问题呀，在常规的数学运算中，1+1 等于 2 呀。不过在一些特殊的情境或概念下，1+1 也可能有不同的结果呢，比如在二进制中 1+1 等于 10；在一些团队合作等概念里，1+1 可能会大于 2，表示产生了额外的协同效应；而在某些错误的计算或者特殊定义下，可能会得出其他奇怪的结果，但从基本的算术意义上来说就是 2。"
	doubaoAnswerMap["李白是什么朝代的"] = "李白是唐朝的著名诗人呀。"
	doubaoAnswerMap["你是谁，你来自哪里"] = "我叫豆包呀，我是由字节跳动独立开发和训练的呀，并不来自某个具体的地理位置呢。我通过不断学习大量的各种类型的文本数据，来获取知识和提升能力，从而可以在这里随时为大家提供各种有用的信息、解答各种问题、进行各种有趣的交流互动呀。"
	xinghuoModel = dto.NewXingHuoModel(xinghuoAnswerMap)
	doubaoModel = dto.NewDoubaoModel(doubaoAnswerMap)
}

type HelloBusiness struct {
}

func NewHelloBusiness() *HelloBusiness {
	return &HelloBusiness{}
}

func (h *HelloBusiness) GetResponse(request *dto.ModelQueryRequest) *dto.ModelQueryResponse {
	response := &dto.ModelQueryResponse{}
	if request.ModelType == "xinghuo" {
		response.Response = xinghuoModel.GetAnswer(request.Question)
	} else {
		response.Response = doubaoModel.GetAnswer(request.Question)
	}
	return response
}
