package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/sovikc/duration"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("---------------------------------------------------------")
	fmt.Println("Please enter experiment dates to get the days elapsed.")
	fmt.Println("To proceed please input valid start and end dates")
	fmt.Println("beginning with the start date, followed by the end date")
	fmt.Println("between 01/01/1901 and 31/12/2999 in a DD/MM/YYYY format")
	fmt.Println("---------------------------------------------------------")

	var datesEntered int
	var start, end string

	for {
		fmt.Print("-> ")
		option, _ := reader.ReadString('\n')
		datesEntered++

		// convert CRLF to LF
		option = strings.Replace(option, "\n", "", -1)

		if datesEntered == 1 {
			//fmt.Println("start date", option)
			start = option
		}

		if datesEntered == 2 {
			//fmt.Println("end date", option)
			end = option
		}

		if datesEntered == 2 {
			datesEntered = 0
			d, err := duration.New(start, end)
			if err != nil {
				fmt.Println(err)
				return
			}

			elapsed, err := d.GetDays()
			if err != nil {
				fmt.Println(err)
				return
			}

			fmt.Println("========================")
			fmt.Println(elapsed, "days")
			fmt.Println("========================")

		}

	}

}
