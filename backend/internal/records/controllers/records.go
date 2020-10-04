package records

import (
	"context"
	"net/http"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"github.com/roaugusto/kohobalance/config"
	dto "github.com/roaugusto/kohobalance/internal/records/dtos"
	rep "github.com/roaugusto/kohobalance/internal/records/repositories"
	ser "github.com/roaugusto/kohobalance/internal/records/services"
)

//RecordHandler a record handler
type RecordHandler struct {
	Repo rep.IRecordAccount
}

var (
	cfg config.Properties
)

func init() {
	if err := cleanenv.ReadEnv(&cfg); err != nil {
		log.Fatalf("Configuration cannot be read: %v ", err)
	}
}

// CreateRecordsFromFile godoc
// @Tags Records
// @Summary Processes Funds from a specific file
// @Description Processes Funds of customers from a specific file
// @Accept  multipart/form-data
// @Produce  text/plain
// @Param   file formData file true  "input.txt"
// @Success 200 {string} string "ok"
// @Router /api/funds [post]
func (h *RecordHandler) CreateRecordsFromFile(c echo.Context) error {

	records, err := ser.ReadRecords(c)
	if err != nil {
		return err
	}

	errValidate := ser.ValidateRecords(c, records)
	if errValidate != nil {
		return errValidate
	}

	res, errLoad := ser.LoadFunds(context.Background(), records, h.Repo, false)
	if errLoad != nil {
		return errLoad
	}
	return c.String(http.StatusCreated, res)

}

// CreateRecordsFromFileDb godoc
// @Tags Records
// @Summary Processes Funds from a specific file and write result on MongoDB
// @Description Processes Funds of customers from a specific file and write result on MongoDB
// @Accept  multipart/form-data
// @Produce  text/plain
// @Param   file formData file true  "input.txt"
// @Success 200 {string} string "ok"
// @Router /api/funds-write-result-db [post]
func (h *RecordHandler) CreateRecordsFromFileDb(c echo.Context) error {

	records, err := ser.ReadRecords(c)
	if err != nil {
		return err
	}

	errValidate := ser.ValidateRecords(c, records)
	if errValidate != nil {
		return errValidate
	}

	res, errLoad := ser.LoadFunds(context.Background(), records, h.Repo, true)
	if errLoad != nil {
		return errLoad
	}
	return c.String(http.StatusCreated, res)

}

// CreateRecordsBodyRequest godoc
// @Tags Records
// @Summary Processes Funds from body json
// @Description Processes Funds of customers from body json
// @Accept  json
// @Param   data      body records.RecordAccountList true  "List of Load Funds of customers"
// @Produce  text/plain
// @Success 200 {string} string "ok"
// @Router /api/funds-body-req [post]
func (h *RecordHandler) CreateRecordsBodyRequest(c echo.Context) error {
	var records []dto.RecordAccount

	if err := c.Bind(&records); err != nil {
		log.Errorf("Unable to bind : %v ", err)
		return echo.NewHTTPError(http.StatusBadRequest, "unable to parse request payload")
	}

	err := ser.ValidateRecords(c, records)
	if err != nil {
		return err
	}

	res, err := ser.LoadFunds(context.Background(), records, h.Repo, false)
	if err != nil {
		return err
	}

	return c.String(http.StatusCreated, res)
}

// GetFile godoc
// @Tags Records
// @Summary Downloads last result of load funds file
// @Description Downloads the last result of loading the file of Load Funds of customers
// @Produce  text/plain
// @Success 200 {string} string "ok"
// @Router /api/funds/download [get]
func (h *RecordHandler) GetFile(c echo.Context) error {
	fileN := cfg.AppHome + "/assets/files/output.txt"
	return c.Attachment(fileN, "output.txt")
}

// GetRecordsFromDB godoc
// @Tags Records
// @Summary Lists the last result of load funds file that was written on MongoDB
// @Description Lists the last result of load funds file that was written on MongoDB
// @Produce  json
// @Success 200 {object} records.RecordProcessedList
// @Router /api/funds/result [get]
func (h *RecordHandler) GetRecordsFromDB(c echo.Context) error {
	records, err := ser.FindRecords(context.Background(), c.QueryParams(), h.Repo)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, records)
}
