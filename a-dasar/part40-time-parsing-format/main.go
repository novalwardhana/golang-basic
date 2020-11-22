package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Part 40 Time Parsing Format")
	fmt.Println("")

	fmt.Println("Part 40.1 Penggunaan time.Time")
	var01A := time.Now()
	fmt.Printf("var01A: %v\n", var01A)
	var01B := time.Date(2020, 9, 9, 11, 11, 23, 0, time.UTC)
	fmt.Printf("var01B: %s\n", var01B)
	var01C := time.Date(2019, 1, 2, 10, 15, 50, 0, time.UTC)
	fmt.Printf("var01C: %s\n", var01C)
	var01D := time.Date(2019, 12, 5, 14, 30, 0, 0, time.UTC)
	fmt.Printf("var01D: %s\n", var01D)
	fmt.Println("")

	fmt.Println("Part 40.2 Method milik time.Time")
	var02Year := time.Now().Year()
	fmt.Printf("var02 Year: %d\n", var02Year)
	var02YearDay := time.Now().YearDay() // Hari ke-? dalam tahun ini
	fmt.Printf("var02 Year Day: %d\n", var02YearDay)
	var02Month := time.Now().Month()
	fmt.Printf("var02 Month: %s\n", var02Month)
	var02WeekDay := time.Now().Weekday()
	fmt.Printf("Var02 Week Day: %s\n", var02WeekDay)
	var02IsoWeekYear, var02IsoWeek := time.Now().ISOWeek()
	fmt.Printf("var02 Iso Week Year: %d\n", var02IsoWeekYear)
	fmt.Printf("var02 Iso Week: %d\n", var02IsoWeek)
	var02Day := time.Now().Day()
	fmt.Printf("var02 Day: %d\n", var02Day)
	var02Hour := time.Now().Hour()
	fmt.Printf("var02 Hour: %d\n", var02Hour)
	var02Minute := time.Now().Minute()
	fmt.Printf("var02 Minute: %d\n", var02Minute)
	var02Second := time.Now().Minute()
	fmt.Printf("var02 Second: %d\n", var02Second)
	var02Nanosecond := time.Now().Nanosecond()
	fmt.Printf("var02 Nanosecond: %v\n", var02Nanosecond)
	var02Local := time.Now().Local()
	fmt.Printf("var02 Local: %s\n", var02Local)
	var02Location := time.Now().Location()
	fmt.Printf("var02 Location: %v\n", var02Location)
	var02TimeZone, var02TimeZoneOffset := time.Now().Zone()
	fmt.Printf("var02 Zone: %v\n", var02TimeZone)
	fmt.Printf("var02 Time Zone Offset: %v\n", var02TimeZoneOffset)
	var02IsZero := time.Now().IsZero()
	fmt.Printf("var02 Is Zero: %v\n", var02IsZero)
	var02UTC := time.Now().UTC()
	fmt.Printf("var02 UTC: %v\n", var02UTC)
	var02Unix := time.Now().Unix()
	fmt.Printf("var02 Unix: %v\n", var02Unix)
	var02UnixNano := time.Now().UnixNano()
	fmt.Printf("var02 Unixnano: %v\n", var02UnixNano)
	var02String := time.Now().String()
	fmt.Printf("var02 String: %s\n", var02String)
	fmt.Println("")

	fmt.Println("Part 40.3 Parsing dari string ke time.Time")
	var03A, _ := time.Parse("2006-01-02 15:04:05", "2020-09-09 11:48:56")
	fmt.Printf("var03A: %v\n", var03A)
	var03B, _ := time.Parse("2006-01-02", "2019-12-05")
	fmt.Printf("var03B: %v\n", var03B)
	var03C, _ := time.Parse("2006/January/02 15:04:05", "2020/September/09 11:52:58")
	fmt.Printf("var03C: %v\n", var03C)
	var03D, _ := time.Parse("2006/01/02 MST", "2019/05/20 WIB")
	fmt.Printf("var03D: %v\n", var03D)
	var03E, _ := time.Parse("02/January/2006 15:04:05 MST", "02/November/2019 11:59:59 WIB")
	fmt.Printf("var02E: %v\n", var03E)
	fmt.Println("")

	fmt.Println("Part 40.4 Predefined layout format")
	var04A, _ := time.Parse(time.RFC822, "09 Sep 20 22:01 WIB")
	fmt.Printf("var04A: %v\n", var04A)
	fmt.Println("")

	fmt.Println("Part 40.5 Parsing dari time.Time ke string")
	var05A, _ := time.Parse("2006/02/01", "2020/09/09")
	var05AStr := var05A.Format("Monday 02, January 2006")
	fmt.Printf("var05AStr: %s\n", var05AStr)
	var05B, _ := time.Parse("2006/02/02 15:04:05 MST", "2019/11/20 13:05:50 WIB")
	var05BStr := var05B.Format("Monday 02/January/06, 15:04:05 WIB")
	fmt.Printf("var05BStr: %s\n", var05BStr)
	var05C, _ := time.Parse("02/01/2006 15:04:05 MST", "20/05/2019 15:30:22 WIB")
	var05CStr := var05C.Format("Monday, 02-January-06, 15:04:05 WIB")
	fmt.Printf("var05CStr: %s\n", var05CStr)
	var05D, _ := time.Parse("01/02/06 15/04/05 MST", "11/19/18 10/22/50 WIB")
	var05DStr := var05D.Format("Monday 02, January 2006, 15:04:05 MST")
	fmt.Printf("var05DStr: %s\n", var05DStr)
	fmt.Println("")

	fmt.Println("Part 38.6 Handling parsing error time.Time")
	if var05A, err := time.Parse("2006-01-02", "2020/09/08"); err != nil {
		fmt.Printf("var05A err: %v\n", err)
	} else {
		fmt.Printf("var05A: %v\n", var05A)
	}
	if var05B, err := time.Parse("Monday, 02 January 06", "Wednesday, 09 September 2020"); err != nil {
		fmt.Printf("var05B err: %v\n", err)
	} else {
		fmt.Printf("var05B: %v\n", var05B)
	}
	if var05C, err := time.Parse("15:04:05 MST", "22:52:1000 WIBB"); err != nil {
		fmt.Printf("var05C err: %v\n", err)
	} else {
		fmt.Printf("var05C: %v\n", var05C)
	}
	if var05D, err := time.Parse("02/01/06", "Wednesday 2020/09/09"); err != nil {
		fmt.Printf("var05D err: %v\n", err)
	} else {
		fmt.Printf("var05D: %v\n", var05D)
	}
	if var05E, err := time.Parse("2006.01.02_15.04.05_MST", "2020.09.08_22.57.30_WIB"); err != nil {
		fmt.Printf("var05E err: %v\n", err)
	} else {
		fmt.Printf("var05E: %s\n", var05E.Format("2006.01.02_15.04.05_MST"))
	}

}
