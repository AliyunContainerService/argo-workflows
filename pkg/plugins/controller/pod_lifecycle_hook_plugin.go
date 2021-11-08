package controller

import (
	corev1 "k8s.io/api/core/v1"

	wfv1 "github.com/argoproj/argo-workflows/v3/pkg/apis/workflow/v1alpha1"
)

type PodPreCreateArgs struct {
	Workflow *wfv1.Workflow `json:"workflow"`
	Template *wfv1.Template `json:"template"`
	Pod      *corev1.Pod    `json:"pod"`
}

type PodPreCreateReply struct {
	Pod *corev1.Pod `json:"pod,omitempty"`
}

type PodPostCreateArgs struct {
	Workflow *wfv1.Workflow `json:"workflow"`
	Template *wfv1.Template `json:"template"`
	Pod      *corev1.Pod    `json:"pod"`
}

type PodPostCreateReply struct {
}

type PodLifecycleHook interface {
	// PodPreCreate is called before creating a pod.
	PodPreCreate(args PodPreCreateArgs, reply *PodPreCreateReply) error
	// PodPostCreate is called after creating a pod.
	PodPostCreate(args PodPostCreateArgs, reply *PodPostCreateReply) error
}