package uwuid

import (
    "math/rand"
)
var uwuids = [6]string{"UwU", "OwO", "TwT", ":3", ">w<", "^w^"}

func random() int {
    return rand.Intn(len(uwuids))
}

func New() string {
    uwid := ""
    for i := 0; i < 6; i++ {
        uwid += uwuids[random()]
        if i == 0 {
            uwid += uwuids[random()]
        }
        if i < 5 {
            uwid += "-"
        } else {
            uwid += uwuids[random()]
            uwid += uwuids[random()]
        }
    }
    return uwid
}
