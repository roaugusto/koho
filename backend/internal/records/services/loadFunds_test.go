package records

import (
	"context"
	"fmt"
	"log"
	"testing"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/roaugusto/kohobalance/config"
	rep "github.com/roaugusto/kohobalance/internal/records/repositories"
	util "github.com/roaugusto/kohobalance/utils"
	"github.com/stretchr/testify/assert"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//RecordHandler a record handler
type RecordHandler struct {
	Repo rep.IRecordAccount
}

var (
	c   *mongo.Client
	db  *mongo.Database
	col *mongo.Collection
	cfg config.PropertiesTest
	h   RecordHandler
)

func init() {
	if err := cleanenv.ReadEnv(&cfg); err != nil {
		log.Fatalf("Configuration cannot be read: %v ", err)
	}

	connectURI := fmt.Sprintf("mongodb://%s:%s", cfg.DBHost, cfg.DBPort)
	c, err := mongo.Connect(context.Background(), options.Client().ApplyURI(connectURI))
	if err != nil {
		log.Fatalf("Enable to connect to database: %v ", err)
	}

	db = c.Database(cfg.DBName)
	col = db.Collection(cfg.CollectionName)
}

func TestLoadFunds(t *testing.T) {

	t.Run("test load funds", func(t *testing.T) {

		records, _ := util.ReadRecordsFromFile("/test/input2.txt")

		rec := RecordHandler{Repo: col}

		_, errLoad := LoadFunds(context.Background(), records, rec.Repo, true)
		if errLoad != nil {
			log.Fatalf("Error on LoadFunds: %v ", errLoad)
		}

		assert.Nil(t, errLoad)

	})
}
