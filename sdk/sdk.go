package sdk

import "github.com/flyingdice/whack-sdk/sdk/exports"

const namespace = "whack"

// Exports returns the namespace and functions that should be exported
// as part of the SDK.
func Exports() (string, []Function){
    return namespace, exportsToFunctions(exports.Exports)
}

// exportsToFunctions creates a slice of sdk.Function instances for the
// given map of exports.
func exportsToFunctions(exported map[string]exports.Export) []Function {
    functions := make([]Function, len(exported))
    for name, export := range exported {
        functions = append(functions, NewFunction(name, export))
    }
    return functions
}
