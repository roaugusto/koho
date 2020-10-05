package records

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"

	guuid "github.com/google/uuid"
	dto "github.com/roaugusto/kohobalance/internal/records/dtos"
	rep "github.com/roaugusto/kohobalance/internal/records/repositories"
	util "github.com/roaugusto/kohobalance/utils"
)

//LoadFunds - Responsible to load a record account
func LoadFunds(ctx context.Context, records []dto.RecordAccount, repo rep.IRecordAccount, writeDb bool) (dto.LoadFundsResponse, *echo.HTTPError) {

	var results string
	resLoad := dto.LoadFundsResponse{
		ProcessID: "",
	}

	InitializeMap()
	uuid := guuid.New()

	//repo.DeleteMany(ctx, bson.M{})
	for _, record := range records {

		res, resDb, _, err := ValidateLoadRecord(record, uuid.String())
		if err != nil {
			return resLoad, echo.NewHTTPError(http.StatusInternalServerError, "unable to validate the record")
		}

		if writeDb {
			_, errInsert := repo.InsertOne(ctx, resDb)
			if errInsert != nil {
				log.Errorf("Unable to insert : %v", errInsert)
				return resLoad, echo.NewHTTPError(http.StatusInternalServerError, "unable to insert to database")
			}
		}

		res2, _ := json.Marshal(res)
		results += string(res2) + "\n"
	}

	fileN := fmt.Sprintf("%s.txt", uuid.String())
	err := util.SaveToFile(fileN, results)
	if err != nil {
		log.Errorf("Error to write output file: %v", err)
		return resLoad, echo.NewHTTPError(http.StatusInternalServerError, "unable to write the output file")
	}

	resLoad = dto.LoadFundsResponse{
		ProcessID: uuid.String(),
	}
	//return results, nil
	return resLoad, nil
}
