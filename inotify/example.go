package main

import (
  "log"
  "github.com/howeyc/fsnotify"
)

func main() {
  // fsnotify初始化
  watcher, err := fsnotify.NewWatcher()
  if err != nil {
    log.Fatal(err)
  }

  // 创建chan用于通知
  done := make(chan bool)

  // 开了一个协程for ever监听event等待事件
  go func() {
    for {
      select {
      case ev := <-watcher.Event:
        log.Println("event:",ev)
      case err := <-watcher.Error:
        log.Println("error:",err)
      }
    }
  }()

  // 指定监控目录
  err = watcher.Watch("testDir")
  if err != nil {
    log.Fatal(err)
  }

  <- done // 等待程序结束

  watcher.Close()
}
