// +build !ignore_autogenerated

// Copyright (c) 2019 SAP SE or an SAP affiliate company. All rights reserved. This file is licensed under the Apache Software License, v. 2 except as noted otherwise in the LICENSE file.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	"k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BackupSpec) DeepCopyInto(out *BackupSpec) {
	*out = *in
	if in.Port != nil {
		in, out := &in.Port, &out.Port
		*out = new(int32)
		**out = **in
	}
	if in.TLS != nil {
		in, out := &in.TLS, &out.TLS
		*out = new(TLSConfig)
		**out = **in
	}
	if in.Image != nil {
		in, out := &in.Image, &out.Image
		*out = new(string)
		**out = **in
	}
	if in.Store != nil {
		in, out := &in.Store, &out.Store
		*out = new(StoreSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.Resources != nil {
		in, out := &in.Resources, &out.Resources
		*out = new(v1.ResourceRequirements)
		(*in).DeepCopyInto(*out)
	}
	if in.FullSnapshotSchedule != nil {
		in, out := &in.FullSnapshotSchedule, &out.FullSnapshotSchedule
		*out = new(string)
		**out = **in
	}
	if in.GarbageCollectionPolicy != nil {
		in, out := &in.GarbageCollectionPolicy, &out.GarbageCollectionPolicy
		*out = new(GarbageCollectionPolicy)
		**out = **in
	}
	if in.GarbageCollectionPeriod != nil {
		in, out := &in.GarbageCollectionPeriod, &out.GarbageCollectionPeriod
		*out = new(metav1.Duration)
		**out = **in
	}
	if in.DeltaSnapshotPeriod != nil {
		in, out := &in.DeltaSnapshotPeriod, &out.DeltaSnapshotPeriod
		*out = new(metav1.Duration)
		**out = **in
	}
	if in.DeltaSnapshotMemoryLimit != nil {
		in, out := &in.DeltaSnapshotMemoryLimit, &out.DeltaSnapshotMemoryLimit
		x := (*in).DeepCopy()
		*out = &x
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BackupSpec.
func (in *BackupSpec) DeepCopy() *BackupSpec {
	if in == nil {
		return nil
	}
	out := new(BackupSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Condition) DeepCopyInto(out *Condition) {
	*out = *in
	in.LastTransitionTime.DeepCopyInto(&out.LastTransitionTime)
	in.LastUpdateTime.DeepCopyInto(&out.LastUpdateTime)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Condition.
func (in *Condition) DeepCopy() *Condition {
	if in == nil {
		return nil
	}
	out := new(Condition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CrossVersionObjectReference) DeepCopyInto(out *CrossVersionObjectReference) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CrossVersionObjectReference.
func (in *CrossVersionObjectReference) DeepCopy() *CrossVersionObjectReference {
	if in == nil {
		return nil
	}
	out := new(CrossVersionObjectReference)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Etcd) DeepCopyInto(out *Etcd) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Etcd.
func (in *Etcd) DeepCopy() *Etcd {
	if in == nil {
		return nil
	}
	out := new(Etcd)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Etcd) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EtcdConfig) DeepCopyInto(out *EtcdConfig) {
	*out = *in
	if in.Quota != nil {
		in, out := &in.Quota, &out.Quota
		x := (*in).DeepCopy()
		*out = &x
	}
	if in.DefragmentationSchedule != nil {
		in, out := &in.DefragmentationSchedule, &out.DefragmentationSchedule
		*out = new(string)
		**out = **in
	}
	if in.ServerPort != nil {
		in, out := &in.ServerPort, &out.ServerPort
		*out = new(int32)
		**out = **in
	}
	if in.ClientPort != nil {
		in, out := &in.ClientPort, &out.ClientPort
		*out = new(int32)
		**out = **in
	}
	if in.Image != nil {
		in, out := &in.Image, &out.Image
		*out = new(string)
		**out = **in
	}
	if in.AuthSecretRef != nil {
		in, out := &in.AuthSecretRef, &out.AuthSecretRef
		*out = new(v1.SecretReference)
		**out = **in
	}
	if in.Metrics != nil {
		in, out := &in.Metrics, &out.Metrics
		*out = new(MetricsLevel)
		**out = **in
	}
	if in.Resources != nil {
		in, out := &in.Resources, &out.Resources
		*out = new(v1.ResourceRequirements)
		(*in).DeepCopyInto(*out)
	}
	if in.TLS != nil {
		in, out := &in.TLS, &out.TLS
		*out = new(TLSConfig)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EtcdConfig.
func (in *EtcdConfig) DeepCopy() *EtcdConfig {
	if in == nil {
		return nil
	}
	out := new(EtcdConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EtcdList) DeepCopyInto(out *EtcdList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Etcd, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EtcdList.
func (in *EtcdList) DeepCopy() *EtcdList {
	if in == nil {
		return nil
	}
	out := new(EtcdList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *EtcdList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EtcdSpec) DeepCopyInto(out *EtcdSpec) {
	*out = *in
	if in.Selector != nil {
		in, out := &in.Selector, &out.Selector
		*out = new(metav1.LabelSelector)
		(*in).DeepCopyInto(*out)
	}
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Annotations != nil {
		in, out := &in.Annotations, &out.Annotations
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	in.Etcd.DeepCopyInto(&out.Etcd)
	in.Backup.DeepCopyInto(&out.Backup)
	if in.PriorityClassName != nil {
		in, out := &in.PriorityClassName, &out.PriorityClassName
		*out = new(string)
		**out = **in
	}
	if in.StorageClass != nil {
		in, out := &in.StorageClass, &out.StorageClass
		*out = new(string)
		**out = **in
	}
	if in.StorageCapacity != nil {
		in, out := &in.StorageCapacity, &out.StorageCapacity
		x := (*in).DeepCopy()
		*out = &x
	}
	if in.VolumeClaimTemplate != nil {
		in, out := &in.VolumeClaimTemplate, &out.VolumeClaimTemplate
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EtcdSpec.
func (in *EtcdSpec) DeepCopy() *EtcdSpec {
	if in == nil {
		return nil
	}
	out := new(EtcdSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EtcdStatus) DeepCopyInto(out *EtcdStatus) {
	*out = *in
	if in.ObservedGeneration != nil {
		in, out := &in.ObservedGeneration, &out.ObservedGeneration
		*out = new(int64)
		**out = **in
	}
	out.Etcd = in.Etcd
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ServiceName != nil {
		in, out := &in.ServiceName, &out.ServiceName
		*out = new(string)
		**out = **in
	}
	if in.LastError != nil {
		in, out := &in.LastError, &out.LastError
		*out = new(string)
		**out = **in
	}
	if in.Ready != nil {
		in, out := &in.Ready, &out.Ready
		*out = new(bool)
		**out = **in
	}
	if in.LabelSelector != nil {
		in, out := &in.LabelSelector, &out.LabelSelector
		*out = new(metav1.LabelSelector)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EtcdStatus.
func (in *EtcdStatus) DeepCopy() *EtcdStatus {
	if in == nil {
		return nil
	}
	out := new(EtcdStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LastOperation) DeepCopyInto(out *LastOperation) {
	*out = *in
	in.LastUpdateTime.DeepCopyInto(&out.LastUpdateTime)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LastOperation.
func (in *LastOperation) DeepCopy() *LastOperation {
	if in == nil {
		return nil
	}
	out := new(LastOperation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StoreSpec) DeepCopyInto(out *StoreSpec) {
	*out = *in
	if in.Container != nil {
		in, out := &in.Container, &out.Container
		*out = new(string)
		**out = **in
	}
	if in.Provider != nil {
		in, out := &in.Provider, &out.Provider
		*out = new(StorageProvider)
		**out = **in
	}
	if in.SecretRef != nil {
		in, out := &in.SecretRef, &out.SecretRef
		*out = new(v1.SecretReference)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StoreSpec.
func (in *StoreSpec) DeepCopy() *StoreSpec {
	if in == nil {
		return nil
	}
	out := new(StoreSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TLSConfig) DeepCopyInto(out *TLSConfig) {
	*out = *in
	out.ServerTLSSecretRef = in.ServerTLSSecretRef
	out.ClientTLSSecretRef = in.ClientTLSSecretRef
	out.TLSCASecretRef = in.TLSCASecretRef
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TLSConfig.
func (in *TLSConfig) DeepCopy() *TLSConfig {
	if in == nil {
		return nil
	}
	out := new(TLSConfig)
	in.DeepCopyInto(out)
	return out
}
