// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: networking/v1beta1/sidecar.proto

// `Sidecar` describes the configuration of the sidecar proxy that mediates
// inbound and outbound communication to the workload instance it is attached to. By
// default, Istio will program all sidecar proxies in the mesh with the
// necessary configuration required to reach every workload instance in the mesh, as
// well as accept traffic on all the ports associated with the
// workload. The `Sidecar` configuration provides a way to fine tune the set of
// ports, protocols that the proxy will accept when forwarding traffic to
// and from the workload. In addition, it is possible to restrict the set
// of services that the proxy can reach when forwarding outbound traffic
// from workload instances.
//
// Services and configuration in a mesh are organized into one or more
// namespaces (e.g., a Kubernetes namespace or a CF org/space). A `Sidecar`
// configuration in a namespace will apply to one or more workload instances in the same
// namespace, selected using the `workloadSelector` field. In the absence of a
// `workloadSelector`, it will apply to all workload instances in the same
// namespace. When determining the `Sidecar` configuration to be applied to a
// workload instance, preference will be given to the resource with a
// `workloadSelector` that selects this workload instance, over a `Sidecar` configuration
// without any `workloadSelector`.
//
// **NOTE 1**: *_Each namespace can have only one `Sidecar`
// configuration without any `workloadSelector`_ that specifies the
// default for all pods in that namespace*. It is recommended to use
// the name `default` for the namespace-wide sidecar. The behavior of
// the system is undefined if more than one selector-less `Sidecar`
// configurations exist in a given namespace. The behavior of the
// system is undefined if two or more `Sidecar` configurations with a
// `workloadSelector` select the same workload instance.
//
// **NOTE 2**: *_A `Sidecar` configuration in the `MeshConfig`
// [root namespace](https://istio.io/docs/reference/config/istio.mesh.v1alpha1/#MeshConfig)
// will be applied by default to all namespaces without a `Sidecar`
// configuration_*. This global default `Sidecar` configuration should not have
// any `workloadSelector`.
//
// The example below declares a global default `Sidecar` configuration
// in the root namespace called `istio-config`, that configures
// sidecars in all namespaces to allow egress traffic only to other
// workloads in the same namespace as well as to services in the
// `istio-system` namespace.
//
// {{<tabset category-name="example">}}
// {{<tab name="v1alpha3" category-value="v1alpha3">}}
// ```yaml
// apiVersion: networking.istio.io/v1alpha3
// kind: Sidecar
// metadata:
//   name: default
//   namespace: istio-config
// spec:
//   egress:
//   - hosts:
//     - "./*"
//     - "istio-system/*"
// ```
// {{</tab>}}
//
// {{<tab name="v1beta1" category-value="v1beta1">}}
// ```yaml
// apiVersion: networking.istio.io/v1beta1
// kind: Sidecar
// metadata:
//   name: default
//   namespace: istio-config
// spec:
//   egress:
//   - hosts:
//     - "./*"
//     - "istio-system/*"
// ```
// {{</tab>}}
// {{</tabset>}}
//
// The example below declares a `Sidecar` configuration in the
// `prod-us1` namespace that overrides the global default defined
// above, and configures the sidecars in the namespace to allow egress
// traffic to public services in the `prod-us1`, `prod-apis`, and the
// `istio-system` namespaces.
//
// {{<tabset category-name="example">}}
// {{<tab name="v1alpha3" category-value="v1alpha3">}}
// ```yaml
// apiVersion: networking.istio.io/v1alpha3
// kind: Sidecar
// metadata:
//   name: default
//   namespace: prod-us1
// spec:
//   egress:
//   - hosts:
//     - "prod-us1/*"
//     - "prod-apis/*"
//     - "istio-system/*"
// ```
// {{</tab>}}
//
// {{<tab name="v1beta1" category-value="v1beta1">}}
// ```yaml
// apiVersion: networking.istio.io/v1beta1
// kind: Sidecar
// metadata:
//   name: default
//   namespace: prod-us1
// spec:
//   egress:
//   - hosts:
//     - "prod-us1/*"
//     - "prod-apis/*"
//     - "istio-system/*"
// ```
// {{</tab>}}
// {{</tabset>}}
//
// The following example declares a `Sidecar` configuration in the
// `prod-us1` namespace for all pods with labels `app: ratings`
// belonging to the `ratings.prod-us1` service.  The workload accepts
// inbound HTTP traffic on port 9080. The traffic is then forwarded to
// the attached workload instance listening on a Unix domain
// socket. In the egress direction, in addition to the `istio-system`
// namespace, the sidecar proxies only HTTP traffic bound for port
// 9080 for services in the `prod-us1` namespace.
//
// {{<tabset category-name="example">}}
// {{<tab name="v1alpha3" category-value="v1alpha3">}}
// ```yaml
// apiVersion: networking.istio.io/v1alpha3
// kind: Sidecar
// metadata:
//   name: ratings
//   namespace: prod-us1
// spec:
//   workloadSelector:
//     labels:
//       app: ratings
//   ingress:
//   - port:
//       number: 9080
//       protocol: HTTP
//       name: somename
//     defaultEndpoint: unix:///var/run/someuds.sock
//   egress:
//   - port:
//       number: 9080
//       protocol: HTTP
//       name: egresshttp
//     hosts:
//     - "prod-us1/*"
//   - hosts:
//     - "istio-system/*"
// ```
// {{</tab>}}
//
// {{<tab name="v1beta1" category-value="v1beta1">}}
// ```yaml
// apiVersion: networking.istio.io/v1beta1
// kind: Sidecar
// metadata:
//   name: ratings
//   namespace: prod-us1
// spec:
//   workloadSelector:
//     labels:
//       app: ratings
//   ingress:
//   - port:
//       number: 9080
//       protocol: HTTP
//       name: somename
//     defaultEndpoint: unix:///var/run/someuds.sock
//   egress:
//   - port:
//       number: 9080
//       protocol: HTTP
//       name: egresshttp
//     hosts:
//     - "prod-us1/*"
//   - hosts:
//     - "istio-system/*"
// ```
// {{</tab>}}
// {{</tabset>}}
//
// If the workload is deployed without IPTables-based traffic capture,
// the `Sidecar` configuration is the only way to configure the ports
// on the proxy attached to the workload instance. The following
// example declares a `Sidecar` configuration in the `prod-us1`
// namespace for all pods with labels `app: productpage` belonging to
// the `productpage.prod-us1` service. Assuming that these pods are
// deployed without IPtable rules (i.e. the `istio-init` container)
// and the proxy metadata `ISTIO_META_INTERCEPTION_MODE` is set to
// `NONE`, the specification, below, allows such pods to receive HTTP
// traffic on port 9080 (wrapped inside Istio mutual TLS) and forward
// it to the application listening on `127.0.0.1:8080`. It also allows
// the application to communicate with a backing MySQL database on
// `127.0.0.1:3306`, that then gets proxied to the externally hosted
// MySQL service at `mysql.foo.com:3306`.
//
// {{<tabset category-name="example">}}
// {{<tab name="v1alpha3" category-value="v1alpha3">}}
// ```yaml
// apiVersion: networking.istio.io/v1alpha3
// kind: Sidecar
// metadata:
//   name: no-ip-tables
//   namespace: prod-us1
// spec:
//   workloadSelector:
//     labels:
//       app: productpage
//   ingress:
//   - port:
//       number: 9080 # binds to proxy_instance_ip:9080 (0.0.0.0:9080, if no unicast IP is available for the instance)
//       protocol: HTTP
//       name: somename
//     defaultEndpoint: 127.0.0.1:8080
//     captureMode: NONE # not needed if metadata is set for entire proxy
//   egress:
//   - port:
//       number: 3306
//       protocol: MYSQL
//       name: egressmysql
//     captureMode: NONE # not needed if metadata is set for entire proxy
//     bind: 127.0.0.1
//     hosts:
//     - "*/mysql.foo.com"
// ```
// {{</tab>}}
//
// {{<tab name="v1beta1" category-value="v1beta1">}}
// ```yaml
// apiVersion: networking.istio.io/v1beta1
// kind: Sidecar
// metadata:
//   name: no-ip-tables
//   namespace: prod-us1
// spec:
//   workloadSelector:
//     labels:
//       app: productpage
//   ingress:
//   - port:
//       number: 9080 # binds to proxy_instance_ip:9080 (0.0.0.0:9080, if no unicast IP is available for the instance)
//       protocol: HTTP
//       name: somename
//     defaultEndpoint: 127.0.0.1:8080
//     captureMode: NONE # not needed if metadata is set for entire proxy
//   egress:
//   - port:
//       number: 3306
//       protocol: MYSQL
//       name: egressmysql
//     captureMode: NONE # not needed if metadata is set for entire proxy
//     bind: 127.0.0.1
//     hosts:
//     - "*/mysql.foo.com"
// ```
// {{</tab>}}
// {{</tabset>}}
//
// And the associated service entry for routing to `mysql.foo.com:3306`
//
// {{<tabset category-name="example">}}
// {{<tab name="v1alpha3" category-value="v1alpha3">}}
// ```yaml
// apiVersion: networking.istio.io/v1alpha3
// kind: ServiceEntry
// metadata:
//   name: external-svc-mysql
//   namespace: ns1
// spec:
//   hosts:
//   - mysql.foo.com
//   ports:
//   - number: 3306
//     name: mysql
//     protocol: MYSQL
//   location: MESH_EXTERNAL
//   resolution: DNS
// ```
// {{</tab>}}
//
// {{<tab name="v1beta1" category-value="v1beta1">}}
// ```yaml
// apiVersion: networking.istio.io/v1beta1
// kind: ServiceEntry
// metadata:
//   name: external-svc-mysql
//   namespace: ns1
// spec:
//   hosts:
//   - mysql.foo.com
//   ports:
//   - number: 3306
//     name: mysql
//     protocol: MYSQL
//   location: MESH_EXTERNAL
//   resolution: DNS
// ```
// {{</tab>}}
// {{</tabset>}}
//
// It is also possible to mix and match traffic capture modes in a single
// proxy. For example, consider a setup where internal services are on the
// `192.168.0.0/16` subnet. So, IP tables are setup on the VM to capture all
// outbound traffic on `192.168.0.0/16` subnet. Assume that the VM has an
// additional network interface on `172.16.0.0/16` subnet for inbound
// traffic. The following `Sidecar` configuration allows the VM to expose a
// listener on `172.16.1.32:80` (the VM's IP) for traffic arriving from the
// `172.16.0.0/16` subnet.
//
// **NOTE**: The `ISTIO_META_INTERCEPTION_MODE` metadata on the
// proxy in the VM should contain `REDIRECT` or `TPROXY` as its value,
// implying that IP tables based traffic capture is active.
//
// {{<tabset category-name="example">}}
// {{<tab name="v1alpha3" category-value="v1alpha3">}}
// ```yaml
// apiVersion: networking.istio.io/v1alpha3
// kind: Sidecar
// metadata:
//   name: partial-ip-tables
//   namespace: prod-us1
// spec:
//   workloadSelector:
//     labels:
//       app: productpage
//   ingress:
//   - bind: 172.16.1.32
//     port:
//       number: 80 # binds to 172.16.1.32:80
//       protocol: HTTP
//       name: somename
//     defaultEndpoint: 127.0.0.1:8080
//     captureMode: NONE
//   egress:
//     # use the system detected defaults
//     # sets up configuration to handle outbound traffic to services
//     # in 192.168.0.0/16 subnet, based on information provided by the
//     # service registry
//   - captureMode: IPTABLES
//     hosts:
//     - "*/*"
// ```
// {{</tab>}}
//
// {{<tab name="v1beta1" category-value="v1beta1">}}
// ```yaml
// apiVersion: networking.istio.io/v1beta1
// kind: Sidecar
// metadata:
//   name: partial-ip-tables
//   namespace: prod-us1
// spec:
//   workloadSelector:
//     labels:
//       app: productpage
//   ingress:
//   - bind: 172.16.1.32
//     port:
//       number: 80 # binds to 172.16.1.32:80
//       protocol: HTTP
//       name: somename
//     defaultEndpoint: 127.0.0.1:8080
//     captureMode: NONE
//   egress:
//     # use the system detected defaults
//     # sets up configuration to handle outbound traffic to services
//     # in 192.168.0.0/16 subnet, based on information provided by the
//     # service registry
//   - captureMode: IPTABLES
//     hosts:
//     - "*/*"
// ```
// {{</tab>}}
// {{</tabset>}}
//

package v1beta1

import (
	bytes "bytes"
	fmt "fmt"
	github_com_gogo_protobuf_jsonpb "github.com/gogo/protobuf/jsonpb"
	proto "github.com/gogo/protobuf/proto"
	_ "istio.io/gogo-genproto/googleapis/google/api"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// MarshalJSON is a custom marshaler for Sidecar
func (this *Sidecar) MarshalJSON() ([]byte, error) {
	str, err := SidecarMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for Sidecar
func (this *Sidecar) UnmarshalJSON(b []byte) error {
	return SidecarUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for IstioIngressListener
func (this *IstioIngressListener) MarshalJSON() ([]byte, error) {
	str, err := SidecarMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for IstioIngressListener
func (this *IstioIngressListener) UnmarshalJSON(b []byte) error {
	return SidecarUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for IstioEgressListener
func (this *IstioEgressListener) MarshalJSON() ([]byte, error) {
	str, err := SidecarMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for IstioEgressListener
func (this *IstioEgressListener) UnmarshalJSON(b []byte) error {
	return SidecarUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for WorkloadSelector
func (this *WorkloadSelector) MarshalJSON() ([]byte, error) {
	str, err := SidecarMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for WorkloadSelector
func (this *WorkloadSelector) UnmarshalJSON(b []byte) error {
	return SidecarUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for OutboundTrafficPolicy
func (this *OutboundTrafficPolicy) MarshalJSON() ([]byte, error) {
	str, err := SidecarMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for OutboundTrafficPolicy
func (this *OutboundTrafficPolicy) UnmarshalJSON(b []byte) error {
	return SidecarUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

var (
	SidecarMarshaler   = &github_com_gogo_protobuf_jsonpb.Marshaler{}
	SidecarUnmarshaler = &github_com_gogo_protobuf_jsonpb.Unmarshaler{}
)
