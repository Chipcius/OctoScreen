package utils

import (
    "fmt"
    "time"
)

func FormatDuration(dInput time.Duration) string {
    duration := dInput.Round(time.Second)
    d        := duration / (time.Hour * 24)
    duration -= d * (time.Hour * 24)
    h        := duration / time.Hour
    duration -= h * time.Hour
    m        := duration / time.Minute
    duration -= m * time.Minute
    s        := duration / time.Second
    if(d > 1){
    	return fmt.Sprintf("%02dd %02dh %02dm %02ds", d, h, m, s)
    }
    if(h > 1){
    	return fmt.Sprintf("%02dh %02dm %02ds", h, m, s)
    }

    return fmt.Sprintf("%02dm %02ds", m, s)
}