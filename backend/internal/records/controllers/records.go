package records

import (
	"context"
	"fmt"
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
// @Produce  json
// @Param   file formData file true  "input.txt"
// @Success 200 {string} string "ok"
// @Success 200 {object} records.LoadFundsResponse
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
	return c.JSON(http.StatusCreated, res)

}

// CreateRecordsFromFileDb godoc
// @Tags Records
// @Summary Processes Funds from a specific file and write result on MongoDB
// @Description Processes Funds of customers from a specific file and write result on MongoDB
// @Accept  multipart/form-data
// @Produce  json
// @Param   file formData file true  "input.txt"
// @Success 200 {object} records.LoadFundsResponse
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
	return c.JSON(http.StatusCreated, res)

}

// CreateRecordsBodyRequest godoc
// @Tags Records
// @Summary Processes Funds from body json
// @Description Processes Funds of customers from body json
// @Accept  json
// @Param   data      body records.RecordAccountList true  "List of Load Funds of customers"
// @Produce json
// @Success 200 {object} records.LoadFundsResponse
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

	return c.JSON(http.StatusCreated, res)
}

// GetFile godoc
// @Tags Records
// @Summary Downloads last result of load funds file
// @Description Downloads the last result of loading the file of Load Funds of customers
// @Produce  text/plain
// @Param   uuid_file		query  string  true   "Process ID"
// @Success 200 {file} file "A txt file"
// @Router /api/funds/download [get]
func (h *RecordHandler) GetFile(c echo.Context) error {
	fileName := c.QueryParam("uuid_file")
	if fileName == "" {
		return c.JSON(http.StatusBadRequest, "uuid_file needs to be informed")
	}

	fileN := fmt.Sprintf("%s/assets/files/%v.txt", cfg.AppHome, fileName)
	return c.Attachment(fileN, "output.txt")
}

// GetRecordsFromDB godoc
// @Tags Records
// @Summary Lists the last result of load funds file that was written on MongoDB
// @Description Lists the last result of load funds file that was written on MongoDB
// @Produce  json
// @Param   process_id		query  string  true   "Process ID"
// @Param   id     				query  string  false   "Transaction ID"
// @Param   customer_id   query  string  false   "Customer ID"
// @Param   accepted      query  string  false   "Accepted"
// @Param   cod_error     query  string  false   "Error Code: 10"
// @Success 200 {object} records.RecordProcessedList
// @Router /api/funds/result [get]
func (h *RecordHandler) GetRecordsFromDB(c echo.Context) error {
	fileName := c.QueryParam("process_id")
	if fileName == "" {
		return c.JSON(http.StatusBadRequest, "process_id needs to be informed")
	}

	records, err := ser.FindRecords(context.Background(), c.QueryParams(), h.Repo)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, records)
}
