package records

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"go.mongodb.org/mongo-driver/bson"

	dto "github.com/roaugusto/kohobalance/internal/records/dtos"
	rep "github.com/roaugusto/kohobalance/internal/records/repositories"
	util "github.com/roaugusto/kohobalance/utils"
)

//LoadFunds - Responsible to load a record account
func LoadFunds(ctx context.Context, records []dto.RecordAccount, repo rep.IRecordAccount, writeDb bool) (string, *echo.HTTPError) {

	//var results []interface{}
	var results string

	InitializeMap()
	// location, err := time.LoadLocation("")

	repo.DeleteMany(ctx, bson.M{})

	for _, record := range records {

		res, resDb, _, err := ValidateLoadRecord(record)
		if err != nil {
			return "", echo.NewHTTPError(http.StatusInternalServerError, "unable to validate the record")
		}

		if writeDb {
			_, errInsert := repo.InsertOne(ctx, resDb)
			if errInsert != nil {
				log.Errorf("Unable to insert : %v", errInsert)
				return "nil", echo.NewHTTPError(http.StatusInternalServerError, "unable to insert to database")
			}
		}

		res2, _ := json.Marshal(res)
		results += string(res2) + "\n"
	}

	err := util.SaveToFile("output.txt", results)
	if err != nil {
		log.Errorf("Error to write output file: %v", err)
		return "nil", echo.NewHTTPError(http.StatusInternalServerError, "unable to write the output file")
	}

	return results, nil
}
