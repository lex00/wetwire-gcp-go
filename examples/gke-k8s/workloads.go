package gke_k8s

// Workloads documentation for GKE K8s-native example.
//
// Sample workloads would be defined using wetwire-k8s-go for standard
// Kubernetes resources (Deployments, Services, etc.). These are cloud-agnostic
// and can be deployed to any Kubernetes cluster.
//
// Example usage with wetwire-k8s-go:
//
//	import (
//		appsv1 "github.com/lex00/wetwire-k8s-go/resources/apps/v1"
//		corev1 "github.com/lex00/wetwire-k8s-go/resources/core/v1"
//		autoscalingv2 "github.com/lex00/wetwire-k8s-go/resources/autoscaling/v2"
//	)
//
//	var SampleDeployment = appsv1.Deployment{
//		ObjectMeta: metav1.ObjectMeta{
//			Name:      "sample-app",
//			Namespace: "default",
//		},
//		Spec: appsv1.DeploymentSpec{
//			Replicas: 3,
//			Selector: metav1.LabelSelector{
//				MatchLabels: map[string]string{"app": "sample"},
//			},
//			Template: corev1.PodTemplateSpec{
//				Spec: corev1.PodSpec{
//					Containers: []corev1.Container{
//						{
//							Name:  "app",
//							Image: "nginx:1.25",
//							Ports: []corev1.ContainerPort{
//								{ContainerPort: 80},
//							},
//						},
//					},
//				},
//			},
//		},
//	}
//
// GKE-specific workload considerations:
//
// 1. Workload Identity: Use GKE Metadata Server for pod authentication
//
//	spec:
//	  serviceAccountName: my-ksa
//	  # Workload Identity annotation on KSA maps to GCP SA
//
// 2. Node Selection: Target specific node pools
//
//	spec:
//	  nodeSelector:
//	    pool-type: application
//	  tolerations:
//	    - key: cloud.google.com/gke-spot
//	      operator: Equal
//	      value: "true"
//	      effect: NoSchedule
//
// 3. Pod Disruption Budgets: For high availability
//
//	var SamplePDB = policyv1.PodDisruptionBudget{
//	  Spec: policyv1.PodDisruptionBudgetSpec{
//	    MinAvailable: intstr.FromInt(2),
//	    Selector: metav1.LabelSelector{
//	      MatchLabels: map[string]string{"app": "sample"},
//	    },
//	  },
//	}
//
// 4. Ingress: Use GKE Ingress with Cloud Load Balancing
//
//	var SampleIngress = networkingv1.Ingress{
//	  ObjectMeta: metav1.ObjectMeta{
//	    Annotations: map[string]string{
//	      "kubernetes.io/ingress.class": "gce",
//	    },
//	  },
//	}
