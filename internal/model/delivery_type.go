package model

import (
	"errors"
	"fmt"
	"time"
)

type DeliveryTypeCode string

const (
	DELIVERY_TYPE_CODE_DELIVERY DeliveryTypeCode = "delivery"
	DELIVERY_TYPE_CODE_PICKUP   DeliveryTypeCode = "pickup"
)

func (c DeliveryTypeCode) Validate() error {
	switch c {
	case DELIVERY_TYPE_CODE_DELIVERY, DELIVERY_TYPE_CODE_PICKUP:
		return nil
	default:
		return errors.New("incorrect delivery type")
	}
}

type DeliveryType struct {
	CreatedAt        time.Time
	UpdatedAt        time.Time
	DeliveryTypeID   uint             `gorm:"primaryKey;column:delivery_type_id" json:"delivery_type_id"`
	DeliveryTypeName string           `gorm:"column:delivery_type_name" json:"delivery_type_name"`
	DeliveryTypeCode DeliveryTypeCode `gorm:"column:delivery_type_code" json:"delivery_type_code"`
}

func (d DeliveryType) TableName() string {
	return fmt.Sprintf("%s.%s", "public", "delivery_type")
}
