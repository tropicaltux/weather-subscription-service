package models

import (
	"gorm.io/gorm"
)

// SubscriptionFrequency defines the frequency of weather updates
type SubscriptionFrequency string

const (
	FrequencyHourly SubscriptionFrequency = "hourly"
	FrequencyDaily  SubscriptionFrequency = "daily"
)

// Subscription represents a user subscription for weather updates
type Subscription struct {
	gorm.Model
	ID        string                `gorm:"primaryKey;type:uuid"`
	Email     string                `gorm:"uniqueIndex;not null"`
	City      string                `gorm:"not null"`
	Frequency SubscriptionFrequency `gorm:"type:varchar(10);not null"`
	Token     string                `gorm:"uniqueIndex;not null;type:varchar(128)"` // 128-character random string
	Confirmed bool                  `gorm:"default:false"`
}
