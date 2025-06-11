package uwuid

var uwuids = [6]string{"UwU", "OwO", "TwT", ":3", ">w<", "^w^"}

var seed uint32 = 1

func random() int {
	seed = (seed*1664525 + 1013904223) % 6
	return int(seed)
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
