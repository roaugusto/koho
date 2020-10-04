package utils

import (
	"testing"

	"github.com/labstack/gommon/log"
)

func TestSaveToFile(t *testing.T) {
	t.Run("test save file", func(t *testing.T) {
		fileContent := "Test content"

		err := SaveToFile("teste.txt", fileContent)
		if err != nil {
			log.Errorf("Error to write output file: %v", err)
		}

	})

}
