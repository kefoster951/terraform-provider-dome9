package variable

// AWS resource/data source
const (
	CloudAccountAWSCreationResourceName    = "test_cloudaccount_aws"
	CloudAccountAWSVendor                  = "aws"
	CloudAccountAWSOriginalAccountName     = "original_cloud_account_name_before_change"
	CloudAccountAWSUpdatedAccountName      = "updated_cloud_account_name"
	CloudAccountAWSFetchedRegion           = "us_east_1"
	CloudAccountAWSReadOnlyGroupBehavior   = "ReadOnly"
	CloudAccountAWSFullManageGroupBehavior = "FullManage"
)

// Azure resource/data source
const (
	CloudAccountAzureCreationResourceName = "test_cloudaccount_azure"
	CloudAccountAzureUpdatedAccountName   = "updated_cloud_account_name"
	CloudAccountAzureOperationMode        = "Read"
	CloudAccountAzureVendor               = "azure"
)

// GCP resource/data source
const (
	CloudAccountGCPCreationResourceName    = "test_cloudaccount_gcp"
	CloudAccountGCPUpdatedAccountName      = "updated_cloud_account_name"
	CloudAccountGCPType                    = "service_account"
	CloudAccountGCPVendor                  = "google"
	CloudAccountGCPAuthURL                 = "https://accounts.google.com/o/oauth2/auth"
	CloudAccountGCPTokenURL                = "https://oauth2.googleapis.com/token"
	CloudAccountGCPAuthProviderX509CertURL = "https://www.googleapis.com/oauth2/v1/certs"
)

// IpList resource/data source
const (
	IPListCreationResourceName      = "test_iplist"
	IPListDescriptionResource       = "acceptance-test"
	IPListUpdateDescriptionResource = "update-acceptance-test"
)