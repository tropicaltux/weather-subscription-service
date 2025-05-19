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
	Token     string                `gorm:"uniqueIndex;not null"`
	Confirmed bool                  `gorm:"default:false"`
}
