package worker

import (
	"context"
	"fmt"
	"work-project/internal/repository"
)

type AirTableSync struct {
	airTable repository.AirTable
}

func NewAirTableSync(airTable repository.AirTable) *AirTableSync {
	return &AirTableSync{airTable: airTable}
}

func (h *AirTableSync) Run() (err error) {
	ctx := context.Background()

	products, err := h.airTable.GetProducts(ctx)
	if err != nil {
		return err
	}
	fmt.Println(products)

	return
}
