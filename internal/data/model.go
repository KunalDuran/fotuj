package data

import "time"

type Bucket struct {
	Name           string    `json:"name" bson:"name"`
	Path           string    `json:"path" bson:"path"`
	Key            string    `json:"key" bson:"key"`
	Images         []Image   `json:"images" bson:"images"`
	ShareableLink  string    `json:"shareable_link" bson:"shareable_link"`
	StorageType    string    `json:"storage_type" bson:"storage_type"`
	PhotographerID string    `json:"photographer_id" bson:"photographer_id"`
	ClientID       string    `json:"client_id" bson:"client_id"`
	CreatedAt      time.Time `json:"created_at" bson:"created_at"`
	UpdatedAt      string    `json:"updated_at" bson:"updated_at"`
}

type Image struct {
	AbsolutePath string    `json:"absolute_path" bson:"absolute_path"`
	Path         string    `json:"path" bson:"path"`
	Status       int       `json:"status" bson:"status"`
	UpdatedAt    time.Time `json:"updated_at" bson:"updated_at"`
}
