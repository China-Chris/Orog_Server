// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameMarketKline15m = "market_kline_15m"

// MarketKline15m mapped from table <market_kline_15m>
type MarketKline15m struct {
	MarketAddress string    `gorm:"column:market_address" json:"market_address"`
	O             float64   `gorm:"column:o" json:"o"`
	H             float64   `gorm:"column:h" json:"h"`
	L             float64   `gorm:"column:l" json:"l"`
	C             float64   `gorm:"column:c" json:"c"`
	V             float64   `gorm:"column:v" json:"v"`
	Timestamp     time.Time `gorm:"column:timestamp" json:"timestamp"`
	UpdatedAt     time.Time `gorm:"column:updated_at" json:"updated_at"`
	CreatedAt     time.Time `gorm:"column:created_at" json:"created_at"`
}

// TableName MarketKline15m's table name
func (*MarketKline15m) TableName() string {
	return TableNameMarketKline15m
}
