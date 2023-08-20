package db

import (
	"context"
	"database/sql"
	"log"
	"time"
	"webapp/pkg/data"
)

const dbTimeout = time.Second * 3

type PostgresConn struct {
	DB *sql.DB
}

// AllUsers returns all users as a slice of *data.User
func (p *PostgresConn) AllUsers() ([]*data.User, error) {
	// create context
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `SELECT id, first_name, last_name, email, password, is_admin, created_at, updated_at FROM users
	order by last_name`

	rows, err := p.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []*data.User
	for rows.Next() {
		var user data.User
		if err := rows.Scan(&user.ID,
			&user.FirstName,
			&user.LastName,
			&user.Email,
			&user.Password,
			&user.IsAdmin,
			&user.CreatedAt,
			&user.UpdatedAt); err != nil {
			log.Printf("Error Scanning %v", err)
			return nil, err
		}

		users = append(users, &user)
	}

	return users, nil
}
