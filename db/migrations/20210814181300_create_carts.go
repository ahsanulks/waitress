package migrations

import "github.com/go-rel/rel"

// MigrateCreateCarts create table carts
func MigrateCreateCarts(schema *rel.Schema) {
	schema.Exec(rel.Raw(
		`CREATE TABLE carts (
			id serial primary key,
			user_id int unique not null,
			created_at timestamp not null,
			updated_at timestamp not null
		)`),
	)
}

// RollbackCreateCarts drop table carts
func RollbackCreateCarts(schema *rel.Schema) {
	schema.DropTable("carts")
}
