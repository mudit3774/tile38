package server

func NewServerHooks(server Server) (SHooks, error) {
	// Wire the hook that you want, NoOp is wired by default
	return NewNoOpHook(server), nil
}
