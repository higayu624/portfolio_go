package main

import (
	"fmt"
	"github.com/higayu624/portfolio_go/src/app/infrastructure"
)

func main(){
	fmt.Println("server start")
	infrastructure.Router.Run()
}

