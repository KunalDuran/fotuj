package data

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type SQLiteDB struct {
	db *sql.DB
}

func Boot(db *sql.DB) {
	_, err := db.Exec(`CREATE TABLE IF NOT EXISTS project (
			id integer not null primary key,
			key text,
			name text,
			path text,
			vendor_id text,
			client_id text,
			link text,
			created_at text,
			updated_at text
		)`)
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS image (
		id integer not null primary key,
		key text,
		absolute_path text,
		path text,
		status text,
		updated_at text
	)`)
	if err != nil {
		log.Fatal(err)
	}
}

func NewSQLiteDB() *SQLiteDB {
	db, err := sql.Open("sqlite3", "./app.sqlite3")
	if err != nil {
		log.Fatal(err)
	}

	Boot(db)

	return &SQLiteDB{db: db}
}

func (s SQLiteDB) UpdateImageStatus(pKey, image, status string) error {
	return nil
}

func (s SQLiteDB) GetProjects(vendorID string) ([]Project, error) {
	var projects []Project

	rows, err := s.db.Query("SELECT key, name, link, client_id FROM project")
	if err != nil {
		log.Println(err)
		return projects, err
	}

	for rows.Next() {
		var p Project
		err = rows.Scan(
			&p.Key,
			&p.Name,
			&p.Link,
			&p.ClientID,
		)

		projects = append(projects, p)
	}
	return projects, nil
}

func (s SQLiteDB) GetProjectByKey(key string) (Project, error) {
	var p Project

	stmt, err := s.db.Prepare("SELECT * FROM project WHERE key=?")
	if err != nil {
		return p, err
	}

	err = stmt.QueryRow(key).Scan(&p)
	if err != nil {
		return p, err
	}

	return p, nil
}

func (s SQLiteDB) AddProject(p Project) error {
	stmt, err := s.db.Prepare(`INSERT INTO project (
		key,
		name,
		path,
		vendor_id,
		client_id,
		link,
		created_at,
		updated_at
	) VALUES(?,?,?,?,?,?,?,?)`)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(p.Key, p.Name, p.Path, p.VendorID, p.ClientID, p.Link, p.CreatedAt, p.UpdatedAt)
	if err != nil {
		return err
	}

	stmt2, err := s.db.Prepare(`INSERT INTO image(
		key,
		absolute_path,
		path,
		status,
		updated_at
	) VALUES(?,?,?,?,?)`)
	if err != nil {
		return err
	}

	for _, img := range p.Images {
		_, err = stmt2.Exec(p.Key, img.AbsolutePath, img.Path, img.Status, img.UpdatedAt)
		if err != nil {
			log.Println(err)
		}
	}

	return nil
}
