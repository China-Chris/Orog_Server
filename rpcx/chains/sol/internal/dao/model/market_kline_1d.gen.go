// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameMarketKline1d = "market_kline_1d"

// MarketKline1d mapped from table <market_kline_1d>
type MarketKline1d struct {
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

// TableName MarketKline1d's table name
func (*MarketKline1d) TableName() string {
	return TableNameMarketKline1d
}
