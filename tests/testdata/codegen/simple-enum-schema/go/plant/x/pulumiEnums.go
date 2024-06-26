// Code generated by test DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package plant

// The log_name to populate in the Cloud Audit Record. This is added to regress pulumi/pulumi issue #7913
type CloudAuditOptionsLogName string

const (
	// Default. Should not be used.
	CloudAuditOptionsLogNameCloudAuditOptionsLogNameUnspecifiedLogName = CloudAuditOptionsLogName("UNSPECIFIED_LOG_NAME")
	// Corresponds to "cloudaudit.googleapis.com/activity"
	CloudAuditOptionsLogNameCloudAuditOptionsLogNameAdminActivity = CloudAuditOptionsLogName("ADMIN_ACTIVITY")
	// Corresponds to "cloudaudit.googleapis.com/data_access"
	CloudAuditOptionsLogNameCloudAuditOptionsLogNameDataAccess = CloudAuditOptionsLogName("DATA_ACCESS")
	// What if triple quotes """ are used in the description
	CloudAuditOptionsLogNameCloudAuditOptionsLogNameSynthetic = CloudAuditOptionsLogName("SYNTHETIC")
	CloudAuditOptionsLogName_CloudAuditOptionsLogName_NO_NAME = CloudAuditOptionsLogName("_NO_NAME")
)

type ContainerBrightness float64

const (
	ContainerBrightnessContainerBrightnessZeroPointOne = ContainerBrightness(0.1)
	ContainerBrightnessContainerBrightnessOne          = ContainerBrightness(1)
)

// plant container colors
type ContainerColor string

const (
	ContainerColorContainerColorRed    = ContainerColor("red")
	ContainerColorContainerColorBlue   = ContainerColor("blue")
	ContainerColorContainerColorYellow = ContainerColor("yellow")
)

// plant container sizes
type ContainerSize int

const (
	ContainerSizeContainerSizeFourInch = ContainerSize(4)
	ContainerSizeContainerSizeSixInch  = ContainerSize(6)
	// Deprecated: Eight inch pots are no longer supported.
	ContainerSizeContainerSizeEightInch = ContainerSize(8)
)
