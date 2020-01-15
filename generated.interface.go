package domain

import (
	"github.com/nortonlifelock/connection"

	"time"
)

//**********************************************************
// GENERATED CODE - DO NOT CHANGE
// This file is generated using scaffolding. Any changes to
// this file will be overwritten on the next build
//**********************************************************

// GeneratedDatabaseConnection outlines all the stored procedures used by the application. This interface can
// be used to mock the connection for the purpose of testing
type GeneratedDatabaseConnection interface {
	connection.DatabaseConnection

	CleanUp() (id int, affectedRows int, err error)
	CreateAssetGroup(inOrgID string, _GroupID int, _ScannerSourceID string) (id int, affectedRows int, err error)
	CreateAssetWithIPInstanceID(_State string, _IP string, _MAC string, _SourceID string, _InstanceID string, _Region string, _OrgID string, _OS string, _OsTypeID int) (id int, affectedRows int, err error)
	CreateCategory(_Category string) (id int, affectedRows int, err error)
	CreateDBLog(_User string, _Command string, _Endpoint string) (id int, affectedRows int, err error)
	CreateDetection(_OrgID string, _SourceID string, _DeviceID string, _VulnID string, _AlertDate time.Time, _Proof string, _Port int, _Protocol string, _DetectionStatusID int, _TimesSeen int) (id int, affectedRows int, err error)
	CreateDetectionActiveKernel(_OrgID string, _SourceID string, _DeviceID string, _VulnID string, _AlertDate time.Time, _Proof string, _Port int, _Protocol string, _ActiveKernel int, _DetectionStatusID int, _TimesSeen int) (id int, affectedRows int, err error)
	CreateDetectionWithIgnore(_OrgID string, _SourceID string, _DeviceID string, _VulnID string, _IgnoreID string, _AlertDate time.Time, _Proof string, _Port int, _Protocol string, _DetectionStatusID int, _TimesSeen int) (id int, affectedRows int, err error)
	CreateDetectionWithIgnoreActiveKernel(_OrgID string, _SourceID string, _DeviceID string, _VulnID string, _IgnoreID string, _AlertDate time.Time, _Proof string, _Port int, _Protocol string, _ActiveKernel int, _DetectionStatusID int, _TimesSeen int) (id int, affectedRows int, err error)
	CreateDevice(_AssetID string, _SourceID string, _Ip string, _Hostname string, _MAC string, _GroupID int, _OrgID string, _OS string, _OSTypeID int) (id int, affectedRows int, err error)
	CreateException(inSourceID string, inOrganizationID string, inTypeID int, inVulnerabilityID string, inDeviceID string, inDueDate time.Time, inApproval string, inActive bool, inPort string, inCreatedBy string) (id int, affectedRows int, err error)
	CreateJobConfig(_JobID int, _OrganizationID string, _PriorityOverride int, _Continuous bool, _WaitInSeconds int, _MaxInstances int, _AutoStart bool, _CreatedBy string, _DataInSourceID string, _DataOutSourceID string) (id int, affectedRows int, err error)
	CreateJobConfigWPayload(_JobID int, _OrganizationID string, _PriorityOverride int, _Continuous bool, _WaitInSeconds int, _MaxInstances int, _AutoStart bool, _CreatedBy string, _DataInSourceID string, _DataOutSourceID string, _Payload string) (id int, affectedRows int, err error)
	CreateJobHistory(_JobID int, _ConfigID string, _StatusID int, _Priority int, _Identifier string, _CurrentIteration int, _Payload string, _ThreadID string, _PulseDate time.Time, _CreatedBy string) (id int, affectedRows int, err error)
	CreateJobHistoryWithParentID(_JobID int, _ConfigID string, _StatusID int, _Priority int, _Identifier string, _CurrentIteration int, _Payload string, _ThreadID string, _PulseDate time.Time, _CreatedBy string, _ParentID string) (id int, affectedRows int, err error)
	CreateOrganization(_Code string, _Description string, _TimeZoneOffset float32, _UpdatedBy string) (id int, affectedRows int, err error)
	CreateOrganizationWithPayloadEkey(_Code string, _Description string, _TimeZoneOffset float32, _Payload string, _EKEY string, _UpdatedBy string) (id int, affectedRows int, err error)
	CreateScanSummary(_SourceID string, _OrgID string, _ScanID string, _ScanStatus string, _ScanClosePayload string, _ParentJobID string) (id int, affectedRows int, err error)
	CreateSourceConfig(_Source string, _SourceID string, _OrganizationID string, _Address string, _Port string, _Username string, _Password string, _PrivateKey string, _ConsumerKey string, _Token string, _Payload string) (id int, affectedRows int, err error)
	CreateTag(_DeviceID string, _TagKeyID string, _Value string) (id int, affectedRows int, err error)
	CreateTagKey(_KeyValue string) (id int, affectedRows int, err error)
	CreateTagMap(_TicketingSourceID string, _TicketingTag string, _CloudSourceID string, _CloudTag string, _Options string, _OrganizationID string) (id int, affectedRows int, err error)
	CreateTicket(_Title string, _Status string, _DetectionID string, _OrganizationID string, _DueDate time.Time, _UpdatedDate time.Time, _ResolutionDate time.Time) (id int, affectedRows int, err error)
	CreateTicketingJob(GroupID int, OrgID string, ScanStartDate string) (id int, affectedRows int, err error)
	CreateUser(_Username string, _FirstName string, _LastName string, _Email string) (id int, affectedRows int, err error)
	CreateUserPermissions(_UserID string, _OrgID string) (id int, affectedRows int, err error)
	CreateUserSession(_UserID string, _OrgID string, _SessionKey string) (id int, affectedRows int, err error)
	CreateVulnInfo(_SourceVulnID string, _Title string, _SourceID string, _CVSSScore float32, _CVSS3Score float32, _Description string, _Solution string, _Software string, _DetectionInformation string) (id int, affectedRows int, err error)
	CreateVulnInfoNoCVSS3(_SourceVulnID string, _Title string, _SourceID string, _CVSSScore float32, _Description string, _Solution string, _Software string, _DetectionInformation string) (id int, affectedRows int, err error)
	CreateVulnRef(_VulnInfoID string, _SourceID string, _Reference string, _RefType int) (id int, affectedRows int, err error)
	DeleteIgnoreForDevice(_sourceID string, _devID string, _orgID string) (id int, affectedRows int, err error)
	DeleteSessionByToken(_SessionKey string) (id int, affectedRows int, err error)
	DeleteTagMap(_TicketingSourceID string, _TicketingTag string, _CloudSourceID string, _CloudTag string, _OrganizationID string) (id int, affectedRows int, err error)
	DeleteUserByUsername(_Username string) (id int, affectedRows int, err error)
	DisableIgnore(inSourceID string, inDevID string, inOrgID string, inVulnID string, inPortID string, inUpdatedBy string) (id int, affectedRows int, err error)
	DisableJobConfig(_ID string, _UpdatedBy string) (id int, affectedRows int, err error)
	DisableOrganization(_ID string, _UpdatedBy string) (id int, affectedRows int, err error)
	DisableSource(_ID string, _OrgID string, _UpdatedBy string) (id int, affectedRows int, err error)
	GetAllExceptions(_offset int, _limit int, _sourceID string, _orgID string, _typeID int, _vulnID string, _devID string, _dueDate time.Time, _port string, _approval string, _active bool, _dBCreatedDate time.Time, _dBUpdatedDate time.Time, _updatedBy string, _createdBy string, _sortField string, _sortOrder string) ([]Ignore, error)
	GetAllJobConfigs(_OrgID string) ([]JobConfig, error)
	GetAllJobConfigsWithOrder(_offset int, _limit int, _configID string, _jobid int, _dataInSourceConfigID string, _dataOutSourceConfigID string, _priorityOverride int, _continuous bool, _Payload string, _waitInSeconds int, _maxInstances int, _autoStart bool, _OrgID string, _updatedBy string, _createdBy string, _sortField string, _sortOrder string, _updatedDate time.Time, _createdDate time.Time, _lastJobStart time.Time, _ID string) ([]JobConfig, error)
	GetAssetGroup(inOrgID string, _GroupID int, _ScannerSourceID string) (AssetGroup, error)
	GetAssetGroupsByCloudSource(inOrgID string, inCloudSourceID string) ([]AssetGroup, error)
	GetAssignmentGroupByIP(_SourceID string, _OrganizationID string, _IP string) ([]AssignmentGroup, error)
	GetAssignmentGroupByOrgIP(_OrganizationID string, _IP string) ([]AssignmentGroup, error)
	GetAutoStartJobs() ([]JobConfig, error)
	GetCISAssignments(_OrganizationID string) ([]CISAssignments, error)
	GetCancelledJobs() ([]JobHistory, error)
	GetCategoryByName(_Name string) ([]Category, error)
	GetDetectionInfo(_DeviceID string, _VulnerabilityID string) (DetectionInfo, error)
	GetDetectionInfoAfter(_After time.Time, _OrgID string) ([]DetectionInfo, error)
	GetDetectionInfoBySourceVulnID(_SourceDeviceID string, _SourceVulnerabilityID string) (DetectionInfo, error)
	GetDetectionStatusByID(_ID int) (DetectionStatus, error)
	GetDetectionStatusByName(_Name string) (DetectionStatus, error)
	GetDetectionStatuses() ([]DetectionStatus, error)
	GetDetectionsInfoForDevice(_DeviceID string) ([]DetectionInfo, error)
	GetDeviceInfoByAssetOrgID(inAssetID string, inOrgID string) (DeviceInfo, error)
	GetDeviceInfoByCloudSourceIDAndIP(_IP string, _CloudSourceID string, _OrgID string) (DeviceInfo, error)
	GetDeviceInfoByGroupIP(inIP string, inGroupID int, inOrgID string) (DeviceInfo, error)
	GetDeviceInfoByIP(_IP string, _OrgID string) (DeviceInfo, error)
	GetDeviceInfoByIPMACAndRegion(_IP string, _MAC string, _Region string, _OrgID string) (DeviceInfo, error)
	GetDeviceInfoByInstanceID(_InstanceID string, _OrgID string) (DeviceInfo, error)
	GetDeviceInfoByScannerSourceID(_IP string, _GroupID int, _OrgID string) (DeviceInfo, error)
	GetDevicesInfoByCloudSourceID(_CloudSourceID string, _OrgID string) ([]DeviceInfo, error)
	GetDevicesInfoBySourceID(_SourceID string, _OrgID string) ([]DeviceInfo, error)
	GetExceptionByVulnIDOrg(_DeviceID string, _VulnID string, _OrgID string) (Ignore, error)
	GetExceptionTypes() ([]ExceptionType, error)
	GetExceptionsDueNext30Days() ([]CERF, error)
	GetExceptionsLength(inSourceID string, inOrgID string, inTypeID int, inVulnID string, inDevID string, inDueDate time.Time, inPort string, inApproval string, inActive bool, inDBCreatedDate time.Time, inDBUpdatedDate time.Time, inUpdatedBy string, inCreatedBy string) (QueryData, error)
	GetJobByID(_ID int) (JobRegistration, error)
	GetJobConfig(_ID string) (JobConfig, error)
	GetJobConfigAudit(inJobConfigID string, inOrgID string) ([]JobConfigAudit, error)
	GetJobConfigByID(_ID string, _OrgID string) (JobConfig, error)
	GetJobConfigByJobHistoryID(_JobHistoryID string) (JobConfig, error)
	GetJobConfigByOrgIDAndJobID(_OrgID string, _JobID int) ([]JobConfig, error)
	GetJobConfigLength(_configID string, _jobID int, _dataInSourceConfigID string, _dataOutSourceConfigID string, _priorityOverride int, _continuous bool, _Payload string, _waitInSeconds int, _maxInstances int, _autoStart bool, _OrgID string, _updatedBy string, _createdBy string, _updatedDate time.Time, _createdDate time.Time, _lastJobStart time.Time, _ID string) (QueryData, error)
	GetJobHistories(_offset int, _limit int, _jobID int, _jobconfig string, _status int, _Payload string, _OrgID string) ([]JobHistory, error)
	GetJobHistoryByID(_ID string) (JobHistory, error)
	GetJobHistoryLength(_jobid int, _jobconfig string, _status int, _Payload string, _orgid string) (QueryData, error)
	GetJobQueueByStatusID(_StatusID int) ([]JobHistory, error)
	GetJobs() ([]JobRegistration, error)
	GetJobsByStruct(_Struct string) (JobRegistration, error)
	GetLeafOrganizationsForUser(_UserID string) ([]Organization, error)
	GetLogTypes() ([]LogType, error)
	GetLogsByParams(_MethodOfDiscovery string, _jobType int, _logType int, _jobHistoryID string, _fromDate time.Time, _toDate time.Time, _OrgID string) ([]DBLog, error)
	GetMatchedVulns() ([]VulnerabilityMatch, error)
	GetOperatingSystemType(_OS string) (OperatingSystemType, error)
	GetOrganizationByCode(Code string) (Organization, error)
	GetOrganizationByID(ID string) (Organization, error)
	GetOrganizations() ([]Organization, error)
	GetPendingActiveRescanJob(_OrgID string) ([]JobHistory, error)
	GetPermissionByUserOrgID(_UserID string, _OrgID string) (Permission, error)
	GetPermissionOfLeafOrgByUserID(_UserID string) (Permission, error)
	GetRecentlyUpdatedScanSummaries(_OrgID string) ([]ScanSummary, error)
	GetScanSummariesBySourceName(_OrgID string, _SourceName string) ([]ScanSummary, error)
	GetScanSummary(_SourceID string, _OrgID string, _ScanID string) (ScanSummary, error)
	GetScanSummaryBySourceKey(_SourceKey string) (ScanSummary, error)
	GetScheduledJobsToStart(_LastChecked time.Time) ([]JobSchedule, error)
	GetSessionByToken(_SessionKey string) (Session, error)
	GetSourceByID(_ID string) (Source, error)
	GetSourceByName(_Source string) (Source, error)
	GetSourceConfigByID(_ID string) (SourceConfig, error)
	GetSourceConfigByNameOrg(_Source string, _OrgID string) ([]SourceConfig, error)
	GetSourceConfigByOrgID(_OrgID string) ([]SourceConfig, error)
	GetSourceConfigBySourceID(_OrgID string, _SourceID string) ([]SourceConfig, error)
	GetSourceInsByJobID(inJob int, inOrgID string) ([]SourceConfig, error)
	GetSourceOauthByOrgURL(_URL string, _OrgID string) (SourceConfig, error)
	GetSourceOauthByURL(_URL string) (SourceConfig, error)
	GetSourceOutsByJobID(inJob int, inOrgID string) ([]SourceConfig, error)
	GetSources() ([]Source, error)
	GetTagByDeviceAndTagKey(_DeviceID string, _TagKeyID string) (Tag, error)
	GetTagKeyByID(_ID string) (TagKey, error)
	GetTagKeyByKey(_KeyValue string) (TagKey, error)
	GetTagMapsByOrg(_OrganizationID string) ([]TagMap, error)
	GetTagMapsByOrgCloudSourceID(_CloudID string, _OrganizationID string) ([]TagMap, error)
	GetTagsForDevice(_DeviceID string) ([]Tag, error)
	GetTicketByDeviceIDVulnID(inDeviceID string, inVulnID string, inOrgID string) (TicketSummary, error)
	GetTicketByTitle(_Title string, _OrgID string) (TicketSummary, error)
	GetUnfinishedScanSummariesBySourceOrgID(_SourceID string, _OrgID string) ([]ScanSummary, error)
	GetUnmatchedVulns(_SourceID int) ([]VulnerabilityInfo, error)
	GetUserAnyOrg(_ID string) (User, error)
	GetUserByID(_ID string, _OrgID string) (User, error)
	GetUserByUsername(_Username string) (User, error)
	GetUsersByOrg(_OrgID string) ([]User, error)
	GetVulnInfoByID(_ID string) (VulnerabilityInfo, error)
	GetVulnInfoBySource(_Source string) ([]VulnerabilityInfo, error)
	GetVulnInfoBySourceID(_SourceID string) ([]VulnerabilityInfo, error)
	GetVulnInfoBySourceVulnID(_SourceVulnID string) (VulnerabilityInfo, error)
	GetVulnInfoBySourceVulnIDSourceID(_SourceVulnID string, _SourceID string, _Modified time.Time) (VulnerabilityInfo, error)
	GetVulnRefInfo(_VulnInfoID string, _SourceID string, _Reference string) (VulnerabilityReferenceInfo, error)
	GetVulnRefInfoVendor(_VulnInfoID string, _SourceID string) ([]VulnerabilityReferenceInfo, error)
	GetVulnReferencesInfo(_VulnInfoID string, _SourceID string) ([]VulnerabilityReferenceInfo, error)
	GetVulnReferencesInfoBySourceAndRef(_SourceID string, _Reference string) ([]VulnerabilityReferenceInfo, error)
	HasDecommissioned(_devID string, _sourceID string, _orgID string) (Ignore, error)
	HasExceptionOrFalsePositive(_sourceID string, _vulnID string, _devID string, _orgID string, _port string, _OS string) ([]Ignore, error)
	HasIgnore(inSourceID string, inVulnID string, inDevID string, inOrgID string, inPort string, inMostCurrentDetection time.Time) (Ignore, error)
	PulseJob(_JobHistoryID string) (id int, affectedRows int, err error)
	SaveAssignmentGroup(_SourceID string, _OrganizationID string, _IpAddress string, _GroupName string) (id int, affectedRows int, err error)
	SaveIgnore(_SourceID string, _OrganizationID string, _TypeID int, _VulnerabilityID string, _DeviceID string, _DueDate time.Time, _Approval string, _Active bool, _port string) (id int, affectedRows int, err error)
	SaveScanSummary(_ScanID string, _ScanStatus string) (id int, affectedRows int, err error)
	SetScheduleLastRun(_ID string) (id int, affectedRows int, err error)
	UpdateAssetIDOsTypeIDOfDevice(_ID string, _AssetID string, _ScannerSourceID string, _GroupID int, _OS string, _HostName string, _OsTypeID int, _OrgID string) (id int, affectedRows int, err error)
	UpdateDetectionTimesSeen(_DeviceID string, _VulnID string, _TimesSeen int, _StatusID int) (id int, affectedRows int, err error)
	UpdateExpirationDateByCERF(_CERForm string, _OrganizationID string, _DueDate time.Time) (id int, affectedRows int, err error)
	UpdateInstanceIDOfDevice(_ID string, _InstanceID string, _CloudSourceID string, _State string, _Region string, _OrgID string) (id int, affectedRows int, err error)
	UpdateJobConfig(_ID string, _DataInSourceID string, _DataOutSourceID string, _Autostart bool, _PriorityOverride int, _Continuous bool, _WaitInSeconds int, _MaxInstances int, _UpdatedBy string, _OrgID string) (id int, affectedRows int, err error)
	UpdateJobConfigLastRun(_ID string, _LastRun time.Time) (id int, affectedRows int, err error)
	UpdateJobHistory(_ID string, _Status int, _ConfigID string, _Payload string, _UpdatedBy string) (id int, affectedRows int, err error)
	UpdateJobHistoryStatus(_ID string, _Status int) (id int, affectedRows int, err error)
	UpdateJobHistoryStatusDetailed(_ID string, _Status int, _UpdatedBy string) (id int, affectedRows int, err error)
	UpdateOrganization(_ID string, _Description string, _TimezoneOffset float32, _UpdatedBy string) (id int, affectedRows int, err error)
	UpdatePermissionsByUserOrgID(_UserID string, _OrgID string, _Admin bool, _Manager bool, _Reader bool, _Reporter bool) (id int, affectedRows int, err error)
	UpdateSourceConfig(_ID string, _OrgID string, _Address string, _Username string, _Password string, _PrivateKey string, _ConsumerKey string, _Token string, _Port string, _Payload string, _UpdatedBy string) (id int, affectedRows int, err error)
	UpdateSourceConfigToken(_ID string, _Token string) (id int, affectedRows int, err error)
	UpdateStateOfDevice(_ID string, _State string, _OrgID string) (id int, affectedRows int, err error)
	UpdateTag(_DeviceID string, _TagKeyID string, _Value string) (id int, affectedRows int, err error)
	UpdateTagMap(_TicketingSourceID string, _TicketingTag string, _CloudSourceID string, _CloudTag string, _Options string, _OrganizationID string) (id int, affectedRows int, err error)
	UpdateTicket(_Title string, _Status string, _OrganizationID string, _UpdatedDate time.Time, _ResolutionDate time.Time) (id int, affectedRows int, err error)
	UpdateUserByID(_ID string, _FirstName string, _LastName string, _Email string, _Disabled bool) (id int, affectedRows int, err error)
	UpdateVulnByID(_ID string, _SourceVulnID string, _Title string, _SourceID string, _CVSSScore float32, _CVSS3Score float32, _Description string, _Solution string, _Software string, _DetectionInformation string) (id int, affectedRows int, err error)
	UpdateVulnByIDNoCVSS3(_ID string, _SourceVulnID string, _Title string, _SourceID string, _CVSSScore float32, _Description string, _Solution string, _Software string, _DetectionInformation string) (id int, affectedRows int, err error)
	UpdateVulnInfoID(_VulnInfoID string, _VulnID string, _MatchConfidence int, _MatchReasons string) (id int, affectedRows int, err error)
}
