package libxdgdirs

import (
    "path/filepath"
    "testing"
)

func TestCheckDir(t *testing.T) {
    tests := []struct {
        path string
        result bool
    }{
        {".", true},
        {"/", true},
        {"cat", false},
    }

    for _, test := range tests {
        result, _:= checkDir(test.path)
        if result != test.result {
            t.Error(test.path, "expected", test.result, "but got", result)
        }
    }
}

func TestLoadEnv(t *testing.T) {
    tests := []struct {
        name string
        err error
    }{
        {"cat", envNotSet},
        {"SHELL", nil},
    }

    for _, test := range tests {
        result, err := loadEnv(test.name)
        if err != test.err {
            t.Error(test.name, "expected", test.err, "but got", err)
        }
        if result == "" && err == nil {
            t.Error(test.name, "expected nonempty result")
        }
    }
}

func TestLoadEnvOrDefault(t *testing.T) {
    tests := []struct {
        name string
        def string
    }{
        {"cat", "default"},
    }

    for _, test := range tests {
        result := loadEnvOrDefault(test.name, test.def)
        if result != test.def {
            t.Error(test.name, "expected", test.def, "but got", result)
        }
    }
}

func TestLoadHome(t *testing.T) {
    result := getHome()
    if result == "" {
        t.Error("expected nonempty result")
    }
}

func TestXDG_DATA_HOME(t *testing.T) {
    result := XDG_DATA_HOME()
    if result != filepath.Join(getHome(), ".local", "share") {
        t.Error("expected", filepath.Join(getHome(), ".local", "share"), "but got", result)
    }
}

func TestXDG_CONFIG_HOME(t *testing.T) {
    result := XDG_CONFIG_HOME()
    if result != filepath.Join(getHome(), ".config") {
        t.Error("expected", filepath.Join(getHome(), ".config"), "but got", result)
    }
}

func TestXDG_CACHE_HOME(t *testing.T) {
    result := XDG_CACHE_HOME()
    if result != filepath.Join(getHome(), ".cache") {
        t.Error("expected", filepath.Join(getHome(), ".cache"), "but got", result)
    }
}
