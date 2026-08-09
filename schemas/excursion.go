package schemas

import (
	"time"

	"gorm.io/gorm"
)

// Excursion represents an excursion with fields for image, title, description, buy link, and find more link.
type Excursion struct {
  gorm.Model

  ExcursionID string `gorm:"uniqueIndex"`
  Image string
  Title string
  Description string
  Buy string
  FindMore string
}

// ExcursionResponse represents the response structure for an excursion, including its ID, timestamps, and relevant fields.
type ExcursionResponse struct {
  ExcursionID string `json:"excursionId"`
  CreatedAt time.Time `json:"createdAt"`
  UpdatedAt time.Time `json:"updatedAt"`
  DeletedAt time.Time `json:"deletedAt,omitempty"`
  Image string `json:"img"`
  Title string `json:"title"`
  Description string `json:"desc"`
  Buy string `json:"buy"`
  FindMore string `json:"findMore"`
}
