package socket_route

const (
	JoinWorldReq = 10001

	JoinWorldResp = 20001
)

// RouteMap string为Method名
var RouteMap = map[int]string{
	JoinWorldReq: "JoinWorldReq",

	JoinWorldResp: "JoinWorldResp",
}
