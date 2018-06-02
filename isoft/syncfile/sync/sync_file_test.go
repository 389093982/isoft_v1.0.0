package sync

import (
	"testing"
	"os"
	"path/filepath"
)

// go test –file sync_file_test.go -run='Test_SyncFile_Static'
func Test_SyncFile_Static(t *testing.T) {
	gopath := os.Getenv("GOPATH")
	SyncFile := ReadSyncFile( filepath.Join(gopath, "src/isoft/syncfile/sync/static.xml"))
	StartAllSyncFile(SyncFile, "")
}
