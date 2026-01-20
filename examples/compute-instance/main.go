// Example: Define a GCP Compute Instance using wetwire-gcp
package gcp

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	computev1beta1 "github.com/lex00/wetwire-gcp-go/resources/cnrm/compute/v1beta1"
)

// MyInstance defines a Compute Engine VM instance
var MyInstance = computev1beta1.ComputeInstance{
	ObjectMeta: metav1.ObjectMeta{
		Name:      "my-instance",
		Namespace: "default",
	},
	Spec: computev1beta1.ComputeInstanceSpec{
		MachineType: "n1-standard-1",
		Zone:        "us-central1-a",
		BootDisk: map[string]interface{}{
			"initializeParams": map[string]interface{}{
				"image": "debian-cloud/debian-11",
			},
		},
		NetworkInterface: []map[string]interface{}{
			{
				"network": "default",
				"accessConfig": []map[string]interface{}{
					{
						"networkTier": "PREMIUM",
					},
				},
			},
		},
		Tags: []string{"http-server", "https-server"},
	},
}
