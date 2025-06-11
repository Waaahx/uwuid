package uwuid

import (
    "math/rand"
)
var uwuids = [6]string{"UwU", "OwO", "TwT", ":3", ">w<", "^w^"}

func random() int {
    return rand.Intn(len(uwuids))
}

func New() string {
    uwuid := ""
    for i := 0; i < 6; i++ {
        uwuid += uwuids[random()]
        if i == 0 {
            uwuid += uwuids[random()]
        }
        if i < 5 {
            uwuid += "-"
        } else {
            uwuid += uwuids[random()]
            uwuid += uwuids[random()]
        }
    }
    return uwuid
}
