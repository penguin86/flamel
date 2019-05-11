package mage

import (
	"encoding/json"
	"errors"
	"golang.org/x/net/context"
)

var ErrMageNoJSON = errors.New("inputs do not contain json data")

func InputsFromContext(ctx context.Context) RequestInputs {
	inputs := ctx.Value(KeyRequestInputs).(RequestInputs)
	return inputs
}

type Controller interface {
	//the page logic is executed here
	//Process method consumes the context -> context variations, i.e. appengine.Namespace
	//can be used INSIDE the Process function
	Process(ctx context.Context, out *ResponseOutput) Redirect
	//called to release resources
	OnDestroy(ctx context.Context)
}

// Convenience method to recover all json inputs
// Returns user json inputs as a map string -> interface{}
func ParseJSONInputs(ctx context.Context) (map[string]interface{}, error) {
	inputs := InputsFromContext(ctx)
	if inputs == nil {
		return nil, ErrMageNoJSON
	}

	jin := []byte(inputs[KeyRequestJSON].Value())

	var data interface{}
	err := json.Unmarshal(jin, &data)

	if err != nil {
		return nil, err
	}

	d := data.(map[string]interface{})

	return d, nil
}
