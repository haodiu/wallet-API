package routes

import (
	"context"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"log"
	"wallet/configs"
	"wallet/controllers"
	middleware "wallet/middlewares"
	"wallet/services"
)

var (
	as services.AffiliateService
	ac controllers.AffiliateController

	cs services.CustomerService
	cc controllers.CustomerController

	ds services.DocumentService
	dc controllers.DocumentController

	ts services.TransactionService
	tc controllers.TransactionController

	us services.UserService
	uc controllers.UserController
)

func init()  {
	ctx := context.TODO()
	db, err := configs.DbConn()
	if err != nil {
		log.Fatal(err.Error())
	}

	as = services.NewAffiliateService(db, ctx)
	ac = controllers.NewAffiliateController(as)

	cs = services.NewCustomerService(db, ctx)
	cc = controllers.NewCustomerController(cs)

	ds = services.NewDocumentService(db, ctx)
	dc = controllers.NewDocumentController(ds)

	ts = services.NewTransactionService(db, ctx)
	tc = controllers.NewTransactionController(ts)

	us = services.NewUserService(db, ctx)
	uc = controllers.NewUserController(us)
}

func Engine() *gin.Engine{
	//gin.New() - new gin Instance with no middlewares
	//goGonicEngine.Use(gin.Logger())
	//goGonicEngine.Use(gin.Recovery())
	route := gin.Default() // gin with the Logger and Recovery Middlewares attached
	route.Use(cors.Default())
	public := route.Group("/api")
	uc.RegisterUserRoutes(public.Group("/users"))

	protected := route.Group("/api/auth")
	protected.Use(middleware.UserLoaderMiddleware())
	ac.RegisterAffiliateRoute(protected.Group("/affiliates"))
	cc.RegisterCustomerRoute(protected.Group("/customers"))
	dc.RegisterDocumentRoute(protected.Group("/documents"))
	tc.RegisterTransactionRoute(protected.Group("/transactions"))

	return route
}
