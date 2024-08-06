package api_test

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/ory/dockertest/v3"
	"github.com/ory/dockertest/v3/docker"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/jdejesus007/gt-api-project/internal/db"
)

type mockDBExecutor struct {
	database *db.GtDB
}

var pool *dockertest.Pool
var resource *dockertest.Resource

// func (databaseExecutor *mockDBExecutor) GetConn() *gorm.DB {
//   if databaseExecutor.database != nil {
//     return databaseExecutor.database.GetConn()
//   }
//
//   mockDb, _, _ := sqlmock.New()
//   dialector := postgres.New(postgres.Config{
//     Conn:       mockDb,
//     DriverName: "postgres",
//   })
//   dataDB, _ := gorm.Open(dialector, &gorm.Config{})
//   databaseExecutor.database = db.NewGtDB(dataDB)
//   return dataDB
// }

func (databaseExecutor *mockDBExecutor) GetConn() *gorm.DB {
	fmt.Println("MOCK DATABASE -- GET CONNECTION CALLED")

	if databaseExecutor.database != nil {
		return databaseExecutor.database.GetConn()
	}

	var (
		database *gorm.DB
	)

	// uses a sensible default on windows (tcp/http) and linux/osx (socket)
	pool, err = dockertest.NewPool("")
	if err != nil {
		log.Fatalf("Could not construct pool: %s", err)
	}

	err = pool.Client.Ping()
	if err != nil {
		log.Fatalf("Could not connect to Docker: %s", err)
	}

	pgUser := "test"
	pgPwd := "test"
	pgDBName := "gtdb"

	// pulls an image, creates a container based on it and runs it
	resource, err = pool.RunWithOptions(&dockertest.RunOptions{
		Repository: "postgres",
		Tag:        "14",
		Env: []string{
			"POSTGRES_PASSWORD=" + pgPwd,
			"POSTGRES_USER=" + pgUser,
			"POSTGRES_DB=" + pgDBName,
			"listen_addresses = '*'",
		},
	}, func(config *docker.HostConfig) {
		// set AutoRemove to true so that stopped container goes away by itself
		config.AutoRemove = true
		config.RestartPolicy = docker.RestartPolicy{Name: "no"}
	})
	if err != nil {
		log.Fatalf("Could not start resource: %s", err)
	}

	hostAndPort := resource.GetHostPort("5432/tcp")
	hostPort := strings.Split(hostAndPort, ":")
	databaseUrl := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		hostPort[0], pgUser, pgPwd, pgDBName, hostPort[1])

	log.Println("Connecting to database on url: ", databaseUrl)

	// Small test suite
	resource.Expire(30) // Tell docker to hard kill the container in 30 seconds

	// exponential backoff-retry, because the application in the container might not be ready to accept connections yet
	pool.MaxWait = 10 * time.Second
	if err = pool.Retry(func() error {
		database, err = gorm.Open(postgres.Open(databaseUrl), &gorm.Config{
			PrepareStmt: false,
		})

		if err != nil {
			return err
		}
		return database.Error
	}); err != nil {
		log.Fatalf("Could not connect to docker: %s", err)
	}

	databaseExecutor.database = db.NewGtDB(database)
	return databaseExecutor.GetConn()
}

func purgeDBResource() {
	// You can't defer this because os.Exit doesn't care for defer
	if err = pool.Purge(resource); resource != nil && err != nil {
		log.Fatalf("Could not purge resource: %s", err)
	}
}
