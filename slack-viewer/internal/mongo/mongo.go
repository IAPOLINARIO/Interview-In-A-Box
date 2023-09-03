package mongo

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"sync"
	"time"

<<<<<<<< HEAD:internal/mongo/mongo.go
	"github.com/slack-viewer/internal/dtos"
========
	"github.com/slack-viewer/pkg/dtos"
>>>>>>>> 42df753928d6046fdabc880dfc8804a9fa34e9c9:pkg/mongo/mongo.go
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	// Path to the AWS CA file
	caFilePath = "../certs/global-bundle.pem"

	// Timeout operations after N seconds
	connectTimeout  = 10
	queryTimeout    = 30
	username        = ""
	password        = ""
	clusterEndpoint = "sampledb.com:27017"

	// Which instances to read from
	readPreference = "secondaryPreferred"

	connectionStringTemplate = "mongodb://%s:%s@%s/sampledb?tls=true&replicaSet=rs0&readpreference=%s"
)

var (
	mongoMgt *MongoManager
)

type MongoManager struct {
	Client                *mongo.Client
	Ctx                   context.Context
	DbName                string
	MessagesColletionName string
	UsersColletionName    string
	GroupsColletionName   string
	DmsColletionName      string
	MpimsColletionName    string
	ChannelsColletionName string
}

// NewMongoManager creates a new instance of the MongoManager struct and returns it.
// If an instance of MongoManager already exists, it will be returned instead of creating a new one.
func NewMongoManager() (mgt *MongoManager, err error) {
	if mongoMgt != nil {
		return mongoMgt, nil
	}

	ctx := context.Background()
	mgt = &MongoManager{
		Client:                NewMongoClient(ctx),
		Ctx:                   ctx,
		MessagesColletionName: "messages",
		DbName:                "slack_",
		UsersColletionName:    "users",
		GroupsColletionName:   "groups",
		DmsColletionName:      "dms",
		MpimsColletionName:    "mpims",
		ChannelsColletionName: "channels",
	}

	return mgt, nil
}

// Disconnect disconnects from the MongoDB client.
// It logs an error if there is any and returns it.
func (m *MongoManager) Disconnect() (err error) {
	if m.Client == nil {
		return
	}

	err = m.Client.Disconnect(m.Ctx)
	if err != nil {
		log.Fatal(err)
	}

	// TODO optional you can log your closed MongoDB client
	fmt.Println("Connection to MongoDB closed.")

	return err
}

// DeleteDatabase deletes a MongoDB database by the given name.
// If there's an error, it logs a message and returns it.
func (m *MongoManager) DeleteDatabase(dbName string) (err error) {

	err = m.Client.Database(dbName).Drop(m.Ctx)

	if err != nil {
		fmt.Println("Error droping the databse: ", err)
	} else {
		fmt.Printf("Database %s deleted \n", dbName)
	}

	return err
}

func getCustomTLSConfig(caFile string) (*tls.Config, error) {
	tlsConfig := new(tls.Config)
	certs, err := ioutil.ReadFile(caFile)

	if err != nil {
		return tlsConfig, err
	}

	tlsConfig.RootCAs = x509.NewCertPool()
	ok := tlsConfig.RootCAs.AppendCertsFromPEM(certs)

	if !ok {
		return tlsConfig, errors.New("Failed parsing pem file")
	}

	return tlsConfig, nil
}

// NewMongoClient creates a new MongoDB client and returns it.
// It uses the default MongoDB URI or one specified in an environment variable.
func NewMongoClient(ctx context.Context) (client *mongo.Client) {
	var err error
	connectionURI := fmt.Sprintf(connectionStringTemplate, username, password, clusterEndpoint, readPreference)
	fmt.Println("Connection URI:", connectionURI)

	tlsConfig, err := getCustomTLSConfig(caFilePath)
	if err != nil {
		panic(fmt.Errorf("Failed getting TLS configuration: %v", err))
	}

	client, err = mongo.NewClient(options.Client().ApplyURI(connectionURI).SetTLSConfig(tlsConfig))
	if err != nil {
		panic(fmt.Errorf("Failed to create client: %v", err))
	}

	ctx, cancel := context.WithTimeout(ctx, connectTimeout*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		panic(fmt.Errorf("Failed to connect to cluster: %v", err))
	}

	// Force a connection to verify our connection string
	err = client.Ping(ctx, nil)
	if err != nil {
		panic(fmt.Errorf("Failed to ping cluster: %v", err))
	}

	fmt.Println("Connected to DocumentDB!")

	return client
}

// Connect returns the current MongoDB client.
// If there isn't one, it creates a new client and returns it.
func (m *MongoManager) Connect() *mongo.Client {
	if m.Client != nil {
		return m.Client
	} else {
		m.Client = NewMongoClient(m.Ctx)
	}

	return m.Client
}

// saveRecordsToMongoDB saves records to a MongoDB collection.
// It logs an error if there is any and returns it.
func (m *MongoManager) saveRecordsToMongoDB(records []interface{}, collectionName string) error {
	_, insertErr := m.Client.Database(m.DbName).Collection(collectionName).InsertMany(m.Ctx, records)

	if insertErr != nil {
		fmt.Println("InsertMany() ERROR:", insertErr)
	}

	return insertErr
}

// FindUserByName finds a user in the MongoDB database by their name and returns it.
// It logs messages for filter, database name, and collection name.
func (m *MongoManager) FindUserByName(userName string) (result dtos.SlackUser, err error) {
	filter := bson.D{{Key: "name", Value: userName}}
	fmt.Printf("Mongo filter: %s \n", filter)
	fmt.Printf("Database: %s, colletion: %s, username: %s \n", m.DbName, m.UsersColletionName, userName)

	err = m.Client.Database(m.DbName).Collection(m.UsersColletionName).FindOne(m.Ctx, filter).Decode(&result)

	return result, err
}

// FindMessagesByUserId finds messages in the MongoDB database by the given user ID and returns them.
// It logs messages for filter, database name, and collection name.
// It also decodes and logs each result.
func (m *MongoManager) FindMessagesByUserId(userID string) (result []dtos.SlackMessage, err error) {
	filter := bson.D{{Key: "user", Value: userID}}

	fmt.Printf("Mongo filter: %s \n", filter)
	fmt.Printf("Database: %s, colletion: %s, userid: %s \n", m.DbName, m.MessagesColletionName, userID)

	cursor, err := m.Client.Database(m.DbName).Collection(m.MessagesColletionName).Find(m.Ctx, filter)

	if err != nil {
		fmt.Println("ERROR findind Messages:", err)
	}
	// end find

	var results []dtos.SlackMessage
	if err = cursor.All(m.Ctx, &results); err != nil {
		fmt.Println("ERROR parsing Messages:", err)
	}

	for _, result := range results {
		cursor.Decode(&result)
		output, err := json.MarshalIndent(result, "", "    ")
		if err != nil {
			fmt.Println("ERROR decoding Messages:", err)
		}
		fmt.Printf("%s\n", output)
	}

	return results, err
}

func ReadSlackFilesParallel(folderPath string, mongoManager MongoManager, maxParallel int) (totalRecords int) {

	folders, _ := ioutil.ReadDir(folderPath)
	fmt.Println("DB name: ", mongoManager.DbName)

	for _, folder := range folders {
		if folder.IsDir() {
			fmt.Printf("Importing %s...\n", folder.Name())
			files, err := ioutil.ReadDir(fmt.Sprintf("%s/%s", folderPath, folder.Name()))

			if err != nil {
				fmt.Println(err)
				return
			}

			var wg sync.WaitGroup
			semaphore := make(chan struct{}, maxParallel)

			if err != nil {
				fmt.Println("Erro: ", err)
				return
			}

			for _, file := range files {
				wg.Add(1)
				semaphore <- struct{}{}

				go func(file os.FileInfo) {
					defer wg.Done()

					jsonPath := fmt.Sprintf("%s/%s/%s", folderPath, folder.Name(), file.Name())
					totalRecords += PersistSlackMessages(jsonPath, folder.Name(), mongoManager)
					<-semaphore
				}(file)
			}

			wg.Wait()

		}
	}

	return totalRecords
}
