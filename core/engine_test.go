package core

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	"uptoc/uploader"
)

type mockUploader struct {
}

func (m mockUploader) ListObjects() ([]uploader.Object, error) {
	return []uploader.Object{
		{Key: "/abc1.txt", ETag: "abc123"},
		{Key: "/abc4.txt", ETag: "aaa123"},
	}, nil
}

func (m mockUploader) Upload(object, rawPath string) error {
	return nil
}

func (m mockUploader) Delete(object string) error {
	return nil
}

func TestEngine(t *testing.T) {
	// init
	tmp := "/tmp/uptoc/"
	files := map[string]string{
		"abc1.txt": "abcabcabc",
		"abc2.txt": "112233",
		"abc3.txt": "445566",
	}
	assert.NoError(t, os.Mkdir(tmp, os.FileMode(0755)))
	for name, content := range files {
		assert.NoError(t, ioutil.WriteFile(tmp+name, []byte(content), os.FileMode(0644)))
	}

	// test
	e := NewEngine(&mockUploader{})
	assert.NoError(t, e.LoadAndCompareObjects("/tmp/uptoc"))
	assert.NoError(t, e.Sync())

	// clean the test files.
	assert.NoError(t, os.RemoveAll(tmp))
}
