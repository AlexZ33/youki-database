package orm

import "time"

// Config mysql config
type Config struct {
	DSN         string        // data source name.
	Active      int           // pool
	Idle        int           // pool
	IdleTimeout time.Duration //connect max life time.
}

type ormLog struct{}
