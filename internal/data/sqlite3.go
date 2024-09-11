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
		path text,
		name text,
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

func (s SQLiteDB) UpdateImageStatus(key, image, status string) error {
	_, err := s.db.Exec("UPDATE image SET status=? WHERE key=? AND name=?", status, key, image)
	if err != nil {
		return err
	}
	return nil
}

func (s SQLiteDB) GetProjects() ([]Project, error) {
	var projects []Project

	rows, err := s.db.Query("SELECT key, name, link FROM project")
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

func (s SQLiteDB) GetImagesByKey(key string) ([]Image, error) {
	var result []Image

	stmt, err := s.db.Prepare(`SELECT 
		name,
		path,
		status FROM image WHERE key=?`)
	if err != nil {
		return result, err
	}

	rows, err := stmt.Query(key)
	for rows.Next() {
		var img Image
		rows.Scan(
			&img.Name,
			&img.Path,
			&img.Status,
		)
		result = append(result, img)
	}
	if err != nil {
		return result, err
	}

	return result, nil
}

func (s SQLiteDB) AddProject(p Project) error {
	stmt, err := s.db.Prepare(`INSERT INTO project (
		key,
		name,
		path,
		link,
		created_at,
		updated_at
	) VALUES(?,?,?,?,?,?)`)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(p.Key, p.Name, p.Path, p.Link, p.CreatedAt, p.UpdatedAt)
	if err != nil {
		return err
	}

	stmt2, err := s.db.Prepare(`INSERT INTO image(
		key,
		path,
		name,
		status,
		updated_at
	) VALUES(?,?,?,?,?)`)
	if err != nil {
		return err
	}

	for _, img := range p.Images {
		_, err = stmt2.Exec(p.Key, img.Path, img.Name, img.Status, img.UpdatedAt)
		if err != nil {
			log.Println(err)
		}
	}

	return nil
}
