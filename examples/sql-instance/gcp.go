// Example: Define a Cloud SQL PostgreSQL instance with database and user
package gcp

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	sqlv1beta1 "github.com/lex00/wetwire-gcp-go/resources/cnrm/sql/v1beta1"
)

// PostgresInstance is a Cloud SQL PostgreSQL instance
var PostgresInstance = sqlv1beta1.SQLInstance{
	ObjectMeta: metav1.ObjectMeta{
		Name:      "my-postgres",
		Namespace: "default",
	},
	Spec: sqlv1beta1.SQLInstanceSpec{
		DatabaseVersion: "POSTGRES_15",
		Region:          "us-central1",
		Settings: map[string]interface{}{
			"tier":             "db-f1-micro",
			"availabilityType": "ZONAL",
			"diskSize":         10,
			"diskType":         "PD_SSD",
			"backupConfiguration": map[string]interface{}{
				"enabled":                    true,
				"startTime":                  "03:00",
				"pointInTimeRecoveryEnabled": true,
			},
			"ipConfiguration": map[string]interface{}{
				"ipv4Enabled": true,
				"authorizedNetworks": []map[string]interface{}{
					{
						"name":  "office",
						"value": "203.0.113.0/24",
					},
				},
			},
			"maintenanceWindow": map[string]interface{}{
				"day":  7,
				"hour": 3,
			},
		},
	},
}

// AppDatabase is the application database
var AppDatabase = sqlv1beta1.SQLDatabase{
	ObjectMeta: metav1.ObjectMeta{
		Name:      "app-db",
		Namespace: "default",
	},
	Spec: sqlv1beta1.SQLDatabaseSpec{
		InstanceRef: map[string]interface{}{
			"name": PostgresInstance.ObjectMeta.Name,
		},
		Charset:   "UTF8",
		Collation: "en_US.UTF8",
	},
}

// AppUser is the application database user
var AppUser = sqlv1beta1.SQLUser{
	ObjectMeta: metav1.ObjectMeta{
		Name:      "app-user",
		Namespace: "default",
	},
	Spec: sqlv1beta1.SQLUserSpec{
		InstanceRef: map[string]interface{}{
			"name": PostgresInstance.ObjectMeta.Name,
		},
		Password: map[string]interface{}{
			"valueFrom": map[string]interface{}{
				"secretKeyRef": map[string]interface{}{
					"name": "db-credentials",
					"key":  "password",
				},
			},
		},
	},
}
