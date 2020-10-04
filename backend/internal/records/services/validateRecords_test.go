package records

import (
	"log"
	"net/http/httptest"
	"strings"
	"testing"

	util "github.com/roaugusto/kohobalance/utils"
	"github.com/stretchr/testify/assert"

	"github.com/labstack/echo/v4"
)

func TestValidateRecords(t *testing.T) {

	t.Run("test validation records", func(t *testing.T) {

		records, _ := util.ReadRecordsFromFile("/test/input2.txt")
		body := ""
		req := httptest.NewRequest("POST", "/api/funds", strings.NewReader(body))
		res := httptest.NewRecorder()
		e := echo.New()
		c := e.NewContext(req, res)

		errValidate := ValidateRecords(c, records)

		if errValidate != nil {
			log.Fatalf("Error on ValidateRecords: %v ", errValidate)
		}

		assert.Nil(t, errValidate)

	})
}
