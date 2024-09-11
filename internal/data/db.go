package data

import "time"

type Project struct {
	Key       string    `json:"key" bson:"key"`
	Name      string    `json:"name" bson:"name"`
	Path      string    `json:"path" bson:"path"`
	Images    []Image   `json:"images" bson:"images"`
	Link      string    `json:"link" bson:"link"`
	CreatedAt time.Time `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time `json:"updated_at" bson:"updated_at"`
}

type Image struct {
	Path      string    `json:"path" bson:"path"`
	Name      string    `json:"name" bson:"name"`
	Status    int       `json:"status" bson:"status"`
	UpdatedAt time.Time `json:"updated_at" bson:"updated_at"`
}

type Database interface {
	AddProject(p Project) error
	GetProjects() ([]Project, error)
	GetProjectByKey(key string) (Project, error)
	GetImagesByKey(key string) ([]Image, error)

	UpdateImageStatus(pKey, image, status string) error
}
