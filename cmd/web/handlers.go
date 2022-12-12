package main

import "net/http"


func (app *application) VirtualTerminal(w http.ResponseWriter, r *http.Request) {
	if err := app.renderTemplate(w, r, "terminal", nil); err != nil {
		app.errorLog.Println(err)
	}
}


func (app *application) succededPayment(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		app.errorLog.Println(err)
		return
	}

	// read posted data 
	cardHolder := r.Form.Get("cardholder-name")
	email := r.Form.Get("email")
	paymentIntent := r.Form.Get("payment_intent")
	paymentMethod := r.Form.Get("payment_method")
	paymentAmount := r.Form.Get("payment_amount")
	Currency := r.Form.Get("payment_currency")


	data := make(map[string]interface{})
	data["cardholder"] = cardHolder
	data["email"] = email
	data["payment_intent"] = paymentIntent
	data["payment_method"] = paymentMethod
	data["payment_amount"] = paymentAmount
	data["Currency"] = Currency


	if err := app.renderTemplate(w, r, "succeeded", &templateData{
		Data: data,
	}); err != nil {
		app.errorLog.Println(err)
	}
}