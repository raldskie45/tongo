package main

import (
	"context"
	"fmt"
	"github.com/tonkeeper/tongo"
	"github.com/tonkeeper/tongo/liteapi"
)

func main() {
	// options, err := config.ParseConfigFile("path/to/config.json")
	tongoClient, err := liteapi.NewClientWithDefaultTestnet()
	if err != nil {
		fmt.Printf("Unable to create tongo client: %v", err)
	}
	accountId := tongo.MustParseAccountID("0:E2D41ED396A9F1BA03839D63C5650FAFC6FCFB574FD03F2E67D6555B61A3ACD9")
	state, err := tongoClient.GetAccountState(context.Background(), accountId)
	if err != nil {
		fmt.Printf("Get account state error: %v", err)
	}
	acc, _ := tongo.GetAccountInfo(state.Account)
	fmt.Printf("Account status: %v\nBalance: %v\n", acc.Status, acc.Balance)
}
