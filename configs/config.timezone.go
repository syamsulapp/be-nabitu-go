package configs

import "time"

var location *time.Location

func InitConfigTimeZone(TimeZone string) error {
	locTimezone, err := time.LoadLocation(TimeZone)
	if err != nil {
		panic(err.Error())
	}

	location = locTimezone
	return nil
}

func GetConfigTimeZone(t time.Time) time.Time {
	return t.In(location)
}
