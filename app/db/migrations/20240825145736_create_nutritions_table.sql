-- +goose Up
create table if not exists nutritions(
	id integer primary key autoincrement,
	name text not null,
	unit text not null,
	amount integer not null,
	created_at datetime not null,
	updated_at datetime not null,
	ingredient_id integer not null,

	FOREIGN KEY(ingredient_id) REFERENCES ingredients(id)
);

-- +goose Down
drop table if exists nutritions;
