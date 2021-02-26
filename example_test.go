package entspatial

import (
	"context"
	"fmt"
	"log"

	"github.com/a8m/entspatial/ent"
	"github.com/a8m/entspatial/ent/schema"

	"entgo.io/ent/dialect"
	_ "github.com/go-sql-driver/mysql"
)

func Example_Point() {
	host, port := "localhost", 3308
	client, err := ent.Open(dialect.MySQL, fmt.Sprintf("root:pass@tcp(%s:%d)/test?parseTime=True", host, port))
	if err != nil {
		log.Fatalf("failed opening connection to mysql: %v", err)
	}
	defer client.Close()
	ctx := context.Background()
	// Run the auto migration tool.
	if err := client.Schema.Create(ctx); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
	tlv := client.Location.
		Create().
		SetName("TLV").
		SetCoords(schema.Point{32.109333, 34.855499}).
		SaveX(ctx)
	fmt.Println(tlv.Name, tlv.Coords)
	office := client.Location.
		Create().
		SetName("FB").
		SetCoords(schema.Point{32.072184, 34.78471}).
		SetParent(tlv).
		SaveX(ctx)
	fmt.Println(office.Name, office.Coords)

	// Output:
	// TLV [32.109333 34.855499]
	// FB [32.072184 34.78471]
}
