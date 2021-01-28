package objects

type GroupRoute struct {
	Prefix         string
	middlewareName string
	Route          []Route
}
