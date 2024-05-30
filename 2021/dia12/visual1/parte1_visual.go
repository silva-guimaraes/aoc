package main

import (
	"bufio"
	"errors"
	"fmt"
	"math"
	"os"
	"strings"
	"unicode"

	"github.com/veandco/go-sdl2/sdl"
)

var win_height int32 = 500
var win_width int32 = 500
var win_title = "foobar"
var PI = 3.14159265359

var cont context

type graph map[string][]string

type rgb struct {
    r, g, b uint8
}

type node struct {
    x, y int32
    selected bool
    label string
}


type context struct {
    nodes []node
    renderer *sdl.Renderer
    window *sdl.Window
    node_radius int
    paths []string
}

func (g graph) add_edge(from, to string) {
    g[from] = append(g[from], to)
    g[to] = append(g[to], from)
}

func duplicates(visited []string) bool {
    count := make(map[string]int)
    for i := range visited {
        count[visited[i]]++
        if count[visited[i]] > 1 {
            return true
        }
    }
    return false
}

func did_visit(visited []string, current string) bool {
    for i := range visited {
        if visited[i] == current {
            return true
        }
    }
    return false
}

func find_node(nodes []node, label string) *node {
    for i := range nodes {
        if nodes[i].label == label {
            return &nodes[i]
        }
    }
    panic("find_node")
}

func (c context) select_node(current string) {
    find_node(c.nodes, current).selected = true
}

func (c context) clear_selection() {
    for i := range c.nodes {
        c.nodes[i].selected = false
    }
}

func (c context) draw() {

    c.renderer.SetDrawColor(0, 0, 0, 255)
    c.renderer.Clear()

    for i := range c.nodes {
        node := c.nodes[i]

        if node.selected {
            c.renderer.SetDrawColor(0, 255, 0, 255)
        } else {
            c.renderer.SetDrawColor(255, 255, 255, 255)
        }
        draw_circle(c.renderer, c.nodes[i].x, c.nodes[i].y, float64(c.node_radius))
    }

    for i := 0; i < len(c.paths); i += 2 {
        to := find_node(c.nodes, c.paths[i])
        from := find_node(c.nodes, c.paths[i+1])

        if to.selected && from.selected {
            c.renderer.SetDrawColor(0, 0, 255, 255)
        } else {
            c.renderer.SetDrawColor(60, 60, 60, 255)
        }
        c.renderer.DrawLine(from.x, from.y, to.x, to.y)
    }
    c.renderer.Present()
    sdl.Delay(50)
}
    
func (g graph) traverse(current string, visited []string) []string {

    cont.clear_selection()
    find_node(cont.nodes, current).selected = true
    cont.draw()

    
    if current == "end" {
        return []string{current} 
    }
    if unicode.IsLower(rune(current[0])) {
        if current == "start" && did_visit(visited, "start") {
            return []string{}
        }
        visited = append(visited, current)
    }

    var paths []string

    for _, node := range g[current] {

        if !duplicates(visited) || !did_visit(visited, node) {

            find_node(cont.nodes, node).selected = true
            cont.draw()
            find_node(cont.nodes, node).selected = false

            new_paths := g.traverse(node, visited)
            paths = append(paths, new_paths...)
        }
    }
    return paths
}

func  draw_circle(r *sdl.Renderer, x, y int32, radius float64) {
    segments := 50

    var points []sdl.Point

    for i := 1; i <= segments; i++ {
        angle := 2 * PI * float64(i) / float64(segments)
        cos := math.Cos(angle)
        sin := math.Sin(angle)
        points = append(points, sdl.Point{
            x + int32(radius * cos), 
            y + int32(radius * sin),
        })
    }
    points = append(points, points[0])
    r.DrawLines(points)
}

func initialize(graph graph, paths []string) (context, error) {

    var window *sdl.Window
    var renderer *sdl.Renderer

    window, err := sdl.CreateWindow(win_title, sdl.WINDOWPOS_UNDEFINED,
    sdl.WINDOWPOS_UNDEFINED, win_width, win_height, sdl.WINDOW_SHOWN)
    if err != nil {
        // fmt.Fprintf(os.Stderr, "Failed to create window: %s\n", err)
        return context{}, errors.New("Failed to create window: %s\n")
    }
    renderer, err = sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
    if err != nil {
        // fmt.Fprintf(os.Stderr, "Failed to create renderer\n", err)
        return context{}, errors.New("Failed to create renderer\n")
    }

    var labels []string
    for k := range graph {
        labels = append(labels, k)
    }

    // running := true
    // node_radius := 2
    radius := 200.0
    points := len(graph)


    var nodes []node
    for i := 1; i <= points; i++ {
        angle := 2.0 * PI * float64(i) / float64(points)
        x := int32(float64(win_width) / 2 + radius * math.Cos(angle))
        y := int32(float64(win_height) / 2 + radius * math.Sin(angle))

        nodes = append(nodes, node{
            x: x, 
            y: y, 
            label: labels[i-1],
            selected: false,
        })
    }

    return context{ 
        nodes: nodes, 
        renderer: renderer, 
        node_radius: 48,
        window: window,
        paths: paths,
    }, nil

}

func main() {

    file, err := os.Open("../input.txt")
    if err != nil { panic(err) }
    defer file.Close()

    scanner := bufio.NewScanner(file)

    var input []string
    for scanner.Scan() {
        input = append(input, strings.Split(scanner.Text(), "-")...)
    }
    
    paths := make(graph)

    for i := 0; i < len(input); i += 2 {
        paths.add_edge(input[i], input[i+1])
    }

    cont, err = initialize(paths, input)
    if err != nil { panic(err) }
    defer cont.window.Destroy()
    defer cont.renderer.Destroy()

    // fmt.Println(cont.window.GetFlags())
    ret := paths.traverse("start", []string{})
    fmt.Println(len(ret))
}
