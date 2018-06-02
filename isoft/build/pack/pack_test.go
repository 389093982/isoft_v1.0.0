package pack

import (
	"testing"
	"path/filepath"
	"os"
)

func Test_StartAllPack(t *testing.T) {
	gopath := os.Getenv("GOPATH")
	packApps := ReadPackApp(filepath.Join(gopath, "src/isoft/isoft/build/pack/pack.xml"))
	StartAllPack(&packApps, "")
}

func Test_StartOnePack(t *testing.T) {
	gopath := os.Getenv("GOPATH")
	packApps := ReadPackApp(filepath.Join(gopath, "src/isoft/isoft/build/pack/pack.xml"))
	StartAllPack(&packApps, "isoft_learning_web")
}

