package records

import (
	"log"
	"testing"

	guuid "github.com/google/uuid"
	util "github.com/roaugusto/kohobalance/utils"
	"github.com/stretchr/testify/assert"
)

func TestValidateLoad(t *testing.T) {

	t.Run("test reached a maximum of 3 charges per day.", func(t *testing.T) {

		InitializeMap()
		uuid := guuid.New()
		accepted := true
		records, _ := util.ReadRecordsFromFile("/test/input3.txt")

		for i, record := range records {
			res, _, _, err := ValidateLoadRecord(record, uuid.String())
			if err != nil {
				log.Fatalf("Error validating maximum 3 charges per day: %v ", err)
			}

			if i == 3 {
				accepted = res.Accepted
			}
		}

		assert.Equal(t, false, accepted)

	})

	t.Run("test a maximum load of $5,000 per day.", func(t *testing.T) {

		InitializeMap()
		uuid := guuid.New()
		accepted := true
		records, _ := util.ReadRecordsFromFile("/test/input4.txt")

		for i, record := range records {
			res, _, _, err := ValidateLoadRecord(record, uuid.String())
			if err != nil {
				log.Fatalf("Error validating maximum load of $5,000 per day.: %v ", err)
			}

			if i == 0 {
				accepted = res.Accepted
			}
		}

		assert.Equal(t, false, accepted)

	})

	t.Run("test Load ID more than once for a customer is not acceptable.", func(t *testing.T) {

		InitializeMap()
		uuid := guuid.New()
		accepted := true
		records, _ := util.ReadRecordsFromFile("/test/input5.txt")

		for i, record := range records {
			res, _, _, err := ValidateLoadRecord(record, uuid.String())
			if err != nil {
				log.Fatalf("Error validating Load ID more than once for a customer: %v ", err)
			}

			if i == 1 {
				accepted = res.Accepted
			}
		}

		assert.Equal(t, false, accepted)

	})

	t.Run("test a maximum load of $20,000 per week.", func(t *testing.T) {

		InitializeMap()
		uuid := guuid.New()
		accepted := true
		records, _ := util.ReadRecordsFromFile("/test/input6.txt")

		for i, record := range records {
			res, _, _, err := ValidateLoadRecord(record, uuid.String())
			if err != nil {
				log.Fatalf("Error validating maximum load of $20,000 per week.: %v ", err)
			}

			if i == 3 {
				accepted = res.Accepted
			}
		}

		assert.Equal(t, false, accepted)

	})

}
