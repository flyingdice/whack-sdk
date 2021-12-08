package exports

import "log"

// success reads WASM memory and translates to golang success value.
func success(args ...int32) (interface{}, error) {
    log.Println("success called!")
    log.Println(args)
    return nil, nil
}

// error reads WASM memory and translate to golang error.
func err(args ...int32) (interface{}, error) {
    log.Println("err called!")
    log.Println(args)
    return nil, nil
}
