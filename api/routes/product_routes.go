package routes

import (
	"net/http"
	"project/api/controllers"
)

type ProductRoutes interface{
	Routes()[]*Route
}


type productsRoutesImpl struct {
	productsController controllers.ProductsController
}

func NewProductRoutes(productsController controllers.ProductsController) *productsRoutesImpl{
		return &productsRoutesImpl{productsController}
}

func (r *productsRoutesImpl) Routes() []*Route {
	return []*Route{
		&Route{
				Path: "/products",
				Method: http.MethodPost,
				Handler: r.productsController.Post,
		},
	}
}