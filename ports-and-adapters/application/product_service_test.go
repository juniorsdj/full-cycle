package application_test

import (
	"github.com/codeedu/go-hexagonal/application"
	mock_application "github.com/codeedu/go-hexagonal/application/mocks"
	"testing"
	"github.com/stretchr/testify/require"
	"github.com/golang/mock/gomock"
)

func TestProductService_Read(t *testing.T){

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	product := mock_application.NewMockProductInterface(ctrl)
	persistence := mock_application.NewMockProductPersistenceInterface(ctrl)
	persistence.EXPECT().Read(gomock.Any()).Return(product, nil).AnyTimes()

	service := application.ProductService{
		Persistence: persistence,
	}

	result, err := service.Read("abc")
	require.Nil(t, err)
	require.Equal(t, product, result)

}
func TestProductService_Create(t *testing.T){

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	product := mock_application.NewMockProductInterface(ctrl)
	persistence := mock_application.NewMockProductPersistenceInterface(ctrl)
	persistence.EXPECT().Write(gomock.Any()).Return(product, nil).AnyTimes()

	service := application.ProductService{
		Persistence: persistence,
	}

	result, err := service.Create("hello", 10)

	require.Nil(t, err)
	require.Equal(t, product, result)

}

func TestProductService_EnableDisable(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	product := mock_application.NewMockProductInterface(ctrl)
	product.EXPECT().Enable().Return(nil)
	product.EXPECT().Disable().Return(nil)

	persistence := mock_application.NewMockProductPersistenceInterface(ctrl)
	persistence.EXPECT().Write(gomock.Any()).Return(product, nil).AnyTimes()
	service := application.ProductService{Persistence: persistence}

	result, err := service.Enable(product)
	require.Nil(t, err)
	require.Equal(t, product, result)

	result, err = service.Disable(product)
	require.Nil(t, err)
	require.Equal(t, product, result)
}