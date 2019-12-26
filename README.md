# Eggcrate

[![Build Status](https://travis-ci.org/devplayg/eggcrate.svg?branch=master)](https://travis-ci.org/devplayg/eggcrate)

EggCrate embeds static files into Go source encoded base64

Encode

```go
_, err := Encode("/data", "js,css,png", "asset.go")
```

Decode

```
var fileMap map[string][]byte
fileMap, _ := Decode("H4sIAAAAAAAA/6p...")
```

Directory structure

```
/data
    assets
        ├─css
        │      custom.css
        ├─img
        │      logo.png
        └─js
               custom.js
```

Map structure

|Key (string)|Value ([]byte)|
|---|---|
|/assets/css/custom.css | []byte |
|/assets/img/logo.png | []byte |
|/assets/js/custom.js | []byte |