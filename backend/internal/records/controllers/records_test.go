package records

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http/httptest"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	config "github.com/roaugusto/kohobalance/config"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

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

func TestCreateRecordsBodyRequest(t *testing.T) {
	t.Run("test load records from request body", func(t *testing.T) {
		body := `
		[{"id":"456","customer_id":"123","load_amount":"$123.45","time":"2000-01-01T00:00:00Z"}]`
		req := httptest.NewRequest("POST", "/api/funds", strings.NewReader(body))
		res := httptest.NewRecorder()
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		e := echo.New()
		c := e.NewContext(req, res)

		rec := RecordHandler{Repo: col}

		err := rec.CreateRecordsBodyRequest(c)
		assert.Nil(t, err)

	})

}

func TestCreateRecordsFromFile(t *testing.T) {
	t.Run("test load records from a specific file", func(t *testing.T) {

		path, err := filepath.Abs("../../../test/input.txt")
		if err != nil {
			t.Error(err)
		}

		file, err := os.Open(path)
		if err != nil {
			t.Error(err)
		}

		defer file.Close()
		body := &bytes.Buffer{}
		writer := multipart.NewWriter(body)
		part, err := writer.CreateFormFile("file", filepath.Base(path))
		if err != nil {
			writer.Close()
			t.Error(err)
		}
		io.Copy(part, file)
		writer.Close()

		req := httptest.NewRequest("POST", "/upload", body)
		req.Header.Set("Content-Type", writer.FormDataContentType())
		res := httptest.NewRecorder()

		e := echo.New()
		c := e.NewContext(req, res)

		rec := RecordHandler{Repo: col}

		err = rec.CreateRecordsFromFile(c)
		assert.Nil(t, err)

	})

}

func TestConvertDate(t *testing.T) {
	str := "2014-11-12T11:45:26.371Z"
	date, err := time.Parse(time.RFC3339, str)
	if err != nil {
		t.Error(err)
	}

	year, week := date.ISOWeek()
	fmt.Println(year, week)

}
