package timeout

import "time"

type Config struct {
	Duration time.Duration
}

var ConfigDefault = Config{
	Duration: 30 * time.Second,
}

type Option func(*Config)

func WithDuration(duration time.Duration) Option {
	return func(c *Config) {
		c.Duration = duration
	}
}
