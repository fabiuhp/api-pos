package database

import (
	"math/rand/v2"
	"strconv"
	"testing"

	"github.com/fabiuhp/api-pos/internal/entity"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestCreateNewProduct(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}
	db.AutoMigrate(&entity.Product{})
	product, err := entity.NewProduct("Product 1", 10.5)
	assert.Nil(t, err)
	productDB := NewProduct(db)
	err = productDB.Create(product)
	assert.NoError(t, err)
	assert.NotEmpty(t, product.ID)
}

func TestFindAllProducts(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}
	db.AutoMigrate(&entity.Product{})
	productDB := NewProduct(db)

	for i := range 10 {
		product, err := entity.NewProduct("Product"+strconv.Itoa(i), rand.Float64()*100)
		assert.Nil(t, err)
		db.Create(product)
	}

	products, err := productDB.FindAll(1, 10, "asc")
	assert.Nil(t, err)
	assert.Len(t, products, 10)
	assert.Equal(t, "Product0", products[0].Name)
	assert.Equal(t, "Product9", products[9].Name)

	products, err = productDB.FindAll(2, 5, "asc")
	assert.Nil(t, err)
	assert.Len(t, products, 5)
	assert.Equal(t, "Product5", products[0].Name)
	assert.Equal(t, "Product9", products[4].Name)
}
