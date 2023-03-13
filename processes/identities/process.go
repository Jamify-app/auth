package identities

import (
	"context"
	"fmt"
	"os"

	"github.com/Jamify-app/auth/processes/identities/repositories"
	"github.com/Jamify-app/auth/processes/identities/services"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Process struct {
	group   *gin.RouterGroup
	store   *mongo.Client
	service services.IService
}

func NewProcess() *Process {
	return &Process{}
}

func (p *Process) Start(ctx context.Context, engine *gin.Engine) {
	var err error
	p.group = engine.Group("/identities")
	p.store, err = mongo.NewClient(options.Client().ApplyURI(mongoDBURI()))
	if err != nil {
		panic(err)
	}

	err = p.store.Connect(ctx)
	if err != nil {
		panic(err)
	}

	repository := repositories.NewRepository(p.store)
	p.service = services.NewService(repository)

	RegisterRoutes(p.group, p.service)
}

func mongoDBURI() string {
	port := os.Getenv("MONGODB_PORT")
	username := os.Getenv("MONGODB_USERNAME")
	password := os.Getenv("MONGODB_PASSWORD")
	return fmt.Sprintf("mongodb://%s:%s@localhost:%s/?ssl=false", username, password, port)
}

func (p *Process) Stop(ctx context.Context) error {
	return p.store.Disconnect(ctx)
}
