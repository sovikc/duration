# duration
Time interval in days between two dates

1. The code is written using Golang without using ```time``` package. So it should be installed before one can execute the code. I have used ```go.mod``` to manage dependencies.
2. The code calculates the difference in days between two dates using Julian day number. The logic and the background is in the comment of the ```duration.go``` file.
3. I have built a terminal user interface to guide the user through the process of entering dates and showing the interval in days between those two dates.
4. To run the program create a github layout structure, and use the command ```go run cmd/main.go``` as shown in the screenshot of the terminal below.

![alt text](https://github.com/sovikc/duration/blob/master/tui_screenshot.png)
