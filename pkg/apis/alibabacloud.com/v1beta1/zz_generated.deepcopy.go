//go:build !ignore_autogenerated

/*
Copyright 2021 Terway Authors.

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

package v1beta1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AllocationType) DeepCopyInto(out *AllocationType) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AllocationType.
func (in *AllocationType) DeepCopy() *AllocationType {
	if in == nil {
		return nil
	}
	out := new(AllocationType)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PodEIP) DeepCopyInto(out *PodEIP) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PodEIP.
func (in *PodEIP) DeepCopy() *PodEIP {
	if in == nil {
		return nil
	}
	out := new(PodEIP)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PodEIP) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PodEIPList) DeepCopyInto(out *PodEIPList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]PodEIP, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PodEIPList.
func (in *PodEIPList) DeepCopy() *PodEIPList {
	if in == nil {
		return nil
	}
	out := new(PodEIPList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PodEIPList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PodEIPSpec) DeepCopyInto(out *PodEIPSpec) {
	*out = *in
	out.AllocationType = in.AllocationType
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PodEIPSpec.
func (in *PodEIPSpec) DeepCopy() *PodEIPSpec {
	if in == nil {
		return nil
	}
	out := new(PodEIPSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PodEIPStatus) DeepCopyInto(out *PodEIPStatus) {
	*out = *in
	in.PodLastSeen.DeepCopyInto(&out.PodLastSeen)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PodEIPStatus.
func (in *PodEIPStatus) DeepCopy() *PodEIPStatus {
	if in == nil {
		return nil
	}
	out := new(PodEIPStatus)
	in.DeepCopyInto(out)
	return out
}
