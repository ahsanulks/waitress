package migrations

import "github.com/go-rel/rel"

// MigrateCreateProducts create table products
func MigrateCreateProducts(schema *rel.Schema) {
	schema.Exec(rel.Raw(
		`CREATE TABLE products (
			id serial primary key,
			name varchar(255) not null,
			seller_id int not null,
			price int not null,
			active boolean,
			stock int not null,
			weight int not null,
			created_at timestamp not null,
			updated_at timestamp not null,
			check (
				stock >= 0
				AND price >= 0
			)
		)`),
	)
}

// RollbackCreateProducts drop table products
func RollbackCreateProducts(schema *rel.Schema) {
	schema.DropTable("claims")
}
