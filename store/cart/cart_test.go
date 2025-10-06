package cart

import (
	"sportsstore/models"
	"testing"
)

// Тест для BasicCart (реализация Cart интерфейса)
// Проверяем базовые операции: добавление, подсчёт суммы, удаление, сброс

func TestBasicCart_AddProduct(t *testing.T) {
	// Создаём пустую корзину
	cart := &BasicCart{}

	// Тестовый продукт
	product1 := models.Product{ID: 1, Price: 10.0}
	product2 := models.Product{ID: 2, Price: 20.0}

	// Добавляем первый продукт
	cart.AddProduct(product1)
	if len(cart.GetLines()) != 1 {
		t.Errorf("Expected 1 line after adding first product, got %d", len(cart.GetLines()))
	}

	// Добавляем тот же продукт — количество должно вырасти
	cart.AddProduct(product1)
	if len(cart.GetLines()) != 1 || cart.GetLines()[0].Quantity != 2 {
		t.Errorf("Expected quantity 2 for product1, got %d", cart.GetLines()[0].Quantity)
	}

	// Добавляем второй продукт
	cart.AddProduct(product2)
	if len(cart.GetLines()) != 2 {
		t.Errorf("Expected 2 lines after adding second product, got %d", len(cart.GetLines()))
	}
}
