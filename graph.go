package main

import (
	"fmt"
	"image/color"

	"git.sr.ht/~sbinet/gg"
)

func createGraphVisualization(pages map[string]int, baseURL string) error {
    dc := gg.NewContext(800, 600)
    dc.SetColor(color.RGBA{255, 255, 255, 255})
    dc.Clear()

    dc.SetColor(color.RGBA{0, 0, 0, 255})
    dc.SetLineWidth(2)

    // Draw nodes
    for url, _ := range pages {
        dc.DrawString(url, 10, float64(10+20*len(pages)))
    }

    // Draw edges
    for _, _ = range pages {
        dc.DrawLine(10, float64(10+20*len(pages)), float64(10+100), float64(10+20*len(pages)))
    }

    dc.Stroke()

    err := dc.SavePNG("graph.png")
    if err!= nil {
        return fmt.Errorf("could not save graph as PNG: %w", err)
    }

    return nil
}