package main

import (
    "math/rand"
    "github.com/brianvoe/gofakeit"
    "os"
    "fmt"
    "github.com/pquerna/ffjson/ffjson"
)

type Event struct {
    Step string `json:"step"`
    Substep string `json:"substep"`
    Ip string
    Payload interface{}
}

type NeEventsPayload struct {
    NeToken string `json:"ne_token"`
    Event string `json:"event"`
}

type MkEventsPayload struct {
    NeToken string `json:"ne_token"`
}

const (
    NeEvent = "ne_event"
    MkEvent = "mk_event"
)

func main() {
    rand.Seed(42)
    count := 1000000
    neEventsCountMax := count / 100
    types := []string{
        NeEvent,
        MkEvent,
    }
    neEventsCount := 0
    var events []Event
    for i := 0; i < count; i++ {
        eventType := types[rand.Intn(len(types))]
        if eventType == NeEvent && neEventsCount < neEventsCountMax {
            substep := gofakeit.JobTitle()
            event := Event{
                Step:eventType,
                Substep:substep,
                Ip:gofakeit.IPv4Address(),
                Payload:NeEventsPayload{
                    NeToken:gofakeit.Password(true,false,true,false,false,64),
                    Event:substep,
                },
            }
            events = append(events, event)
            neEventsCount++
        } else {
            substep := gofakeit.JobTitle()
            event := Event{
                Step:MkEvent,
                Substep:substep,
                Ip:gofakeit.IPv4Address(),
                Payload:MkEventsPayload{
                    NeToken:gofakeit.Password(true,false,true,false,false,64),
                },
            }
            events = append(events, event)
        }
    }

    f, err := os.Create(fmt.Sprintf("events.%d.json",count))
    if err != nil {
        panic(err)
    }
    defer f.Close()
    bytes, err := ffjson.Marshal(&events)
    if err != nil {
        panic(err)
    }
    n, err := f.Write(bytes)
    if err != nil {
        panic(err)
    }

    fmt.Printf("\nSuccessfully write %d bytes\n", n)
}
