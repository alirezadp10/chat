package configs

func JWT() map[string]string {
    return map[string]string{
        "secret":        getEnv("JWT_SECRET", ""),
        "tokenLifeTime": getEnv("JWT_TOKEN_LIFE_TIME", "72"), //Hour
    }
}
