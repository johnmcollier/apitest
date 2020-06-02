// +build !ignore_autogenerated

// Code generated by operator-sdk. DO NOT EDIT.

package v1alpha1

import (
	v1 "k8s.io/api/core/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BaseCommand) DeepCopyInto(out *BaseCommand) {
	*out = *in
	if in.Group != nil {
		in, out := &in.Group, &out.Group
		*out = new(CommandGroup)
		**out = **in
	}
	if in.Attributes != nil {
		in, out := &in.Attributes, &out.Attributes
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BaseCommand.
func (in *BaseCommand) DeepCopy() *BaseCommand {
	if in == nil {
		return nil
	}
	out := new(BaseCommand)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BaseComponent) DeepCopyInto(out *BaseComponent) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BaseComponent.
func (in *BaseComponent) DeepCopy() *BaseComponent {
	if in == nil {
		return nil
	}
	out := new(BaseComponent)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Command) DeepCopyInto(out *Command) {
	*out = *in
	if in.Exec != nil {
		in, out := &in.Exec, &out.Exec
		*out = new(ExecCommand)
		(*in).DeepCopyInto(*out)
	}
	if in.VscodeTask != nil {
		in, out := &in.VscodeTask, &out.VscodeTask
		*out = new(VscodeConfigurationCommand)
		(*in).DeepCopyInto(*out)
	}
	if in.VscodeLaunch != nil {
		in, out := &in.VscodeLaunch, &out.VscodeLaunch
		*out = new(VscodeConfigurationCommand)
		(*in).DeepCopyInto(*out)
	}
	if in.Composite != nil {
		in, out := &in.Composite, &out.Composite
		*out = new(CompositeCommand)
		(*in).DeepCopyInto(*out)
	}
	if in.Custom != nil {
		in, out := &in.Custom, &out.Custom
		*out = new(CustomCommand)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Command.
func (in *Command) DeepCopy() *Command {
	if in == nil {
		return nil
	}
	out := new(Command)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CommandGroup) DeepCopyInto(out *CommandGroup) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CommandGroup.
func (in *CommandGroup) DeepCopy() *CommandGroup {
	if in == nil {
		return nil
	}
	out := new(CommandGroup)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CommonProjectSource) DeepCopyInto(out *CommonProjectSource) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CommonProjectSource.
func (in *CommonProjectSource) DeepCopy() *CommonProjectSource {
	if in == nil {
		return nil
	}
	out := new(CommonProjectSource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Component) DeepCopyInto(out *Component) {
	*out = *in
	if in.Container != nil {
		in, out := &in.Container, &out.Container
		*out = new(ContainerComponent)
		(*in).DeepCopyInto(*out)
	}
	if in.Volume != nil {
		in, out := &in.Volume, &out.Volume
		*out = new(VolumeComponent)
		**out = **in
	}
	if in.Plugin != nil {
		in, out := &in.Plugin, &out.Plugin
		*out = new(PluginComponent)
		(*in).DeepCopyInto(*out)
	}
	if in.Kubernetes != nil {
		in, out := &in.Kubernetes, &out.Kubernetes
		*out = new(KubernetesComponent)
		**out = **in
	}
	if in.Openshift != nil {
		in, out := &in.Openshift, &out.Openshift
		*out = new(OpenshiftComponent)
		**out = **in
	}
	if in.Custom != nil {
		in, out := &in.Custom, &out.Custom
		*out = new(CustomComponent)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Component.
func (in *Component) DeepCopy() *Component {
	if in == nil {
		return nil
	}
	out := new(Component)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ComponentOverride) DeepCopyInto(out *ComponentOverride) {
	*out = *in
	if in.Container != nil {
		in, out := &in.Container, &out.Container
		*out = new(ContainerComponent)
		(*in).DeepCopyInto(*out)
	}
	if in.Volume != nil {
		in, out := &in.Volume, &out.Volume
		*out = new(VolumeComponent)
		**out = **in
	}
	if in.Kubernetes != nil {
		in, out := &in.Kubernetes, &out.Kubernetes
		*out = new(KubernetesComponent)
		**out = **in
	}
	if in.Openshift != nil {
		in, out := &in.Openshift, &out.Openshift
		*out = new(OpenshiftComponent)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ComponentOverride.
func (in *ComponentOverride) DeepCopy() *ComponentOverride {
	if in == nil {
		return nil
	}
	out := new(ComponentOverride)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CompositeCommand) DeepCopyInto(out *CompositeCommand) {
	*out = *in
	in.LabeledCommand.DeepCopyInto(&out.LabeledCommand)
	if in.Commands != nil {
		in, out := &in.Commands, &out.Commands
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CompositeCommand.
func (in *CompositeCommand) DeepCopy() *CompositeCommand {
	if in == nil {
		return nil
	}
	out := new(CompositeCommand)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Container) DeepCopyInto(out *Container) {
	*out = *in
	if in.Env != nil {
		in, out := &in.Env, &out.Env
		*out = make([]EnvVar, len(*in))
		copy(*out, *in)
	}
	if in.VolumeMounts != nil {
		in, out := &in.VolumeMounts, &out.VolumeMounts
		*out = make([]VolumeMount, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Container.
func (in *Container) DeepCopy() *Container {
	if in == nil {
		return nil
	}
	out := new(Container)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ContainerComponent) DeepCopyInto(out *ContainerComponent) {
	*out = *in
	out.BaseComponent = in.BaseComponent
	in.Container.DeepCopyInto(&out.Container)
	if in.Endpoints != nil {
		in, out := &in.Endpoints, &out.Endpoints
		*out = make([]Endpoint, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ContainerComponent.
func (in *ContainerComponent) DeepCopy() *ContainerComponent {
	if in == nil {
		return nil
	}
	out := new(ContainerComponent)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CustomCommand) DeepCopyInto(out *CustomCommand) {
	*out = *in
	in.LabeledCommand.DeepCopyInto(&out.LabeledCommand)
	in.EmbeddedResource.DeepCopyInto(&out.EmbeddedResource)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CustomCommand.
func (in *CustomCommand) DeepCopy() *CustomCommand {
	if in == nil {
		return nil
	}
	out := new(CustomCommand)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CustomComponent) DeepCopyInto(out *CustomComponent) {
	*out = *in
	in.EmbeddedResource.DeepCopyInto(&out.EmbeddedResource)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CustomComponent.
func (in *CustomComponent) DeepCopy() *CustomComponent {
	if in == nil {
		return nil
	}
	out := new(CustomComponent)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CustomProjectSource) DeepCopyInto(out *CustomProjectSource) {
	*out = *in
	in.EmbeddedResource.DeepCopyInto(&out.EmbeddedResource)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CustomProjectSource.
func (in *CustomProjectSource) DeepCopy() *CustomProjectSource {
	if in == nil {
		return nil
	}
	out := new(CustomProjectSource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DevWorkspace) DeepCopyInto(out *DevWorkspace) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DevWorkspace.
func (in *DevWorkspace) DeepCopy() *DevWorkspace {
	if in == nil {
		return nil
	}
	out := new(DevWorkspace)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DevWorkspace) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DevWorkspaceList) DeepCopyInto(out *DevWorkspaceList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]DevWorkspace, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DevWorkspaceList.
func (in *DevWorkspaceList) DeepCopy() *DevWorkspaceList {
	if in == nil {
		return nil
	}
	out := new(DevWorkspaceList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DevWorkspaceList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DevWorkspaceSpec) DeepCopyInto(out *DevWorkspaceSpec) {
	*out = *in
	in.Template.DeepCopyInto(&out.Template)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DevWorkspaceSpec.
func (in *DevWorkspaceSpec) DeepCopy() *DevWorkspaceSpec {
	if in == nil {
		return nil
	}
	out := new(DevWorkspaceSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DevWorkspaceStatus) DeepCopyInto(out *DevWorkspaceStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]WorkspaceCondition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DevWorkspaceStatus.
func (in *DevWorkspaceStatus) DeepCopy() *DevWorkspaceStatus {
	if in == nil {
		return nil
	}
	out := new(DevWorkspaceStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DevWorkspaceTemplate) DeepCopyInto(out *DevWorkspaceTemplate) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DevWorkspaceTemplate.
func (in *DevWorkspaceTemplate) DeepCopy() *DevWorkspaceTemplate {
	if in == nil {
		return nil
	}
	out := new(DevWorkspaceTemplate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DevWorkspaceTemplate) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DevWorkspaceTemplateList) DeepCopyInto(out *DevWorkspaceTemplateList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]DevWorkspaceTemplate, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DevWorkspaceTemplateList.
func (in *DevWorkspaceTemplateList) DeepCopy() *DevWorkspaceTemplateList {
	if in == nil {
		return nil
	}
	out := new(DevWorkspaceTemplateList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DevWorkspaceTemplateList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DevWorkspaceTemplateSpec) DeepCopyInto(out *DevWorkspaceTemplateSpec) {
	*out = *in
	if in.Parent != nil {
		in, out := &in.Parent, &out.Parent
		*out = new(Parent)
		(*in).DeepCopyInto(*out)
	}
	in.DevWorkspaceTemplateSpecContent.DeepCopyInto(&out.DevWorkspaceTemplateSpecContent)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DevWorkspaceTemplateSpec.
func (in *DevWorkspaceTemplateSpec) DeepCopy() *DevWorkspaceTemplateSpec {
	if in == nil {
		return nil
	}
	out := new(DevWorkspaceTemplateSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DevWorkspaceTemplateSpecContent) DeepCopyInto(out *DevWorkspaceTemplateSpecContent) {
	*out = *in
	if in.Commands != nil {
		in, out := &in.Commands, &out.Commands
		*out = make([]Command, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	in.Events.DeepCopyInto(&out.Events)
	if in.Projects != nil {
		in, out := &in.Projects, &out.Projects
		*out = make([]Project, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Components != nil {
		in, out := &in.Components, &out.Components
		*out = make([]Component, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DevWorkspaceTemplateSpecContent.
func (in *DevWorkspaceTemplateSpecContent) DeepCopy() *DevWorkspaceTemplateSpecContent {
	if in == nil {
		return nil
	}
	out := new(DevWorkspaceTemplateSpecContent)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Endpoint) DeepCopyInto(out *Endpoint) {
	*out = *in
	if in.Configuration != nil {
		in, out := &in.Configuration, &out.Configuration
		*out = new(EndpointConfiguration)
		**out = **in
	}
	if in.Attributes != nil {
		in, out := &in.Attributes, &out.Attributes
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Endpoint.
func (in *Endpoint) DeepCopy() *Endpoint {
	if in == nil {
		return nil
	}
	out := new(Endpoint)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EndpointConfiguration) DeepCopyInto(out *EndpointConfiguration) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EndpointConfiguration.
func (in *EndpointConfiguration) DeepCopy() *EndpointConfiguration {
	if in == nil {
		return nil
	}
	out := new(EndpointConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EnvVar) DeepCopyInto(out *EnvVar) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EnvVar.
func (in *EnvVar) DeepCopy() *EnvVar {
	if in == nil {
		return nil
	}
	out := new(EnvVar)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Events) DeepCopyInto(out *Events) {
	*out = *in
	in.WorkspaceEvents.DeepCopyInto(&out.WorkspaceEvents)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Events.
func (in *Events) DeepCopy() *Events {
	if in == nil {
		return nil
	}
	out := new(Events)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExecCommand) DeepCopyInto(out *ExecCommand) {
	*out = *in
	in.LabeledCommand.DeepCopyInto(&out.LabeledCommand)
	if in.Env != nil {
		in, out := &in.Env, &out.Env
		*out = make([]EnvVar, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExecCommand.
func (in *ExecCommand) DeepCopy() *ExecCommand {
	if in == nil {
		return nil
	}
	out := new(ExecCommand)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GitLikeProjectSource) DeepCopyInto(out *GitLikeProjectSource) {
	*out = *in
	out.CommonProjectSource = in.CommonProjectSource
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GitLikeProjectSource.
func (in *GitLikeProjectSource) DeepCopy() *GitLikeProjectSource {
	if in == nil {
		return nil
	}
	out := new(GitLikeProjectSource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GitProjectSource) DeepCopyInto(out *GitProjectSource) {
	*out = *in
	out.GitLikeProjectSource = in.GitLikeProjectSource
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GitProjectSource.
func (in *GitProjectSource) DeepCopy() *GitProjectSource {
	if in == nil {
		return nil
	}
	out := new(GitProjectSource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GithubProjectSource) DeepCopyInto(out *GithubProjectSource) {
	*out = *in
	out.GitLikeProjectSource = in.GitLikeProjectSource
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GithubProjectSource.
func (in *GithubProjectSource) DeepCopy() *GithubProjectSource {
	if in == nil {
		return nil
	}
	out := new(GithubProjectSource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ImportReference) DeepCopyInto(out *ImportReference) {
	*out = *in
	in.ImportReferenceUnion.DeepCopyInto(&out.ImportReferenceUnion)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ImportReference.
func (in *ImportReference) DeepCopy() *ImportReference {
	if in == nil {
		return nil
	}
	out := new(ImportReference)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ImportReferenceUnion) DeepCopyInto(out *ImportReferenceUnion) {
	*out = *in
	if in.Kubernetes != nil {
		in, out := &in.Kubernetes, &out.Kubernetes
		*out = new(KubernetesCustomResourceImportReference)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ImportReferenceUnion.
func (in *ImportReferenceUnion) DeepCopy() *ImportReferenceUnion {
	if in == nil {
		return nil
	}
	out := new(ImportReferenceUnion)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *K8sLikeComponent) DeepCopyInto(out *K8sLikeComponent) {
	*out = *in
	out.BaseComponent = in.BaseComponent
	out.K8sLikeComponentLocation = in.K8sLikeComponentLocation
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new K8sLikeComponent.
func (in *K8sLikeComponent) DeepCopy() *K8sLikeComponent {
	if in == nil {
		return nil
	}
	out := new(K8sLikeComponent)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *K8sLikeComponentLocation) DeepCopyInto(out *K8sLikeComponentLocation) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new K8sLikeComponentLocation.
func (in *K8sLikeComponentLocation) DeepCopy() *K8sLikeComponentLocation {
	if in == nil {
		return nil
	}
	out := new(K8sLikeComponentLocation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KubernetesComponent) DeepCopyInto(out *KubernetesComponent) {
	*out = *in
	out.K8sLikeComponent = in.K8sLikeComponent
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KubernetesComponent.
func (in *KubernetesComponent) DeepCopy() *KubernetesComponent {
	if in == nil {
		return nil
	}
	out := new(KubernetesComponent)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KubernetesCustomResourceImportReference) DeepCopyInto(out *KubernetesCustomResourceImportReference) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KubernetesCustomResourceImportReference.
func (in *KubernetesCustomResourceImportReference) DeepCopy() *KubernetesCustomResourceImportReference {
	if in == nil {
		return nil
	}
	out := new(KubernetesCustomResourceImportReference)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LabeledCommand) DeepCopyInto(out *LabeledCommand) {
	*out = *in
	in.BaseCommand.DeepCopyInto(&out.BaseCommand)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LabeledCommand.
func (in *LabeledCommand) DeepCopy() *LabeledCommand {
	if in == nil {
		return nil
	}
	out := new(LabeledCommand)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OpenshiftComponent) DeepCopyInto(out *OpenshiftComponent) {
	*out = *in
	out.K8sLikeComponent = in.K8sLikeComponent
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OpenshiftComponent.
func (in *OpenshiftComponent) DeepCopy() *OpenshiftComponent {
	if in == nil {
		return nil
	}
	out := new(OpenshiftComponent)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Parent) DeepCopyInto(out *Parent) {
	*out = *in
	in.ImportReference.DeepCopyInto(&out.ImportReference)
	in.DevWorkspaceTemplateSpecContent.DeepCopyInto(&out.DevWorkspaceTemplateSpecContent)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Parent.
func (in *Parent) DeepCopy() *Parent {
	if in == nil {
		return nil
	}
	out := new(Parent)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PluginComponent) DeepCopyInto(out *PluginComponent) {
	*out = *in
	out.BaseComponent = in.BaseComponent
	in.ImportReference.DeepCopyInto(&out.ImportReference)
	in.PluginOverrides.DeepCopyInto(&out.PluginOverrides)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PluginComponent.
func (in *PluginComponent) DeepCopy() *PluginComponent {
	if in == nil {
		return nil
	}
	out := new(PluginComponent)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PluginOverrides) DeepCopyInto(out *PluginOverrides) {
	*out = *in
	if in.Components != nil {
		in, out := &in.Components, &out.Components
		*out = make([]ComponentOverride, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Commands != nil {
		in, out := &in.Commands, &out.Commands
		*out = make([]Command, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PluginOverrides.
func (in *PluginOverrides) DeepCopy() *PluginOverrides {
	if in == nil {
		return nil
	}
	out := new(PluginOverrides)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Project) DeepCopyInto(out *Project) {
	*out = *in
	in.ProjectSource.DeepCopyInto(&out.ProjectSource)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Project.
func (in *Project) DeepCopy() *Project {
	if in == nil {
		return nil
	}
	out := new(Project)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProjectSource) DeepCopyInto(out *ProjectSource) {
	*out = *in
	if in.Git != nil {
		in, out := &in.Git, &out.Git
		*out = new(GitProjectSource)
		**out = **in
	}
	if in.Github != nil {
		in, out := &in.Github, &out.Github
		*out = new(GithubProjectSource)
		**out = **in
	}
	if in.Zip != nil {
		in, out := &in.Zip, &out.Zip
		*out = new(ZipProjectSource)
		**out = **in
	}
	if in.Custom != nil {
		in, out := &in.Custom, &out.Custom
		*out = new(CustomProjectSource)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProjectSource.
func (in *ProjectSource) DeepCopy() *ProjectSource {
	if in == nil {
		return nil
	}
	out := new(ProjectSource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Volume) DeepCopyInto(out *Volume) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Volume.
func (in *Volume) DeepCopy() *Volume {
	if in == nil {
		return nil
	}
	out := new(Volume)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VolumeComponent) DeepCopyInto(out *VolumeComponent) {
	*out = *in
	out.BaseComponent = in.BaseComponent
	out.Volume = in.Volume
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VolumeComponent.
func (in *VolumeComponent) DeepCopy() *VolumeComponent {
	if in == nil {
		return nil
	}
	out := new(VolumeComponent)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VolumeMount) DeepCopyInto(out *VolumeMount) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VolumeMount.
func (in *VolumeMount) DeepCopy() *VolumeMount {
	if in == nil {
		return nil
	}
	out := new(VolumeMount)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VscodeConfigurationCommand) DeepCopyInto(out *VscodeConfigurationCommand) {
	*out = *in
	in.BaseCommand.DeepCopyInto(&out.BaseCommand)
	out.VscodeConfigurationCommandLocation = in.VscodeConfigurationCommandLocation
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VscodeConfigurationCommand.
func (in *VscodeConfigurationCommand) DeepCopy() *VscodeConfigurationCommand {
	if in == nil {
		return nil
	}
	out := new(VscodeConfigurationCommand)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VscodeConfigurationCommandLocation) DeepCopyInto(out *VscodeConfigurationCommandLocation) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VscodeConfigurationCommandLocation.
func (in *VscodeConfigurationCommandLocation) DeepCopy() *VscodeConfigurationCommandLocation {
	if in == nil {
		return nil
	}
	out := new(VscodeConfigurationCommandLocation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkspaceCondition) DeepCopyInto(out *WorkspaceCondition) {
	*out = *in
	in.LastTransitionTime.DeepCopyInto(&out.LastTransitionTime)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkspaceCondition.
func (in *WorkspaceCondition) DeepCopy() *WorkspaceCondition {
	if in == nil {
		return nil
	}
	out := new(WorkspaceCondition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkspaceEvents) DeepCopyInto(out *WorkspaceEvents) {
	*out = *in
	if in.PreStart != nil {
		in, out := &in.PreStart, &out.PreStart
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.PostStart != nil {
		in, out := &in.PostStart, &out.PostStart
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.PreStop != nil {
		in, out := &in.PreStop, &out.PreStop
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.PostStop != nil {
		in, out := &in.PostStop, &out.PostStop
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkspaceEvents.
func (in *WorkspaceEvents) DeepCopy() *WorkspaceEvents {
	if in == nil {
		return nil
	}
	out := new(WorkspaceEvents)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkspacePodContributions) DeepCopyInto(out *WorkspacePodContributions) {
	*out = *in
	if in.Volumes != nil {
		in, out := &in.Volumes, &out.Volumes
		*out = make([]v1.Volume, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.InitContainers != nil {
		in, out := &in.InitContainers, &out.InitContainers
		*out = make([]v1.Container, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Containers != nil {
		in, out := &in.Containers, &out.Containers
		*out = make([]v1.Container, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ImagePullSecrets != nil {
		in, out := &in.ImagePullSecrets, &out.ImagePullSecrets
		*out = make([]v1.LocalObjectReference, len(*in))
		copy(*out, *in)
	}
	if in.CommonEnv != nil {
		in, out := &in.CommonEnv, &out.CommonEnv
		*out = make([]v1.EnvVar, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkspacePodContributions.
func (in *WorkspacePodContributions) DeepCopy() *WorkspacePodContributions {
	if in == nil {
		return nil
	}
	out := new(WorkspacePodContributions)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ZipProjectSource) DeepCopyInto(out *ZipProjectSource) {
	*out = *in
	out.CommonProjectSource = in.CommonProjectSource
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ZipProjectSource.
func (in *ZipProjectSource) DeepCopy() *ZipProjectSource {
	if in == nil {
		return nil
	}
	out := new(ZipProjectSource)
	in.DeepCopyInto(out)
	return out
}
