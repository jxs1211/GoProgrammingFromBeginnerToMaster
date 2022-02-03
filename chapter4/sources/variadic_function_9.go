package main

import (
	"fmt"
)

type Room struct {
	style         int
	airCondition  bool
	floorMaterial string
}

type Option2 func(*Room)

func SetStyle(style int) Option2 {
	return func(r *Room) {
		r.style = style
	}
}

func SetAirCondition(isNeed bool) Option2 {
	return func(r *Room) {
		r.airCondition = isNeed
	}
}

func SetFloorMaterial(material string) Option2 {
	return func(r *Room) {
		r.floorMaterial = material
	}
}

func NewRoom(options ...Option2) *Room {
	r := &Room{
		style:         0,
		airCondition:  true,
		floorMaterial: "wall paper",
	}

	for _, option := range options {
		option(r)
	}

	return r
}

func Show2() {
	fmt.Printf("%#v\n", NewRoom()) // default room
	fmt.Printf("%#v\n", NewRoom(
		SetStyle(1),
		SetAirCondition(true),
		SetFloorMaterial("ground-tile"),
	)) // custom room
}

func main() {
	// Show()
	Show2()
}

func Show() {
	fmt.Printf("%+v\n", NewFinishedHouse()) // use default options
	fmt.Printf("%+v\n", NewFinishedHouse(
		WithStyle(1),
		WithFloorMaterial("ground-tile"),
		WithCentralAirConditioning(false)),
	)
}

type FinishedHouse struct {
	style                  int    // 0: Chinese, 1: American, 2: European
	centralAirConditioning bool   // true or false
	floorMaterial          string // "ground-tile" or ”wood"
	wallMaterial           string // "latex" or "paper" or "diatom-mud"
}

type Option func(*FinishedHouse)

func NewFinishedHouse(options ...Option) *FinishedHouse {
	h := &FinishedHouse{
		// default options
		style:                  0,
		centralAirConditioning: true,
		floorMaterial:          "wood",
		wallMaterial:           "paper",
	}

	for _, option := range options {
		option(h)
	}

	return h
}

func WithStyle(style int) Option {
	return func(h *FinishedHouse) {
		h.style = style
	}
}

func WithFloorMaterial(material string) Option {
	return func(h *FinishedHouse) {
		h.floorMaterial = material
	}
}

func WithWallMaterial(material string) Option {
	return func(h *FinishedHouse) {
		h.wallMaterial = material
	}
}

func WithCentralAirConditioning(centralAirConditioning bool) Option {
	return func(h *FinishedHouse) {
		h.centralAirConditioning = centralAirConditioning
	}
}
