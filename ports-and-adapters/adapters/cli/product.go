package cli

import (
	"fmt"
	"github.com/codeedu/go-hexagonal/application"
)
func Run(service application.ProductServiceInterface, action string, productId string, productName string, price float64) (string, error){
	var result = ""

	switch action {
	case "create":
		product, err := service.Create(productName, price)

		if err != nil {
			return result, err
		}
		result = fmt.Sprintf("ProductId %s with the name %s has been created with the price %f and status %s", 
		product.GetID(),
		product.GetName(),
		product.GetPrice(),
		product.GetStatus())
	
	case "enable":
		product, err := service.Get(productId)
		if err != nil {
			return result, err
		}
		_, err = service.Enable(product)
		if err != nil {
			return result, err
		}
		result = fmt.Sprintf("Product %s has been enabled.", 
		product.GetName())
	case "disable":
		product, err := service.Get(productId)
		if err != nil {
			return result, err
		}
		_, err = service.Disable(product)
		if err != nil {
			return result, err
		}
		result = fmt.Sprintf("Product %s has been disabled.", 
		product.GetName())
	default:
		product, err := service.Get(productId)
		if err != nil {
			return result, err
		}
		result = fmt.Sprintf("Product ID: %s\nName: %s\nPrice: %f\nStatus: %s",
		product.GetID(), product.GetName(), product.GetPrice(), product.GetStatus())
	}

	return result, nil

}