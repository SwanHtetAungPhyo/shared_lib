package config_gate

type GateWayRoutes struct {
	UserRoutes    *RoutesConfig `json:"user_routes"`
	ServiceRoutes *RoutesConfig `json:"service_routes"`
	ReviewRoutes  *RoutesConfig `json:"review_routes"`
}

type RoutesConfig struct {
	ServiceName string   `json:"service_name"`
	Routes      []string `json:"routes"`
	Methods     []string `json:"methods"`
	Port        int      `json:"port"`
	Docker      bool     `json:"docker"`
}

func NewRoutesConfig(serviceName string, routes []string, methods []string, port int, docker bool) *RoutesConfig {
	return &RoutesConfig{
		ServiceName: serviceName,
		Routes:      routes,
		Methods:     methods,
		Port:        port,
		Docker:      docker,
	}
}

func NewGateWayRoutes() *GateWayRoutes {
	userUri := []string{"/user"}
	serviceUri := []string{"/service"}
	reviewUri := []string{"/review"}

	methods := []string{"GET", "POST", "PUT", "DELETE"}
	return &GateWayRoutes{
		UserRoutes:    NewRoutesConfig("user-service", userUri, methods, 3001, false),
		ServiceRoutes: NewRoutesConfig("service-service", serviceUri, methods, 3001, false),
		ReviewRoutes:  NewRoutesConfig("review-service", reviewUri, methods, 3001, false),
	}
}
