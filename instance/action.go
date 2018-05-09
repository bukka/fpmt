package instance

import "fmt"

// Action is a generic interface for all actions
type Action interface {
	Type() string
	Run(s *Settings) error
}

// BaseAction is a base for all actions.
type BaseAction struct {
	typeName string
	expect   *Expectation
}

// ServerAction is a base for all server action
type ServerAction struct {
	config ServerConfig
	BaseAction
}

// ServerStartAction starts the server
type ServerStartAction struct {
	ServerAction
}

// Expectation contains expectation for the action
type Expectation struct {
	response *ResponseExpectation
	output   []string
}

// ResponseExpectation is an expectation for the response
type ResponseExpectation struct {
	body string
}

type actionMappingItem struct {
	instance           *BaseAction
	defaultExpectation *Expectation
}

var actionMapping = map[string]actionMappingItem{
	"server-start": actionMappingItem{
		defaultExpectation: &Expectation{
			output: []string{
				"{{date}} NOTICE: {{app}} is running, pid {{pid}}",
				"{{date}} NOTICE: ready to handle connections",
			},
		},
	},
}

// CreateAction creates a new action from generic record.
func CreateAction(record interface{}) (Action, error) {
	var action BaseAction
	switch item := record.(type) {
	case string:
		action = BaseAction{typeName: item}
	case map[string]interface{}:
		typeVal, ok := item["Type"]
		if !ok {
			return nil, fmt.Errorf("Action does not have Type field")
		}
		actionType, ok := typeVal.(string)
		if !ok {
			return nil, fmt.Errorf(
				"Action Type has to be a string and not %T", typeVal)
		}
		var expect *Expectation
		if expectVal, ok := item["Expect"]; ok {
			var err error
			expect, err = createExpectation(expectVal)
			if err != nil {
				return nil, err
			}
		}
		action = BaseAction{typeName: actionType, expect: expect}
	default:
		return nil, fmt.Errorf("Invalid Action type %T", record)
	}

	return &action, nil
}

// Create an expectation
func createExpectation(record interface{}) (*Expectation, error) {
	item, ok := record.(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("Invalid Expect type %T", item)
	}
	var err error
	var response *ResponseExpectation
	if responseVal, ok := item["Response"]; ok {
		response, err = createResponseExpectation(responseVal)
		if err != nil {
			return nil, err
		}
	}

	var output []string
	if outputVal, ok := item["Output"]; ok {
		outputArray, ok := outputVal.([]string)
		if !ok {
			return nil, fmt.Errorf("The output has to be an array of strings and not %T", outputVal)
		}
		output = outputArray
	}

	return &Expectation{response: response, output: output}, nil
}

func createResponseExpectation(record interface{}) (*ResponseExpectation, error) {
	expectation := &ResponseExpectation{}
	switch item := record.(type) {
	case string:
		expectation.body = item
	case map[string]interface{}:
		if bodyVal, ok := item["Body"]; ok {
			if bodyString, ok := bodyVal.(string); ok {
				expectation.body = bodyString
			} else {
				return nil, fmt.Errorf(
					"ResponseExpectation Body has to be a string and not %T", bodyVal)
			}
		}
	default:
		return nil, fmt.Errorf("Invalid ResponseExpectation type %T", record)
	}

	return expectation, nil
}

// Type retrieves type string of the action
func (a BaseAction) Type() string {
	return a.typeName
}

// Run base action
func (a BaseAction) Run(s *Settings) error {
	fmt.Print("BaseAction run")
	return nil
}
