package main

import (
	"myapp/internal/models"
	"net/http"
)


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
	cardHolder := r.Form.Get("cardholder_name")
	email := r.Form.Get("cardholder_email")
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

	// should write this data tp session, and then redirect user to new page

	if err := app.renderTemplate(w, r, "succeeded", &templateData{
		Data: data,
	}); err != nil {
		app.errorLog.Println(err)
	}
}


// ChargeOnce displays the page to buy one widget
func (app *application) ChargeOnce(w http.ResponseWriter, r *http.Request) {
	
	Widget := models.Widget{
		Id: 1,
		Name: "Custom widget",
		Description: "Very nice widget",
		InventoryLevel: 10,
		Price: 1000,
	}
	data := make(map[string]interface{})
	data["widget"] = Widget
	if err := app.renderTemplate(w, r, "buy-once", &templateData{
		Data: data,
	}); err != nil {
		app.errorLog.Println(err)
	}
}


func (app *application) Home(w http.ResponseWriter, r *http.Request) {
	if err := app.renderTemplate(w, r, "home", nil); err != nil {
		app.errorLog.Println(err)
	}
}