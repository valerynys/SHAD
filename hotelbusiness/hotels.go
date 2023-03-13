package hotelbusiness

type Guest struct {
	CheckInDate  int
	CheckOutDate int
}

type Load struct {
	StartDate  int
	GuestCount int
}

func ComputeLoad(guests []Guest) []Load {
	load := []Load{}
	currDate := 0
	currCount := 0

	for _, guest := range guests {
		// check-out
		if guest.CheckOutDate == currDate {
			currCount--
		}

		// check-in
		if guest.CheckInDate == currDate {
			currCount++
		}

		// update load if count changes
		if currCount != 0 {
			load = append(load, Load{StartDate: currDate, GuestCount: currCount})
		}

		// update date
		currDate = guest.CheckOutDate
	}

	// last date
	if currCount != 0 {
		load = append(load, Load{StartDate: currDate, GuestCount: currCount})
	}

	return load
}
