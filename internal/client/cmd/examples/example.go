package main

import (
	"context"
	"fmt"

	"github.com/go-provider-sdk/pkg/client"
	"github.com/go-provider-sdk/pkg/clientconfig"
)

func main() {
	config := clientconfig.NewConfig()

	client := client.NewClient(config)

	res, err := client.GrantKits.ListGrantKits(context.Background())
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v", res)
}
