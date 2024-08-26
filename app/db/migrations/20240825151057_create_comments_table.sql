-- +goose Up
create table if not exists comments(
	id integer primary key autoincrement,
	author text not null,
	content text not null,
	approved boolean not null,
	created_at datetime not null,
	updated_at datetime not null,
	meal_id integer not null,

	FOREIGN KEY(meal_id) REFERENCES meals(id)
);

-- +goose Down
drop table if exists comments;

