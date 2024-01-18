package samplerepository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/google/uuid"
	_ "github.com/lib/pq"
	"initial/domain/sample/samplemodel"
)

type Repository interface {
	InsertData(ctx context.Context, sampleData samplemodel.Sample) (res samplemodel.Sample, err error)
	UpdateData(ctx context.Context, sampleData samplemodel.Sample) (res samplemodel.Sample, err error)
	DeleteData(ctx context.Context, sampleData samplemodel.Sample) (err error)
	GetDataById(ctx context.Context, id string) (res samplemodel.Sample, err error)
	GetAllData(ctx context.Context) (res []samplemodel.Sample, err error)
	Begin(ctx context.Context) (*sql.Tx, error)
}

func New(db *sql.DB) *repository {
	return &repository{DB: db}
}

type repository struct {
	*sql.DB
}

func (r *repository) Begin(ctx context.Context) (*sql.Tx, error) {
	return r.DB.BeginTx(ctx, nil)
}

func (r *repository) InsertData(ctx context.Context, sampleData samplemodel.Sample) (res samplemodel.Sample, err error) {
	sampleData.ID = uuid.New()
	query := "INSERT INTO sample_datas (id, updated_at, deleted_at, created_at) VALUES ($1, $2, $3, $4)"
	result, err := r.DB.ExecContext(ctx, query, sampleData.ID, sampleData.UpdatedAt, sampleData.DeletedAt, sampleData.CreatedAt)
	if err != nil {
		return res, fmt.Errorf("samplerepository: error inserting data, %w", err)
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected != 1 {
		return res, errors.New("samplerepository: no rows affected during insert")
	}

	return sampleData, nil
}

func (r *repository) UpdateData(ctx context.Context, sampleData samplemodel.Sample) (res samplemodel.Sample, err error) {
	query := "UPDATE sample_datas SET updated_at=$1, deleted_at=$2, created_at=$3 WHERE id=$4"
	result, err := r.DB.ExecContext(ctx, query, sampleData.UpdatedAt, sampleData.DeletedAt, sampleData.CreatedAt, sampleData.ID)
	if err != nil {
		return res, fmt.Errorf("samplerepository: error updating data, %w", err)
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected != 1 {
		return res, errors.New("samplerepository: no rows affected during update")
	}

	return sampleData, nil
}

func (r *repository) DeleteData(ctx context.Context, sampleData samplemodel.Sample) (err error) {
	query := "DELETE FROM sample_datas WHERE id=$1"
	result, err := r.DB.ExecContext(ctx, query, sampleData.ID)
	if err != nil {
		return fmt.Errorf("samplerepository: error deleting data, %w", err)
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected != 1 {
		return errors.New("samplerepository: no rows affected during delete")
	}

	return nil
}

func (r *repository) GetDataById(ctx context.Context, id string) (res samplemodel.Sample, err error) {
	var sampleData samplemodel.Sample
	query := "SELECT * FROM sample_datas WHERE id=$1 AND deleted_at IS NULL"
	err = r.DB.QueryRowContext(ctx, query, id).Scan(
		&sampleData.ID, &sampleData.UpdatedAt, &sampleData.DeletedAt, &sampleData.CreatedAt,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return samplemodel.Sample{}, errors.New("Sample Not Found")
		}
		return res, fmt.Errorf("samplerepository: error getting data, %w", err)
	}
	return sampleData, nil
}

func (r *repository) GetAllData(ctx context.Context) (res []samplemodel.Sample, err error) {
	query := "SELECT * FROM sample_datas WHERE deleted_at IS NULL"
	rows, err := r.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("samplerepository: error getting data, %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var sampleData samplemodel.Sample
		err := rows.Scan(
			&sampleData.ID, &sampleData.UpdatedAt, &sampleData.DeletedAt, &sampleData.CreatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("samplerepository: error scanning data, %w", err)
		}
		res = append(res, sampleData)
	}

	return res, nil
}
