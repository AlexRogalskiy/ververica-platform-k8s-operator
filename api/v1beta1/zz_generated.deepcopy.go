// +build !ignore_autogenerated

/*
Copyright 2019 FinTech Studios, Inc.

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
	"k8s.io/api/core/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JSONPatchGeneric) DeepCopyInto(out *JSONPatchGeneric) {
	*out = *in
	if in.Value != nil {
		in, out := &in.Value, &out.Value
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JSONPatchGeneric.
func (in *JSONPatchGeneric) DeepCopy() *JSONPatchGeneric {
	if in == nil {
		return nil
	}
	out := new(JSONPatchGeneric)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NamespaceRoleBinding) DeepCopyInto(out *NamespaceRoleBinding) {
	*out = *in
	if in.Members != nil {
		in, out := &in.Members, &out.Members
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NamespaceRoleBinding.
func (in *NamespaceRoleBinding) DeepCopy() *NamespaceRoleBinding {
	if in == nil {
		return nil
	}
	out := new(NamespaceRoleBinding)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VpArtifact) DeepCopyInto(out *VpArtifact) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VpArtifact.
func (in *VpArtifact) DeepCopy() *VpArtifact {
	if in == nil {
		return nil
	}
	out := new(VpArtifact)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VpDeployment) DeepCopyInto(out *VpDeployment) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VpDeployment.
func (in *VpDeployment) DeepCopy() *VpDeployment {
	if in == nil {
		return nil
	}
	out := new(VpDeployment)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VpDeployment) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VpDeploymentList) DeepCopyInto(out *VpDeploymentList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]VpDeployment, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VpDeploymentList.
func (in *VpDeploymentList) DeepCopy() *VpDeploymentList {
	if in == nil {
		return nil
	}
	out := new(VpDeploymentList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VpDeploymentList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VpDeploymentObjectSpec) DeepCopyInto(out *VpDeploymentObjectSpec) {
	*out = *in
	in.Metadata.DeepCopyInto(&out.Metadata)
	in.Spec.DeepCopyInto(&out.Spec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VpDeploymentObjectSpec.
func (in *VpDeploymentObjectSpec) DeepCopy() *VpDeploymentObjectSpec {
	if in == nil {
		return nil
	}
	out := new(VpDeploymentObjectSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VpDeploymentRestoreStrategy) DeepCopyInto(out *VpDeploymentRestoreStrategy) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VpDeploymentRestoreStrategy.
func (in *VpDeploymentRestoreStrategy) DeepCopy() *VpDeploymentRestoreStrategy {
	if in == nil {
		return nil
	}
	out := new(VpDeploymentRestoreStrategy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VpDeploymentSpec) DeepCopyInto(out *VpDeploymentSpec) {
	*out = *in
	if in.UpgradeStrategy != nil {
		in, out := &in.UpgradeStrategy, &out.UpgradeStrategy
		*out = new(VpDeploymentUpgradeStrategy)
		**out = **in
	}
	if in.RestoreStrategy != nil {
		in, out := &in.RestoreStrategy, &out.RestoreStrategy
		*out = new(VpDeploymentRestoreStrategy)
		**out = **in
	}
	if in.MaxSavepointCreationAttempts != nil {
		in, out := &in.MaxSavepointCreationAttempts, &out.MaxSavepointCreationAttempts
		*out = new(int32)
		**out = **in
	}
	if in.MaxJobCreationAttempts != nil {
		in, out := &in.MaxJobCreationAttempts, &out.MaxJobCreationAttempts
		*out = new(int32)
		**out = **in
	}
	if in.Template != nil {
		in, out := &in.Template, &out.Template
		*out = new(VpDeploymentTemplate)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VpDeploymentSpec.
func (in *VpDeploymentSpec) DeepCopy() *VpDeploymentSpec {
	if in == nil {
		return nil
	}
	out := new(VpDeploymentSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VpDeploymentStatus) DeepCopyInto(out *VpDeploymentStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VpDeploymentStatus.
func (in *VpDeploymentStatus) DeepCopy() *VpDeploymentStatus {
	if in == nil {
		return nil
	}
	out := new(VpDeploymentStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VpDeploymentTarget) DeepCopyInto(out *VpDeploymentTarget) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VpDeploymentTarget.
func (in *VpDeploymentTarget) DeepCopy() *VpDeploymentTarget {
	if in == nil {
		return nil
	}
	out := new(VpDeploymentTarget)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VpDeploymentTarget) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VpDeploymentTargetList) DeepCopyInto(out *VpDeploymentTargetList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]VpDeploymentTarget, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VpDeploymentTargetList.
func (in *VpDeploymentTargetList) DeepCopy() *VpDeploymentTargetList {
	if in == nil {
		return nil
	}
	out := new(VpDeploymentTargetList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VpDeploymentTargetList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VpDeploymentTargetObjectSpec) DeepCopyInto(out *VpDeploymentTargetObjectSpec) {
	*out = *in
	in.Metadata.DeepCopyInto(&out.Metadata)
	out.Spec = in.Spec
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VpDeploymentTargetObjectSpec.
func (in *VpDeploymentTargetObjectSpec) DeepCopy() *VpDeploymentTargetObjectSpec {
	if in == nil {
		return nil
	}
	out := new(VpDeploymentTargetObjectSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VpDeploymentTargetSpec) DeepCopyInto(out *VpDeploymentTargetSpec) {
	*out = *in
	out.Kubernetes = in.Kubernetes
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VpDeploymentTargetSpec.
func (in *VpDeploymentTargetSpec) DeepCopy() *VpDeploymentTargetSpec {
	if in == nil {
		return nil
	}
	out := new(VpDeploymentTargetSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VpDeploymentTemplate) DeepCopyInto(out *VpDeploymentTemplate) {
	*out = *in
	if in.Metadata != nil {
		in, out := &in.Metadata, &out.Metadata
		*out = new(VpDeploymentTemplateMetadata)
		(*in).DeepCopyInto(*out)
	}
	if in.Spec != nil {
		in, out := &in.Spec, &out.Spec
		*out = new(VpDeploymentTemplateSpec)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VpDeploymentTemplate.
func (in *VpDeploymentTemplate) DeepCopy() *VpDeploymentTemplate {
	if in == nil {
		return nil
	}
	out := new(VpDeploymentTemplate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VpDeploymentTemplateMetadata) DeepCopyInto(out *VpDeploymentTemplateMetadata) {
	*out = *in
	if in.Annotations != nil {
		in, out := &in.Annotations, &out.Annotations
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VpDeploymentTemplateMetadata.
func (in *VpDeploymentTemplateMetadata) DeepCopy() *VpDeploymentTemplateMetadata {
	if in == nil {
		return nil
	}
	out := new(VpDeploymentTemplateMetadata)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VpDeploymentTemplateSpec) DeepCopyInto(out *VpDeploymentTemplateSpec) {
	*out = *in
	if in.Artifact != nil {
		in, out := &in.Artifact, &out.Artifact
		*out = new(VpArtifact)
		**out = **in
	}
	if in.Parallelism != nil {
		in, out := &in.Parallelism, &out.Parallelism
		*out = new(int32)
		**out = **in
	}
	if in.NumberOfTaskManagers != nil {
		in, out := &in.NumberOfTaskManagers, &out.NumberOfTaskManagers
		*out = new(int32)
		**out = **in
	}
	if in.Resources != nil {
		in, out := &in.Resources, &out.Resources
		*out = make(map[string]VpResourceSpec, len(*in))
		for key, val := range *in {
			(*out)[key] = *val.DeepCopy()
		}
	}
	if in.FlinkConfiguration != nil {
		in, out := &in.FlinkConfiguration, &out.FlinkConfiguration
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Logging != nil {
		in, out := &in.Logging, &out.Logging
		*out = new(VpLogging)
		(*in).DeepCopyInto(*out)
	}
	if in.Kubernetes != nil {
		in, out := &in.Kubernetes, &out.Kubernetes
		*out = new(VpKubernetesOptions)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VpDeploymentTemplateSpec.
func (in *VpDeploymentTemplateSpec) DeepCopy() *VpDeploymentTemplateSpec {
	if in == nil {
		return nil
	}
	out := new(VpDeploymentTemplateSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VpDeploymentUpgradeStrategy) DeepCopyInto(out *VpDeploymentUpgradeStrategy) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VpDeploymentUpgradeStrategy.
func (in *VpDeploymentUpgradeStrategy) DeepCopy() *VpDeploymentUpgradeStrategy {
	if in == nil {
		return nil
	}
	out := new(VpDeploymentUpgradeStrategy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VpKubernetesOptions) DeepCopyInto(out *VpKubernetesOptions) {
	*out = *in
	if in.Pods != nil {
		in, out := &in.Pods, &out.Pods
		*out = new(VpPodSpec)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VpKubernetesOptions.
func (in *VpKubernetesOptions) DeepCopy() *VpKubernetesOptions {
	if in == nil {
		return nil
	}
	out := new(VpKubernetesOptions)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VpKubernetesTarget) DeepCopyInto(out *VpKubernetesTarget) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VpKubernetesTarget.
func (in *VpKubernetesTarget) DeepCopy() *VpKubernetesTarget {
	if in == nil {
		return nil
	}
	out := new(VpKubernetesTarget)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VpLogging) DeepCopyInto(out *VpLogging) {
	*out = *in
	if in.Log4jLoggers != nil {
		in, out := &in.Log4jLoggers, &out.Log4jLoggers
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VpLogging.
func (in *VpLogging) DeepCopy() *VpLogging {
	if in == nil {
		return nil
	}
	out := new(VpLogging)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VpMetadata) DeepCopyInto(out *VpMetadata) {
	*out = *in
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
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VpMetadata.
func (in *VpMetadata) DeepCopy() *VpMetadata {
	if in == nil {
		return nil
	}
	out := new(VpMetadata)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VpNamespace) DeepCopyInto(out *VpNamespace) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VpNamespace.
func (in *VpNamespace) DeepCopy() *VpNamespace {
	if in == nil {
		return nil
	}
	out := new(VpNamespace)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VpNamespace) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VpNamespaceList) DeepCopyInto(out *VpNamespaceList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]VpNamespace, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VpNamespaceList.
func (in *VpNamespaceList) DeepCopy() *VpNamespaceList {
	if in == nil {
		return nil
	}
	out := new(VpNamespaceList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VpNamespaceList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VpNamespaceSpec) DeepCopyInto(out *VpNamespaceSpec) {
	*out = *in
	if in.RoleBindings != nil {
		in, out := &in.RoleBindings, &out.RoleBindings
		*out = make([]NamespaceRoleBinding, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VpNamespaceSpec.
func (in *VpNamespaceSpec) DeepCopy() *VpNamespaceSpec {
	if in == nil {
		return nil
	}
	out := new(VpNamespaceSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VpNamespaceStatus) DeepCopyInto(out *VpNamespaceStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VpNamespaceStatus.
func (in *VpNamespaceStatus) DeepCopy() *VpNamespaceStatus {
	if in == nil {
		return nil
	}
	out := new(VpNamespaceStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VpPodSpec) DeepCopyInto(out *VpPodSpec) {
	*out = *in
	if in.Annotations != nil {
		in, out := &in.Annotations, &out.Annotations
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.EnvVars != nil {
		in, out := &in.EnvVars, &out.EnvVars
		*out = make([]v1.EnvVar, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.VolumeMounts != nil {
		in, out := &in.VolumeMounts, &out.VolumeMounts
		*out = make([]VpVolumeAndMount, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.NodeSelector != nil {
		in, out := &in.NodeSelector, &out.NodeSelector
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.SecurityContext != nil {
		in, out := &in.SecurityContext, &out.SecurityContext
		*out = new(v1.PodSecurityContext)
		(*in).DeepCopyInto(*out)
	}
	if in.ImagePullSecrets != nil {
		in, out := &in.ImagePullSecrets, &out.ImagePullSecrets
		*out = make([]v1.LocalObjectReference, len(*in))
		copy(*out, *in)
	}
	if in.Affinity != nil {
		in, out := &in.Affinity, &out.Affinity
		*out = new(v1.Affinity)
		(*in).DeepCopyInto(*out)
	}
	if in.Tolerations != nil {
		in, out := &in.Tolerations, &out.Tolerations
		*out = make([]v1.Toleration, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VpPodSpec.
func (in *VpPodSpec) DeepCopy() *VpPodSpec {
	if in == nil {
		return nil
	}
	out := new(VpPodSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VpResourceSpec) DeepCopyInto(out *VpResourceSpec) {
	*out = *in
	out.CPU = in.CPU.DeepCopy()
	if in.Memory != nil {
		in, out := &in.Memory, &out.Memory
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VpResourceSpec.
func (in *VpResourceSpec) DeepCopy() *VpResourceSpec {
	if in == nil {
		return nil
	}
	out := new(VpResourceSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VpSavepoint) DeepCopyInto(out *VpSavepoint) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VpSavepoint.
func (in *VpSavepoint) DeepCopy() *VpSavepoint {
	if in == nil {
		return nil
	}
	out := new(VpSavepoint)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VpSavepoint) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VpSavepointFailure) DeepCopyInto(out *VpSavepointFailure) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VpSavepointFailure.
func (in *VpSavepointFailure) DeepCopy() *VpSavepointFailure {
	if in == nil {
		return nil
	}
	out := new(VpSavepointFailure)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VpSavepointList) DeepCopyInto(out *VpSavepointList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]VpSavepoint, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VpSavepointList.
func (in *VpSavepointList) DeepCopy() *VpSavepointList {
	if in == nil {
		return nil
	}
	out := new(VpSavepointList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VpSavepointList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VpSavepointMetadata) DeepCopyInto(out *VpSavepointMetadata) {
	*out = *in
	in.VpMetadata.DeepCopyInto(&out.VpMetadata)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VpSavepointMetadata.
func (in *VpSavepointMetadata) DeepCopy() *VpSavepointMetadata {
	if in == nil {
		return nil
	}
	out := new(VpSavepointMetadata)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VpSavepointObjectSpec) DeepCopyInto(out *VpSavepointObjectSpec) {
	*out = *in
	in.Metadata.DeepCopyInto(&out.Metadata)
	out.Spec = in.Spec
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VpSavepointObjectSpec.
func (in *VpSavepointObjectSpec) DeepCopy() *VpSavepointObjectSpec {
	if in == nil {
		return nil
	}
	out := new(VpSavepointObjectSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VpSavepointSpec) DeepCopyInto(out *VpSavepointSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VpSavepointSpec.
func (in *VpSavepointSpec) DeepCopy() *VpSavepointSpec {
	if in == nil {
		return nil
	}
	out := new(VpSavepointSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VpSavepointStatus) DeepCopyInto(out *VpSavepointStatus) {
	*out = *in
	out.Failure = in.Failure
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VpSavepointStatus.
func (in *VpSavepointStatus) DeepCopy() *VpSavepointStatus {
	if in == nil {
		return nil
	}
	out := new(VpSavepointStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VpVolumeAndMount) DeepCopyInto(out *VpVolumeAndMount) {
	*out = *in
	if in.Volume != nil {
		in, out := &in.Volume, &out.Volume
		*out = new(v1.Volume)
		(*in).DeepCopyInto(*out)
	}
	if in.VolumeMount != nil {
		in, out := &in.VolumeMount, &out.VolumeMount
		*out = new(v1.VolumeMount)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VpVolumeAndMount.
func (in *VpVolumeAndMount) DeepCopy() *VpVolumeAndMount {
	if in == nil {
		return nil
	}
	out := new(VpVolumeAndMount)
	in.DeepCopyInto(out)
	return out
}
