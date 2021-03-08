package main

import (
	"bufio"
	"fmt"

	// "log"
	// "log/syslog"
	"os"
)

func main() {

	var f float32 = 10.4
	i := int(f)
	fmt.Printf("Value of i %d and type %T ", i, i)
	var y string
	fmt.Println(" \n Enter value")
	fmt.Scanf("%s", &y)
	fmt.Printf("Value of Y %s", y)
	reader := bufio.NewReader(os.Stdin)
	name, _, ok := reader.ReadLine()
	if ok == nil {
		fmt.Printf("Name %s", name)
	}
	// syslogger, _ := syslog.NewLogger(syslog.LOG_ALERT, log.Flags())
	// syslogger("Name %s", name)

}
