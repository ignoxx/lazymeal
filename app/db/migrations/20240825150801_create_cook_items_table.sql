-- +goose Up
create table if not exists cook_items(
	id integer primary key autoincrement,
	name text not null,
	washing_effort integer not null,
	created_at datetime not null,
	updated_at datetime not null,
	meal_id integer not null,

	FOREIGN KEY(meal_id) REFERENCES meals(id)
);

-- +goose Down
drop table if exists cook_items;

