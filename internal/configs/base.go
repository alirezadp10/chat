package configs

import (
    "github.com/joho/godotenv"
    "log"
    "os"
    "strconv"
)

func getEnv(key, defaultValue string) string {
    err := godotenv.Load()

    if err != nil {
        log.Default().Fatal("Error loading .env file: ", err)
    }

    if value, exists := os.LookupEnv(key); exists {
        return value
    }
    return defaultValue
}

func getEnvAsInt(key string, defaultValue int) int {
    valueStr := getEnv(key, strconv.Itoa(defaultValue))
    value, err := strconv.Atoi(valueStr)
    if err != nil {
        return defaultValue
    }
    return value
}

func getEnvAsBool(key string, defaultValue bool) bool {
    valueStr := getEnv(key, strconv.FormatBool(defaultValue))
    value, err := strconv.ParseBool(valueStr)
    if err != nil {
        return defaultValue
    }
    return value
}
