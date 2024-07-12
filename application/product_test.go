package application_test

import (
	"github.com/gui-meireles/go-hexagonal/application"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestProduct_Enable(t *testing.T) {
	// Vamos instanciar o struct de Product
	product := application.Product{}
	product.Name = "Tomate"
	product.Status = application.DISABLED
	product.Price = 10

	// Chama a func Enable() - 1° Teste
	err := product.Enable()
	// Verifica se o retorno da func é Nil
	require.Nil(t, err)

	// Teste de erro enviando o price = 0 - 2° Teste
	product.Price = 0
	err = product.Enable()
	require.Equal(t, "The price must be greater than 0 to enable the product", err.Error())
}
