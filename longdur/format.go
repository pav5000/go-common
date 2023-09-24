package longdur

import (
	"strconv"
	"strings"
)

func (d Duration) String() string {
	if d.Years <= 0 && d.Months <= 0 && d.Days <= 0 {
		return ""
	}

	res := strings.Builder{}
	res.Grow(12) // assuming that in 99.9% cases string won't be longer than "12y 34m 56d "

	if d.Years > 0 {
		res.WriteString(strconv.Itoa(d.Years))
		res.WriteString("y ")
	}
	if d.Months > 0 {
		res.WriteString(strconv.Itoa(d.Months))
		res.WriteString("m ")
	}
	if d.Days > 0 {
		res.WriteString(strconv.Itoa(d.Days))
		res.WriteString("d ")
	}

	str := res.String()
	lastIndex := len(str) - 1
	if str[lastIndex] == ' ' {
		str = str[:lastIndex]
	}

	return str
}
