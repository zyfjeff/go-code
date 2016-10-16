package mp

import (
    "fmt"
    "time"
)

type MP3Player struct {
    stat int
    progess int
}

func (p *MP3Player) Play(source string) {
    fmt.Println("Playing MP3 music",source)
    p.progess = 0
    for p.progess < 100 {
        time.Sleep(100 * time.Millisecond) //假装正在播放
        fmt.Print(".")
        p.progess += 10
    }
    fmt.Println("\nFinished playing",source)
}


