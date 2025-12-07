package main

import "github.com/feelinlit/saga-temporal-go/services/order/internal/app"

func main() {
	a, err := app.NewApp()
	if err != nil {
		panic(err)
	}

	if err = a.ListenAndServe(); err != nil {
		panic(err)
	}
}
