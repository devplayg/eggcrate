package eggcrate

import (
	"github.com/devplayg/golibs/compress"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

func TestEncodeContentToBase64(t *testing.T) {
	assert := assert.New(t)
	extJs := ".js"
	extCss := ".css"
	assetFile := "asset.go"

	dir, err := ioutil.TempDir("", "eggcrate-encode")
	assert.Nil(err)
	defer func() {
		os.RemoveAll(dir)
	}()
	dir = filepath.ToSlash(dir)

	jsFile, err := ioutil.TempFile(dir, "eggcrate-encode")
	assert.Nil(err)
	jsText := "alert('hello');"
	_, err = jsFile.WriteString(jsText)
	assert.Nil(err)
	assert.Nil(jsFile.Close())
	assert.Nil(os.Rename(jsFile.Name(), filepath.Join(dir, filepath.Base(jsFile.Name())+extJs)))
	defer func() {
		assert.Nil(os.Remove(jsFile.Name() + extJs))
	}()

	cssFile, err := ioutil.TempFile(dir, "eggcrate-encode")
	assert.Nil(err)
	cssText := "body { color: red; }"
	_, err = cssFile.WriteString(cssText)
	assert.Nil(err)
	assert.Nil(cssFile.Close())
	assert.Nil(os.Rename(cssFile.Name(), filepath.Join(dir, filepath.Base(cssFile.Name())+extCss)))
	defer func() {
		//assert.Nil(os.Remove(jsFile.Name()+extCss))
	}()

	config := Config{
		Dir:        dir,
		OutFile:    filepath.Join(dir, assetFile),
		UriPrefix:  "",
		Extensions: "js,css",
	}
	_, err = Encode(&config)
	assert.Nil(err)

	fileMap, err := Decode(`Dv+BBAEC/4IAAQwBCgAA/5T/ggACHS9lZ2djcmF0ZS1lbmNvZGU0ODA0MzU1NDAuY3NzLB+LCAAAAAAAAP9Kyk
+pVKhWSM7PyS+yUihKTbFWqAUEAAD//4Gv1I8UAAAAHC9lZ2djcmF0ZS1lbmNvZGU4MDg1ODQwNzMuanMnH4sIAAAAAAAA/0rMSS0q0VDPSM3JyVfXtAYEAAD//y80DPcPAAAA`)
	assert.Nil(err)
	key := "/eggcrate-encode480435540.css"
	found, exists := fileMap[key]
	if !exists {
		assert.Fail("key does not exists; " + key)
	}

	expected, err := compress.Compress([]byte(cssText), compress.GZIP)
	if err != nil {
		assert.Fail("compress error")
	}
	assert.Equal(found, expected)
}
