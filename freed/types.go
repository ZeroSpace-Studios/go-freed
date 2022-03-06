package main

import "net"

type freed_data struct {
	pan   float64
	tilt  float64
	roll  float64
	x     float64
	y     float64
	z     float64
	zoom  int32
	focus int32
}

type FreeDClient struct {
	uDPServer *net.UDPConn
	Port int
	ListenAddr string
	onPacket packethandler
}

type packethandler func(freed_data)