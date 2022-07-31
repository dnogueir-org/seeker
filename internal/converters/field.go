package converters

import (
	"fmt"

	"github.com/dnogueir-org/seeker/internal/entity"
)

func ConvertFilterField(filterFields []string) []entity.Field {
	var fields []entity.Field

	for _, ff := range filterFields {
		field, err := entity.MakeField(ff)
		if err != nil {
			fmt.Println("filter not recognized: ", ff)
			continue
		}
		fields = append(fields, field)
	}

	return fields
}
