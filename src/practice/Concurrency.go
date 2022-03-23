package practice

func AddValue(i *int64, trigger *bool, channel chan<- string) {
	for n := 0; n < 100; n++ {
		if *trigger {
			*i = *i + 1
			*trigger = false
		} else {
			*trigger = true
		}
	}
	channel <- "ok"
}

func HaveATest(i *int64, trigger *bool, channel chan<- string) {
	go AddValue(i, trigger, channel)
	go AddValue(i, trigger, channel)
	go AddValue(i, trigger, channel)
}
