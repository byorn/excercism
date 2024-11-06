package clock

import (
	"strconv"
)

// Define the Clock type here.
type Clock struct {
    h int
    m int
}

func New(h, m int) Clock {
    hour := 0
    if h>24{
            hr := h / 24
            hour = hr
    } else {
		    if h == 24 {
                hour = 0
            }
            hour = h
    }

    if m > 60 {

    }
}

func (c Clock) Add(m int) Clock {
	if (m<60){
        c.m += m
    }else{
        hr := m / 60
        mi := m % 60
        c.h += hr
        c.m += mi
    }
    return c
}

func (c Clock) Subtract(m int) Clock {
   if (m<60){
       c.m -= m
   }else{
       hr := m / 60
       mi := m % 60
       c.h -= hr
       c.m -= mi
   }
    return c
}

func (c Clock) String() string {
    str:=""
    if c.h < 10{
        str="0"+strconv.Itoa(c.h)
    }else{

        str = strconv.Itoa(c.h)


    }
    if c.m < 10{
        str+=":0"+strconv.Itoa(c.m)
    }else{
        str += ":"+strconv.Itoa(c.m)
    }
    return str
}
