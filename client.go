package ray

import (
	"database/sql"
	"fmt"
)

type Client struct {
	driver Driver
}

func Open(driverName, dataSourceName string) (*Client, error) {
	switch driverName {
	case MySQL, Postgres, SQLite:
		_, err := sql.Open(driverName, dataSourceName)
		if err != nil {
			return nil, err
		}
		return &Client{}, nil
	default:
		return nil, fmt.Errorf("unsupported driver: %q", driverName)
	}
}
