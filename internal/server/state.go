package server

type State int32

//go:generate stringer -type=State -linecomment
const (
	StateInit       State = iota // Init
	StateStarting                // Starting
	StateRunning                 // Running
	StateStopping                // Stopping
	StateStopped                 // Stopped
	StateCrashed                 // Crashed
	StateRestarting              // Restarting
)
