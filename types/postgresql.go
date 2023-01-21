package types

import (
	"fmt"
	"strings"
)

type ColType string

const (
	Bigint       ColType = "bigint"
	Bigserial    ColType = "bigserial"
	Bit          ColType = "bit"
	Varbit       ColType = "varbit"
	Bool         ColType = "bool"
	Box          ColType = "box"
	Bytea        ColType = "bytea"
	Char         ColType = "char"
	Varchar      ColType = "varchar"
	Cidr         ColType = "cidr"
	Circle       ColType = "circle"
	Date         ColType = "date"
	Float8       ColType = "float8"
	Inet         ColType = "inet"
	Int          ColType = "int"
	Interval     ColType = "interval"
	Json         ColType = "json"
	Jsonb        ColType = "jsonb"
	Line         ColType = "line"
	Lseg         ColType = "lseg"
	Macaddr      ColType = "macaddr"
	Macaddr8     ColType = "macaddr8"
	Money        ColType = "money"
	Decimal      ColType = "decimal"
	Path         ColType = "path"
	PgLsn        ColType = "pg_lsn"
	PgSnapshot   ColType = "pg_snapshot"
	Point        ColType = "point"
	Polygon      ColType = "polygon"
	Real         ColType = "real"
	Smallint     ColType = "smallint"
	Smallserial  ColType = "smallserial"
	Serial       ColType = "serial"
	Text         ColType = "text"
	Time         ColType = "time"
	Timetz       ColType = "timetz"
	Timestamp    ColType = "timestamp"
	Timestamptz  ColType = "timestamptz"
	Tsquery      ColType = "tsquery"
	Tsvector     ColType = "tsvector"
	TxidSnapshot ColType = "txid_snapshot"
	Uuid         ColType = "uuid"
	Xml          ColType = "xml"
)

func (c ColType) Options(options ...string) ColType {
	additional := strings.Join(options, ",")

	return ColType(fmt.Sprintf("%s(%s)", c, additional))
}
