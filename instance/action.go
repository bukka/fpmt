package instance

import "fmt"

// Action is a generic interface for all actions
type Action interface {
	Type() string
	Run() error
}

// BaseAction is a base for all actions.
type BaseAction struct {
	typeName string
	expect   *Expectation
}

// Expectation contains expectation for the action
type Expectation struct {
	response *ResponseExpectation
	output   *StringExpectation
}

// ResponseExpectation is an expectation for the response
type ResponseExpectation struct {
	body string
}

// StringExpectation is an expaction for a string.
type StringExpectation struct {
	regexp string
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
			return nil, fmt.Errorf("Action Type has to be a string")
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
		return nil, fmt.Errorf("Invalid Action type")
	}

	return action, nil
}

// Create an expectation
func createExpectation(record interface{}) (*Expectation, error) {
	item, ok := record.(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("Invalid Expect type")
	}
	var err error
	var response *ResponseExpectation
	if responseVal, ok := item["Response"]; ok {
		response, err = createResponseExpectation(responseVal)
		if err != nil {
			return nil, err
		}
	}

	var output *StringExpectation
	if outputVal, ok := item["Output"]; ok {
		output, err = createStringExpectation(outputVal)
		if err != nil {
			return nil, err
		}
	}

	return &Expectation{response: response, output: output}, nil
}

func createResponseExpectation(record interface{}) (*ResponseExpectation, error) {
	item, ok := record.(map[string]string)
	if !ok {
		return nil, fmt.Errorf("Invalid ResponseExpectation type")
	}
	expectation := &ResponseExpectation{}
	if bodyVal, ok := item["Body"]; ok {
		expectation.body = bodyVal
	}

	return expectation, nil
}

func createStringExpectation(record interface{}) (*StringExpectation, error) {
	item, ok := record.(map[string]string)
	if !ok {
		return nil, fmt.Errorf("Invalid StringExpectation type")
	}
	expectation := &StringExpectation{}
	if regexpVal, ok := item["Regexp"]; ok {
		expectation.regexp = regexpVal
	}

	return expectation, nil
}

// Type retrieves type string of the action
func (a BaseAction) Type() string {
	return a.typeName
}

// Run base action
func (a BaseAction) Run() error {
	fmt.Print("BaseAction run")
	return nil
}
