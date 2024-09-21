package configs

func Crypto() map[string]interface{} {
    return map[string]interface{}{
        "key": App()["key"].(string),
    }
}
