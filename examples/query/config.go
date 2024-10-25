package query

// Type is an enum that represents the type of query to be made to the Pythnet Hermes API.
type Type string

const (
	// LatestSync is a query type that fetches all latest prices together & synchronously.
	LatestSync Type = "latest-sync"
	// LatestAsync is a query type that fetches all latest prices individually, in parallel.
	LatestAsync Type = "latest-async"
	// Stream is a query type that subscribes to price updates and returns the latest cached price.
	StreamCached Type = "stream"
)

// Settings is a struct that holds the settings for querying Hermes using the Pythnet client.
type Settings struct {
	UseEma           bool
	DesiredPrecision int
	RequestType      Type
	SingleUpdateFee  uint
}
