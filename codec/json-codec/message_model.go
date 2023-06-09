package json_codec

type JoinWorldRequest struct {
	PlayerId string
}

type JoinWorldResponse struct {
	Code int
	X    int
	Y    int
}
