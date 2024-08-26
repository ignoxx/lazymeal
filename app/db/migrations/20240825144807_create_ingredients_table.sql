-- +goose Up
create table if not exists ingredients(
	id integer primary key autoincrement,
	name text not null,
	unit text not null,
	amount integer not null,
	calories integer not null,
	washing_effort integer not null,
	cutting_effort integer not null,
	pelling_effort integer not null,
	created_at datetime not null,
	updated_at datetime not null,
	meal_id integer not null,

	FOREIGN KEY(meal_id) REFERENCES meals(id)
);

-- +goose Down
drop table if exists ingredients;
