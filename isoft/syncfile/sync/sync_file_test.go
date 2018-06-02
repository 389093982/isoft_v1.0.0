package sync

import (
	"testing"
	"os"
	"path/filepath"
)

func Test_SyncFile_Static(t *testing.T) {
	gopath := os.Getenv("GOPATH")
	SyncFile := ReadSyncFile( filepath.Join(gopath, "src/isoft/isoft/syncfile/sync/static.xml"))
	StartAllSyncFile(gopath, SyncFile, "")
}

func Test_SyncOneFile_Static(t *testing.T) {
	gopath := os.Getenv("GOPATH")
	SyncFile := ReadSyncFile( filepath.Join(gopath, "src/isoft/isoft/syncfile/sync/static.xml"))
	StartAllSyncFile(gopath, SyncFile, "isoft_deploy_web")
}
