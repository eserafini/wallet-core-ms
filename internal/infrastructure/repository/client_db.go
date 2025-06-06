package repository

import (
	"database/sql"

	"github.com/eserafini/wallet-core-ms/internal/entity"
)

type ClientDB struct {
	db *sql.DB
}

func NewSqClientDB(client *sql.DB) *ClientDB {
	return &ClientDB{db: client}
}

func (c *ClientDB) FindByID(id string) (*entity.Client, error) {
	client := entity.Client{}
	stmt, err := c.db.Prepare("SELECT id, name, email, created_at, updated_at FROM clients WHERE id = ?")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	if err := stmt.QueryRow(id).Scan(&client.ID, &client.Name, &client.Email, &client.CreatedAt, &client.UpdatedAt); err != nil {
		return nil, err
	}

	return &client, nil
}

func (c *ClientDB) Save(client *entity.Client) error {
	stmt, err := c.db.Prepare("INSERT INTO clients (id, name, email, created_at, updated_at) VALUES (?, ?, ?, ?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	if _, err := stmt.Exec(client.ID, client.Name, client.Email, client.CreatedAt, client.UpdatedAt); err != nil {
		return err
	}

	return nil
}
