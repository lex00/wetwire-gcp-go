package registry

import (
	"testing"
)

func TestNewRegistry(t *testing.T) {
	r := NewRegistry()
	if r == nil {
		t.Fatal("NewRegistry returned nil")
	}
	if r.types == nil {
		t.Error("types map not initialized")
	}
	if r.packages == nil {
		t.Error("packages map not initialized")
	}
	if r.kinds == nil {
		t.Error("kinds map not initialized")
	}
}

func TestRegistry_Register(t *testing.T) {
	r := NewRegistry()

	info := TypeInfo{
		Package:    "computev1beta1",
		Group:      "compute.cnrm.cloud.google.com",
		Version:    "v1beta1",
		Kind:       "ComputeInstance",
		APIVersion: "compute.cnrm.cloud.google.com/v1beta1",
	}

	r.Register(info)

	// Verify type was registered
	if !r.IsKnownType("computev1beta1.ComputeInstance") {
		t.Error("type not registered correctly")
	}

	// Verify package was registered
	if !r.IsKnownPackage("computev1beta1") {
		t.Error("package not registered correctly")
	}

	// Verify kind was registered
	if !r.IsKnownType("ComputeInstance") {
		t.Error("kind not registered correctly")
	}
}

func TestRegistry_IsKnownType(t *testing.T) {
	r := NewRegistry()
	r.Register(TypeInfo{
		Package:    "storagev1beta1",
		Group:      "storage.cnrm.cloud.google.com",
		Version:    "v1beta1",
		Kind:       "StorageBucket",
		APIVersion: "storage.cnrm.cloud.google.com/v1beta1",
	})

	tests := []struct {
		name     string
		typeName string
		want     bool
	}{
		{"full type name", "storagev1beta1.StorageBucket", true},
		{"kind only", "StorageBucket", true},
		{"unknown type", "unknownv1.Unknown", false},
		{"unknown kind", "NotAResource", false},
		{"partial match package", "storagev1beta1.NotAType", true}, // package is known
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := r.IsKnownType(tt.typeName); got != tt.want {
				t.Errorf("IsKnownType(%q) = %v, want %v", tt.typeName, got, tt.want)
			}
		})
	}
}

func TestRegistry_IsKnownPackage(t *testing.T) {
	r := NewRegistry()
	r.Register(TypeInfo{
		Package:    "containerv1beta1",
		Group:      "container.cnrm.cloud.google.com",
		Version:    "v1beta1",
		Kind:       "ContainerCluster",
		APIVersion: "container.cnrm.cloud.google.com/v1beta1",
	})

	tests := []struct {
		name string
		pkg  string
		want bool
	}{
		{"known package", "containerv1beta1", true},
		{"unknown package", "unknownv1alpha1", false},
		{"empty string", "", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := r.IsKnownPackage(tt.pkg); got != tt.want {
				t.Errorf("IsKnownPackage(%q) = %v, want %v", tt.pkg, got, tt.want)
			}
		})
	}
}

func TestRegistry_GetTypeInfo(t *testing.T) {
	r := NewRegistry()
	info := TypeInfo{
		Package:    "sqlv1beta1",
		Group:      "sql.cnrm.cloud.google.com",
		Version:    "v1beta1",
		Kind:       "SQLInstance",
		APIVersion: "sql.cnrm.cloud.google.com/v1beta1",
	}
	r.Register(info)

	t.Run("by full type name", func(t *testing.T) {
		got, ok := r.GetTypeInfo("sqlv1beta1.SQLInstance")
		if !ok {
			t.Fatal("expected to find type info")
		}
		if got.Kind != "SQLInstance" {
			t.Errorf("Kind = %q, want %q", got.Kind, "SQLInstance")
		}
	})

	t.Run("by kind only", func(t *testing.T) {
		got, ok := r.GetTypeInfo("SQLInstance")
		if !ok {
			t.Fatal("expected to find type info by kind")
		}
		if got.Package != "sqlv1beta1" {
			t.Errorf("Package = %q, want %q", got.Package, "sqlv1beta1")
		}
	})

	t.Run("unknown type", func(t *testing.T) {
		_, ok := r.GetTypeInfo("NotAType")
		if ok {
			t.Error("expected not to find unknown type")
		}
	})
}

func TestRegistry_APIVersionForPackage(t *testing.T) {
	r := NewRegistry()
	r.Register(TypeInfo{
		Package:    "dnsv1beta1",
		Group:      "dns.cnrm.cloud.google.com",
		Version:    "v1beta1",
		Kind:       "DNSManagedZone",
		APIVersion: "dns.cnrm.cloud.google.com/v1beta1",
	})

	tests := []struct {
		name string
		pkg  string
		want string
	}{
		{"known package", "dnsv1beta1", "dns.cnrm.cloud.google.com/v1beta1"},
		{"unknown package", "unknownv1", ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := r.APIVersionForPackage(tt.pkg); got != tt.want {
				t.Errorf("APIVersionForPackage(%q) = %q, want %q", tt.pkg, got, tt.want)
			}
		})
	}
}

func TestRegistry_ListPackages(t *testing.T) {
	r := NewRegistry()
	r.Register(TypeInfo{Package: "computev1beta1", Kind: "ComputeInstance"})
	r.Register(TypeInfo{Package: "storagev1beta1", Kind: "StorageBucket"})
	r.Register(TypeInfo{Package: "computev1beta1", Kind: "ComputeDisk"}) // Same package

	pkgs := r.ListPackages()
	if len(pkgs) != 2 {
		t.Errorf("ListPackages() returned %d packages, want 2", len(pkgs))
	}

	found := make(map[string]bool)
	for _, p := range pkgs {
		found[p] = true
	}

	if !found["computev1beta1"] || !found["storagev1beta1"] {
		t.Errorf("ListPackages() = %v, expected computev1beta1 and storagev1beta1", pkgs)
	}
}

func TestRegistry_RegisterBulk(t *testing.T) {
	r := NewRegistry()

	types := []TypeInfo{
		{Package: "pubsubv1beta1", Kind: "PubSubTopic"},
		{Package: "pubsubv1beta1", Kind: "PubSubSubscription"},
		{Package: "redisv1beta1", Kind: "RedisInstance"},
	}

	r.RegisterBulk(types)

	if !r.IsKnownType("PubSubTopic") {
		t.Error("PubSubTopic not registered")
	}
	if !r.IsKnownType("PubSubSubscription") {
		t.Error("PubSubSubscription not registered")
	}
	if !r.IsKnownType("RedisInstance") {
		t.Error("RedisInstance not registered")
	}
}

func TestRegistry_Concurrency(t *testing.T) {
	r := NewRegistry()

	// Register types concurrently
	done := make(chan bool)
	for i := 0; i < 100; i++ {
		go func(n int) {
			r.Register(TypeInfo{
				Package: "testv1",
				Kind:    "TestResource",
			})
			_ = r.IsKnownType("testv1.TestResource")
			_ = r.IsKnownPackage("testv1")
			_ = r.ListPackages()
			done <- true
		}(i)
	}

	for i := 0; i < 100; i++ {
		<-done
	}

	// Verify state is consistent
	if !r.IsKnownType("testv1.TestResource") {
		t.Error("type not found after concurrent operations")
	}
}

func TestDefaultRegistry(t *testing.T) {
	// DefaultRegistry should be initialized and usable
	if DefaultRegistry == nil {
		t.Fatal("DefaultRegistry is nil")
	}

	// Should have types registered from builtins.go
	pkgs := DefaultRegistry.ListPackages()
	if len(pkgs) == 0 {
		t.Error("DefaultRegistry has no packages registered")
	}
}
