package data

import "time"

type Bucket struct {
	Name          string              `json:"name" bson:"name"`
	Key           string              `json:"key" bson:"key"`
	Images        []string            `json:"images" bson:"images"`
	Selected      map[string][]string `json:"selected" bson:"selected"`
	Rejected      map[string][]string `json:"rejected" bson:"rejected"`
	ShareableLink string              `json:"shareable_link" bson:"shareable_link"`
	StorageType   string              `json:"storage_type" bson:"storage_type"`
	VendorID      string              `json:"vendor_id" bson:"vendor_id"`
	ClientID      string              `json:"client_id" bson:"client_id"`
	Selectors     []string            `json:"selectors" bson:"selectors"`
	CreatedAt     time.Time           `json:"created_at" bson:"created_at"`
	UpdatedAt     string              `json:"updated_at" bson:"updated_at"`
}
