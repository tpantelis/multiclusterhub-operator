package v1

import (
	"errors"
	"fmt"
	"os"
)

// Component related to MultiClusterHub (MCH)
const (
	Appsub                    string = "app-lifecycle"
	ClusterBackup             string = "cluster-backup"
	ClusterLifecycle          string = "cluster-lifecycle"
	ClusterPermission         string = "cluster-permission"
	Console                   string = "console"
	GRC                       string = "grc"
	Insights                  string = "insights"
	ManagementIngress         string = "management-ingress"
	MultiClusterEngine        string = "multicluster-engine"
	MultiClusterObservability string = "multicluster-observability"
	Repo                      string = "multiclusterhub-repo"
	Search                    string = "search"
	SubmarinerAddon           string = "submariner-addon"
	Volsync                   string = "volsync"
)

// Component related to MultiCluster Engine (MCE)
const (
	MCEAssistedService              string = "assisted-service"
	MCEClusterLifecycle             string = "cluster-lifecycle-mce"
	MCEClusterManager               string = "cluster-manager"
	MCEClusterProxyAddon            string = "cluster-proxy-addon"
	MCEConsole                      string = "console-mce"
	MCEDiscovery                    string = "discovery"
	MCEHive                         string = "hive"
	MCEHypershiftLocalHosting       string = "hypershift-local-hosting"
	MCEHypershiftPreview            string = "hypershift-preview"
	MCEHypershift                   string = "hypershift"
	MCELocalCluster                 string = "local-cluster"
	MCEManagedServiceAccount        string = "managedserviceaccount"
	MCEManagedServiceAccountPreview string = "managedserviceaccount-preview"
	MCEServerFoundation             string = "server-foundation"
)

// allComponents is a slice containing all the component names from both "MCH" and "MCE" categories.
var allComponents = []string{
	Appsub,
	ClusterBackup,
	ClusterLifecycle,
	ClusterPermission,
	Console,
	GRC,
	Insights,
	ManagementIngress,
	MultiClusterEngine,
	MultiClusterObservability,
	Repo,
	Search,
	SubmarinerAddon,
	Volsync,

	MCEAssistedService,
	MCEClusterLifecycle,
	MCEClusterManager,
	MCEClusterProxyAddon,
	MCEConsole,
	MCEDiscovery,
	MCEHive,
	MCEServerFoundation,
	MCEConsole,
	MCEManagedServiceAccount,
	MCEManagedServiceAccountPreview,
	MCEHypershift,
	MCEHypershiftLocalHosting,
	MCEHypershiftPreview,
	MCEManagedServiceAccount,
	MCEServerFoundation,
}

// MCHComponents is a slice containing component names specific to the "MCH" category.
var MCHComponents = []string{
	Appsub,
	ClusterBackup,
	ClusterLifecycle,
	ClusterPermission,
	Console,
	GRC,
	Insights,
	// MultiClusterEngine,
	MultiClusterObservability,
	//Repo,
	Search,
	SubmarinerAddon,
	Volsync,
}

// MCEComponents is a slice containing component names specific to the "MCE" category.
var MCEComponents = []string{
	MCEAssistedService,
	MCEClusterLifecycle,
	MCEClusterManager,
	MCEConsole,
	MCEDiscovery,
	MCEHive,
	MCEServerFoundation,
	MCEConsole,
	MCEManagedServiceAccount,
	MCEManagedServiceAccountPreview,
	MCEHypershift,
	MCEHypershiftLocalHosting,
	MCEHypershiftPreview,
	MCEManagedServiceAccount,
	MCEServerFoundation,
}

// clusterManagementAddOns is a map that associates certain component names with their corresponding add-ons.
var clusterManagementAddOns = map[string]string{
	SubmarinerAddon: "submariner",
	// Add other components here when clusterManagementAddOns is required.
}

/*
GetDefaultEnabledComponents returns a slice of default enabled component names.
It is expected to be used to get a list of components that are enabled by default.
*/
func GetDefaultEnabledComponents() ([]string, error) {
	var defaultEnabledComponents = []string{
		//Repo,
		Appsub,
		ClusterLifecycle,
		ClusterPermission,
		Console,
		GRC,
		Insights,
		MultiClusterEngine,
		MultiClusterObservability,
		Search,
		SubmarinerAddon,
		Volsync,
	}

	return defaultEnabledComponents, nil
}

/*
GetDefaultDisabledComponents returns a slice of default disabled component names.
It is expected to be used to get a list of components that are disabled by default.
*/
func GetDefaultDisabledComponents() ([]string, error) {
	var defaultDisabledComponents = []string{
		ClusterBackup,
	}
	return defaultDisabledComponents, nil
}

/*
GetDefaultHostedComponents returns a slice of default hosted components.
These are components that are hosted within the system.
*/
func GetDefaultHostedComponents() []string {
	var defaultHostedComponents = []string{
		MultiClusterEngine,
		// Add other components here when added to hostedmode
	}

	return defaultHostedComponents
}

// GetClusterManagementAddonName returns the name of the ClusterManagementAddOn based on the provided component name.
func GetClusterManagementAddonName(component string) (string, error) {
	if val, ok := clusterManagementAddOns[component]; !ok {
		return val, fmt.Errorf("failed to find ClusterManagementAddOn name for: %s component", component)
	} else {
		return val, nil
	}
}

/*
ComponentPresent checks if a specific component is present based on the provided component name in the
MultiClusterHub struct.
*/
func (mch *MultiClusterHub) ComponentPresent(s string) bool {
	if mch.Spec.Overrides == nil {
		return false
	}
	for _, c := range mch.Spec.Overrides.Components {
		if c.Name == s {
			return true
		}
	}
	return false
}

// Enabled checks if a specific component is enabled based on the provided component name in the MultiClusterHub struct.
func (mch *MultiClusterHub) Enabled(s string) bool {
	if mch.Spec.Overrides == nil {
		return false
	}
	for _, c := range mch.Spec.Overrides.Components {
		if c.Name == s {
			return c.Enabled
		}
	}

	return false
}

// Enable enables a specific component based on the provided component name in the MultiClusterHub struct.
func (mch *MultiClusterHub) Enable(s string) {
	if mch.Spec.Overrides == nil {
		mch.Spec.Overrides = &Overrides{}
	}
	for i, c := range mch.Spec.Overrides.Components {
		if c.Name == s {
			mch.Spec.Overrides.Components[i].Enabled = true
			return
		}
	}
	mch.Spec.Overrides.Components = append(mch.Spec.Overrides.Components, ComponentConfig{
		Name:    s,
		Enabled: true,
	})
}

// Disable disables a specific component based on the provided component name in the MultiClusterHub struct.
func (mch *MultiClusterHub) Disable(s string) {
	if mch.Spec.Overrides == nil {
		mch.Spec.Overrides = &Overrides{}
	}
	for i, c := range mch.Spec.Overrides.Components {
		if c.Name == s {
			mch.Spec.Overrides.Components[i].Enabled = false
			return
		}
	}
	mch.Spec.Overrides.Components = append(mch.Spec.Overrides.Components, ComponentConfig{
		Name:    s,
		Enabled: false,
	})
}

/*
Prune removes a specific component from the component list in the MultiClusterHub struct.
It returns true if changes were made.
*/
func (mch *MultiClusterHub) Prune(s string) bool {
	if mch.Spec.Overrides == nil {
		return false
	}
	pruned := false
	prunedList := []ComponentConfig{}
	for _, c := range mch.Spec.Overrides.Components {
		if c.Name == s {
			pruned = true
		} else {
			prunedList = append(prunedList, c)
		}
	}

	if pruned {
		mch.Spec.Overrides.Components = prunedList
		return true
	}
	return false
}

// ValidComponent checks if a given component configuration is valid by comparing its name to the known component names.
func ValidComponent(c ComponentConfig) bool {
	for _, name := range allComponents {
		if c.Name == name {
			return true
		}
	}
	return false
}

// IsCommunity checks to see if the operator is running in community mode
func IsCommunity() (bool, error) {
	packageName := os.Getenv("OPERATOR_PACKAGE")
	if packageName == "advanced-cluster-management" {
		return false, nil
	} else if (packageName == "stolostron") || (packageName == "") {
		return true, nil
	} else {
		err := errors.New("There is an illegal value set for OPERATOR_PACKAGE")
		return true, err
	}
}

// IsInHostedMode returns true if mch is configured for hosted mode
func (mch *MultiClusterHub) IsInHostedMode() bool {
	a := mch.GetAnnotations()
	if a == nil {
		return false
	}
	if a["deploymentmode"] == string(ModeHosted) {
		return true
	}
	return false
}

// AvailabilityConfigIsValid returns true is the availability type is a recognized value
func AvailabilityConfigIsValid(config AvailabilityType) bool {
	switch config {
	case HAHigh, HABasic:
		return true
	default:
		return false
	}
}
