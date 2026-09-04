package tree_agent

import (
	"fmt"
	"strings"

	"github.com/tealeg/xlsx"
	"ril.api-ia/internal/domain/entity"
)

func parseExcelContent(fileBytes []byte) ([]entity.QuestionTree, []string, []string, error) {
	xlFile, err := xlsx.OpenBinary(fileBytes)
	if err != nil {
		return nil, nil, nil, err
	}

	if len(xlFile.Sheets) == 0 {
		return nil, nil, nil, fmt.Errorf("el excel no tiene hojas")
	}
	sheet := xlFile.Sheets[0]

	var nuevasPreguntas []entity.QuestionTree
	dimMap := make(map[string]bool)
	tagMap := make(map[string]bool)

	for i, row := range sheet.Rows {
		// Saltamos las 4 filas de headers
		if i < 4 {
			continue
		}

		if len(row.Cells) == 0 || row.Cells[0].String() == "" {
			continue
		}

		getCell := func(idx int) string {
			if idx < len(row.Cells) {
				return strings.TrimSpace(row.Cells[idx].String())
			}
			return ""
		}

		// Armamos tags (separados por coma o "·")
		tagsRaw := getCell(15)
		var tagsRag []string
		if tagsRaw != "" {
			tagsCleaned := strings.NewReplacer(
				"·", ",",
				"|", ",",
			).Replace(tagsRaw)
			for _, t := range strings.Split(tagsCleaned, ",") {
				t = strings.TrimSpace(t)
				if t != "" {
					tagsRag = append(tagsRag, t)
					tagMap[t] = true
				}
			}
		}

		dimension := getCell(1)
		if dimension != "" {
			dimMap[dimension] = true
		}

		pregunta := entity.QuestionTree{
			ID:        getCell(0),
			Dimension: dimension,
			Pregunta:  getCell(2),
			Opciones:  getCell(3),
			Bajo: entity.Level{
				Descripcion: getCell(4),
				ODM:         getCell(5),
				Acciones:    getCell(6),
			},
			Intermedio: entity.Level{
				Descripcion: getCell(7),
				ODM:         getCell(8),
				Acciones:    getCell(9),
			},
			Avanzado: entity.Level{
				Descripcion: getCell(10),
				ODM:         getCell(11),
				Acciones:    getCell(12),
			},
			AjusteChica:     getCell(13),
			AjusteMedianaGr: getCell(14),
			TagsRAG:         tagsRag,
			EsAncla:         strings.ToLower(getCell(16)) == "sí" || strings.ToLower(getCell(16)) == "si",
			Padre:           strings.TrimPrefix(getCell(17), "P"),
			NotaFacilitador: getCell(18),
		}

		nuevasPreguntas = append(nuevasPreguntas, pregunta)
	}

	// Convertir mapas a slices
	var dims []string
	for d := range dimMap {
		dims = append(dims, d)
	}

	var tags []string
	for t := range tagMap {
		tags = append(tags, t)
	}

	return nuevasPreguntas, dims, tags, nil
}
