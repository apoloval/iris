package karen

// Option is an application option
type Option func(*Config) error

func applyOptions(cfg *Config, opts []Option) error {
	for _, o := range opts {
		if err := o(cfg); err != nil {
			return err
		}
	}
	return nil
}
