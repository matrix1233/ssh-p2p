package signaling

// URI default signaling server
const URI = "https://stun.l.google.com"

// ConnectInfo SDP by offer or answer
type ConnectInfo struct {
	Source string `json:"source"`
	SDP    string `json:"sdp"`
}
