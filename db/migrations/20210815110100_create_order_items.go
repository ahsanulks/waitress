package migrations

import "github.com/go-rel/rel"

// MigrateCreateOrderItems create table order_items
func MigrateCreateOrderItems(schema *rel.Schema) {
	schema.Exec(rel.Raw(
		`CREATE TABLE order_items (
			id serial primary key,
			order_id int not null,
			product_id int not null,
			cart_item_id int unique not null,
			quantity int not null,
			price int not null,
			weight int not null,
			created_at timestamp not null,
			updated_at timestamp not null,
			check (
				quantity > 0
				AND price > 0
				AND weight > 0
			),
			CONSTRAINT fk_order FOREIGN KEY(order_id) REFERENCES orders(id),
			CONSTRAINT fk_product FOREIGN KEY(product_id) REFERENCES products(id),
			CONSTRAINT fk_cart_item FOREIGN KEY(cart_item_id) REFERENCES cart_items(id)
		)`),
	)
}

// RollbackCreateOrderItems drop table order_items
func RollbackCreateOrderItems(schema *rel.Schema) {
	schema.DropTable("order_items")
}
