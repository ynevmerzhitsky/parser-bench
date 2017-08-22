package main

import (
    "io/ioutil"
    "os"
    "testing"

    "fmt"
    "github.com/pquerna/ffjson/ffjson"
    "encoding/json"
)

type EventRawMessage struct {
    Step string `json:"step"`
    Substep string `json:"substep"`
    Ip string
    Payload *json.RawMessage
}

type EventInterfaceRawMessage struct {
    Step string `json:"step"`
    Substep string `json:"substep"`
    Ip string
    PayloadObj interface{} `json:"-"`
    Payload *json.RawMessage
}

var codeJSON []byte
var TestFunc = UnmarshalInterfaceWithRawMessage

func codeInit(n int) {
    f, err := os.Open(fmt.Sprintf("events.%d.json", n))
    if err != nil {
        panic(err)
    }
    defer f.Close()
    data, err := ioutil.ReadAll(f)
    if err != nil {
        panic(err)
    }

    codeJSON = data
}

func UnmarshalInterface() {
    var entities []Event
    if err := ffjson.Unmarshal(codeJSON, &entities); err != nil {
        panic(err)
    }
}

func UnmarshalRawMessage() {
    var entities []EventRawMessage
    if err := ffjson.Unmarshal(codeJSON, &entities); err != nil {
        panic(err)
    }
    var finalEntities []Event
    for _, item := range entities {
        event := Event{
            Step:item.Step,
            Substep:item.Substep,
            Ip:item.Ip,
        }
        switch item.Step {
        case NeEvent:
            var payload NeEventsPayload
            err := ffjson.Unmarshal(*item.Payload, &payload)
            if err != nil {
                panic(err)
            }
            event.Payload = payload
        case MkEvent:
            var payload MkEventsPayload
            err := ffjson.Unmarshal(*item.Payload, &payload)
            if err != nil {
                panic(err)
            }
            event.Payload = payload
        }
        finalEntities = append(finalEntities, event)
    }
}

func UnmarshalInterfaceWithRawMessage() {
    var entities []EventInterfaceRawMessage
    if err := ffjson.Unmarshal(codeJSON, &entities); err != nil {
        panic(err)
    }
    for i, item := range entities {
        switch item.Step {
        case NeEvent:
            var payload NeEventsPayload
            err := ffjson.Unmarshal(*item.Payload, &payload)
            if err != nil {
                panic(err)
            }
            item.PayloadObj = payload
            entities[i] = item
        case MkEvent:
            var payload MkEventsPayload
            err := ffjson.Unmarshal(*item.Payload, &payload)
            if err != nil {
                panic(err)
            }
            item.PayloadObj = payload
            entities[i] = item
        }
    }
}

func BenchmarkUnmarshalDynamicJson100(b *testing.B) {
    if codeJSON == nil {
        b.StopTimer()
        codeInit(100)
        b.StartTimer()
    }
    for i:=0;i<b.N; i++ {
        TestFunc()
    }
    codeJSON = nil
}
func BenchmarkUnmarshalDynamicJson1000(b *testing.B) {
    if codeJSON == nil {
        b.StopTimer()
        codeInit(1000)
        b.StartTimer()
    }

    for i:=0;i<b.N; i++ {
        TestFunc()
    }
    codeJSON = nil
}
func BenchmarkUnmarshalDynamicJson10000(b *testing.B) {
    if codeJSON == nil {
        b.StopTimer()
        codeInit(10000)
        b.StartTimer()
    }

    for i:=0;i<b.N; i++ {
        TestFunc()
    }
    codeJSON = nil
}
func BenchmarkUnmarshalDynamicJson100000(b *testing.B) {
    if codeJSON == nil {
        b.StopTimer()
        codeInit(100000)
        b.StartTimer()
    }

    for i:=0;i<b.N; i++ {
        TestFunc()
    }
    codeJSON = nil
}
func BenchmarkUnmarshalDynamicJson1000000(b *testing.B) {
    if codeJSON == nil {
        b.StopTimer()
        codeInit(1000000)
        b.StartTimer()
    }

    for i:=0;i<b.N; i++ {
        TestFunc()
    }
    codeJSON = nil
}
