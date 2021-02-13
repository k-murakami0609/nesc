const createPcAndChannel = async () => {
    const pc = new RTCPeerConnection({
        iceServers: [
            {
                urls: 'stun:stun.l.google.com:19302'
            }
        ]
    })

    sendChannel = pc.createDataChannel('foo')
    sendChannel.onclose = () => console.log('sendChannel has closed')
    sendChannel.onopen = () => console.log('sendChannel has opened')
    sendChannel.onmessage = e => log(`Message from DataChannel '${sendChannel.label}' payload '${e.data}'`)

    pc.ontrack = async (event) => {
        var el = document.createElement(event.track.kind)
        el.srcObject = event.streams[0]
        el.autoplay = true
        el.controls = true

        await document.getElementById('remoteVideos').appendChild(el)
    }
    pc.addTransceiver('audio', { 'direction': 'sendrecv' })
    pc.addTransceiver('video', { 'direction': 'sendrecv' })
    pc.addTransceiver('video', { 'direction': 'sendrecv' })

    const d = await pc.createOffer()
    await pc.setLocalDescription(d)

    return [pc, sendChannel]
}

const connect = (pc) => {
    const ws = new WebSocket("ws://localhost:8000/connect");
    ws.onopen = function (evt) {
        ws.send(btoa(JSON.stringify(pc.localDescription)));
    }
    ws.onmessage = function (evt) {
        try {
            pc.setRemoteDescription(new RTCSessionDescription(JSON.parse(atob(evt.data))))
        } catch (e) {
            alert(e)
        }
    }
    ws.onerror = function (evt) {
        console.log(evt.data)
    }
}

window.start = async () => {
    document.getElementById("open").remove()
    const [pc, sendChannel] = await createPcAndChannel()
    connect(pc)
}

window.sendMessage = (color) => {
    sendChannel.send(color)
}
