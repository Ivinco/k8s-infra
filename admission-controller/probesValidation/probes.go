package probesValidation

import (
	"encoding/json"

	"admissioncontroller"

	v1 "k8s.io/api/apps/v1"
	core "k8s.io/api/core/v1"
)

// NewValidationHook creates a new instance of deployment validation hook
func NewValidationHook() admissioncontroller.Hook {
	return admissioncontroller.Hook{
		Create: validateCreate(),
		Update: validateUpdate(),
	}
}

func parseObject(object []byte) (*v1.ReplicaSet, error) {
	var objectToParse v1.ReplicaSet

	if err := json.Unmarshal(object, &objectToParse); err != nil {
		return nil, err
	}
	return &objectToParse, nil
}

func hasProbes(spec core.PodSpec) bool {
	for _, container := range spec.Containers {
		if container.ReadinessProbe == nil || container.LivenessProbe == nil || container.StartupProbe == nil {
			return false
		}
	}
	return true
}
