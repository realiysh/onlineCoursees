package config

// Config holds the configuration for the application
type Config struct {
	Port     string
	Database string
}

// NewConfig creates a new Config instance
func NewConfig() *Config {
	return &Config{
		Port:     "8084",
		Database: "your_database_name",
	}
}

// ApplyConfig applies the configuration to the application
func (c *Config) ApplyConfig() {
	// Apply configuration to the application
}

// GetDatabase returns the database connection
func (c *Config) GetDatabase() string {
	return c.Database
}

// GetPort returns the port for the application
func (c *Config) GetPort() string {
	return c.Port
}
