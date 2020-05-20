// Code generated by goa v3.1.2, DO NOT EDIT.
//
// jenga HTTP server types
//
// Command:
// $ goa gen github.com/wondenge/listeners/design

package server

import (
	jenga "github.com/wondenge/listeners/gen/jenga"
	goa "goa.design/goa/v3/pkg"
)

// PublishRequestBody is the type of the "jenga" service "publish" endpoint
// HTTP request body.
type PublishRequestBody struct {
	// recipient phone number
	MobileNumber *string `form:"mobileNumber,omitempty" json:"mobileNumber,omitempty" xml:"mobileNumber,omitempty"`
	// recipient name
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
	// reference number
	Rrn *string `form:"rrn,omitempty" json:"rrn,omitempty" xml:"rrn,omitempty"`
	// for M-Pesa the value is M-Pesa
	Service *string `form:"service,omitempty" json:"service,omitempty" xml:"service,omitempty"`
	// M-Pesa receipt number
	TranID *string `form:"tranID,omitempty" json:"tranID,omitempty" xml:"tranID,omitempty"`
	// M-Pesa return code
	ResultCode *string `form:"resultCode,omitempty" json:"resultCode,omitempty" xml:"resultCode,omitempty"`
	// M-Pesa return code description
	ResultDescription *string `form:"resultDescription,omitempty" json:"resultDescription,omitempty" xml:"resultDescription,omitempty"`
	// Additional information
	AdditionalInfo *string `form:"additionalInfo,omitempty" json:"additionalInfo,omitempty" xml:"additionalInfo,omitempty"`
}

// AlertsRequestBody is the type of the "jenga" service "alerts" endpoint HTTP
// request body.
type AlertsRequestBody struct {
	Name         *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	MobileNumber *string `form:"mobileNumber,omitempty" json:"mobileNumber,omitempty" xml:"mobileNumber,omitempty"`
	// S2596405
	Reference *string `form:"reference,omitempty" json:"reference,omitempty" xml:"reference,omitempty"`
	// date
	Date *string `form:"date,omitempty" json:"date,omitempty" xml:"date,omitempty"`
	// paymentMode
	PaymentMode *string `form:"paymentMode,omitempty" json:"paymentMode,omitempty" xml:"paymentMode,omitempty"`
	// amount
	Amount *string `form:"amount,omitempty" json:"amount,omitempty" xml:"amount,omitempty"`
	Till   *string `form:"till,omitempty" json:"till,omitempty" xml:"till,omitempty"`
	// billNumber
	BillNumber *string `form:"billNumber,omitempty" json:"billNumber,omitempty" xml:"billNumber,omitempty"`
	// orderAmount
	OrderAmount *string `form:"orderAmount,omitempty" json:"orderAmount,omitempty" xml:"orderAmount,omitempty"`
	// serviceCharge
	ServiceCharge *string `form:"serviceCharge,omitempty" json:"serviceCharge,omitempty" xml:"serviceCharge,omitempty"`
	ServedBy      *string `form:"servedBy,omitempty" json:"servedBy,omitempty" xml:"servedBy,omitempty"`
	// additionalInfo
	AdditionalInfo *string `form:"additionalInfo,omitempty" json:"additionalInfo,omitempty" xml:"additionalInfo,omitempty"`
	// transactionType
	TransactionType *string `form:"transactionType,omitempty" json:"transactionType,omitempty" xml:"transactionType,omitempty"`
	// account
	Account *string `form:"account,omitempty" json:"account,omitempty" xml:"account,omitempty"`
}

// PublishResponseBody is the type of the "jenga" service "publish" endpoint
// HTTP response body.
type PublishResponseBody struct {
	Code    *string `form:"Code,omitempty" json:"Code,omitempty" xml:"Code,omitempty"`
	Message *string `form:"Message,omitempty" json:"Message,omitempty" xml:"Message,omitempty"`
}

// NewPublishResponseBody builds the HTTP response body from the result of the
// "publish" endpoint of the "jenga" service.
func NewPublishResponseBody(res *jenga.JengaCallbackMedia) *PublishResponseBody {
	body := &PublishResponseBody{
		Code:    res.Code,
		Message: res.Message,
	}
	return body
}

// NewPublishJengaCallbackPayload builds a jenga service publish endpoint
// payload.
func NewPublishJengaCallbackPayload(body *PublishRequestBody, auth *string) *jenga.JengaCallbackPayload {
	v := &jenga.JengaCallbackPayload{
		MobileNumber:      body.MobileNumber,
		Message:           body.Message,
		Rrn:               body.Rrn,
		Service:           body.Service,
		TranID:            body.TranID,
		ResultCode:        body.ResultCode,
		ResultDescription: body.ResultDescription,
		AdditionalInfo:    body.AdditionalInfo,
	}
	v.Auth = auth

	return v
}

// NewAlertsAccountAlerts builds a jenga service alerts endpoint payload.
func NewAlertsAccountAlerts(body *AlertsRequestBody, auth *string) *jenga.AccountAlerts {
	v := &jenga.AccountAlerts{
		Name:            body.Name,
		MobileNumber:    body.MobileNumber,
		Reference:       body.Reference,
		Date:            body.Date,
		PaymentMode:     body.PaymentMode,
		Amount:          body.Amount,
		Till:            body.Till,
		BillNumber:      body.BillNumber,
		OrderAmount:     body.OrderAmount,
		ServiceCharge:   body.ServiceCharge,
		ServedBy:        body.ServedBy,
		AdditionalInfo:  body.AdditionalInfo,
		TransactionType: body.TransactionType,
		Account:         body.Account,
	}
	v.Auth = auth

	return v
}

// ValidateAlertsRequestBody runs the validations defined on AlertsRequestBody
func ValidateAlertsRequestBody(body *AlertsRequestBody) (err error) {
	if body.Date != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("body.date", *body.Date, goa.FormatDateTime))
	}
	return
}