package main

import (
	sdk "github.com/Ccheers/crawlab-go-sdk"
	"github.com/Ccheers/crawlab-go-sdk/entity"
	"time"
)

func main() {
	item := entity.Result{
		"hello": "world",
	}
	sdk.SaveItem(item)
	time.Sleep(5 * time.Second)
}
