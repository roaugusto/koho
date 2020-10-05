package records

import (
	"strconv"
	"strings"
	"time"

	"github.com/labstack/gommon/log"
	dto "github.com/roaugusto/kohobalance/internal/records/dtos"
	util "github.com/roaugusto/kohobalance/utils"
)

//RecordBalance Information about the balance of a customer
type RecordBalance struct {
	customerID   string
	dayBalance   float64
	weekBalance  float64
	dayCount     int
	lastLoadDate time.Time
	loadIDs      []string
}

var (
	m = map[string]RecordBalance{}
)

//ValidateLoadRecord - Responsible to validate the business rules of the record
func ValidateLoadRecord(record dto.RecordAccount, uuid string) (*dto.RecordsResponse, *dto.RecordProcessed, bool, error) {

	var cID = record.CustomerID
	var currentBalance = m[cID]

	// When the custumor exist in the memory map, the validation needs to
	// verify the business rules
	if currentBalance.customerID != "" {

		ret := existingCustomerLoadBalance(record)
		if ret != "00" {
			data, dataDb := formatReturnData(uuid, record, ret, false)
			return data, dataDb, false, nil
		}

	} else {

		// First time that customer was found in the process
		// Create a reference for this new customer
		ret := newCustomerLoadBalance(record)
		if ret != "00" {
			data, dataDb := formatReturnData(uuid, record, ret, false)
			return data, dataDb, false, nil
		}

	}

	// Record was validate with success!
	data, dataDb := formatReturnData(uuid, record, "00", true)
	return data, dataDb, true, nil
}

//newCustomerLoadBalance - Responsible to create a new reference of the customer in memory map variable
func newCustomerLoadBalance(record dto.RecordAccount) string {
	const maxLoadAmountDay = 5000
	var newBalance RecordBalance
	var cID = record.CustomerID

	loadAmount, err := parseLoadAmount(record.LoadAmount)
	if err != nil {
		return "99"
	}

	newBalance.customerID = cID
	newBalance.dayBalance = loadAmount
	newBalance.weekBalance = loadAmount
	newBalance.lastLoadDate = record.Time
	newBalance.dayCount = 1
	newBalance.loadIDs = append(newBalance.loadIDs, record.ID)

	if loadAmount > maxLoadAmountDay {
		return "30"
	}

	m[cID] = newBalance

	return "00"

}

//existingCustomerLoadBalance - Responsible to update information of the customer in memory map variable
func existingCustomerLoadBalance(record dto.RecordAccount) string {

	var cID = record.CustomerID
	var currentBalance = m[cID]

	loadAmount, err := parseLoadAmount(record.LoadAmount)
	if err != nil {
		return "99"
	}

	// Business rule:
	// If a load ID is observed more than once for a particular user,
	// all but the first instance can be ignored
	// Return codes:
	// 10 - The Load ID more than once for a customer is not acceptable.
	_, found := util.Find(currentBalance.loadIDs, record.ID)
	if found == true {
		return "10"
	}

	// Business rule:
	// A maximum of $20,000 can be loaded per week
	calculateLoadAmountWeek, ret := validationWeekLoad(currentBalance.lastLoadDate, record.Time, currentBalance.weekBalance, loadAmount)
	if ret != "00" {
		return ret
	}

	// Business rule:
	// A maximum of $5,000 can be loaded per day
	// A maximum of 3 loads can be performed per day, regardless of amount
	calculateLoadAmountDay, dayCount, ret := validationDayLoad(currentBalance.lastLoadDate, record.Time, currentBalance.dayBalance, loadAmount, currentBalance.dayCount)
	if ret != "00" {
		return ret
	}

	// Updating information about the customer load amount
	currentBalance.weekBalance = calculateLoadAmountWeek
	currentBalance.dayBalance = calculateLoadAmountDay
	currentBalance.dayCount = dayCount

	currentBalance.loadIDs = append(currentBalance.loadIDs, record.ID)
	currentBalance.lastLoadDate = record.Time
	m[cID] = currentBalance

	return "00"
}

//parseLoadAmount - Responsible to parse the loadAmount value
func parseLoadAmount(loadAmount string) (float64, error) {
	loadAmountParsed, err := strconv.ParseFloat(strings.ReplaceAll(loadAmount, "$", ""), 64)
	if err != nil {
		log.Errorf("Unable to get the loadAmount ")
		return 0, err
	}
	return loadAmountParsed, nil
}

//validationWeekLoad - Validation of the week amount load, based on the business rules
func validationWeekLoad(lastDate time.Time, dateRecord time.Time, weekBalance float64, loadAmount float64) (float64, string) {
	// Return codes:
	// 40 - The record reached a maximum load of $20,000 per week.

	const maxLoadAmountWeek = 20000

	weekC := util.GetWeekYear(lastDate)
	weekNew := util.GetWeekYear(dateRecord)

	calculateLoadAmountWeek := util.RoundUp(loadAmount, 2)
	if weekC == weekNew {
		calculateLoadAmountWeek = util.RoundUp((weekBalance + loadAmount), 2)

		// Business rule:
		// A maximum of $20,000 can be loaded per week
		if calculateLoadAmountWeek > maxLoadAmountWeek {
			return 0, "40"
		}
	}
	return calculateLoadAmountWeek, "00"
}

//validationDayLoad - Validation of the day amount load, based on the business rules
func validationDayLoad(lastDate time.Time, dateRecord time.Time, dayBalance float64, loadAmount float64, dayCount int) (float64, int, string) {
	// Return codes:
	// 20 - The record reached a maximum of 3 charges per day.
	// 30 - The record reached a maximum load of $5,000 per day.

	const maxCountLoadDay = 3
	const maxLoadAmountDay = 5000
	var countD = dayCount

	calculateLoadAmountDay := util.RoundUp(loadAmount, 2)

	yearC, monthC, dayC := lastDate.Date()
	yearNew, monthNew, dayNew := dateRecord.Date()

	if (yearC == yearNew) &&
		(monthC == monthNew) &&
		(dayC == dayNew) {

		// Business rule:
		// A maximum of 3 loads can be performed per day, regardless of amount
		if dayCount == maxCountLoadDay {
			return 0, countD, "20"
		}

		calculateLoadAmountDay = util.RoundUp((dayBalance + loadAmount), 2)
		countD++
	} else {
		countD = 1
	}

	// Business rule:
	// A maximum of $5,000 can be loaded per day
	if calculateLoadAmountDay > maxLoadAmountDay {
		return 0, 0, "30"
	}

	return calculateLoadAmountDay, countD, "00"
}

//formatReturnData - Responsible to format the return date of the function validateLoadRecord
func formatReturnData(uuid string, record dto.RecordAccount, codErr string, accepted bool) (*dto.RecordsResponse, *dto.RecordProcessed) {

	var message string
	var ac string

	if accepted {
		ac = "true"
	} else {
		ac = "false"
	}

	switch codErr {
	case "00":
		message = ""
	case "10":
		message = "The Load ID more than once for a customer is not acceptable."
	case "20":
		message = "The record reached a maximum of 3 charges per day."
	case "30":
		message = "The record reached a maximum load of $5,000 per day."
	case "40":
		message = "The record reached a maximum load of $20,000 per week."
	default:
		message = "Processing error or wrong input layout."
	}

	res := &dto.RecordsResponse{
		ID:         record.ID,
		CustomerID: record.CustomerID,
		Accepted:   accepted,
	}

	res2 := &dto.RecordProcessed{
		ProcessID:  uuid,
		ID:         record.ID,
		CustomerID: record.CustomerID,
		LoadAmount: record.LoadAmount,
		Time:       record.Time,
		Accepted:   ac,
		CodError:   codErr,
		Message:    message,
	}

	return res, res2
}

//InitializeMap Responsible to initialize the memory map variable
func InitializeMap() {
	m = map[string]RecordBalance{}
}
