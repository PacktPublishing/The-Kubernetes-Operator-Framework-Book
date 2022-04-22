package v1alpha1

import (
	"github.com/sample/nginx-operator/api/v1alpha2"
	v1 "k8s.io/api/core/v1"
	"k8s.io/utils/pointer"
	"sigs.k8s.io/controller-runtime/pkg/conversion"
)

// ConvertTo converts v1alpha1 to v1alpha2
func (src *NginxOperator) ConvertTo(dst conversion.Hub) error {
	objV1alpha2 := dst.(*v1alpha2.NginxOperator)
	objV1alpha2.ObjectMeta = src.ObjectMeta
	objV1alpha2.Status.Conditions = src.Status.Conditions

	if src.Spec.Replicas != nil {
		objV1alpha2.Spec.Replicas = src.Spec.Replicas
	}
	if len(src.Spec.ForceRedploy) > 0 {
		objV1alpha2.Spec.ForceRedploy = src.Spec.ForceRedploy
	}
	if src.Spec.Port != nil {
		objV1alpha2.Spec.Ports = make([]v1.ContainerPort, 0, 1)
		objV1alpha2.Spec.Ports = append(objV1alpha2.Spec.Ports, v1.ContainerPort{ContainerPort: *src.Spec.Port})
	}

	return nil
}

// ConvertFrom converts v1alpha2 to v1alpha1
func (dst *NginxOperator) ConvertFrom(src conversion.Hub) error {
	objV1alpha2 := src.(*v1alpha2.NginxOperator)
	dst.ObjectMeta = objV1alpha2.ObjectMeta
	dst.Status.Conditions = objV1alpha2.Status.Conditions

	if objV1alpha2.Spec.Replicas != nil {
		dst.Spec.Replicas = objV1alpha2.Spec.Replicas
	}
	if len(objV1alpha2.Spec.ForceRedploy) > 0 {
		dst.Spec.ForceRedploy = objV1alpha2.Spec.ForceRedploy
	}

	if len(objV1alpha2.Spec.Ports) > 0 {
		dst.Spec.Port = pointer.Int32(objV1alpha2.Spec.Ports[0].ContainerPort)
	}

	return nil
}
