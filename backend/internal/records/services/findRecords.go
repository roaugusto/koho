package records

import (
	"context"
	"net/http"
	"net/url"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"

	dto "github.com/roaugusto/kohobalance/internal/records/dtos"
	rep "github.com/roaugusto/kohobalance/internal/records/repositories"
)

//FindRecords Returns the result of the processing file with the load funds of clients
func FindRecords(ctx context.Context, q url.Values, repo rep.IRecordAccount) ([]dto.RecordProcessed, error) {
	var records []dto.RecordProcessed

	filter := make(map[string]interface{})

	for k, v := range q {
		filter[k] = v[0]
	}

	cursor, err := repo.Find(ctx, filter)
	if err != nil {
		log.Errorf("Unable to find the records: %v", err)
		return records, echo.NewHTTPError(http.StatusNotFound, "unable to find records")
	}
	err = cursor.All(ctx, &records)
	if err != nil {
		log.Errorf("Unable to read the cursor: %v", err)
		return records, echo.NewHTTPError(http.StatusInternalServerError, "unable to parse retrieve records")
	}
	return records, nil

}
