package schema

import (
	"database/sql/driver"
	"fmt"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"github.com/paulmach/orb"
	"github.com/paulmach/orb/encoding/wkb"
)

// A Point consists of (X,Y) or (Lat, Lon) coordinates
// and it is stored in MySQL the POINT spatial data type.
type Point [2]float64

// Scan implements the Scanner interface.
func (p *Point) Scan(value interface{}) error {
	bin, ok := value.([]byte)
	if !ok {
		return fmt.Errorf("invalid binary value for point")
	}
	var op orb.Point
	if err := wkb.Scanner(&op).Scan(bin[4:]); err != nil {
		return err
	}
	p[0], p[1] = op.X(), op.Y()
	return nil
}

// Value implements the driver Valuer interface.
func (p Point) Value() (driver.Value, error) {
	op := orb.Point{p[0], p[1]}
	return wkb.Value(op).Value()
}

// FormatParam implements the sql.ParamFormatter interface to tell the SQL
// builder that the placeholder for a Point parameter needs to be formatted.
func (p Point) FormatParam(placeholder string, info *sql.StmtInfo) string {
	if info.Dialect == dialect.MySQL {
		return "ST_GeomFromWKB(" + placeholder + ")"
	}
	return placeholder
}

// SchemaType defines the schema-type of the Point object.
func (Point) SchemaType() map[string]string {
	return map[string]string{
		dialect.MySQL: "POINT",
	}
}
