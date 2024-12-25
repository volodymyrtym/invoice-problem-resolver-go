package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
)

const policyTableName = "day_off_policies"
const policyIdColumnName = "id"
const policyUserIdColumnName = "user_id"
const policyHalfDayColumnName = "half_day"
const policyApprovableColumnName = "approvable"
const policyNameColumnName = "name"

type PolicyRepository struct {
	db  *sql.DB
	ctx context.Context
}

func NewPolicyRepository(db *sql.DB, ctx context.Context) *PolicyRepository {
	return &PolicyRepository{db: db, ctx: ctx}
}

func (r *PolicyRepository) DeletePolicy(id string) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE %s = $1", policyTableName, policyIdColumnName)
	_, err := r.db.ExecContext(r.ctx, query, id)
	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}

type PolicyCreateDTO struct {
	Id         string
	Name       string
	HalfDay    bool
	Approvable bool
	UserId     string
}

func (r *PolicyRepository) SavePolicy(dto *PolicyCreateDTO) error {
	query := fmt.Sprintf("INSERT INTO %s (%s, %s, %s, %s, %s) VALUES ($1, $2, $3, $4, $5)", policyTableName, policyIdColumnName, policyUserIdColumnName, policyHalfDayColumnName, policyApprovableColumnName, policyNameColumnName)
	_, err := r.db.Exec(query, dto.Id, dto.UserId, dto.HalfDay, dto.Approvable, dto.Name)
	return err
}

func (r *PolicyRepository) IsPolicyOwner(id string, userId string) (bool, error) {
	query := fmt.Sprintf("SELECT 1 FROM %s WHERE id = $1 AND user_id = $2 LIMIT 1", policyTableName)
	var exists int
	err := r.db.QueryRowContext(r.ctx, query, id, userId).Scan(&exists)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return false, nil
		}
		return false, err
	}

	return true, nil
}
