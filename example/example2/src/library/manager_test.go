package library

import (
    "testing"
)


func TestOps(t *testing.T) {
    //测试MusicManager创建
    mm := NewMusicManager()
    if mm == nil {
        t.Error("NewMusicManager failed.")
    }
    //测试Len方法
    if mm.Len() != 0 {
        t.Error("NewMusicManager failed,not empty.")
    }


    m0 := &MusicEntry {
        "1","My Heart Will Go On","Celion Dion","Pop",
        "Http://qbox.me/24501234","MP3"}
    //测试Add方法
    mm.Add(m0)
    if mm.Len() != 1 {
        t.Error("MusicManager.Add() failed.")
    }
    //测试Find方法
    m := mm.Find(m0.Name)
    if m == nil {
        t.Error("MusicManager.Find() failed.")
    }

    if m.Id != m0.Id || m.Artist != m0.Artist ||
       m.Name != m0.Name || m.Genre != m0.Genre ||
       m.Source != m0.Source || m.Type != m0.Type {
        t.Error("MusicManager.Find() failed.Found item mismatch.")
    }
    //测试Remove方法
    m = mm.Remove(0)
    if m == nil || mm.Len() != 0 {
        t.Error("MusicManager.Remove() failed.",err)
    }
}
