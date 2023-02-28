package provider

import (
	"fmt"
	"net/http"
)

const (
	DefaultHost = "https://api.abbey.so"
	TypePrefix  = "abbey"
)

func NewTypeName(name string) string {
	return fmt.Sprintf("%s_%s", TypePrefix, name)
}

type ResourceData struct {
	Client *http.Client
	Host   string
	Token  string
}
