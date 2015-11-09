package uuid

import (
    "time"
    "strconv"
)

type UUID struct {
    prefix string
    c *SyncCounter
}

var globalCount=&SyncCounter{
    counter: 0,
}

func AUUID(nodeid int) *UUID {
    return &UUID{
        prefix: strconv.Itoa(nodeid)+"~"+strconv.FormatInt(globalCount.Inc(), 36)+
            "~"+strconv.FormatInt(time.Now().UnixNano(), 36)+"~",
        c: &SyncCounter{
            counter: 0,
        },
    }
}

func (this *UUID)Get() string {
    return this.prefix+strconv.FormatInt(this.c.Inc(), 36)
}
