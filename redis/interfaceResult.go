package redis

type InterfaceResult struct {
	Result interface{}
	Error error
}

func NewInterfaceResult(result interface{}, error error) *InterfaceResult {
	return &InterfaceResult{Result: result, Error: error}
}

func(this *InterfaceResult) Unwrap() interface{} {
	if this.Error != nil {
		panic(this.Error)
	}
	return this.Result
}

func(this *InterfaceResult) UnwrapOr(val interface{}) interface{} {
	if this.Error != nil {
		return val
	}
	return this.Result
}

