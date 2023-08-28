package repository

import (
	"context"
	"database/sql"
	"demo/entity"
	"errors"
)

type Branch struct {
	DB *sql.DB
}

func (b Branch) FindAll() (branches []entity.Branch) {
	query := `SELECT * FROM branches`
	result, err := b.DB.QueryContext(context.Background(), query)
	if err != nil {
		panic(err)
	}

	for result.Next() {
		b := entity.Branch{}
		err := result.Scan(&b.BranchId, &b.Name, &b.Location)
		if err != nil {
			panic(err)
		}

		branches = append(branches, b)
	}

	return branches
}

func (b Branch) FindById(id int) (branch entity.Branch, err error) {
	query := `SELECT * FROM branches WHERE branch_id = ?`
	result, err := b.DB.QueryContext(context.Background(), query, id)

	if err != nil {
		panic(err)
	}

	if !result.Next() {
		return branch, errors.New("branch not found")
	}

	err = result.Scan(&branch.BranchId, &branch.Name, &branch.Location)
	if err != nil {
		panic(err)
	}

	return branch, err
}

func (b Branch) Create(name, location string) (branch entity.Branch) {
	query := `
		INSERT INTO branches (name, location) VALUES
		(?, ?)
	`

	result, err := b.DB.ExecContext(context.Background(), query, name, location)
	if err != nil {
		panic(err)
	}

	newId, _ := result.LastInsertId()
	branch = entity.Branch{
		BranchId: int(newId),
		Name:     name,
		Location: location,
	}

	return branch
}
