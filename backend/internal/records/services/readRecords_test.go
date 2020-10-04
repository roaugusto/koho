package records

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

type recordAccount struct {
	ID         string    `json:"id" bson:"id" example:"1234"`
	CustomerID string    `json:"customer_id" bson:"customer_id" example:"4567"`
	LoadAmount string    `json:"load_amount" bson:"load_amount" example:"$3456.78"`
	Time       time.Time `json:"time" bson:"time" example:"2000-01-01T00:00:00Z"`
}

func TestReadRecords(t *testing.T) {

	t.Run("test reading records from file", func(t *testing.T) {
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

		records, err := ReadRecords(c)

		res2, _ := json.Marshal(records)
		if err != nil {
			t.Error(err)
		}

		assert.Equal(t, `[{"id":"15887","customer_id":"528","load_amount":"$3318.47","time":"2000-01-01T00:00:00Z"}]`,
			fmt.Sprint(string(res2)))
	})
}
