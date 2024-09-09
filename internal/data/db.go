package data

import "time"

type Project struct {
	Key       string    `json:"key" bson:"key"`
	Name      string    `json:"name" bson:"name"`
	Path      string    `json:"path" bson:"path"`
	VendorID  string    `json:"vendor_id" bson:"vendor_id"`
	ClientID  string    `json:"client_id" bson:"client_id"`
	Images    []Image   `json:"images" bson:"images"`
	Link      string    `json:"link" bson:"link"`
	CreatedAt time.Time `json:"created_at" bson:"created_at"`
	UpdatedAt string    `json:"updated_at" bson:"updated_at"`
}

type Image struct {
	AbsolutePath string    `json:"absolute_path" bson:"absolute_path"`
	Path         string    `json:"path" bson:"path"`
	Status       int       `json:"status" bson:"status"`
	UpdatedAt    time.Time `json:"updated_at" bson:"updated_at"`
}

type Database interface {
	AddProject(p Project) error
	GetProjects(vendorID string) ([]Project, error)
	GetProjectByKey(key string) (Project, error)

	UpdateImageStatus(pKey, image, status string) error
}
