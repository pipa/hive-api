package cockroach

// Config contains the options for the cockroach middleware
// can be optionally be passed to the `New`
type Config struct {
	Conn *gnorm.DB
	User string
	Pass string
	Host string
	Port string
	Name string
}

// DefaultConfig returns a default config
func DefaultConfig() Config {
	return Config{
		User: "root",
		Pass: "",
		Host: "localhost",
		Port: "26257",
		Name: "hive",
	}
}
