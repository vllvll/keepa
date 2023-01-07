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

		create table bank_cards
		(
			id         serial
			    constraint bank_cards_pk
				primary key,
			number     text      not null,
			holder     text      not null,
			cvv        text      not null,
			meta       text,
			user_id    integer   not null
				constraint bank_cards_users_id_fk
					references users,
			updated_at timestamp not null
		);

		alter table bank_cards
			owner to postgres;

		create unique index bank_cards_id_uindex
			on bank_cards (id);

		create table texts
		(
			id         serial
			    constraint texts_pk
				primary key,
			content    text      not null,
			meta       text,
			user_id    integer   not null
				constraint texts_users_id_fk
					references users,
			updated_at timestamp not null
		);

		alter table texts
			owner to postgres;

		create unique index texts_id_uindex
			on texts (id);

		create table binaries
		(
			id         serial
			    constraint binaries_pk
				primary key,
			content    bytea      not null,
			meta       text,
			user_id    integer   not null
				constraint binaries_users_id_fk
					references users,
			updated_at timestamp not null
		);

		alter table binaries
			owner to postgres;

		create unique index binaries_id_uindex
			on binaries (id);
	`)

	if err != nil {
		return db, nil
	}

	return db, nil
}
