package entity

import (
	"errors"
	"strings"

	"github.com/dnogueir-org/seeker/internal/common"
)

var validFieldLabels = []string{
	"colecao",
	"ofertas",
	"esportes",
	"marca",
	"linha",
	"produto",
	"categorias",
	"idade",
	"genero",
	"promocao",
	"modalidade",
	"vendidopor",
	"desconto",
	"tamanho",
	"cor",
	"promocoes",
	"preco",
	"avaliacao",
	"personalizado",
	"lojas",
	"alturadocano",
	"caimento",
	"estilotop",
	"suportetops",
	"modelos",
	"soladodachuteira",
	"niveldechuteiras",
	"tecidos",
	"tecnologia",
	"terreno",
	"times",
	"snkrs",
	"datadelancamento",
	"emestoque",
	"tendencia",
	"tipodeproduto",
	"tipodejaquetasmoletons",
	"tipodecamiseta",
	"tipodecalcasleggings",
	"tipodebolsasmochilas",
}

type Field struct {
	Label string
	Value string
}

var notValidField = errors.New("this is not a valid filter field")

func MakeField(fieldQueryParam string) (Field, error) {
	splittedField := strings.Split(fieldQueryParam, ":")
	field := Field{
		Label: splittedField[0],
		Value: splittedField[1],
	}

	isFieldValid := field.isFieldValid()
	if !isFieldValid {
		return Field{}, notValidField
	}
	return field, nil
}

func (field *Field) isFieldValid() bool {
	return common.IsStringInSlice(field.Label, validFieldLabels)
}
