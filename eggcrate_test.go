package eggcrate

import (
    "bytes"
    "github.com/devplayg/goutils"
    "io/ioutil"
    "os"
    "path/filepath"
    "testing"
)

func TestEncodeContentToBase64(t *testing.T) {
    extJs := ".js"
    extCss := ".css"
    assetFile := "asset.go"

    dir, err := ioutil.TempDir("", "eggcrate-encode")
    if err != nil {
        t.Error(err)
    }
    defer func() {
        os.RemoveAll(dir)
    }()
    dir = filepath.ToSlash(dir)

    jsFile, err := ioutil.TempFile(dir, "eggcrate-encode")
    if err != nil {
        t.Error(err)
    }
    jsText := "alert('hello');"
    _, err = jsFile.WriteString(jsText)
    if err != nil {
        t.Error(err)
    }
    if err := jsFile.Close(); err != nil {
        t.Error(err)
    }
    if err := os.Rename(jsFile.Name(), filepath.Join(dir, filepath.Base(jsFile.Name())+extJs)); err != nil {
        t.Error(err)
    }
    defer func() {
        if err := os.Remove(filepath.Join(dir, filepath.Base(jsFile.Name())+extJs)); err != nil {
           t.Error(err)
        }
    }()

    cssFile, err := ioutil.TempFile(dir, "eggcrate-encode")
    if err != nil {
        t.Error(err)
    }
    cssText := "body { color: red; }"
    _, err = cssFile.WriteString(cssText)
    if err != nil {
        t.Error(err)
    }
    if err := cssFile.Close(); err != nil {
        t.Error(err)
    }
    if err := os.Rename(cssFile.Name(), filepath.Join(dir, filepath.Base(cssFile.Name())+extCss)); err != nil {
        t.Error(err)
    }
    defer func() {
        if err := os.Remove(filepath.Join(dir, filepath.Base(cssFile.Name())+extCss)); err != nil {
           t.Error(err)
        }
    }()

    config := Config{
        Dir:        dir,
        OutFile:    filepath.Join(dir, assetFile),
        UriPrefix:  "",
        Extensions: "js,css",
    }
    if _, err = Encode(&config); err != nil {
        t.Error(err)
    }

    fileMap, err := Decode(`Dv+BBAEC/4IAAQwBCgAA/5T/ggACHS9lZ2djcmF0ZS1lbmNvZGU0ODA0MzU1NDAuY3NzLB+LCAAAAAAAAP9Kyk
+pVKhWSM7PyS+yUihKTbFWqAUEAAD//4Gv1I8UAAAAHC9lZ2djcmF0ZS1lbmNvZGU4MDg1ODQwNzMuanMnH4sIAAAAAAAA/0rMSS0q0VDPSM3JyVfXtAYEAAD//y80DPcPAAAA`)
    if err != nil {
        t.Error(err)
    }
    key := "/eggcrate-encode480435540.css"
    found, exists := fileMap[key]
    if !exists {
        t.Error("key does not exists; " + key)
    }

    expected, err := goutils.Gzip([]byte(cssText))
    if err != nil {
        t.Error("compress error")
    }
    if !bytes.Equal(found, expected) {
        t.Error("not same")
    }
}
