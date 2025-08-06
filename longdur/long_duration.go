package longdur

import "time"

//nolint:recvcheck
type Duration struct {
	Years  int
	Months int
	Days   int
}

// AddTo adds duration to the provided time using standard time.AddDate function.
func (d Duration) AddTo(t time.Time) time.Time {
	return t.AddDate(d.Years, d.Months, d.Days)
}

// SubFrom subtracts duration from the provided time using standard time.AddDate function.
func (d Duration) SubFrom(t time.Time) time.Time {
	return t.AddDate(-d.Years, -d.Months, -d.Days)
}
