package migrations

import "github.com/go-rel/rel"

// MigrateCreateCarts create table carts
func MigrateCreateCarts(schema *rel.Schema) {
	schema.Exec(rel.Raw(
		`CREATE TABLE carts (
			id serial primary key,
			user_id int not null,
			created_at timestamp not null,
			updated_at timestamp not null
		)`),
	)

	schema.Exec(rel.Raw(`CREATE INDEX carts_user_id_idx ON carts (user_id)`))
}

// RollbackCreateCarts drop table carts
func RollbackCreateCarts(schema *rel.Schema) {
	schema.Exec(rel.Raw(`DROP INDEX carts_user_id_idx`))

	schema.DropTable("carts")
}
