package user

import (
	"context"
	"database/sql"
	"golang-websocket/api/models"
	"golang-websocket/api/repository"
	"strconv"
	"strings"
	"time"
)

type mysqlUserRepository struct {
	Conn *sql.DB
}

func NewUserRepository(Conn *sql.DB) repository.UserRepository {
	return &mysqlUserRepository{Conn}
}

func (m *mysqlUserRepository) List(ctx context.Context) ([]*models.User, error) {
	var users []*models.User

	query := `select id, nama, username, email, password, created_at, updated_at FROM user`

	rows, err := m.Conn.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}

	defer func() {
		err := rows.Close()
		if err != nil {
			panic(err)
		}
	}()

	for rows.Next() {
		row := new(models.User)
		err := rows.Scan(
			&row.ID,
			&row.Nama,
			&row.Username,
			&row.Email,
			&row.Password,
			&row.CreatedAt,
			&row.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		users = append(users, row)
	}

	return users, nil
}

func (m *mysqlUserRepository) Detail(ctx context.Context, id int) (*models.User, error) {
	var user models.User

	ids := strconv.Itoa(id)
	query := `SELECT id, nama, username, email, password, created_at, updated_at FROM user WHERE id = ` + ids
	err := m.Conn.QueryRowContext(ctx, query).Scan(&user.ID, &user.Nama, &user.Username, &user.Email, &user.Password, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (m *mysqlUserRepository) Insert(ctx context.Context, user models.User) (*models.User, error) {
	query := `INSERT INTO user (nama, username, email, password, created_at, updated_at) VALUES(?, ?, ?, ?, ?, ?)`
	stmt, err := m.Conn.PrepareContext(ctx, query)
	if err != nil {
		return nil, err
	}

	res, err := stmt.ExecContext(ctx, user.Nama, user.Username, user.Email, user.Password, time.Now(), time.Now())
	if err != nil {
		return nil, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}

	user.ID = int(id)
	return &user, nil
}

func (m *mysqlUserRepository) Update(ctx context.Context, datas map[string]interface{}, id int) error {
	query := `UPDATE user SET`

	var data []interface{}
	for index, value := range datas {
		if value != "" {
			query = query + ` ` + index + ` = ?,`
			data = append(data, value)
		}
	}
	data = append(data, id)

	query = strings.TrimRight(query, ", ")
	query = query + ` WHERE id = ?`

	stmt, err := m.Conn.PrepareContext(ctx, query)
	if err != nil {
		return err
	}

	_, err = stmt.ExecContext(ctx, data...)
	if err != nil {
		return err
	}
	return nil
}

func (m *mysqlUserRepository) Delete(ctx context.Context, id int) error {
	ids := strconv.Itoa(id)
	query := `DELETE FROM user WHERE id = ?`

	stmt, err := m.Conn.PrepareContext(ctx, query)
	if err != nil {
		return err
	}

	_, err = stmt.ExecContext(ctx, ids)
	if err != nil {
		return err
	}

	return nil
}

func (m *mysqlUserRepository) Login(ctx context.Context, username string, password string) (*models.User, error) {
	var user models.User
	query := `SELECT * FROM user WHERE username = ? AND password = ?`

	err := m.Conn.QueryRowContext(ctx, query, username, password).Scan(&user.ID, &user.Nama, &user.Username, &user.Email, &user.Password, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (m *mysqlUserRepository) CheckUser(ctx context.Context, username string) error {
	var user models.User
	query := `SELECT * FROM user WHERE username = ?`

	err := m.Conn.QueryRowContext(ctx, query, username).Scan(&user.ID, &user.Nama, &user.Username, &user.Email, &user.Password, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return err
	}
	return nil
}
