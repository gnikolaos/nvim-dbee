package adapters

import (
	"database/sql"
	"fmt"
	nurl "net/url"

	_ "github.com/SAP/go-hdb/driver"

	"github.com/kndndrj/nvim-dbee/dbee/core"
	"github.com/kndndrj/nvim-dbee/dbee/core/builders"
)

func init() {
	_ = register(&Hana{}, "hana", "sap-hana")
}

var _ core.Adapter = (*Hana)(nil)

type Hana struct{}

func (h *Hana) Connect(url string) (core.Driver, error) {
	u, err := nurl.Parse(url)
	if err != nil {
		return nil, fmt.Errorf("Could not parse SAP HANA connection string: %w: ", err)
	}

	db, err := sql.Open("hdb", u.String())
	if err != nil {
		return nil, fmt.Errorf("Unable to connect to SAP HANA database: %v", err)
	}

	return &hanaDriver{
		c: builders.NewClient(db),
	}, nil
}

func (*Hana) GetHelpers(opts *core.TableOptions) map[string]string {
	return map[string]string{
		"List":         fmt.Sprintf("SELECT * FROM \"%s\".\"%s\" LIMIT 500", opts.Schema, opts.Table),
		"Columns":      fmt.Sprintf("SELECT COLUMN_NAME, DATA_TYPE_NAME FROM SYS.TABLE_COLUMNS WHERE SCHEMA_NAME = '%s' AND TABLE_NAME = '%s'", opts.Schema, opts.Table),
		"Indexes":      fmt.Sprintf("SELECT INDEX_NAME FROM SYS.INDEXES WHERE SCHEMA_NAME = '%s' AND TABLE_NAME = '%s'", opts.Schema, opts.Table),
		"Foreign Keys": fmt.Sprintf("SELECT CONSTRAINT_NAME FROM SYS.CONSTRAINTS WHERE SCHEMA_NAME = '%s' AND TABLE_NAME = '%s' AND CONSTRAINT_TYPE = 'FOREIGN KEY'", opts.Schema, opts.Table),
		"Primary Keys": fmt.Sprintf("SELECT CONSTRAINT_NAME FROM SYS.CONSTRAINTS WHERE SCHEMA_NAME = '%s' AND TABLE_NAME = '%s' AND CONSTRAINT_TYPE = 'PRIMARY KEY'", opts.Schema, opts.Table),
	}
}
