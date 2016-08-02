

//----------------------------------------------------------------------------------------
//
//                                                         BlitzHere-Server : Scheduler.go
//                                                        Periodicaly runs scheduled tasks
//
//                                                                 E.B. Smith, August 2016
//                        -©- Copyright © 2014-2016 Edward Smith, all rights reserved. -©-
//
//----------------------------------------------------------------------------------------


package main


import (
    "sort"
    "time"
    "violent.blue/GoKit/Log"
)


//  Not done... Not thread safe (yet)...


//----------------------------------------------------------------------------------------
//
//                                                                               Scheduler
//
//----------------------------------------------------------------------------------------


//----------------------------------------------------------------------------------------
//                                                                           ScheduledItem
//----------------------------------------------------------------------------------------


type ScheduledItem struct {
    Interval    time.Duration
    Task        func()
    nextTime    time.Time
}
type ScheduledItems []ScheduledItem


func (si ScheduledItems) Len() int {
    return len(si)
}

func (si ScheduledItems) Swap(i, j int) {
    t := si[i]
    si[i] = si[j]
    si[j] = t
}

func (si ScheduledItems) Less(i, j int) bool {
    return si[i].nextTime.Before(si[j].nextTime)
}


//----------------------------------------------------------------------------------------
//                                                                       Scheduler Control
//----------------------------------------------------------------------------------------


var schedulerChannel chan bool
var scheduledItems   ScheduledItems    //  Switch array to heap....


func scheduler() {
    Log.LogFunctionName()
    defer Log.Debugf("=> Exit Scheduler <=")

    //  Runs a single scheduled item at a time.

    runScheduledItem := func(item *ScheduledItem) {
        item.Task()
        item.nextTime = time.Now().Add(item.Interval)
    }

    var shouldContinue bool = true
    for shouldContinue {

        var waitTime time.Duration = time.Second
        if len(scheduledItems) > 0 {
            sort.Sort(scheduledItems)
            waitTime = time.Since(scheduledItems[0].nextTime) * -1
        }

        if waitTime < 0 {
            runScheduledItem(&scheduledItems[0])
            continue
        }

        var timer *time.Timer = time.NewTimer(waitTime)
        select {
            case shouldContinue = <- schedulerChannel:
                Log.Debugf("Scheduler should continue: %v.", shouldContinue)

            case <- timer.C:
                if len(scheduledItems) > 0 {
                    runScheduledItem(&scheduledItems[0])
                }
        }
    }
}


func StartScheduler() {
    Log.LogFunctionName()
    schedulerChannel = make(chan bool)
    go scheduler()
}


func StopScheduler() {
    Log.LogFunctionName()
    schedulerChannel <- false
}


func ScheduleTask(frequency time.Duration, task func()) {
    item := ScheduledItem {
        Interval:   frequency,
        Task:       task,
        nextTime:   time.Now().Add(frequency),
    }
    scheduledItems = append(scheduledItems, item)
}

