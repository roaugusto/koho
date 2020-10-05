package records

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"strings"
	"testing"

	config "github.com/roaugusto/kohobalance/config"
	dto "github.com/roaugusto/kohobalance/internal/records/dtos"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	c       *mongo.Client
	db      *mongo.Database
	col     *mongo.Collection
	cfgTest config.PropertiesTest
	h       RecordHandler
)

func init() {
	if err := cleanenv.ReadEnv(&cfgTest); err != nil {
		log.Fatalf("Configuration cannot be read: %v ", err)
	}

	connectURI := fmt.Sprintf("mongodb://%s:%s", cfgTest.DBHost, cfgTest.DBPort)
	c, err := mongo.Connect(context.Background(), options.Client().ApplyURI(connectURI))
	if err != nil {
		log.Fatalf("Enable to connect to database: %v ", err)
	}

	db = c.Database(cfgTest.DBName)
	col = db.Collection(cfgTest.CollectionName)
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

func TestGetFile(t *testing.T) {
	t.Run("test getting download result data after load", func(t *testing.T) {

		path, err := filepath.Abs("../../../test/input1.txt")
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
		err = rec.CreateRecordsFromFileDb(c)

		var resLoad dto.LoadFundsResponse
		err = json.Unmarshal(res.Body.Bytes(), &resLoad)
		assert.Nil(t, err)

		fmt.Println(resLoad)
		req2 := httptest.NewRequest("GET", fmt.Sprintf("/api/funds/download?uuid_file=%s", resLoad.ProcessID), nil)
		res2 := httptest.NewRecorder()
		req2.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		e2 := echo.New()
		c2 := e2.NewContext(req2, res2)

		err = rec.GetFile(c2)
		assert.Nil(t, err)
		assert.Equal(t, http.StatusOK, res2.Code)

		result := string(res2.Body.Bytes())
		//err = json.Unmarshal(res2.Body.Bytes(), &records)
		fmt.Println(result)
		assert.Equal(t, fmt.Sprintln(`{"id":"15887","customer_id":"528","accepted":true}`), result)

	})

}

func TestGetRecordsFromDB(t *testing.T) {
	t.Run("test getting records from MongoDB", func(t *testing.T) {

		path, err := filepath.Abs("../../../test/input1.txt")
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
		err = rec.CreateRecordsFromFileDb(c)

		var resLoad dto.LoadFundsResponse
		err = json.Unmarshal(res.Body.Bytes(), &resLoad)
		assert.Nil(t, err)

		req2 := httptest.NewRequest("GET", fmt.Sprintf("/api/funds/result?process_id=%s", resLoad.ProcessID), nil)

		res2 := httptest.NewRecorder()
		req2.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		e2 := echo.New()
		c2 := e2.NewContext(req2, res2)

		err = rec.GetRecordsFromDB(c2)
		assert.Nil(t, err)
		assert.Equal(t, http.StatusOK, res2.Code)

		var records dto.RecordProcessedList

		err = json.Unmarshal(res2.Body.Bytes(), &records)
		assert.Nil(t, err)
		for _, record := range records {
			assert.Equal(t, "15887", record.ID)
		}

	})

}
