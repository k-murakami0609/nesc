package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"strconv"

	"github.com/gorilla/websocket"
	"github.com/pion/webrtc/v3"
)

func decode(message []byte) webrtc.SessionDescription {
	recvOnlyOffer := webrtc.SessionDescription{}
	b, err := base64.StdEncoding.DecodeString(string(message))
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(b, &recvOnlyOffer)
	if err != nil {
		panic(err)
	}

	return recvOnlyOffer
}

func encode(obj interface{}) string {
	b, err := json.Marshal(obj)
	if err != nil {
		panic(err)
	}

	return base64.StdEncoding.EncodeToString(b)
}

func createConnection(c *websocket.Conn, videoTrack *webrtc.TrackLocalStaticSample, ki *keyInput) {
	mt, message, err := c.ReadMessage()
	if err != nil {
		log.Println("read:", err)
		return
	}

	recvOnlyOffer := decode(message)
	if err != nil {
		panic(err)
	}

	peerConnection, err := webrtc.NewPeerConnection(config)
	if err != nil {
		panic(err)
	}

	peerConnection.OnICEConnectionStateChange(func(connectionState webrtc.ICEConnectionState) {
		fmt.Printf("Connection State has changed %s \n", connectionState.String())
	})

	peerConnection.OnDataChannel(func(d *webrtc.DataChannel) {
		fmt.Printf("New DataChannel %s %d\n", d.Label(), d.ID())

		d.OnMessage(func(msg webrtc.DataChannelMessage) {
			fmt.Printf("Message from DataChannel '%s': '%s'\n", d.Label(), string(msg.Data))
			tmp, err := strconv.Atoi(string(msg.Data))
			if err != nil {
				log.Printf("failed change")
				ki.selector = 0
			} else {
				log.Printf("%d", tmp)
				ki.selector = tmp
			}
		})
	})

	_, err = peerConnection.AddTrack(videoTrack)
	if err != nil {
		panic(err)
	}

	err = peerConnection.SetRemoteDescription(recvOnlyOffer)
	if err != nil {
		panic(err)
	}

	answer, err := peerConnection.CreateAnswer(nil)
	if err != nil {
		panic(err)
	}

	gatherComplete := webrtc.GatheringCompletePromise(peerConnection)

	err = peerConnection.SetLocalDescription(answer)
	if err != nil {
		panic(err)
	}

	<-gatherComplete

	ds := encode(*peerConnection.LocalDescription())

	c.WriteMessage(mt, []byte(ds))
}
