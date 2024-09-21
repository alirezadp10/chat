package configs

func Cors() map[string]interface{} {
    allowOrigins := []string{"*"}

    if App()["env"] != "production" {
        allowOrigins = append(allowOrigins, "http://localhost:63342")
    }

    allowMethods := []string{"GET", "POST", "PUT", "DELETE"}

    return map[string]interface{}{
        "allowOrigins":     allowOrigins,
        "allowMethods":     allowMethods,
        "allowCredentials": true,
    }
}
