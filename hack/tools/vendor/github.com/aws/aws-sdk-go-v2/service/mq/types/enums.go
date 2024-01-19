// Code generated by smithy-go-codegen DO NOT EDIT.

package types

type AuthenticationStrategy string

// Enum values for AuthenticationStrategy
const (
	AuthenticationStrategySimple AuthenticationStrategy = "SIMPLE"
	AuthenticationStrategyLdap   AuthenticationStrategy = "LDAP"
)

// Values returns all known values for AuthenticationStrategy. Note that this can
// be expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (AuthenticationStrategy) Values() []AuthenticationStrategy {
	return []AuthenticationStrategy{
		"SIMPLE",
		"LDAP",
	}
}

type BrokerState string

// Enum values for BrokerState
const (
	BrokerStateCreationInProgress     BrokerState = "CREATION_IN_PROGRESS"
	BrokerStateCreationFailed         BrokerState = "CREATION_FAILED"
	BrokerStateDeletionInProgress     BrokerState = "DELETION_IN_PROGRESS"
	BrokerStateRunning                BrokerState = "RUNNING"
	BrokerStateRebootInProgress       BrokerState = "REBOOT_IN_PROGRESS"
	BrokerStateCriticalActionRequired BrokerState = "CRITICAL_ACTION_REQUIRED"
	BrokerStateReplica                BrokerState = "REPLICA"
)

// Values returns all known values for BrokerState. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (BrokerState) Values() []BrokerState {
	return []BrokerState{
		"CREATION_IN_PROGRESS",
		"CREATION_FAILED",
		"DELETION_IN_PROGRESS",
		"RUNNING",
		"REBOOT_IN_PROGRESS",
		"CRITICAL_ACTION_REQUIRED",
		"REPLICA",
	}
}

type BrokerStorageType string

// Enum values for BrokerStorageType
const (
	BrokerStorageTypeEbs BrokerStorageType = "EBS"
	BrokerStorageTypeEfs BrokerStorageType = "EFS"
)

// Values returns all known values for BrokerStorageType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (BrokerStorageType) Values() []BrokerStorageType {
	return []BrokerStorageType{
		"EBS",
		"EFS",
	}
}

type ChangeType string

// Enum values for ChangeType
const (
	ChangeTypeCreate ChangeType = "CREATE"
	ChangeTypeUpdate ChangeType = "UPDATE"
	ChangeTypeDelete ChangeType = "DELETE"
)

// Values returns all known values for ChangeType. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (ChangeType) Values() []ChangeType {
	return []ChangeType{
		"CREATE",
		"UPDATE",
		"DELETE",
	}
}

type DataReplicationMode string

// Enum values for DataReplicationMode
const (
	DataReplicationModeNone DataReplicationMode = "NONE"
	DataReplicationModeCrdr DataReplicationMode = "CRDR"
)

// Values returns all known values for DataReplicationMode. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (DataReplicationMode) Values() []DataReplicationMode {
	return []DataReplicationMode{
		"NONE",
		"CRDR",
	}
}

type DayOfWeek string

// Enum values for DayOfWeek
const (
	DayOfWeekMonday    DayOfWeek = "MONDAY"
	DayOfWeekTuesday   DayOfWeek = "TUESDAY"
	DayOfWeekWednesday DayOfWeek = "WEDNESDAY"
	DayOfWeekThursday  DayOfWeek = "THURSDAY"
	DayOfWeekFriday    DayOfWeek = "FRIDAY"
	DayOfWeekSaturday  DayOfWeek = "SATURDAY"
	DayOfWeekSunday    DayOfWeek = "SUNDAY"
)

// Values returns all known values for DayOfWeek. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (DayOfWeek) Values() []DayOfWeek {
	return []DayOfWeek{
		"MONDAY",
		"TUESDAY",
		"WEDNESDAY",
		"THURSDAY",
		"FRIDAY",
		"SATURDAY",
		"SUNDAY",
	}
}

type DeploymentMode string

// Enum values for DeploymentMode
const (
	DeploymentModeSingleInstance       DeploymentMode = "SINGLE_INSTANCE"
	DeploymentModeActiveStandbyMultiAz DeploymentMode = "ACTIVE_STANDBY_MULTI_AZ"
	DeploymentModeClusterMultiAz       DeploymentMode = "CLUSTER_MULTI_AZ"
)

// Values returns all known values for DeploymentMode. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (DeploymentMode) Values() []DeploymentMode {
	return []DeploymentMode{
		"SINGLE_INSTANCE",
		"ACTIVE_STANDBY_MULTI_AZ",
		"CLUSTER_MULTI_AZ",
	}
}

type EngineType string

// Enum values for EngineType
const (
	EngineTypeActivemq EngineType = "ACTIVEMQ"
	EngineTypeRabbitmq EngineType = "RABBITMQ"
)

// Values returns all known values for EngineType. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (EngineType) Values() []EngineType {
	return []EngineType{
		"ACTIVEMQ",
		"RABBITMQ",
	}
}

type PromoteMode string

// Enum values for PromoteMode
const (
	PromoteModeSwitchover PromoteMode = "SWITCHOVER"
	PromoteModeFailover   PromoteMode = "FAILOVER"
)

// Values returns all known values for PromoteMode. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (PromoteMode) Values() []PromoteMode {
	return []PromoteMode{
		"SWITCHOVER",
		"FAILOVER",
	}
}

type SanitizationWarningReason string

// Enum values for SanitizationWarningReason
const (
	SanitizationWarningReasonDisallowedElementRemoved     SanitizationWarningReason = "DISALLOWED_ELEMENT_REMOVED"
	SanitizationWarningReasonDisallowedAttributeRemoved   SanitizationWarningReason = "DISALLOWED_ATTRIBUTE_REMOVED"
	SanitizationWarningReasonInvalidAttributeValueRemoved SanitizationWarningReason = "INVALID_ATTRIBUTE_VALUE_REMOVED"
)

// Values returns all known values for SanitizationWarningReason. Note that this
// can be expanded in the future, and so it is only as up to date as the client.
// The ordering of this slice is not guaranteed to be stable across updates.
func (SanitizationWarningReason) Values() []SanitizationWarningReason {
	return []SanitizationWarningReason{
		"DISALLOWED_ELEMENT_REMOVED",
		"DISALLOWED_ATTRIBUTE_REMOVED",
		"INVALID_ATTRIBUTE_VALUE_REMOVED",
	}
}
