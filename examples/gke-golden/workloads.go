package gke_golden

// This file documents the sample Kubernetes workloads that should be deployed
// to the GKE cluster. These are cloud-agnostic and can be managed with wetwire-k8s-go.

/*
Sample Deployment (nginx):

apiVersion: apps/v1
kind: Deployment
metadata:
  name: sample-app
  namespace: default
  labels:
    app: sample-app
spec:
  replicas: 3
  selector:
    matchLabels:
      app: sample-app
  template:
    metadata:
      labels:
        app: sample-app
    spec:
      # Prefer application node pool
      nodeSelector:
        pool-type: application
      # Tolerate spot nodes for cost savings
      tolerations:
        - key: "cloud.google.com/gke-spot"
          operator: "Equal"
          value: "true"
          effect: "NoSchedule"
      containers:
        - name: nginx
          image: nginx:1.25
          ports:
            - containerPort: 80
          resources:
            requests:
              cpu: "100m"
              memory: "128Mi"
            limits:
              cpu: "500m"
              memory: "256Mi"
          readinessProbe:
            httpGet:
              path: /
              port: 80
            initialDelaySeconds: 5
            periodSeconds: 10
          livenessProbe:
            httpGet:
              path: /
              port: 80
            initialDelaySeconds: 15
            periodSeconds: 20

---
apiVersion: v1
kind: Service
metadata:
  name: sample-app
  namespace: default
spec:
  type: ClusterIP
  selector:
    app: sample-app
  ports:
    - port: 80
      targetPort: 80

---
apiVersion: autoscaling/v2
kind: HorizontalPodAutoscaler
metadata:
  name: sample-app
  namespace: default
spec:
  scaleTargetRef:
    apiVersion: apps/v1
    kind: Deployment
    name: sample-app
  minReplicas: 3
  maxReplicas: 10
  metrics:
    - type: Resource
      resource:
        name: cpu
        target:
          type: Utilization
          averageUtilization: 70
    - type: Resource
      resource:
        name: memory
        target:
          type: Utilization
          averageUtilization: 80

---
# GKE Ingress with Google Cloud Load Balancer
apiVersion: networking.k8s.io/v1
kind: Ingress
metadata:
  name: sample-app
  namespace: default
  annotations:
    kubernetes.io/ingress.class: "gce"
    kubernetes.io/ingress.global-static-ip-name: "sample-app-ip"
    networking.gke.io/managed-certificates: "sample-app-cert"
    kubernetes.io/ingress.allow-http: "false"
spec:
  defaultBackend:
    service:
      name: sample-app
      port:
        number: 80
  rules:
    - host: sample.example.com
      http:
        paths:
          - path: /
            pathType: Prefix
            backend:
              service:
                name: sample-app
                port:
                  number: 80

---
# BackendConfig for health checks and connection draining
apiVersion: cloud.google.com/v1
kind: BackendConfig
metadata:
  name: sample-app-backend-config
  namespace: default
spec:
  healthCheck:
    type: HTTP
    requestPath: /
    port: 80
  connectionDraining:
    drainingTimeoutSec: 60
  logging:
    enable: true
    sampleRate: 0.5

---
# Managed Certificate for HTTPS
apiVersion: networking.gke.io/v1
kind: ManagedCertificate
metadata:
  name: sample-app-cert
  namespace: default
spec:
  domains:
    - sample.example.com
*/
