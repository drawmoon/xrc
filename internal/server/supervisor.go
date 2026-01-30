package server

// Supervisor manages the kernel process.
type Supervisor struct {
	bin    string
	config string
}

// NewSupervisor creates a new Supervisor with the given binary and config paths.
func NewSupervisor(bin string, config string) *Supervisor {
	return &Supervisor{bin: bin, config: config}
}
