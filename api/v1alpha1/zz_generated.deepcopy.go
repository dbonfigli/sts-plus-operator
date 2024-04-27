//go:build !ignore_autogenerated

/*
Copyright 2024.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	"k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Check) DeepCopyInto(out *Check) {
	*out = *in
	out.Query = in.Query
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Check.
func (in *Check) DeepCopy() *Check {
	if in == nil {
		return nil
	}
	out := new(Check)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PhasedRollout) DeepCopyInto(out *PhasedRollout) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PhasedRollout.
func (in *PhasedRollout) DeepCopy() *PhasedRollout {
	if in == nil {
		return nil
	}
	out := new(PhasedRollout)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PhasedRollout) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PhasedRolloutList) DeepCopyInto(out *PhasedRolloutList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]PhasedRollout, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PhasedRolloutList.
func (in *PhasedRolloutList) DeepCopy() *PhasedRolloutList {
	if in == nil {
		return nil
	}
	out := new(PhasedRolloutList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PhasedRolloutList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PhasedRolloutSpec) DeepCopyInto(out *PhasedRolloutSpec) {
	*out = *in
	out.Check = in.Check
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PhasedRolloutSpec.
func (in *PhasedRolloutSpec) DeepCopy() *PhasedRolloutSpec {
	if in == nil {
		return nil
	}
	out := new(PhasedRolloutSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PhasedRolloutStatus) DeepCopyInto(out *PhasedRolloutStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	in.RolloutStartTime.DeepCopyInto(&out.RolloutStartTime)
	in.RolloutEndTime.DeepCopyInto(&out.RolloutEndTime)
	if in.RollingPodStatus != nil {
		in, out := &in.RollingPodStatus, &out.RollingPodStatus
		*out = new(RollingPodStatus)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PhasedRolloutStatus.
func (in *PhasedRolloutStatus) DeepCopy() *PhasedRolloutStatus {
	if in == nil {
		return nil
	}
	out := new(PhasedRolloutStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PrometheusQuery) DeepCopyInto(out *PrometheusQuery) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PrometheusQuery.
func (in *PrometheusQuery) DeepCopy() *PrometheusQuery {
	if in == nil {
		return nil
	}
	out := new(PrometheusQuery)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RollingPodStatus) DeepCopyInto(out *RollingPodStatus) {
	*out = *in
	in.AnalisysStartTime.DeepCopyInto(&out.AnalisysStartTime)
	in.LastCheckTime.DeepCopyInto(&out.LastCheckTime)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RollingPodStatus.
func (in *RollingPodStatus) DeepCopy() *RollingPodStatus {
	if in == nil {
		return nil
	}
	out := new(RollingPodStatus)
	in.DeepCopyInto(out)
	return out
}
