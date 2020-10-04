package records

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	dto "github.com/roaugusto/kohobalance/internal/records/dtos"
	"gopkg.in/go-playground/validator.v9"
)

var (
	v = validator.New()
)

//RecordValidator a record validator
type RecordValidator struct {
	validator *validator.Validate
}

//Validate validates a product
func (r *RecordValidator) Validate(i interface{}) error {
	return r.validator.Struct(i)
}

//ValidateRecords Realize a validation of the data
func ValidateRecords(c echo.Context, records []dto.RecordAccount) *echo.HTTPError {

	c.Echo().Validator = &RecordValidator{validator: v}

	for _, record := range records {
		if err := c.Validate(record); err != nil {
			log.Errorf("Unable to validate the records %+v %v", record, err)
			return echo.NewHTTPError(http.StatusBadRequest, "unable to validate request payload")
		}
	}

	return nil
}
