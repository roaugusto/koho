package records

import (
	"bufio"
	"encoding/json"

	"github.com/labstack/echo/v4"
	dto "github.com/roaugusto/kohobalance/internal/records/dtos"
)

//ReadRecords - Read each line from file
func ReadRecords(c echo.Context) ([]dto.RecordAccount, error) {
	var records []dto.RecordAccount

	file, err := c.FormFile("file")
	if err != nil {
		return records, err
	}
	src, err := file.Open()
	if err != nil {
		return records, err
	}
	defer src.Close()

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
