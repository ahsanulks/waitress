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
			purchased boolean,
			created_at timestamp not null,
			updated_at timestamp not null,
			check (
				quantity >= 1
			),
			CONSTRAINT fk_cart FOREIGN KEY(cart_id) REFERENCES carts(id),
			CONSTRAINT fk_product FOREIGN KEY(product_id) REFERENCES products(id)
		)`),
	)

	schema.Exec(rel.Raw(`CREATE INDEX cart_items_cart_id_product_id_pruchased_idx ON cart_items(purchased, cart_id, product_id)`))
}

// RollbackCreateCartItems drop table cart_items
func RollbackCreateCartItems(schema *rel.Schema) {
	schema.DropTable("cart_items")
}
