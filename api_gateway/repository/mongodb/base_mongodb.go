package mongodb

import (
	"fmt"

	"github.com/aksan/weplus/apigw/repository"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/net/context"
)

type mongoDb struct {
	Host string
	Port string
	User string
	Pass string
	Db   string
}

func NewMongoDB(host, port, user, pass, db string) repository.RepoConf {
	return &mongoDb{
		Host: host,
		Port: port,
		User: user,
		Pass: pass,
		Db:   db,
	}
}

func (*mongoDb) GetRepoName() string {
	return "Repo mongoDb"
}
func (m *mongoDb) HealthCheck() error {
	return nil
}

func (m *mongoDb) Init(r *repository.Repository) error {
	// connStr := fmt.Sprintf("mongodb://%s:%s@%s/%s?directConnection=true", m.User, m.Pass, m.Host+":"+m.Port, m.Db)
	connStr := fmt.Sprintf("mongodb://%s/", m.Host+":"+m.Port)
	if m.Pass != "" {
		connStr = fmt.Sprintf("mongodb://%s:%s@%s/", m.User, m.Pass, m.Host+":"+m.Port)
	}
	fmt.Println(connStr)
	clientOptions := options.Client().ApplyURI(connStr)
	// connStr = fmt.Sprintf("mongodb://%s:%s@%s/%s?authSource=%s&directConnection=true", m.User, m.Pass, m.Host+":"+m.Port, m.Db, authdb)
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		return err
	}

	// Check the connection
	err = client.Ping(context.Background(), nil)
	if err != nil {
		return err
	}
	// include tables
	r.MongoDb.User = &MongoClient{conn: client, db: m.Db}
	return nil
}
