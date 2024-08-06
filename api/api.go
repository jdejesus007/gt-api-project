package api

import (
	"log"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/jdejesus007/gt-api-project/api/provider"
	"github.com/jdejesus007/gt-api-project/docs"
	"github.com/jdejesus007/gt-api-project/internal/config"
	"github.com/jdejesus007/gt-api-project/internal/constants"
	"github.com/jdejesus007/gt-api-project/internal/db"
	"github.com/jdejesus007/gt-api-project/internal/models"
	"github.com/jdejesus007/gt-api-project/internal/routes"
)

// Make sure that this provider struct matches the correct interface.
var _ provider.RepositoryProvider = &defaultRepositoryProvider{}

type defaultRepositoryProvider struct {
	database db.DBExecutor
}

func (d *defaultRepositoryProvider) Database() db.DBExecutor {
	return d.database
}

func GetRepositories() provider.RepositoryProvider {
	return &defaultRepositoryProvider{
		database: &db.GtDB{},
	}
}

// Describe api
type API interface {
	ListenAndServe() error
	Repositories() provider.RepositoryProvider
	CreateServer() *gin.Engine
}

type Implementation struct {
	repositoryProvider provider.RepositoryProvider
}

func (apiImpl *Implementation) providerMiddleware(c *gin.Context) {
	providers := map[string]interface{}{
		constants.DependencyProviderContextKey: apiImpl.Repositories(),
		constants.DBExecutorContextKey:         apiImpl.Repositories().Database(),
	}
	for key, value := range providers {
		c.Set(key, value)
	}
	c.Next()
}

func (apiImpl *Implementation) Repositories() provider.RepositoryProvider {
	return apiImpl.repositoryProvider
}

func (apiImpl *Implementation) CreateServer() *gin.Engine {
	// Run gorm auto migrations and seed test data
	models.AutoMigrate(apiImpl.Repositories().Database())
	models.AutoSeed(apiImpl.Repositories().Database())

	r := gin.Default()

	// repository (dependency) provider middleware
	// allows for easy switch of db and other providers for testing purposes
	r.Use(apiImpl.providerMiddleware)

	// CORS middleware
	r.Use(cors.Default())

	// Base swagger / openapi doc
	docs.SwaggerInfo.BasePath = "/"

	// Register all routes
	routes.Register(r)

	return r
}

func (apiImpl *Implementation) ListenAndServe() error {
	log.Printf("Current Startup Time: %v - Location: %v\n",
		time.Now(), time.Now().Location())

	r := apiImpl.CreateServer()

	// Listen and Server in 0.0.0.0:3000 - standard api port
	return r.Run(":" + config.GtConfig.Port)
}
