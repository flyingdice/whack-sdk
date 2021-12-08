package exports

// Export represents the func definition of TODO
type Export func(...int32) (interface{}, error)

// Exports is a map of function name to host function export
// for the host SDK.
var Exports = map[string]Export{
	"success": success,
	"err":     err,
}
