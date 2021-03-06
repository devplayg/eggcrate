# eggcrate

[![Build Status](https://travis-ci.org/devplayg/eggcrate.svg?branch=master)](https://travis-ci.org/devplayg/eggcrate)

The `eggcrate` compresses static files and encodes them in Base64 and saves them as Go source files.

## Encode


Static files in the directory

```
`-- /tmp/static
    |-- css
    |   |-- app.css
    |   `-- fa.css
    `-- js
        `-- app.js
```

### Encoding

```go
config := eggcrate.Config{
    Dir:        "/tmp/static",
    OutFile:    "output.go",
    UriPrefix:  "/assets",
    Extensions: "js, css",
}
_, err := eggcrate.Encode(&config)
if err != nil {
    fmt.Printf(err.Error())
}
```

output.go

```go
/*
    /assets/css/app.css
    /assets/css/fa.css
    /assets/js/app.js
*/

package main

var assetData=`Dv+BBAEC/...
```



## Decoding

Basic decoding

```
var fileMap map[string][]byte
fileMap, _ = eggcrate.Decode(assetData) 
```

Decoding and using with http/net

```go
package main

import (
	"github.com/devplayg/eggcrate"
	"net/http"
)

var staticMap map[string][]byte

func main() {
	m, err := eggcrate.Decode(assetData)
	if err != nil {
		panic(err)
	}
	staticMap = m

	http.HandleFunc("/assets/js/app.js", jsHandler)
	http.HandleFunc("/assets/css/app.css", cssHandler)
	http.HandleFunc("/assets/css/fa.css", cssHandler)

	http.ListenAndServe(":8000", nil)

}

func jsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Encoding", "gzip")
	w.Header().Set("Content-Type", "text/javascript")
	w.Write(staticMap[r.RequestURI])
}

func cssHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Encoding", "gzip")
	w.Header().Set("Content-Type", "text/css")
	w.Write(staticMap[r.RequestURI])
}

```


## examples

encoding

- https://github.com/devplayg/eggcrate/blob/master/examples/encode/encode.go

decoding

- https://github.com/devplayg/eggcrate/blob/master/examples/decode/decode.go