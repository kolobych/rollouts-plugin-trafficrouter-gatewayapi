package mocks

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"sigs.k8s.io/gateway-api/apis/v1alpha2"

	"sigs.k8s.io/gateway-api/apis/v1beta1"
)

const (
	HTTPRoute         = "HTTPRoute"
	TCPRoute          = "TCPRoute"
	StableServiceName = "argo-rollouts-stable-service"
	CanaryServiceName = "argo-rollouts-canary-service"
	HTTPRouteName     = "argo-rollouts-http-route"
	TCPRouteName      = "argo-rollouts-tcp-route"
	Namespace         = "default"
)

var (
	port         = v1beta1.PortNumber(80)
	weight int32 = 0
)

var HTTPRouteObj = v1beta1.HTTPRoute{
	ObjectMeta: metav1.ObjectMeta{
		Name:      HTTPRouteName,
		Namespace: Namespace,
	},
	Spec: v1beta1.HTTPRouteSpec{
		CommonRouteSpec: v1beta1.CommonRouteSpec{
			ParentRefs: []v1beta1.ParentReference{
				{
					Name: "argo-rollouts-gateway",
				},
			},
		},
		Rules: []v1beta1.HTTPRouteRule{
			{
				BackendRefs: []v1beta1.HTTPBackendRef{
					{
						BackendRef: v1beta1.BackendRef{
							BackendObjectReference: v1beta1.BackendObjectReference{
								Name: StableServiceName,
								Port: &port,
							},
							Weight: &weight,
						},
					},
					{
						BackendRef: v1beta1.BackendRef{
							BackendObjectReference: v1beta1.BackendObjectReference{
								Name: CanaryServiceName,
								Port: &port,
							},
							Weight: &weight,
						},
					},
				},
			},
		},
	},
}

var TCPPRouteObj = v1alpha2.TCPRoute{
	ObjectMeta: metav1.ObjectMeta{
		Name:      TCPRouteName,
		Namespace: Namespace,
	},
	Spec: v1alpha2.TCPRouteSpec{
		CommonRouteSpec: v1alpha2.CommonRouteSpec{
			ParentRefs: []v1beta1.ParentReference{
				{
					Name: "argo-rollouts-gateway",
				},
			},
		},
		Rules: []v1alpha2.TCPRouteRule{
			{
				BackendRefs: []v1alpha2.BackendRef{
					{
						BackendObjectReference: v1alpha2.BackendObjectReference{
							Name: StableServiceName,
							Port: &port,
						},
						Weight: &weight,
					},
					{
						BackendObjectReference: v1alpha2.BackendObjectReference{
							Name: CanaryServiceName,
							Port: &port,
						},
						Weight: &weight,
					},
				},
			},
		},
	},
}
