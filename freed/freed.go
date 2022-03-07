package main

import (
	"fmt"
	"net"
	"strconv"
)

func (client *FreeDClient) read_freed_data_15(data []byte) float64 {
	var r int = -1
	var p int = int(data[0])
	var t int = int(data[1])
	var s int = int(data[2])

	r <<= 8
	r |= p
	r <<= 8
	r |= t
	r <<= 8
	r |= s
	if !(r&0x00800000 == 1) {
		r &= 0x00ffffff
	}

	return float64(r) / float64(32768.0)
}

func (client *FreeDClient) read_freed_data_6(data []byte) float64 {
	var r int = -1
	var p int = int(data[0])
	var t int = int(data[1])
	var s int = int(data[2])

	r <<= 8
	r |= p
	r <<= 8
	r |= t
	r <<= 8
	r |= s
	if !(r&0x00800000 == 1) {
		r &= 0x00ffffff
	}

	return float64(r) / float64(64.0)
}

func (client *FreeDClient) read_freed_data_2(data []byte) int32 {
	var r int32 = 0

	r <<= 8
	r |= int32(data[0])
	r <<= 8
	r |= int32(data[1])
	r <<= 8
	r |= int32(data[2])

	return r
}


func (client *FreeDClient) Start() {
	fmt.Printf("Server Listening")
	defer client.uDPServer.Close()
	for {
		buffer := make([]byte, 29)
		_, _, err := client.uDPServer.ReadFromUDP(buffer)
		if err != nil {
			fmt.Println(err)
			return
		}
		if buffer[0] != 0xD1 {
			fmt.Println("Invalid Packet Received.")
			continue
		}

		data := freed_data{
			pan:   client.read_freed_data_15(buffer[2:5]),
			tilt:  client.read_freed_data_15(buffer[5:8]),
			roll:  client.read_freed_data_15(buffer[8:11]),
			x:     client.read_freed_data_6(buffer[11:14]),
			y:     client.read_freed_data_6(buffer[14:17]),
			z:     client.read_freed_data_6(buffer[17:20]),
			zoom:  client.read_freed_data_2(buffer[20:23]), //These are unpacked correctly now
			focus: client.read_freed_data_2(buffer[23:26]), //These are unpacked correctly now
		}
		client.onPacket(data)
	}
}


func NewFreeDClient (port int, listen_address string) (client FreeDClient, err error) {
	s, err := net.ResolveUDPAddr("udp4", listen_address + ":" + strconv.Itoa(port))
	if err != nil {
		return FreeDClient{}, err
	}
	c, err := net.ListenUDP("udp4", s)
	if err != nil {
		return FreeDClient{}, err
	}
	return FreeDClient{
		Port: port,
		ListenAddr: listen_address,
		uDPServer: c,
	}, nil
}

