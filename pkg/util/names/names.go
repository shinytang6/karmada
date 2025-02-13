package names

import (
	"fmt"
	"strings"
)

const (
	// KubernetesReservedNSPrefix is the prefix of namespace which reserved by Kubernetes system, such as:
	// - kube-system
	// - kube-public
	// - kube-node-lease
	KubernetesReservedNSPrefix = "kube-"

	// KarmadaReservedNSPrefix is the prefix of namespace which reserved by Karmada system, such as:
	// - karmada-system
	// - karmada-cluster
	// - karmada-es-*
	KarmadaReservedNSPrefix = "karmada-"
)

// executionSpacePrefix is the prefix of execution space
const executionSpacePrefix = "karmada-es-"

// GenerateExecutionSpaceName generates execution space name for the given member cluster
func GenerateExecutionSpaceName(clusterName string) (string, error) {
	if clusterName == "" {
		return "", fmt.Errorf("the member cluster name is empty")
	}
	executionSpace := executionSpacePrefix + clusterName
	return executionSpace, nil
}

// GetClusterName returns member cluster name for the given execution space
func GetClusterName(executionSpaceName string) (string, error) {
	if !strings.HasPrefix(executionSpaceName, executionSpacePrefix) {
		return "", fmt.Errorf("the execution space name is in wrong format")
	}
	return strings.TrimPrefix(executionSpaceName, executionSpacePrefix), nil
}

// GenerateBindingName will generate binding name by namespace, kind and name
func GenerateBindingName(kind, name string) string {
	return strings.ToLower(name + "-" + kind)
}

// GenerateWorkName will generate work name by namespace, kind and name
func GenerateWorkName(kind, name, namespace string) string {
	if len(namespace) == 0 {
		return strings.ToLower(name + "-" + kind)
	}
	return strings.ToLower(namespace + "-" + name + "-" + kind)
}

// GenerateServiceAccountName generates the name of a ServiceAccount.
func GenerateServiceAccountName(clusterName string) string {
	return fmt.Sprintf("%s-%s", "karmada", clusterName)
}

// GenerateRoleName generates the name of a Role or ClusterRole.
func GenerateRoleName(serviceAccountName string) string {
	return fmt.Sprintf("karmada-controller-manager:%s", serviceAccountName)
}
