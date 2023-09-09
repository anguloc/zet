package main

func main() {
	var m IMessage = 

	p :=

}


type IMessage interface {
	GetTid() string
	GetIdentifier() string
	GetParam() *EmailData
}

type EmailData struct {
	// 详细见消息平台协议
	Subject map[string]map[string]string      `json:"Subject"` // [语种 => [字段名 => 字段值]]
	Body    map[string]map[string]interface{} `json:"Body"`    // [语种 => [字段名 => 字段值]]
}
