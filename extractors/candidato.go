package extractors

import (
	"fmt"
	"strconv"
	"strings"

	"go-tse/models"
)

type CandidatoExtractor struct{}

func (e *CandidatoExtractor) Extract(record []string) *models.Candidato {
	fmt.Println(record)
	csvID, _ := strconv.Atoi(strings.Trim(record[29], " "))
	csvName := strings.Trim(record[30], " ")
	csvCargo := strings.Trim(record[17], " ")

	csvTipoVoto := strings.Trim(record[28], " ")

	if csvTipoVoto == "Nominal" {
		return &models.Candidato{ID: csvID, Name: csvName, Cargo: csvCargo}
	}
	return nil
}
