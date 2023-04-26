/*
dependency injection
*/
package di

import (
	"test/gormtransactionerr/app/deliveries/httpDelivery"
	"test/gormtransactionerr/app/repositories"
	"test/gormtransactionerr/app/usecases"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// Injectors from wire.go:

func InitializeAPIs(dbConn *gorm.DB, route_engine *gin.Engine) {

	// repositories
	dummyRepository := repositories.NewDummy(dbConn)

	// usecase
	dummyUseCase := usecases.NewDummyUseCase(
		dbConn,
		dummyRepository,
	)

	// controller
	dummyHandler := httpDelivery.NewDummyHandler(
		route_engine,
		dummyUseCase,
	)

	api := route_engine.Group("/api")
	api.POST("/dummy/update", dummyHandler.PostUpdate)
	api.GET("/dummy/list", dummyHandler.GetList)

}
