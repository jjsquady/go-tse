package extractors

import (
	"strconv"
	"strings"

	"go-tse/models"
)

type PartidoExtractor struct{}

func (e *PartidoExtractor) Extract(record []string) *models.Partido {

	csvPID, _ := strconv.Atoi(strings.Trim(record[18], " "))
	csvSign := strings.Trim(record[19], " ")
	csvPName := strings.Trim(record[20], " ")
	csvTipoVoto := strings.Trim(record[28], " ")

	if csvTipoVoto == "Nominal" || csvTipoVoto == "Legenda" {
		return &models.Partido{ID: csvPID, Sign: csvSign, Name: csvPName}
	}

	return nil
}
