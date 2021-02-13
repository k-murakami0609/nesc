package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	"github.com/notedit/gst"
	"github.com/pion/webrtc/v3"
	"github.com/pion/webrtc/v3/pkg/media"
)

var config = webrtc.Configuration{
	ICEServers: []webrtc.ICEServer{
		{
			URLs: []string{"stun:stun.l.google.com:19302"},
		},
	},
}

type keyInput struct {
	selector int
}

var upgrader = websocket.Upgrader{} // use default options

func generateImageAndPushToPipeLine(pipeline *gst.Pipeline, ki *keyInput) {
	element := pipeline.GetByName("mysource")

	t := time.NewTicker(20 * time.Millisecond)
	for {
		select {
		case <-t.C:
			b := generateImage(ki.selector)
			err := element.PushBuffer(b.Bytes())

			if err != nil {
				fmt.Println("push buffer error")
			}
		}
	}
}

func pullImageAndWriteToVideo(pipeline *gst.Pipeline, videoTrack *webrtc.TrackLocalStaticSample) {
	var duration uint64 = 0
	element := pipeline.GetByName("appsink")
	for {
		sample, err := element.PullSample()
		if err != nil {
			panic(err)
		}
		duration += sample.Duration
		if err := videoTrack.WriteSample(media.Sample{Data: sample.Data, Duration: time.Duration(duration)}); err != nil {
			fmt.Printf("failed to write video data: %+v", err)
		}
	}
}

func startRendering(videoTrack *webrtc.TrackLocalStaticSample, ki *keyInput) {
	pipeline, err := gst.ParseLaunch("appsrc name=mysource is-live=true ! image/jpeg,framerate=60/1 ! jpegparse ! jpegdec ! video/x-raw,format=I420 ! x264enc speed-preset=ultrafast tune=zerolatency key-int-max=20 ! video/x-h264,stream-format=byte-stream ! appsink name=appsink")
	if err != nil {
		panic(err)
	}

	pipeline.SetState(gst.StatePlaying)

	go generateImageAndPushToPipeLine(pipeline, ki)
	go pullImageAndWriteToVideo(pipeline, videoTrack)
}

func main() {
	videoTrack, err := webrtc.NewTrackLocalStaticSample(webrtc.RTPCodecCapability{MimeType: "video/h264"}, "video", "pion2")
	if err != nil {
		panic(err)
	}

	ki := keyInput{}

	startRendering(videoTrack, &ki)

	r := mux.NewRouter()
	r.HandleFunc("/connect", func(w http.ResponseWriter, r *http.Request) {
		c, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Print("upgrade:", err)
			return
		}
		defer c.Close()

		createConnection(c, videoTrack, &ki)
	})
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./public/")))

	srv := &http.Server{
		Handler:      r,
		Addr:         "127.0.0.1:8000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
