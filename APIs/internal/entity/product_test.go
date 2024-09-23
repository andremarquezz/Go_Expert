package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewProduct(t *testing.T) {
	p, err := NewProduct("Iphone 14 Pro max", 100)
	assert.Nil(t, err)
	assert.NotNil(t, p)
	assert.NotEmpty(t, p.ID)
	assert.NotEmpty(t, p.CreatedAt)
	assert.Equal(t, p.Name, "Iphone 14 Pro max")
	assert.Equal(t, p.Price, 100)
}

func TestProductWhenNameIsRequired(t *testing.T) {
	p, err := NewProduct("", 0)
	assert.NotNil(t, err)
	assert.Nil(t, p)
	assert.Equal(t, ErrNameIsRequired, err)
}

func TestProductWhenPriceIsRequired(t *testing.T) {
	p, err := NewProduct("Iphone 14 Pro max", 0)
	assert.NotNil(t, err)
	assert.Nil(t, p)
	assert.Equal(t, ErrPriceIsRequired, err)
}

func TestProductWhenPriceIsInvalid(t *testing.T) {
	p, err := NewProduct("Iphone 14 Pro max", -100)
	assert.NotNil(t, err)
	assert.Nil(t, p)
	assert.Equal(t, ErrInvalidPrice, err)
}
