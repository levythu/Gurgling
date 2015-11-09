package uuid

import (
    "time"
    "strconv"
)

type UUID struct {
    prefix string
    c *SyncCounter
}

func AUUID(nodeid int) *UUID {
    return &UUID{
        prefix: strconv.Itoa(nodeid)+"~"+strconv.FormatInt(time.Now().UnixNano(), 36)+"~",
        c: &SyncCounter{
            counter: 0,
        },
    }
}

func (*UUID)Get() string {
    return prefix+strconv.FormatInt(c.Inc(), 36)
}
