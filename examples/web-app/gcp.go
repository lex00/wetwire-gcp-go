// Example: Define a web application with compute instance, firewall, and external IP
package gcp

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	computev1beta1 "github.com/lex00/wetwire-gcp-go/resources/cnrm/compute/v1beta1"
)

// WebServerIP is the external IP for the web server
var WebServerIP = computev1beta1.ComputeAddress{
	ObjectMeta: metav1.ObjectMeta{
		Name:      "web-server-ip",
		Namespace: "default",
	},
	Spec: computev1beta1.ComputeAddressSpec{
		Location:    "us-central1",
		AddressType: "EXTERNAL",
		NetworkTier: "PREMIUM",
		Description: "External IP for web server",
	},
}

// AllowHTTP allows HTTP traffic to instances with web-server tag
var AllowHTTP = computev1beta1.ComputeFirewall{
	ObjectMeta: metav1.ObjectMeta{
		Name:      "allow-http",
		Namespace: "default",
	},
	Spec: computev1beta1.ComputeFirewallSpec{
		NetworkRef: map[string]interface{}{
			"name": "default",
		},
		Direction:    "INGRESS",
		Priority:     1000,
		SourceRanges: []string{"0.0.0.0/0"},
		TargetTags:   []string{"web-server"},
		Allow: []map[string]interface{}{
			{
				"protocol": "tcp",
				"ports":    []string{"80"},
			},
		},
		Description: "Allow HTTP traffic",
	},
}

// AllowHTTPS allows HTTPS traffic to instances with web-server tag
var AllowHTTPS = computev1beta1.ComputeFirewall{
	ObjectMeta: metav1.ObjectMeta{
		Name:      "allow-https",
		Namespace: "default",
	},
	Spec: computev1beta1.ComputeFirewallSpec{
		NetworkRef: map[string]interface{}{
			"name": "default",
		},
		Direction:    "INGRESS",
		Priority:     1000,
		SourceRanges: []string{"0.0.0.0/0"},
		TargetTags:   []string{"web-server"},
		Allow: []map[string]interface{}{
			{
				"protocol": "tcp",
				"ports":    []string{"443"},
			},
		},
		Description: "Allow HTTPS traffic",
	},
}

// AllowSSH allows SSH access from specific IP range
var AllowSSH = computev1beta1.ComputeFirewall{
	ObjectMeta: metav1.ObjectMeta{
		Name:      "allow-ssh-admin",
		Namespace: "default",
	},
	Spec: computev1beta1.ComputeFirewallSpec{
		NetworkRef: map[string]interface{}{
			"name": "default",
		},
		Direction:    "INGRESS",
		Priority:     1000,
		SourceRanges: []string{"203.0.113.0/24"}, // Admin office IP range
		TargetTags:   []string{"web-server"},
		Allow: []map[string]interface{}{
			{
				"protocol": "tcp",
				"ports":    []string{"22"},
			},
		},
		Description: "Allow SSH from admin network",
	},
}

// WebServer is the compute instance running the web application
var WebServer = computev1beta1.ComputeInstance{
	ObjectMeta: metav1.ObjectMeta{
		Name:      "web-server",
		Namespace: "default",
	},
	Spec: computev1beta1.ComputeInstanceSpec{
		MachineType: "e2-medium",
		Zone:        "us-central1-a",
		BootDisk: map[string]interface{}{
			"initializeParams": map[string]interface{}{
				"image": "debian-cloud/debian-11",
				"size":  20,
				"type":  "pd-ssd",
			},
		},
		NetworkInterface: []map[string]interface{}{
			{
				"network": "default",
				"accessConfig": []map[string]interface{}{
					{
						"natIpRef": map[string]interface{}{
							"name": WebServerIP.ObjectMeta.Name,
						},
						"networkTier": "PREMIUM",
					},
				},
			},
		},
		Tags: []string{"web-server"},
		MetadataStartupScript: `#!/bin/bash
apt-get update
apt-get install -y nginx
systemctl start nginx
`,
		ServiceAccount: map[string]interface{}{
			"scopes": []string{
				"https://www.googleapis.com/auth/logging.write",
				"https://www.googleapis.com/auth/monitoring.write",
			},
		},
	},
}
