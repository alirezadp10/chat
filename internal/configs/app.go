package configs

func App() map[string]interface{} {
    return map[string]interface{}{
        "name":  getEnv("APP_NAME", "Go"),
        "url":   getEnv("APP_URL", "localhost"),
        "debug": getEnvAsBool("APP_DEBUG", false),
        "env":   getEnv("APP_ENV", "production"),
        "key":   getEnv("APP_KEY", ""),
    }
}
