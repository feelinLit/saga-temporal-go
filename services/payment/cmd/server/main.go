package main

import "github.com/feelinlit/saga-temporal-go/services/payment/internal/app"

func main() {
	a := app.NewApp()

	if err := a.ListenAndServe(); err != nil {
		panic(err)
	}
}
