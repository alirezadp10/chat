package configs

func Mariadb() map[string]string {
    return map[string]string{
        "user":     getEnv("DB_USER", "root"),
        "password": getEnv("DB_PASSWORD", "password"),
        "database": getEnv("DB_NAME", "chat"),
        "host":     getEnv("DB_HOST", "localhost"),
        "port":     getEnv("DB_PORT", "3306"),
    }
}
