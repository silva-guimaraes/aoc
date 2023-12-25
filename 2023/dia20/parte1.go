package main

import (
    "fmt"
    "os"
    "strings"
)


type module interface {
    Receive(signal, map[string]module) []signal
    Name() string
    Children() []string
}

type flipflop struct {
    name string
    onOff bool
    children []string
}

type conjunction struct {
    name string
    state map[string]bool
    children []string
}

type signal struct {
    from, to string
    onoff bool
}


func (f *flipflop) Receive(s signal, lookup map[string]module) []signal {
    if s.onoff {
        return nil
    }
    pulse := false
    if !f.onOff { pulse = true }

    signals := []signal{}
    for _, c := range f.children {
        signals = append(signals, signal{f.name, c, pulse})
    }
    f.onOff = !f.onOff
    return signals
}

func (f flipflop) Name() string {
    return f.name
}

func (f flipflop) Children() []string {
    return f.children
}


func (f *conjunction) Receive(s signal, lookup map[string]module) []signal {

    f.state[s.from] = s.onoff

    allHigh := true
    for _, v := range f.state {
        if !v { allHigh = false; break }
    }


    signals := []signal{}
    for _, c := range f.children {
        signals = append(signals, signal{f.name, c, !allHigh})
    }

    return signals

}

func (f conjunction) Name() string {
    return f.name
}

func (f conjunction) Children() []string {
    return f.children
}



func main() {
    file, err := os.ReadFile("input.txt")
    if err != nil {
        panic(err)
    }

    button := []signal{}
    lookup := map[string]module{}

    for _, l :=  range strings.Split(strings.TrimSpace(string(file)), "\n") {
        s := strings.Split(l, " -> ")
        children := strings.Split(s[1], ", ")
        name := s[0][1:]

        if s[0] == "broadcaster" {
            for _, c := range children {
                button = append(button, signal{
                    from: "broadcaster", to: c, onoff: false,
                })
            }
            continue
        }

        var m module
        if s[0][0] == '%' {
            m = &flipflop{name, false, children}
        } else if s[0][0] == '&' {
            m = &conjunction{name, make(map[string]bool), children}
        }
        lookup[name] = m
    }

    for k, v := range lookup {
        for _, label := range v.Children() {
            switch a := lookup[label].(type) {
            case *conjunction:
                a.state[k] = false
            }
        }
    }



    highs, lows := 0, 0
    for times := 0; times < 1000; times++ {
        broadcast := []signal{}
        broadcast = append(broadcast, button...)
        lows++

        for len(broadcast) > 0 {
            b := broadcast[0]
            broadcast = broadcast[1:]

            if b.onoff { highs++ } else { lows++ }

            to, ok := lookup[b.to] 

            if !ok  { continue }

            if b.to == "output" {
            }


            broadcast = append(broadcast, to.Receive(b, lookup)...)
        }
    }

    fmt.Println(highs * lows)
}
