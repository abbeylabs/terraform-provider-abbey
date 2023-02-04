package provider

import (
	"fmt"
	"net/http"
)

const TypePrefix = "abbey"

func NewTypeName(name string) string {
	return fmt.Sprintf("%s_%s", TypePrefix, name)
}

type ResourceData struct {
	Client *http.Client
	Host   string
	Token  string
}
