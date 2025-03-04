package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewProduct(t *testing.T) {
	p, err := NewProduct("Catecismo", 100.0)
	assert.Nil(t, err)
	assert.NotNil(t, p)
	assert.NotEmpty(t, p.ID)
	assert.Equal(t, "Catecismo", p.Name)
	assert.Equal(t, 100.0, p.Price)
}

func TestProductWhenNameIsRequired(t *testing.T) {
	p, err := NewProduct("", 100.0)
	assert.NotNil(t, err)
	assert.Nil(t, p)
	assert.Equal(t, ErrorNameIsRequired, err)
}

func TestProductWhenPriceIsInvalid(t *testing.T) {
	p, err := NewProduct("Catecismo", -100.0)
	assert.NotNil(t, err)
	assert.Nil(t, p)
	assert.Equal(t, ErrorInvalidPrice, err)
}

func TestProductWhenPriceIsRequired(t *testing.T) {
	p, err := NewProduct("Catecismo", 0)
	assert.NotNil(t, err)
	assert.Nil(t, p)
	assert.Equal(t, ErrorPriceIsRequired, err)
}
