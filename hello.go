package main

import (
	"context"
	"fmt"
	"github.com/onflow/flow-go-sdk/client"
	"google.golang.org/grpc"
)

func ExecuteScript(node string, script []byte) {
	ctx := context.Background()
	c, err := client.New(node, grpc.WithInsecure())
	if err != nil {
		panic("failed to connect to node")
	}

	result, err := c.ExecuteScriptAtLatestBlock(ctx, script, nil)
	if err != nil {
		panic(err)
	}

	fmt.Println(result)
}

func main() {
	node := "127.0.0.1:3569"
	script := `
	import HelloWorld from 0x179b6b1cb6755e31

	pub fun main(): String {
	    return HelloWorld.hello()
	}
	`
	ExecuteScript(node, []byte(script))
}
