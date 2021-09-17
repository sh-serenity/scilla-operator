// +build !ignore_autogenerated

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1

import (
	corev1 "k8s.io/api/core/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AlternatorSpec) DeepCopyInto(out *AlternatorSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AlternatorSpec.
func (in *AlternatorSpec) DeepCopy() *AlternatorSpec {
	if in == nil {
		return nil
	}
	out := new(AlternatorSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BackupTaskSpec) DeepCopyInto(out *BackupTaskSpec) {
	*out = *in
	in.SchedulerTaskSpec.DeepCopyInto(&out.SchedulerTaskSpec)
	if in.DC != nil {
		in, out := &in.DC, &out.DC
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Keyspace != nil {
		in, out := &in.Keyspace, &out.Keyspace
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Location != nil {
		in, out := &in.Location, &out.Location
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.RateLimit != nil {
		in, out := &in.RateLimit, &out.RateLimit
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.SnapshotParallel != nil {
		in, out := &in.SnapshotParallel, &out.SnapshotParallel
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.UploadParallel != nil {
		in, out := &in.UploadParallel, &out.UploadParallel
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BackupTaskSpec.
func (in *BackupTaskSpec) DeepCopy() *BackupTaskSpec {
	if in == nil {
		return nil
	}
	out := new(BackupTaskSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BackupTaskStatus) DeepCopyInto(out *BackupTaskStatus) {
	*out = *in
	in.BackupTaskSpec.DeepCopyInto(&out.BackupTaskSpec)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BackupTaskStatus.
func (in *BackupTaskStatus) DeepCopy() *BackupTaskStatus {
	if in == nil {
		return nil
	}
	out := new(BackupTaskStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DatacenterSpec) DeepCopyInto(out *DatacenterSpec) {
	*out = *in
	if in.Racks != nil {
		in, out := &in.Racks, &out.Racks
		*out = make([]RackSpec, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DatacenterSpec.
func (in *DatacenterSpec) DeepCopy() *DatacenterSpec {
	if in == nil {
		return nil
	}
	out := new(DatacenterSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GenericUpgradeSpec) DeepCopyInto(out *GenericUpgradeSpec) {
	*out = *in
	out.PollInterval = in.PollInterval
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GenericUpgradeSpec.
func (in *GenericUpgradeSpec) DeepCopy() *GenericUpgradeSpec {
	if in == nil {
		return nil
	}
	out := new(GenericUpgradeSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Network) DeepCopyInto(out *Network) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Network.
func (in *Network) DeepCopy() *Network {
	if in == nil {
		return nil
	}
	out := new(Network)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PlacementSpec) DeepCopyInto(out *PlacementSpec) {
	*out = *in
	if in.NodeAffinity != nil {
		in, out := &in.NodeAffinity, &out.NodeAffinity
		*out = new(corev1.NodeAffinity)
		(*in).DeepCopyInto(*out)
	}
	if in.PodAffinity != nil {
		in, out := &in.PodAffinity, &out.PodAffinity
		*out = new(corev1.PodAffinity)
		(*in).DeepCopyInto(*out)
	}
	if in.PodAntiAffinity != nil {
		in, out := &in.PodAntiAffinity, &out.PodAntiAffinity
		*out = new(corev1.PodAntiAffinity)
		(*in).DeepCopyInto(*out)
	}
	if in.Tolerations != nil {
		in, out := &in.Tolerations, &out.Tolerations
		*out = make([]corev1.Toleration, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PlacementSpec.
func (in *PlacementSpec) DeepCopy() *PlacementSpec {
	if in == nil {
		return nil
	}
	out := new(PlacementSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RackCondition) DeepCopyInto(out *RackCondition) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RackCondition.
func (in *RackCondition) DeepCopy() *RackCondition {
	if in == nil {
		return nil
	}
	out := new(RackCondition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RackSpec) DeepCopyInto(out *RackSpec) {
	*out = *in
	in.Storage.DeepCopyInto(&out.Storage)
	if in.Placement != nil {
		in, out := &in.Placement, &out.Placement
		*out = new(PlacementSpec)
		(*in).DeepCopyInto(*out)
	}
	in.Resources.DeepCopyInto(&out.Resources)
	in.AgentResources.DeepCopyInto(&out.AgentResources)
	if in.Volumes != nil {
		in, out := &in.Volumes, &out.Volumes
		*out = make([]corev1.Volume, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.VolumeMounts != nil {
		in, out := &in.VolumeMounts, &out.VolumeMounts
		*out = make([]corev1.VolumeMount, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.AgentVolumeMounts != nil {
		in, out := &in.AgentVolumeMounts, &out.AgentVolumeMounts
		*out = make([]corev1.VolumeMount, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RackSpec.
func (in *RackSpec) DeepCopy() *RackSpec {
	if in == nil {
		return nil
	}
	out := new(RackSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RackStatus) DeepCopyInto(out *RackStatus) {
	*out = *in
	if in.UpdatedMembers != nil {
		in, out := &in.UpdatedMembers, &out.UpdatedMembers
		*out = new(int32)
		**out = **in
	}
	if in.Stale != nil {
		in, out := &in.Stale, &out.Stale
		*out = new(bool)
		**out = **in
	}
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]RackCondition, len(*in))
		copy(*out, *in)
	}
	if in.ReplaceAddressFirstBoot != nil {
		in, out := &in.ReplaceAddressFirstBoot, &out.ReplaceAddressFirstBoot
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RackStatus.
func (in *RackStatus) DeepCopy() *RackStatus {
	if in == nil {
		return nil
	}
	out := new(RackStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RepairTaskSpec) DeepCopyInto(out *RepairTaskSpec) {
	*out = *in
	in.SchedulerTaskSpec.DeepCopyInto(&out.SchedulerTaskSpec)
	if in.DC != nil {
		in, out := &in.DC, &out.DC
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Keyspace != nil {
		in, out := &in.Keyspace, &out.Keyspace
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Host != nil {
		in, out := &in.Host, &out.Host
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RepairTaskSpec.
func (in *RepairTaskSpec) DeepCopy() *RepairTaskSpec {
	if in == nil {
		return nil
	}
	out := new(RepairTaskSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RepairTaskStatus) DeepCopyInto(out *RepairTaskStatus) {
	*out = *in
	in.RepairTaskSpec.DeepCopyInto(&out.RepairTaskSpec)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RepairTaskStatus.
func (in *RepairTaskStatus) DeepCopy() *RepairTaskStatus {
	if in == nil {
		return nil
	}
	out := new(RepairTaskStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SchedulerTaskSpec) DeepCopyInto(out *SchedulerTaskSpec) {
	*out = *in
	if in.NumRetries != nil {
		in, out := &in.NumRetries, &out.NumRetries
		*out = new(int64)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SchedulerTaskSpec.
func (in *SchedulerTaskSpec) DeepCopy() *SchedulerTaskSpec {
	if in == nil {
		return nil
	}
	out := new(SchedulerTaskSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ScyllaCluster) DeepCopyInto(out *ScyllaCluster) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ScyllaCluster.
func (in *ScyllaCluster) DeepCopy() *ScyllaCluster {
	if in == nil {
		return nil
	}
	out := new(ScyllaCluster)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ScyllaCluster) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ScyllaClusterList) DeepCopyInto(out *ScyllaClusterList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ScyllaCluster, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ScyllaClusterList.
func (in *ScyllaClusterList) DeepCopy() *ScyllaClusterList {
	if in == nil {
		return nil
	}
	out := new(ScyllaClusterList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ScyllaClusterList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ScyllaClusterSpec) DeepCopyInto(out *ScyllaClusterSpec) {
	*out = *in
	if in.Alternator != nil {
		in, out := &in.Alternator, &out.Alternator
		*out = new(AlternatorSpec)
		**out = **in
	}
	if in.GenericUpgrade != nil {
		in, out := &in.GenericUpgrade, &out.GenericUpgrade
		*out = new(GenericUpgradeSpec)
		**out = **in
	}
	in.Datacenter.DeepCopyInto(&out.Datacenter)
	if in.Sysctls != nil {
		in, out := &in.Sysctls, &out.Sysctls
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	out.Network = in.Network
	if in.Repairs != nil {
		in, out := &in.Repairs, &out.Repairs
		*out = make([]RepairTaskSpec, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Backups != nil {
		in, out := &in.Backups, &out.Backups
		*out = make([]BackupTaskSpec, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ImagePullSecrets != nil {
		in, out := &in.ImagePullSecrets, &out.ImagePullSecrets
		*out = make([]corev1.LocalObjectReference, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ScyllaClusterSpec.
func (in *ScyllaClusterSpec) DeepCopy() *ScyllaClusterSpec {
	if in == nil {
		return nil
	}
	out := new(ScyllaClusterSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ScyllaClusterStatus) DeepCopyInto(out *ScyllaClusterStatus) {
	*out = *in
	if in.ObservedGeneration != nil {
		in, out := &in.ObservedGeneration, &out.ObservedGeneration
		*out = new(int64)
		**out = **in
	}
	if in.Racks != nil {
		in, out := &in.Racks, &out.Racks
		*out = make(map[string]RackStatus, len(*in))
		for key, val := range *in {
			(*out)[key] = *val.DeepCopy()
		}
	}
	if in.ManagerID != nil {
		in, out := &in.ManagerID, &out.ManagerID
		*out = new(string)
		**out = **in
	}
	if in.Repairs != nil {
		in, out := &in.Repairs, &out.Repairs
		*out = make([]RepairTaskStatus, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Backups != nil {
		in, out := &in.Backups, &out.Backups
		*out = make([]BackupTaskStatus, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Upgrade != nil {
		in, out := &in.Upgrade, &out.Upgrade
		*out = new(UpgradeStatus)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ScyllaClusterStatus.
func (in *ScyllaClusterStatus) DeepCopy() *ScyllaClusterStatus {
	if in == nil {
		return nil
	}
	out := new(ScyllaClusterStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StorageSpec) DeepCopyInto(out *StorageSpec) {
	*out = *in
	if in.StorageClassName != nil {
		in, out := &in.StorageClassName, &out.StorageClassName
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StorageSpec.
func (in *StorageSpec) DeepCopy() *StorageSpec {
	if in == nil {
		return nil
	}
	out := new(StorageSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UpgradeStatus) DeepCopyInto(out *UpgradeStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UpgradeStatus.
func (in *UpgradeStatus) DeepCopy() *UpgradeStatus {
	if in == nil {
		return nil
	}
	out := new(UpgradeStatus)
	in.DeepCopyInto(out)
	return out
}
