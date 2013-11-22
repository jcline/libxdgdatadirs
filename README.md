libxdgdatadirs
==============

A simple go library implementing the XDG Base Directory Specification

##Usage

Add, to your source:

```go
import (
    xdg "github.com/jcline/libxdgdatadirs"
)
...

func main() {
    conf, err := xdg.LoadOrCreate(xdg.XDG_CONFIG_HOME, "program_name")
    if err != nil {
        return
    }

    conf_path := filepath.Join(conf, "config")
    file, err := ioutil.readFile(conf_path)
    ...
}
```

Then when you run `go build` for your program, it will automatically pull the library and build it.

##Notes

* XDG_*_DIRS aren't fully supported at this time.
  * XDG_CONFIG_HOME does not fall back to XDG_CONFIG_DIRS
  * XDG_DATA_HOME does not fall back to XDG_DATA_DIRS
