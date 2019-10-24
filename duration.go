package duration

//  Subtracting the "start date" from the "end date" will have to return a time interval.
//
//  The challenge of calculating a difference between dates is that there are leap years.
//  These leap years are a result of the fact that a year has 365.25 days and the 0.25 days adds up over time.
//
//  So this happens every 4 years and we end up having 366 days on the 4th year.
//  Hence 3 normal years of 365 and then a leap year of 366 days. This will make the calculation hard.
//
//  As per wikipedia, https://en.wikipedia.org/wiki/Year
//  1 year = 365 days 5 hours 48 mins 45 secs
//
//  Now, 1/4th day could be equal to 6 hours, so a year is actually (365 + 1/4) days - 11 mins and 15 secs
//  So 4 years would be equal to 3 Normal years + 1 leap year - 4 * (11 mins 15 secs)
//
//  4 years = (3 * 365) + 366 days - 45 mins (or 0.75 hours)
//
//  Now 100 years = 25 * 3 Normal years + 25 * 1 leap years - 25 * 0.75 hours
// 							 = 75 Normal years + 25 leap years - 18.75 hours (1 day - 5.25 hours)
// 							 = 76 Normal years + 24 leap years + 5.25 hours
//
//  Then 400 years = 4 * 76 Normal years + 4 * 24 leap years + 4 * 5.25 hours
// 								= 304 Normal years + 96 leap years + 21 hours
// 								= 303 Normal years + 97 leap years - 3 hours
//
//  An easy way to perform the above subtraction of dates would be convert the start and end dates to a so-called "epoch time",
//  which is a number of days that have elapsed since a particular reference time, called the "epoch".
//
//  Julian day is an example of an epoch system. The epoch date in that system is 1st January 4713 BC.
//
//  There are 2 algorithms for converting Gregorian calendars to Julian Day Number
//  1. By Henry F. Fliegel and Thomas C. Van Flandern (also in https://en.wikipedia.org/wiki/Julian_day)
//     JDN = (1461 × (Y + 4800 + (M − 14)/12))/4 +(367 × (M − 2 − 12 × ((M − 14)/12)))/12 − (3 × ((Y + 4900 + (M - 14)/12)/100))/4 + D − 32075
//
//  2. University of Texas http://www.cs.utsa.edu/~cs1063/projects/Spring2011/Project1/jdn-explanation.html and
// 		https://www.tondering.dk/claus/cal/julperiod.php
//
// 		"The algorithm works fine for AD dates. If you want to use it for BC dates, you must first convert the BC year to a negative year (e.g., 10 BC = -9).
// 		The algorithm works correctly for all dates after 4800 BC, i.e. at least for all positive Julian Day."
//
//  Since I am going to support dates from 1900 to 2999 I can easily use the second one.
//
//  The computation of Julian Day number from a calendar date should have 3 parts
//  1. Finding number of days in the whole years upto the date
//  2. Finding number of days in completed months upto the date
//  3. Adding these numbers together
//
//  The number of days elapsed in the year of the date depends, except in months of January and February, on whether or not the year is a leap year.
//  Thus both the year number and the month number are required to evaluate this term. So in case Gregorian calendar, this would be very complex.
//  If, however, the calculations are performed using a calendar with the year beginning on 1st March, these could be pushed to the end of the year.

// Duration represents the elapsed days between two dates
type Duration struct {
	Start string
	End   string
	start *datePart
	end   *datePart
}

// New creates a new Duration object with valid start and end dates
func New(start, end string) (*Duration, error) {
	if validStartDate, err := isValidFormat(start, startDateLabel); !validStartDate {
		return nil, err
	}

	if validEndDate, err := isValidFormat(end, endDateLabel); !validEndDate {
		return nil, err
	}

	startDatePart, err := getDatePart(start, startDateLabel)
	if err != nil {
		return nil, err
	}

	endDatePart, err := getDatePart(end, endDateLabel)
	if err != nil {
		return nil, err
	}

	return &Duration{
		Start: start,
		End:   end,
		start: startDatePart,
		end:   endDatePart,
	}, nil

}

type datePart struct {
	dd   int
	mm   int
	yyyy int
}

// GetDays returns the number of days elapsed between start and end dates
func (d *Duration) GetDays() (int, error) {
	startJulianDay := getJulianDay(d.start.dd, d.start.mm, d.start.yyyy)
	endJulianDay := getJulianDay(d.end.dd, d.end.mm, d.end.yyyy)
	interval := getInterval(startJulianDay, endJulianDay)
	return excludeStartAndEndDays(interval), nil
}

func getInterval(start, end int) int {
	interval := end - start

	if interval < 0 {
		interval = -interval
	}

	return interval
}

func excludeStartAndEndDays(interval int) int {
	return interval - 1
}

func getJulianDay(dd, mm, yyyy int) int {
	a := (14 - mm) / 12
	// will result in a 1 for January (month 1) and February (month 2).  The result is 0 for the other 10 months.

	y := yyyy + 4800 - a
	// adds 4800 to the year so that we will start counting years from the year -4800
	// 1 corresponds to 1 AD, year 0 corresponds to 1 BC, year -1 corresponds to 2 BC, and so on.
	// There is no year between 1 AD and 1 BC

	m := mm + 12*a - 3
	// results in a 10 for January, 11 for February, 0 for March, 1 for April, ..., and a 9 for December.
	// This is because a is 1 for January and February and 0 for the other months.
	// The effect of the combined calculation of y and m is to pretend that the year begins in March and ends in February.

	/*
		  dd =  Adding day, the day of the month. Each increment to day increments the number of days since a fixed day.

			(153*m+2)/5 =
			The integer division (153m+ 2)/5 is a cleverly designed expression to calculate the number of days
			in the previous months (where March corresponds to m=0)

			For example, June corresponds to a value of 3 for m.
			For this value of m, the expression results in a value of 461/5 = 92 using integer division,
			which are the total number of days in March, April, and May

		  y*365 = Each non-leap-year has 365 days

			y/4 - y/100 + y/400 =
			Calculates the number of leap years since the year -4800 (which corresponds to a value of 0 for y).
			There is a leap year every year that is divisible by 4, except for years that are divisible by 100, but not divisible by 400.
			The number of leap years is, of course, the same as the number of leap days that need to be added in.

			-32045 =
			Ensures that the result will be 0 for January 1, 4713 BCE.

	*/

	return dd + (153*m+2)/5 + y*365 + y/4 - y/100 + y/400 - 32045
}
