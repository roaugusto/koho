package utils

import (
	"bufio"
	"encoding/json"
	"io/ioutil"
	"log"
	"math"
	"os"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
	config "github.com/roaugusto/kohobalance/config"
	dto "github.com/roaugusto/kohobalance/internal/records/dtos"
)

var (
	cfg config.Properties
)

func init() {
	if err := cleanenv.ReadEnv(&cfg); err != nil {
		log.Fatalf("Configuration cannot be read: %v ", err)
	}
}

//RoundUp Function to RoundUp a value to a specific precision
func RoundUp(input float64, places int) (newVal float64) {
	var round float64
	pow := math.Pow(10, float64(places))
	digit := pow * input
	round = math.Ceil(digit)
	newVal = round / pow
	return
}

//Find Find a specific value in slice
func Find(slice []string, val string) (int, bool) {
	for i, item := range slice {
		if item == val {
			return i, true
		}
	}
	return -1, false
}

//GetWeekYear Return the number of the year week, where monday it's the first day
func GetWeekYear(valDate time.Time) int {
	date := firstDayOfISOWeek(valDate)
	_, isoWeek := date.ISOWeek()
	return isoWeek
}

func firstDayOfISOWeek(valDate time.Time) time.Time {
	date := valDate
	for date.Weekday() != time.Monday { // iterate back to Monday
		date = date.AddDate(0, 0, -1)
	}

	return date
}

//SaveToFile - Save data to specific file
func SaveToFile(filename string, data string) error {
	path := cfg.AppHome
	fileN := path + "/Assets/files/" + filename
	return ioutil.WriteFile(fileN, []byte(data), 0666)
}

//ReadRecordsFromFile - Read records data from specific file
func ReadRecordsFromFile(filename string) ([]dto.RecordAccount, error) {

	var records []dto.RecordAccount

	path := cfg.AppHome
	fileN := path + filename

	src, err := os.OpenFile(fileN, os.O_RDONLY, 0644)
	if err != nil {
		return records, err
	}

	contentFile := bufio.NewScanner(src)
	for contentFile.Scan() {
		rawIn := json.RawMessage(contentFile.Text())
		bytes, err := rawIn.MarshalJSON()
		if err != nil {
			return records, err
		}

		var r dto.RecordAccount
		err = json.Unmarshal(bytes, &r)
		if err != nil {
			return records, err
		}

		records = append(records, r)
	}

	return records, nil

}
