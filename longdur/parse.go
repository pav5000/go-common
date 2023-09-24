package longdur

import (
	"errors"
	"regexp"
	"strconv"
	"strings"
)

var durRe = regexp.MustCompile(`(\d+)(\S)`)

func Parse(str string) (Duration, error) {
	str = strings.TrimSpace(str)
	if str == "" {
		return Duration{}, nil
	}

	dur := Duration{}

	matches := durRe.FindAllStringSubmatch(str, -1)
	for _, match := range matches {
		if len(match) < 3 {
			continue
		}

		number, err := strconv.Atoi(match[1])
		if err != nil {
			return dur, errors.New("cannot parse number '" + match[1] + "'")
		}

		switch match[2] {
		case "y":
			if dur.Years != 0 {
				return dur, errors.New("duplicate group '" + match[0] + "'")
			}
			dur.Years = number
		case "m":
			if dur.Months != 0 {
				return dur, errors.New("duplicate group '" + match[0] + "'")
			}
			dur.Months = number
		case "d":
			if dur.Days != 0 {
				return dur, errors.New("duplicate group '" + match[0] + "'")
			}
			dur.Days = number
		default:
			return dur, errors.New("unknown suffix '" + match[2] + "' in the group '" + match[0] + "'")
		}
	}

	return dur, nil
}
