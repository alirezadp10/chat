package configs

import (
    "github.com/joho/godotenv"
    "os"
)

func getEnv(key, defaultValue string) string {
    _ = godotenv.Load()
    if value, exists := os.LookupEnv(key); exists {
        return value
    }
    return defaultValue
}
