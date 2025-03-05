package app

import (
	"database/sql"
	"errors"
	"os"
	"time"
)

type ProjectRegister struct {
	projectRepository ProjectRepository
}

func NewProjectRegister(projectRepository ProjectRepository) *ProjectRegister {
	return &ProjectRegister{
		projectRepository: projectRepository,
	}
}

func (r *ProjectRegister) Register(title string, loginUserID int64) int64 {
	return r.projectRepository.Register(title, loginUserID)
}

type ProjectRepository interface {
	Register(title string, loginUserID int64) int64
}

type ProjectDbRepository struct {
	db *sql.DB
}

type ProjectRecord struct {
	ID        int
	Title     string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewProjectDbRepository() (*ProjectDbRepository, error) {
	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))

	if err != nil {
		return nil, errors.New("DBへの接続に失敗しました。")
	}

	return &ProjectDbRepository{db: db}, nil
}

func (p *ProjectDbRepository) Register(title string) (int64, error) {
	var id int64
	err := p.db.QueryRow("INSERT INTO projects (title) VALUES ($1) RETURNING id", title).Scan(&id)

	if err != nil {
		return 0, err
	}

	return id, nil
}

func (p *ProjectDbRepository) Ping() bool {
	err := p.db.Ping()

	if err != nil {
		return false
	}

	return true
}
