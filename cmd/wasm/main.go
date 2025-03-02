package wasm

import (
	"spisok/internal/robokassa"
	"syscall/js"
)

func generatePaymentURL(this js.Value, args []js.Value) interface{} {
	if len(args) < 1 {
		return "Error: Amount is required"
	}

	amount := args[0].Float()

	config := robokassa.Config{
		MerchantLogin: "demo",
		Password1:     "password1",
		Password2:     "password2",
		IsTest:        true,
	}

	robo := robokassa.NewRobokassa(config)
	url := robo.GeneratePaymentURL(amount, "Оплата услуг", "1")

	return map[string]interface{}{
		"url": url,
	}
}

func main() {
	js.Global().Set("generatePaymentURL", js.FuncOf(generatePaymentURL))
	<-make(chan struct{})
}
