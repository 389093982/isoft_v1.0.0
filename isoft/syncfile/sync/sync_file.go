package sync

import (
	"encoding/xml"
	"isoft/isoft/common/fileutil"
	"io/ioutil"
	"log"
	"os"
	"fmt"
)

type SyncFile struct {
	XMLName xml.Name `xml:"syncfile"` // 指定最外层的标签为 syncfile
	Source  string   `xml:"source"`   // 读取source配置项,并将结果保存到Source变量中
	Targets []Target `xml:"target"`   // 读取target标签下的内容,以结构方式
}

type Target struct {
	Name  string `xml:"name"`
	Value string `xml:"value"`
}

func ReadSyncFile(filepath string) (syncFile SyncFile) {
	content, err := ioutil.ReadFile(filepath)
	if err != nil {
		log.Fatal(err)
	}
	err = xml.Unmarshal(content, &syncFile)
	if err != nil {
		log.Fatal(err)
	}
	return syncFile
}

// 同步所有目录
func StartAllSyncFile(syncFile SyncFile, filterTargetName string) {
	source := syncFile.Source
	targets := syncFile.Targets
	for _, target := range targets {
		if filterTargetName == "" || (filterTargetName != "" && filterTargetName == target.Name) {
			// 开启协程执行任务
			go StartOneSyncFile(source, target.Value)
		}

	}
}

// 开始同步一个目录
func StartOneSyncFile(source, target string) {
	// 判断源文件夹是否存在
	if exist, err := fileutil.PathExists(source); exist == false {
		fmt.Println(err.Error())
	}
	// 判断目标文件夹是否存在,存在则进行删除
	if exist, err := fileutil.PathExists(target); exist == true {
		if err = os.RemoveAll(target); err != nil {
			fmt.Println(err.Error())
		} else {
			fmt.Println("clean dir %s", target)
		}
	}
	// 拷贝文件
	err := fileutil.CopyDir(source, target)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("copy dir %s to %s", source, target)
	}
}
