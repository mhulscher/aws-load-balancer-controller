package deploy

import (
	"encoding/json"
	"sigs.k8s.io/aws-alb-ingress-controller/pkg/model/core"
)

// StackMarshaller will marshall a resource stack into JSON.
type StackMarshaller interface {
	Marshal(stack core.Stack) (string, error)
}

func NewDefaultStackMarshaller() *defaultStackMarshaller {
	return &defaultStackMarshaller{}
}

var _ StackMarshaller = &defaultStackMarshaller{}

type defaultStackMarshaller struct{}

func (d *defaultStackMarshaller) Marshal(stack core.Stack) (string, error) {
	builder := NewStackSchemaBuilder()
	if err := stack.TopologicalTraversal(builder); err != nil {
		return "", err
	}
	stackSchema := builder.Build()
	payload, err := json.Marshal(stackSchema)
	if err != nil {
		return "", err
	}
	return string(payload), nil
}