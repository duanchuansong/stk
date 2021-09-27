package web

type Http struct {
	Port  string `json:"port"`
	Debug bool   `json:"debug"`
}

type Static struct {
	Port string `json:"port"`
	Dir  string `json:"dir"`
}
