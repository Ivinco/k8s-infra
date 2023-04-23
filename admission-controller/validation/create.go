package validation

import (
	"admissioncontroller"

	v1 "k8s.io/api/admission/v1"
)

func validateCreate() admissioncontroller.AdmitFunc {
	return func(r *v1.AdmissionRequest) (*admissioncontroller.Result, error) {
		receivedObject, err := parseObject(r.Object.Raw)
		if err != nil {
			return &admissioncontroller.Result{Msg: err.Error()}, err
		}

		if !hasProbes(receivedObject.Spec.Template.Spec) {
			return &admissioncontroller.Result{Msg: "Created resource doesn't have probes set."}, nil
		}

		if !checkImageLatest(receivedObject.Spec.Template.Spec) {
			return &admissioncontroller.Result{Msg: "Created resource uses \"latest\" tag. That is restricted."}, nil
		}

		return &admissioncontroller.Result{Allowed: true}, nil
	}

}

func validateUpdate() admissioncontroller.AdmitFunc {
	return func(r *v1.AdmissionRequest) (*admissioncontroller.Result, error) {
		receivedObject, err := parseObject(r.Object.Raw)
		if err != nil {
			return &admissioncontroller.Result{Msg: err.Error()}, err
		}

		if !hasProbes(receivedObject.Spec.Template.Spec) {
			return &admissioncontroller.Result{Msg: "Updated resource doesn't have probes set."}, nil
		}

		if !checkImageLatest(receivedObject.Spec.Template.Spec) {
			return &admissioncontroller.Result{Msg: "Updated resource uses \"latest\" tag. That is restricted."}, nil
		}

		return &admissioncontroller.Result{Allowed: true}, nil
	}

}
