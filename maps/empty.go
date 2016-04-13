package maps

var (
	emptyMap = make(map[interface{}]interface{})
)

// Empty returns a map with no elements in it
func Empty() map[interface{}]interface{} {
	return emptyMap
}
