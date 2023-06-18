package dto

type Module struct {
	Name string `json:"name"`
}

var DmhyModule = &Module{
	Name: "dmhy",
}
