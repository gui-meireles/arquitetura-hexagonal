package application_test

import (
	"github.com/golang/mock/gomock"
	"github.com/gui-meireles/go-hexagonal/application"
	mock_application "github.com/gui-meireles/go-hexagonal/application/mocks"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestProductService_Get(t *testing.T) {
	ctrl := gomock.NewController(t)

	// Assim que todas as chamadas forem feitas nessa func de teste, o defer é acionado e executa a ação de Finish()
	defer ctrl.Finish()

	// Criação dos mocks
	product := mock_application.NewMockProductInterface(ctrl)
	persistence := mock_application.NewMockProductPersistenceInterface(ctrl)

	// Sempre que o método Get() for chamado, ele retornará um mock de product
	persistence.EXPECT().Get(gomock.Any()).Return(product, nil).AnyTimes()

	service := application.ProductService{
		Persistence: persistence,
	}

	// Chamando o método Get()
	result, err := service.Get("abc")

	// Validando o método Get()
	require.Nil(t, err)
	require.Equal(t, product, result)
}
