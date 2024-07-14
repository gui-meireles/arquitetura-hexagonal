package application_test

import (
	"github.com/gui-meireles/go-hexagonal/application"
	uuid "github.com/satori/go.uuid"
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

func TestProduct_Disable(t *testing.T) {
	// Vamos instanciar o struct de Product
	product := application.Product{}
	product.Name = "Tomate"
	product.Status = application.ENABLED
	product.Price = 0

	err := product.Disable()
	require.Nil(t, err)

	product.Price = 10
	err = product.Disable()
	require.Equal(t, "The price must be 0 in order to have the product disabled", err.Error())
}

func TestProduct_IsValid(t *testing.T) {
	product := application.Product{}
	product.ID = uuid.NewV4().String()
	product.Name = "Abacate"
	product.Status = application.DISABLED
	product.Price = 10

	_, err := product.IsValid()
	require.Nil(t, err)

	// Passando status que não existe
	product.Status = "INVALID"
	_, err = product.IsValid()
	require.Equal(t, "The status must be enabled or disabled", err.Error())

	product.Status = application.ENABLED
	_, err = product.IsValid()
	require.Nil(t, err)

	product.Price = -10
	_, err = product.IsValid()
	require.Equal(t, "The price must be greater or equal zero", err.Error())
}
