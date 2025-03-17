package main

const (
	// Network Status code number
	Ok                  int = iota // 200
	Created             int = iota // 201
	NotContent          int = iota // 204
	MovedPermanently    int = iota // 300
	Found               int = iota // 302
	NotModified         int = iota // 304
	BadRequest          int = iota // 400
	Unauthorized        int = iota // 401
	Forbidden           int = iota // 403
	NotFound            int = iota // 404
	Conflict            int = iota // 409
	InternalServerError int = iota // 500
	ServiceUnavailable  int = iota // 503
)

func main() {

}
