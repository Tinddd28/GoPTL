package repository

import (
	"context"
	"fmt"
	"github.com/Tinddd28/GoPTL/internal/models"
	"github.com/jackc/pgx/v5/pgxpool"
)

type ProjPostgres struct {
	db *pgxpool.Pool
}

func NewProjPostgres(db *pgxpool.Pool) *ProjPostgres {
	return &ProjPostgres{db: db}
}

func (pp *ProjPostgres) CreateProject(proj models.Project) (int, error) {
	query := fmt.Sprintf("INSERT INTO %s (title, description, token_title, amount, cost_per_token, image) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id", projectsTable)
	err := pp.db.QueryRow(context.Background(), query, proj.Title, proj.Description, proj.TokenTitle, proj.Amount, proj.CostPerToken, proj.Image).Scan(&proj.Id)
	if err != nil {
		return 0, err
	}
	return proj.Id, nil
}

func (pp *ProjPostgres) GetProjects(offset int) ([]models.Project, error) {
	query := fmt.Sprintf("SELECT id, title, description, token_title, amount, cost_per_token, image FROM %s LIMIT $1 OFFSET $2", projectsTable)
	rows, err := pp.db.Query(context.Background(), query, limitСount, offset)
	if err != nil {
		return nil, err
	}
	var projects []models.Project
	for rows.Next() {
		var proj models.Project
		err := rows.Scan(&proj.Id, &proj.Title, &proj.Description, &proj.TokenTitle, &proj.Amount, &proj.CostPerToken, &proj.Image)
		if err != nil {
			return nil, err
		}
		projects = append(projects, proj)
	}

	//slog.Info("smth", slog.Any("projects", projects))
	return projects, nil
}

func (pp *ProjPostgres) GetProjectById(id int) (models.Project, error) {
	query := fmt.Sprintf("SELECT id, title, description, token_title, amount, cost_per_token, image FROM %s WHERE id=$1", projectsTable)
	var proj models.Project
	err := pp.db.QueryRow(context.Background(), query, id).Scan(&proj.Id, &proj.Title, &proj.Description, &proj.TokenTitle, &proj.Amount, &proj.CostPerToken, &proj.Image)
	if err != nil {
		return models.Project{}, err
	}

	return proj, nil
}

func (pp *ProjPostgres) UpdateProject(id int, input models.Project) error {
	query := fmt.Sprintf("UPDATE %s SET title=$2, description=$3, token_title=$4, amount=$6, cost_per_token=$6, image=$6 WHERE id=$1", projectsTable)
	_, err := pp.db.Exec(context.Background(), query, id, input.Title, input.Description, input.TokenTitle, input.Amount, input.CostPerToken, input.Image)
	if err != nil {
		return err
	}
	return nil
}

func (pp *ProjPostgres) DeleteProject(id int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id=$1", projectsTable)
	_, err := pp.db.Exec(context.Background(), query, id)
	if err != nil {
		return err
	}
	return nil
}
