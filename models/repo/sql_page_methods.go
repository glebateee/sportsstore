package repo

import "sportsstore/models"

func (repo *SqlRepository) GetProductPage(curPage, pageSize int) (pageProducts []models.Product, totalAmount int) {
	rows, err := repo.Commands.GetPage.QueryContext(repo.Context, pageSize, (curPage*pageSize)-pageSize)
	if err == nil {
		if pageProducts, err = scanProducts(rows); err != nil {
			repo.Logger.Panicf("Cannot scan data: %v", err.Error())
		}
	} else {
		repo.Logger.Panicf("Cannot exec GetProductPage command: %v", err)
	}
	row := repo.Commands.GetPageCount.QueryRowContext(repo.Context)
	if row.Err() == nil {
		if err := row.Scan(&totalAmount); err != nil {
			repo.Logger.Panicf("Cannot scan data: %v", err.Error())
		}
	} else {
		repo.Logger.Panicf("Cannot exec GetPageCount command: %v",
			row.Err().Error())
	}
	return

}

func (repo *SqlRepository) GetProductPageCategory(catID, curPage, pageSize int) (products []models.Product, totalAmount int) {
	if catID == 0 {
		return repo.GetProductPage(curPage, pageSize)
	}
	rows, err := repo.Commands.GetCategoryPage.QueryContext(repo.Context, catID, pageSize, (pageSize*curPage)-pageSize)
	if err == nil {
		if products, err = scanProducts(rows); err != nil {
			repo.Logger.Panicf("Cannot scan data: %v", err.Error())
		}
	} else {
		repo.Logger.Panicf("Cannot exec GetProductPage command: %v", err)
	}
	row := repo.Commands.GetCategoryPageCount.QueryRowContext(repo.Context, catID)
	if row.Err() == nil {
		if err := row.Scan(&totalAmount); err != nil {
			repo.Logger.Panicf("Cannot scan data: %v", err.Error())
		}
	} else {
		repo.Logger.Panicf("Cannot exec GetPageCount command: %v",
			row.Err().Error())
	}
	return
}
