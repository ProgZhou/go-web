package dto

type XingHuoModel struct {
	Type          string `json:"type"`
	defaultAnswer string
	answerMap     map[string]string
}

func NewXingHuoModel(answerMap map[string]string) *XingHuoModel {
	return &XingHuoModel{
		Type:          "xinghuo",
		defaultAnswer: "很抱歉，这个问题我目前没有答案，但我可以和你一起探讨或者帮你寻找其他可能的解决方案。",
		answerMap:     answerMap,
	}
}

func (x *XingHuoModel) GetAnswer(question string) string {
	if x.answerMap == nil || len(x.answerMap) == 0 {
		return x.defaultAnswer
	}
	answer, ok := x.answerMap[question]
	if ok {
		return answer
	} else {
		return x.defaultAnswer
	}
}

type DoubaoModel struct {
	Type          string `json:"type"`
	defaultAnswer string
	answerMap     map[string]string
}

func NewDoubaoModel(answerMap map[string]string) *DoubaoModel {
	return &DoubaoModel{
		Type:          "doubao",
		defaultAnswer: "真不好意思，关于这个问题我确实不太了解呀，我会马上进行学习补充，等我学会了一定第一时间告诉你。",
		answerMap:     answerMap,
	}
}

func (d *DoubaoModel) GetAnswer(question string) string {
	if d.answerMap == nil || len(d.answerMap) == 0 {
		return d.defaultAnswer
	}

	answer, ok := d.answerMap[question]
	if ok {
		return answer
	} else {
		return d.defaultAnswer
	}
}
