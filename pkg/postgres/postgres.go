// Package postgres Функционал для работы с PostgreSQL
package postgres

import (
	"database/sql"

	_ "github.com/jackc/pgx/v4/stdlib"
)

// ConnectDatabase Инициализация базы данных
func ConnectDatabase(dsn string) (*sql.DB, error) {
	var db *sql.DB

	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}

	_, err = db.Exec(`
		create table users
		(
			id            serial
				constraint users_pk
					primary key,
			login         text      not null,
			password_hash text      not null,
			created_at    timestamp not null
		);
		
		alter table users
			owner to postgres;
		
		create unique index users_id_uindex
			on users (id);
		
		create unique index users_login_uindex
			on users (login);

		create table tokens
		(
			id         serial
				constraint tokens_pk
					primary key,
			token      text      not null,
			user_id    integer   not null
				constraint tokens_users_id_fk
					references users,
			last_login timestamp not null
		);
		
		alter table tokens
			owner to postgres;
		
		create unique index tokens_id_uindex
			on tokens (id);
	`)

	if err != nil {
		return db, nil
	}

	return db, nil
}
