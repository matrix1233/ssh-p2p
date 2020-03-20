package signaling

// URI default signaling server
const URI = "https://74.125.140.127"

// ConnectInfo SDP by offer or answer
type ConnectInfo struct {
	Source string `json:"source"`
	SDP    string `json:"sdp"`
}
