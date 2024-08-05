package routes

import (
	"github.com/gin-gonic/gin"
	internal "github.com/jdejesus007/gt-api-project/internal/actions"
	"github.com/jdejesus007/gt-api-project/internal/actions/checkout"
	"github.com/jdejesus007/gt-api-project/internal/actions/customer"
	"github.com/jdejesus007/gt-api-project/internal/actions/order"
)

// ModelRegister defines the common interface to port all associated
// model routes into the given router.
type Model interface {
	// RegisterRoutes receives a gin.Engine object
	// and will add all the associated paths and handler
	// functions into the router.
	RegisterRoutes(*gin.Engine)
}

func Register(ginRouter *gin.Engine) {
	routeHandlers := []Model{
		// Internal general routes
		&internal.PingService{},
		&internal.FavIconService{},
		&internal.SwaggerService{},

		// Business routes
		&customer.CustomerService{},
		&order.OrderService{},
		&checkout.CheckoutService{},
	}

	for _, r := range routeHandlers {
		r.RegisterRoutes(ginRouter)
	}
}
