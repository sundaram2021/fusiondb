package sqlite

// Database represents the overall database configuration.
type Database struct {
    Name      string     `json:"name"`
    Charset   string     `json:"charset"`
    Collation string     `json:"collation"`
    Engine    string     `json:"engine"`
    Tables    []Table    `json:"tables"`
    Partitions []Partition `json:"partitions,omitempty"`
    Options   Options     `json:"options,omitempty"`
}

// Table represents a table within the database.
type Table struct {
    Name        string       `json:"name"`
    Columns     []Column     `json:"columns"`
    Indexes     []Index      `json:"indexes,omitempty"`
    ForeignKeys []ForeignKey  `json:"foreign_keys,omitempty"`
}

// Column represents a column in a table.
type Column struct {
    Name          string `json:"name"`
    Type          string `json:"type"`
    AutoIncrement bool   `json:"auto_increment,omitempty"`
    PrimaryKey    bool   `json:"primary_key,omitempty"`
    Nullable      bool   `json:"nullable,omitempty"`
}

// Index represents an index on a table.
type Index struct {
    Name    string   `json:"name"`
    Columns []string `json:"columns"`
    Unique  bool     `json:"unique,omitempty"`
}

// ForeignKey represents a foreign key constraint.
type ForeignKey struct {
    Name       string   `json:"name"`
    Column     string   `json:"column"`
    References Reference `json:"references"`
    OnDelete   string   `json:"on_delete,omitempty"`
    OnUpdate   string   `json:"on_update,omitempty"`
}

// Reference represents the referenced table and column for foreign keys.
type Reference struct {
    Table  string `json:"table"`
    Column string `json:"column"`
}

// Partition represents partitioning strategy for tables.
type Partition struct {
    Table  string  `json:"table"`
    Type   string  `json:"type"` // e.g., "RANGE"
    Columns []string `json:"columns"`
}

// Options contains global database options.
type Options struct {
    MaxConnections int    `json:"max_connections,omitempty"` // Maximum allowed connections
    SQLMode        string `json:"sql_mode,omitempty"`        // SQL mode for stricter validation (leave empty for SQLite)
    Timezone       string `json:"timezone,omitempty"`       // Timezone setting (not typically used in SQLite)
    JournalMode    string `json:"journal_mode,omitempty"`   // Write-Ahead Logging mode (optional)
    SynchronousMode string `json:"synchronous_mode,omitempty"` // Synchronous setting (optional)
}