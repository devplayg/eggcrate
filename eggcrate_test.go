package eggcrate

import (
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
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
		//os.RemoveAll(dir)
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
		//assert.Nil(os.Remove(jsFile.Name()+extJs))
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

	str, err := Encode(dir, "js,css", filepath.Join(dir, assetFile))
	assert.Nil(err)

	fileMap, err := Decode(*str)
	assert.Nil(err)

	key := strings.TrimPrefix(filepath.ToSlash(jsFile.Name())+extJs, dir)
	jsVal, exists := fileMap[key]
	if !exists {
		assert.Fail("xxx")

	}
	assert.Equal([]byte(jsText), jsVal)

	key = strings.TrimPrefix(filepath.ToSlash(cssFile.Name())+extCss, dir)
	jsCss, exists := fileMap[key]
	if !exists {
		assert.Fail("xxx")

	}
	assert.Equal([]byte(cssText), jsCss)

}
