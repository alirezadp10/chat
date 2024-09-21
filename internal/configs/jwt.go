package configs

func JWT() map[string]interface{} {
    return map[string]interface{}{
        "secret":        getEnv("JWT_SECRET", ""),
        "tokenLifeTime": getEnvAsInt("JWT_TOKEN_LIFE_TIME", 72),
    }
}
