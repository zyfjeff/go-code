package mp

import "testing"

func TestMp(t *testing.T) {
    mp := &MP3Player{0,0}
    mp.play("test")
    if mp.progess != 100 {
        t.Error("MP3Player play error")
    }
}
