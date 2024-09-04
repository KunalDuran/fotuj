package data

import "time"

type Bucket struct {
	Name          string    `json:"name" bson:"name"`
	Key           string    `json:"key" bson:"key"`
	Images        []Image   `json:"images" bson:"images"`
	ShareableLink string    `json:"shareable_link" bson:"shareable_link"`
	StorageType   string    `json:"storage_type" bson:"storage_type"`
	VendorID      string    `json:"vendor_id" bson:"vendor_id"`
	ClientID      string    `json:"client_id" bson:"client_id"`
	CreatedAt     time.Time `json:"created_at" bson:"created_at"`
	UpdatedAt     string    `json:"updated_at" bson:"updated_at"`
}

type Image struct {
	Path      string    `json:"path" bson:"path"`
	Status    int       `json:"status" bson:"status"`
	UpdatedAt time.Time `json:"updated_at" bson:"updated_at"`
}
