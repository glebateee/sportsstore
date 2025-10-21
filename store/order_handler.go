package store

import (
	"encoding/json"
	"platform/http/actionresults"
	"platform/http/handling"
	"platform/sessions"
	"platform/validation"
	"sportsstore/models"
	"sportsstore/store/cart"
	"strings"
)

type OrderHandler struct {
	cart.Cart
	sessions.Session
	Repository   models.Repository
	URLGenerator handling.URLGenerator
	validation.Validator
}

type OrderTemplateContext struct {
	models.ShippingDetails
	ValidationErrors [][]string
	CancelUrl        string
}

func (handler OrderHandler) GetCheckout() actionresults.ActionResult {
	context := OrderTemplateContext{}
	contextJson := handler.Session.GetValueDefault("checkout_details", "")
	if contextJson != nil {
		json.NewDecoder(strings.NewReader(contextJson.(string))).Decode(&context)
	}
	context.CancelUrl = mustGenerateUrl(handler.URLGenerator, CartHandler.GetCart)
	return actionresults.NewTemplateAction("checkout.html", context)
}

func (handler OrderHandler) PostCheckout(details models.ShippingDetails) actionresults.ActionResult {
	valid, errors := handler.Validator.Validate(details)
	if !valid {
		ctx := OrderTemplateContext{
			ShippingDetails:  details,
			ValidationErrors: [][]string{},
		}
		for _, err := range errors {
			ctx.ValidationErrors = append(ctx.ValidationErrors, []string{err.FieldName, err.Error.Error()})
		}
		builder := strings.Builder{}
		json.NewEncoder(&builder).Encode(ctx)
		handler.Session.SetValue("checkout_details", builder.String())
		redirectUrl := mustGenerateUrl(handler.URLGenerator, OrderHandler.GetCheckout)
		return actionresults.NewRedirectAction(redirectUrl)
	}
	handler.Session.SetValue("checkout_details", "")
	lines := handler.Cart.GetLines()
	order := models.Order{
		ShippingDetails: details,
		Products:        make([]models.ProductSelection, 0, len(lines)),
	}
	for _, line := range lines {
		order.Products = append(order.Products, models.ProductSelection{
			Quantity: line.Quantity,
			Product:  line.Product,
		})
	}
	handler.Repository.SaveOrder(&order)
	handler.Cart.Reset()
	targetUrl, _ := handler.URLGenerator.GenerateUrl(OrderHandler.GetSummary, order.ID)
	return actionresults.NewRedirectAction(targetUrl)
}

func (handler OrderHandler) GetSummary(id int) actionresults.ActionResult {
	targetUrl, _ := handler.URLGenerator.GenerateUrl(ProductHandler.GetProducts, 0, 1)
	return actionresults.NewTemplateAction("checkout_summary.html",
		struct {
			ID        int
			TargetUrl string
		}{ID: id, TargetUrl: targetUrl})
}
