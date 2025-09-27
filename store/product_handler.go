package store

import (
	"math"
	"platform/http/actionresults"
	"platform/http/handling"
	"sportsstore/models"
)

const pageSize = 4

type ProductHandler struct {
	Repository   models.Repository
	URLGenerator handling.URLGenerator
}
type ProductTemplateContext struct {
	Products    []models.Product
	Page        int
	PageCount   int
	PageNumbers []int
	PageUrlFunc func(int) string
}

func (handler ProductHandler) GetProducts(page int) actionresults.ActionResult {
	shownProds, totalProds := handler.Repository.GetProductPage(page, pageSize)
	pageCount := int(math.Ceil(float64(totalProds) / float64(pageSize)))
	return actionresults.NewTemplateAction("store_layout.html",
		ProductTemplateContext{
			Products:    shownProds,
			Page:        page,
			PageCount:   pageCount,
			PageNumbers: handler.generatePageNumbers(pageCount),
			PageUrlFunc: handler.createPageUrlFunction(),
		})
}

func (handler ProductHandler) generatePageNumbers(pageCount int) (pages []int) {
	pages = make([]int, pageCount)
	for i := range pageCount {
		pages[i] = i + 1
	}
	return
}

func (handler ProductHandler) createPageUrlFunction() func(int) string {
	return func(pageNum int) string {
		url, _ := handler.URLGenerator.GenerateUrl(ProductHandler.GetProducts, pageNum)
		return url
	}
}
