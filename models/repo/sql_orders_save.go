package repo

import "sportsstore/models"

func (repo *SqlRepository) SaveOrder(order *models.Order) {
	tx, err := repo.DB.Begin()
	if err != nil {
		repo.Logger.Panicf("Cannot create transaction: %v", err.Error())
	}
	result, err := tx.StmtContext(repo.Context, repo.Commands.SaveOrder).Exec(order.Name, order.StreetAddr,
		order.City, order.Zip, order.Country, order.Shipped)
	if err != nil {
		repo.Logger.Panicf("Cannot exec SaveOrder command: %v", err.Error())
		tx.Rollback()
	}
	id, err := result.LastInsertId()
	if err != nil {
		repo.Logger.Panicf("Cannot get inserted ID: %v", err.Error())
		tx.Rollback()
	}
	statement := tx.StmtContext(repo.Context, repo.Commands.SaveOrderLine)
	for _, prod := range order.Products {
		_, err := statement.Exec(id, prod.Product.ID, prod.Quantity)
		if err != nil {
			repo.Logger.Panicf("Cannot exec SaveOrderLine command: %v", err.Error())
			tx.Rollback()
		}

	}
	err = tx.Commit()
	if err != nil {
		repo.Logger.Panicf("Transaction cannot be committed: %v", err.Error())
		err = tx.Rollback()
		if err != nil {
			repo.Logger.Panicf("Transaction cannot be rolled back: %v", err.Error())
		}
	}
	order.ID = int(id)
}
