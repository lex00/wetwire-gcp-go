// Test fixture: Simple resource declarations
package testdata

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	computev1beta1 "github.com/lex00/wetwire-gcp-go/resources/cnrm/compute/v1beta1"
)

var SimpleInstance = computev1beta1.ComputeInstance{
	ObjectMeta: metav1.ObjectMeta{
		Name: "simple-instance",
	},
	Spec: computev1beta1.ComputeInstanceSpec{
		MachineType: "n1-standard-1",
		Zone:        "us-central1-a",
	},
}
