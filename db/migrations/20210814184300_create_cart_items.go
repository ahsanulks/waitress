package migrations

import "github.com/go-rel/rel"

// MigrateCreateCartItems create table cart_items
func MigrateCreateCartItems(schema *rel.Schema) {
	schema.Exec(rel.Raw(
		`CREATE TABLE cart_items (
			id serial primary key,
			cart_id int,
			product_id int,
			quantity int not null,
			weight int not null,
			price int not null,
			purchased boolean,
			created_at timestamp not null,
			updated_at timestamp not null,
			check (
				quantity >= 0
				AND price >= 0
				AND weight >= 0
			),
			CONSTRAINT fk_cart FOREIGN KEY(cart_id) REFERENCES carts(id),
			CONSTRAINT fk_product FOREIGN KEY(product_id) REFERENCES products(id)
		)`),
	)
}

// RollbackCreateCartItems drop table cart_items
func RollbackCreateCartItems(schema *rel.Schema) {
	schema.DropTable("cart_items")
}