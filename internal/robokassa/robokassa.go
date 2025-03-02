package robokassa

import (
	"crypto/md5"
	"fmt"
	"net/url"
)

type Config struct {
	MerchantLogin string
	Password1     string
	Password2     string
	IsTest        bool
}

type Robokassa struct {
	config Config
}

func NewRobokassa(config Config) *Robokassa {
	return &Robokassa{config: config}
}

func (r *Robokassa) GeneratePaymentURL(amount float64, description, invoiceID string) string {
	baseURL := "https://auth.robokassa.ru/Merchant/Index.aspx"
	if r.config.IsTest {
		baseURL = "https://auth.robokassa.ru/Merchant/Index.aspx"
	}

	// Формируем подпись
	signature := fmt.Sprintf("%s:%.2f:%s:%s",
		r.config.MerchantLogin,
		amount,
		invoiceID,
		r.config.Password1)
	signatureHash := fmt.Sprintf("%x", md5.Sum([]byte(signature)))

	// Формируем параметры запроса
	params := url.Values{
		"MerchantLogin":  {r.config.MerchantLogin},
		"OutSum":         {fmt.Sprintf("%.2f", amount)},
		"InvId":          {invoiceID},
		"Description":    {description},
		"SignatureValue": {signatureHash},
	}

	if r.config.IsTest {
		params.Add("IsTest", "1")
	}

	return fmt.Sprintf("%s?%s", baseURL, params.Encode())
}
