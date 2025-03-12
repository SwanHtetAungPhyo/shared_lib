package handlers

type HandlerConfig struct {
	UserHandle *UsersHandler
}

var HandleConfig *HandlerConfig

func Init() {
	userHandler := NewUsersHandler()
	HandleConfig = &HandlerConfig{
		UserHandle: userHandler,
	}
}
