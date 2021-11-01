package eventloop

type Eventloop struct {
	channels []*Channel
	pos      int64
	size     int64
}
