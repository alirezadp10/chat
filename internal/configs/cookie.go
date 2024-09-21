package configs

func Cookie() map[string]interface{} {
    return map[string]interface{}{
        "secure": getEnvAsBool("COOKIE_SECURE", true),
    }
}
