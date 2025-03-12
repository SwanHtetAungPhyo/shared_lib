package handlers

type HandlerConfig struct {
	UserHandle         *UsersHandler
	ServiceBlockHandle *BlocksHandler
}

var HandleConfig *HandlerConfig

func Init() {
	userHandler := NewUsersHandler()
	serviceBlockHandler := NewBlocksHandler()
	HandleConfig = &HandlerConfig{
		UserHandle:         userHandler,
		ServiceBlockHandle: serviceBlockHandler,
	}
}
