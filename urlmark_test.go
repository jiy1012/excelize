package excelize

import (
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"

	_ "golang.org/x/image/tiff"

	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddUrlMark(t *testing.T) {
	f, err := OpenFile(filepath.Join("test", "blank.xlsx"))
	if !assert.NoError(t, err) {
		t.FailNow()
	}

	// Test add url mark to worksheet
	assert.NoError(t, f.AddUrlMark("Sheet1", "JIY1012", "https://ffstudio.free.beeceptor.com/stat"))
	// Test write file to given path.
	assert.NoError(t, f.SaveAs(filepath.Join("test", "TestBlankAddUrlMark.xlsx")))
	assert.NoError(t, f.Close())
}
