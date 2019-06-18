// Config is put into a different package to prevent cyclic imports in case
// it is needed in several locations

package config

type Config struct {
	Option string   `config:"option"`
	OS     string   `config:"os"` // linux, android
	Tags   []string `config:"tags"`
}

var DefaultConfig = Config{
	Option: "",
	OS:     "linux",
	Tags:   []string{},
}
