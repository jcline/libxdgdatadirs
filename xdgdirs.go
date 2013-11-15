package libxdgdirs

import (
    "errors"
    "path/filepath"
    "os"
    //"path/filepath"
)

var (
    envNotSet = errors.New("Environment variable not set")
)

func CheckDir(path string) (bool, error) {
    _, err := os.Stat(path)
    if err != nil {
        return false, err
    }
    return true, nil
}

func loadEnv(name string) (string, error) {
    result := os.Getenv(name)
    if result == "" {
        return "", envNotSet
    }
    return result, nil
}

func loadEnvOrDefault(name string, def string) (string) {
    result, err := loadEnv(name)
    if err != nil {
        return def
    }
    return result
}

func getHome() (string) {
    return loadEnvOrDefault("HOME", "")
}

func XDG_DATA_HOME() (string) {
    return loadEnvOrDefault("XDG_DATA_HOME", filepath.Join(getHome(), ".local", "share"))
}

func XDG_DATA_DIRS() (string) {
    return loadEnvOrDefault("XDG_DATA_DIRS", "/usr/local/share/:/usr/share/")
}

func XDG_CONFIG_HOME() (string) {
    return loadEnvOrDefault("XDG_CONFIG_HOME", filepath.Join(getHome(), ".config"))
}

func XDG_CONFIG_DIRS() (string) {
    return loadEnvOrDefault("XDG_CONFIG_DIRS", "/etc/xdg/")
}

func XDG_CACHE_HOME() (string) {
    return loadEnvOrDefault("XDG_CACHE_HOME", filepath.Join(getHome(), ".cache"))
}

func TryLoad(env func() string) (string, error) {
    dir := env()

    _, err := checkDir(dir)
    return dir, err
}

func LoadOrCreate(env func() string, subpaths ...string) (string, error) {
    dir, err := TryLoad(env)
    if err != nil {
        err = os.MkdirAll(dir, os.ModeDir | 0700)
        if err != nil {
            return "", err
        }
    }

    if len(subpaths) > 0 {
        crap := []string{dir}
        dir = filepath.Join(append(crap, subpaths...)...)
        result, err := checkDir(dir)
        if !result || err != nil {
            err = os.MkdirAll(dir, os.ModeDir | 0700)
            if err != nil {
                return "", err
            }
        }
    }

    return dir, nil
}
