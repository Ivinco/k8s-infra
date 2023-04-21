package validation

import (
	"encoding/json"
	"regexp"

	"admissioncontroller"

	v1 "k8s.io/api/apps/v1"
	core "k8s.io/api/core/v1"
	log "k8s.io/klog/v2"
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
		if container.LivenessProbe == nil {
			log.Errorf("Container %s doesn't have Liveness Probe set", container.Name)
			return false
		}
	}
	return true
}

func checkImageLatest(spec core.PodSpec) bool {
	pattern := regexp.MustCompile(`:latest$`)
	// log.Infof("Checking that image does not contain `latest` tag")
	for _, container := range spec.Containers {
		// log.Infof("Container %s is created with the following image: %s", container.Name, container.Image)
		// result := pattern.MatchString(container.Image)
		// log.Infof("Image uses `latest` tag:", result)
		if pattern.MatchString(container.Image) {
			log.Errorf("Container %s is created with the following image: %s. `latest` tags are restricted. Please use specific image versions", container.Name, container.Image)
			return false
		}
	}
	return true
}
