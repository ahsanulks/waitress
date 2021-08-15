package migrations

import "github.com/go-rel/rel"

// MigrateCreateOrders create table orders
func MigrateCreateOrders(schema *rel.Schema) {
	schema.Exec(rel.Raw(
		`CREATE TABLE orders (
			id serial primary key,
			code varchar(255) unique not null,
			buyer_id int not null,
			seller_id int not null,
			state smallint not null,
			total_price int not null,
			note varchar(255),
			created_at timestamp not null,
			updated_at timestamp not null,
			check (
				total_price > 0
			)
		)`),
	)

	schema.Exec(rel.Raw(`CREATE INDEX orders_seller_id_state_idx ON orders(seller_id, state)`))
	schema.Exec(rel.Raw(`CREATE INDEX orders_buyer_id_state_idx ON orders(buyer_id, state)`))
}

// RollbackCreateOrders drop table orders
func RollbackCreateOrders(schema *rel.Schema) {
	schema.DropTable("orders")
}
