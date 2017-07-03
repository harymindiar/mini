package provider

import (
	"database/sql"
	"errors"
	_ "github.com/go-sql-driver/mysql"
	"github.com/harymindiar/mini/core"
	"log"
	"sync"
)

const (
	// DatabaseProviderName name of config provider
	DatabaseProviderName = "core.database_provider"

	// DatabaseServiceName service name of config
	DatabaseServiceName = "core.database"

	// DatabaseDriver database driver
	DatabaseDriver = "mysql"
)

// DatabaseCollection struct
type DatabaseCollection struct {
	Conn map[string]*sql.DB
	mu   sync.RWMutex
}

// Get database by connection name
func (d *DatabaseCollection) Get(connectionName string) *sql.DB {
	d.mu.RLock()
	conn := d.Conn[connectionName]
	d.mu.RUnlock()

	return conn
}

// NewDatabaseCollection create database collection
func NewDatabaseCollection() *DatabaseCollection {
	return &DatabaseCollection{
		Conn: make(map[string]*sql.DB, 0),
	}
}

// Database provider
type Database struct {
	Base
}

// GetName get provider name
func (m Database) GetName() string {
	return DatabaseProviderName
}

// Boot the provider
func (m *Database) Boot() error {
	config, _ := m.Container.Get(ConfigServiceName)
	stringConnection := config.(core.Config).GetStringMap("database.mysql")
	if stringConnection == nil {
		return errors.New("Database configuration not found")
	}
	databaseCollection := NewDatabaseCollection()
	for connectionName, config := range stringConnection {
		dataSource := config.(string)
		conn, err := sql.Open(DatabaseDriver, dataSource)
		if err != nil {
			log.Fatalln(err.Error())
		}
		databaseCollection.Conn[connectionName] = conn
	}

	m.Container.Set(DatabaseServiceName, databaseCollection)

	return nil
}
