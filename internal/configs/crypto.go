package configs

func Crypto() map[string]string {
    return map[string]string{
        "key": getEnv("APP_KEY", "examplekey123456"),
    }
}
