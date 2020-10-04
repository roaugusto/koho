package records

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/url"
	"testing"

	"github.com/ilyakaznacheev/cleanenv"
	util "github.com/roaugusto/kohobalance/utils"
	"github.com/stretchr/testify/assert"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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

func TestFindRecords(t *testing.T) {

	t.Run("test find records loaded", func(t *testing.T) {

		records, _ := util.ReadRecordsFromFile("/test/input1.txt")

		rec := RecordHandler{Repo: col}

		_, errLoad := LoadFunds(context.Background(), records, rec.Repo, true)
		if errLoad != nil {
			log.Fatalf("Error on LoadFunds: %v ", errLoad)
		}

		var q url.Values

		res, errFind := FindRecords(context.Background(), q, rec.Repo)
		if errFind != nil {
			log.Fatalf("Error on FindRecords: %v ", errFind)
		}

		res2, _ := json.Marshal(res)

		assert.Equal(t, `[{"id":"15887","customer_id":"528","load_amount":"$3318.47","time":"2000-01-01T00:00:00Z","accepted":"true","cod_error":"00","message":""}]`,
			fmt.Sprint(string(res2)))

	})
}
