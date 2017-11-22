package main

import "fmt"
import "time"

func main() {
    p := fmt.Println

    now := time.Now()
    t1 := now.AddDate(0, 0, -15)
    p(now.Format(time.ANSIC))
    p(t1.Format(time.ANSIC))

    p(now.Format(time.RFC3339))

    t2, e := time.Parse(time.RFC3339, "2012-11-01T22:08:41+00:00")
    p(t2)
    p(e)

    p(now.Format("3:04PM"))
    p(now.Format("Mon, Jan _2 15:04:05 2006"))
    p(now.Format("2006-01-02T15:04:05.999999-07:00"))

    form := "3 04 PM"
    t3, e := time.Parse(form, "8 41 PM")
    p(t3)
    p(e)
    fmt.Printf("%d-%02d-%02dT%02d:%02d:%02d-00:00\n",
        now.Year(), now.Month(), now.Day(),
        now.Hour(), now.Minute(), now.Second())
    p("----------------")

    ansic := "Mon Jan _2 15:04:05 2006"
    t4, e := time.Parse(ansic, "Wed Nov 22 11:21:38 2017")
    p(t4)
    p(e)
}
