package setup

// a struct of configuralable options
type Config struct {
	// Domain is the default domain suffix to use.
	Domain string `yaml:"domain"`
}
