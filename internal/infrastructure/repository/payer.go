package repository

import (
	"database/sql"
	"errors"
	_ "github.com/mattn/go-sqlite3"
	"github.com/urabexon/WarikanBot/internal/domain/entity"
	"github.com/urabexon/WarikanBot/internal/domain/valueobject"
)

type PayerRepository struct {
	db *sql.DB
}

func NewPayerRepository(filename string) (*PayerRepository, error) {
	db, err := sql.Open("sqlite3", filename)
	if err != nil {
		return nil, err
	}

	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS payers (
			id TEXT PRIMARY KEY,
			event_id TEXT NOT NULL,
			weight INTEGER NOT NULL
		);
	`)
	if err != nil {
		return nil, err
	}

	return &PayerRepository{
		db: db,
	}, nil
}

func (r *PayerRepository) Create(payer *entity.Payer) error {
	_, err := r.db.Exec("INSERT INTO payers (id, event_id, weight) VALUES (?, ?, ?)",
		payer.ID.String(),
		payer.EventID.String(),
		payer.Weight.Int(),
	)
	if sqliteErr := new(sqlite3.Error); errors.As(err, sqliteErr) {
		if sqliteErr.ExtendedCode == sqlite3.ErrConstraintPrimaryKey || sqliteErr.ExtendedCode == sqlite3.ErrConstraintUnique {
			return valueobject.NewErrorAlreadyExists("payer already exists", err)
		}
	}
	return err
}

func (r *PayerRepository) CreateIfNotExists(payer *entity.Payer) error {
	_, err := r.db.Exec("INSERT OR IGNORE INTO payers (id, event_id, weight) VALUES (?, ?, ?)",
		payer.ID.String(),
		payer.EventID.String(),
		payer.Weight.Int(),
	)
	return err
}

func (r *PayerRepository) FindByEventID(eventID valueobject.EventID) ([]*entity.Payer, error) {
	rows, err := r.db.Query("SELECT id, event_id, weight FROM payers WHERE event_id = ?", eventID.String())
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var payers []*entity.Payer
	for rows.Next() {
		var rawID, rawEventID string
		var weight int
		if err := rows.Scan(&rawID, &rawEventID, &weight); err != nil {
			return nil, err
		}
		var payer entity.Payer
		payer.ID = valueobject.NewPayerID(rawID)
		payer.EventID = valueobject.NewEventID(rawEventID)
		payer.Weight, err = valueobject.NewPercent(weight)
		if err != nil {
			return nil, err
		}
		payers = append(payers, &payer)
	}
	return payers, nil
}
