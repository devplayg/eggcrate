# eggcrate
EggCrate embeds static files into Go source encoded base64

Encode

```go
_, err := Encode("/data", "js,css,png", "asset.go")
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


Decode

```
fileMap, err := Decode("H4sIAAAAAAAA/6p...")
```

fileMap: map[string][]byte

|Key (string)|Value ([]byte)|
|---|---|
|/assets/css/custom.css | []byte |
|/assets/img/logo.png | []byte |
|/assets/js/custom.js | []byte |