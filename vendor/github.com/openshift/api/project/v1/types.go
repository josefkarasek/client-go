package v1

import (
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ProjectList is a list of Project objects.
//
// Compatibility level 1: Stable within a major release for a minimum of 12 months or 3 minor releases (whichever is longer).
// +openshift:compatibility-gen:level=1
type ProjectList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`

	// Items is the list of projects
	Items []Project `json:"items" protobuf:"bytes,2,rep,name=items"`
}

const (
	// These are internal finalizer values to Origin
	FinalizerOrigin corev1.FinalizerName = "openshift.io/origin"
	// ProjectNodeSelector is an annotation that holds the node selector;
	// the node selector annotation determines which nodes will have pods from this project scheduled to them
	ProjectNodeSelector = "openshift.io/node-selector"

	// ProjectRequesterAnnotation is the username that requested a given project.  Its not guaranteed to be present,
	// but it is set by the default project template.
	ProjectRequesterAnnotation = "openshift.io/requester"
)

// ProjectSpec describes the attributes on a Project
type ProjectSpec struct {
	// Finalizers is an opaque list of values that must be empty to permanently remove object from storage
	Finalizers []corev1.FinalizerName `json:"finalizers,omitempty" protobuf:"bytes,1,rep,name=finalizers,casttype=k8s.io/api/core/v1.FinalizerName"`
}

// ProjectStatus is information about the current status of a Project
type ProjectStatus struct {
	// Phase is the current lifecycle phase of the project
	// +optional
	Phase corev1.NamespacePhase `json:"phase,omitempty" protobuf:"bytes,1,opt,name=phase,casttype=k8s.io/api/core/v1.NamespacePhase"`

	// Represents the latest available observations of the project current state.
	// +optional
	// +patchMergeKey=type
	// +patchStrategy=merge
	Conditions []corev1.NamespaceCondition `json:"conditions,omitempty" patchStrategy:"merge" patchMergeKey:"type" protobuf:"bytes,2,rep,name=conditions"`
}

// +genclient
// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Projects are the unit of isolation and collaboration in OpenShift. A project has one or more members,
// a quota on the resources that the project may consume, and the security controls on the resources in
// the project. Within a project, members may have different roles - project administrators can set
// membership, editors can create and manage the resources, and viewers can see but not access running
// containers. In a normal cluster project administrators are not able to alter their quotas - that is
// restricted to cluster administrators.
//
// Listing or watching projects will return only projects the user has the reader role on.
//
// An OpenShift project is an alternative representation of a Kubernetes namespace. Projects are exposed
// as editable to end users while namespaces are not. Direct creation of a project is typically restricted
// to administrators, while end users should use the requestproject resource.
//
// Compatibility level 1: Stable within a major release for a minimum of 12 months or 3 minor releases (whichever is longer).
// +openshift:compatibility-gen:level=1
type Project struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`

	// Spec defines the behavior of the Namespace.
	Spec ProjectSpec `json:"spec,omitempty" protobuf:"bytes,2,opt,name=spec"`

	// Status describes the current status of a Namespace
	// +optional
	Status ProjectStatus `json:"status,omitempty" protobuf:"bytes,3,opt,name=status"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +genclient
// +genclient:nonNamespaced
// +genclient:skipVerbs=get,list,create,update,patch,delete,deleteCollection,watch
// +genclient:method=Create,verb=create,result=Project

// ProjectRequest is the set of options necessary to fully qualify a project request
//
// Compatibility level 1: Stable within a major release for a minimum of 12 months or 3 minor releases (whichever is longer).
// +openshift:compatibility-gen:level=1
type ProjectRequest struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`

	// DisplayName is the display name to apply to a project
	DisplayName string `json:"displayName,omitempty" protobuf:"bytes,2,opt,name=displayName"`
	// Description is the description to apply to a project
	Description string `json:"description,omitempty" protobuf:"bytes,3,opt,name=description"`
}

// +genclient
// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ProjectRequestLimit is the configuration for the project request limit plug-in
// It contains an ordered list of limits based on user label selectors. Selectors will
// be checked in order and the first one that applies will be used as the limit.
//
// Compatibility level 1: Stable within a major release for a minimum of 12 months or 3 minor releases (whichever is longer).
// +openshift:compatibility-gen:level=1
type ProjectRequestLimit struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`

	Limits []ProjectLimitBySelector `json:"limits,omitempty" protobuf:"bytes,2,opt,name=limits"`

	// MaxProjectsForSystemUsers controls how many projects a certificate user may have.  Certificate
	// users do not have any labels associated with them for more fine grained control
	MaxProjectsForSystemUsers *int `json:"maxProjectsForSystemUsers,omitempty" protobuf:"bytes,3,opt,name=maxProjectsForSystemUsers"`

	// MaxProjectsForServiceAccounts controls how many projects a service account may have.  Service
	// accounts can't create projects by default, but if they are allowed to create projects, you cannot
	// trust any labels placed on them since project editors can manipulate those labels
	MaxProjectsForServiceAccounts *int `json:"maxProjectsForServiceAccounts,omitempty" protobuf:"bytes,4,opt,name=maxProjectsForServiceAccounts"`
}

// ProjectLimitBySelector specifies the maximum number of projects allowed for a given user label selector
type ProjectLimitBySelector struct {
	// Selector is a user label selector. An empty selector selects everything.
	Selector map[string]string `json:"selector,omitempty" protobuf:"bytes,1,opt,name=selector"`

	// MaxProjects is the number of projects allowed for this class of users. If MaxProjects is nil,
	// there is no limit to the number of projects users can request. An unlimited number of projects
	// is useful in the case a limit is specified as the default for all users and only users with a
	// specific set of labels should be allowed unlimited project creation.
	MaxProjects *int `json:"maxProjects,omitempty" protobuf:"bytes,2,opt,name=maxProjects"`
}
