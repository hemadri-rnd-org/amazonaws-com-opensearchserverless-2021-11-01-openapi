package models

import (
	"context"
	"github.com/mark3labs/mcp-go/mcp"
)

type Tool struct {
	Definition mcp.Tool
	Handler    func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error)
}

// DeleteCollectionResponse represents the DeleteCollectionResponse schema from the OpenAPI specification
type DeleteCollectionResponse struct {
	Deletecollectiondetail interface{} `json:"deleteCollectionDetail,omitempty"`
}

// Tag represents the Tag schema from the OpenAPI specification
type Tag struct {
	Value interface{} `json:"value"`
	Key interface{} `json:"key"`
}

// CreateSecurityPolicyResponse represents the CreateSecurityPolicyResponse schema from the OpenAPI specification
type CreateSecurityPolicyResponse struct {
	Securitypolicydetail interface{} `json:"securityPolicyDetail,omitempty"`
}

// UntagResourceResponse represents the UntagResourceResponse schema from the OpenAPI specification
type UntagResourceResponse struct {
}

// CreateVpcEndpointResponse represents the CreateVpcEndpointResponse schema from the OpenAPI specification
type CreateVpcEndpointResponse struct {
	Createvpcendpointdetail interface{} `json:"createVpcEndpointDetail,omitempty"`
}

// CreateAccessPolicyResponse represents the CreateAccessPolicyResponse schema from the OpenAPI specification
type CreateAccessPolicyResponse struct {
	Accesspolicydetail interface{} `json:"accessPolicyDetail,omitempty"`
}

// CreateSecurityConfigResponse represents the CreateSecurityConfigResponse schema from the OpenAPI specification
type CreateSecurityConfigResponse struct {
	Securityconfigdetail interface{} `json:"securityConfigDetail,omitempty"`
}

// DeleteSecurityPolicyRequest represents the DeleteSecurityPolicyRequest schema from the OpenAPI specification
type DeleteSecurityPolicyRequest struct {
	Name interface{} `json:"name"`
	TypeField interface{} `json:"type"`
	Clienttoken interface{} `json:"clientToken,omitempty"`
}

// BatchGetCollectionResponse represents the BatchGetCollectionResponse schema from the OpenAPI specification
type BatchGetCollectionResponse struct {
	Collectiondetails interface{} `json:"collectionDetails,omitempty"`
	Collectionerrordetails interface{} `json:"collectionErrorDetails,omitempty"`
}

// CapacityLimits represents the CapacityLimits schema from the OpenAPI specification
type CapacityLimits struct {
	Maxindexingcapacityinocu interface{} `json:"maxIndexingCapacityInOCU,omitempty"`
	Maxsearchcapacityinocu interface{} `json:"maxSearchCapacityInOCU,omitempty"`
}

// CreateVpcEndpointDetail represents the CreateVpcEndpointDetail schema from the OpenAPI specification
type CreateVpcEndpointDetail struct {
	Id interface{} `json:"id,omitempty"`
	Name interface{} `json:"name,omitempty"`
	Status interface{} `json:"status,omitempty"`
}

// UpdateCollectionRequest represents the UpdateCollectionRequest schema from the OpenAPI specification
type UpdateCollectionRequest struct {
	Id interface{} `json:"id"`
	Clienttoken interface{} `json:"clientToken,omitempty"`
	Description interface{} `json:"description,omitempty"`
}

// CreateAccessPolicyRequest represents the CreateAccessPolicyRequest schema from the OpenAPI specification
type CreateAccessPolicyRequest struct {
	Name interface{} `json:"name"`
	Policy interface{} `json:"policy"`
	TypeField interface{} `json:"type"`
	Clienttoken interface{} `json:"clientToken,omitempty"`
	Description interface{} `json:"description,omitempty"`
}

// DeleteVpcEndpointResponse represents the DeleteVpcEndpointResponse schema from the OpenAPI specification
type DeleteVpcEndpointResponse struct {
	Deletevpcendpointdetail interface{} `json:"deleteVpcEndpointDetail,omitempty"`
}

// CollectionFilters represents the CollectionFilters schema from the OpenAPI specification
type CollectionFilters struct {
	Status interface{} `json:"status,omitempty"`
	Name interface{} `json:"name,omitempty"`
}

// VpcEndpointDetail represents the VpcEndpointDetail schema from the OpenAPI specification
type VpcEndpointDetail struct {
	Securitygroupids interface{} `json:"securityGroupIds,omitempty"`
	Status interface{} `json:"status,omitempty"`
	Subnetids interface{} `json:"subnetIds,omitempty"`
	Vpcid interface{} `json:"vpcId,omitempty"`
	Createddate interface{} `json:"createdDate,omitempty"`
	Id interface{} `json:"id,omitempty"`
	Name interface{} `json:"name,omitempty"`
}

// ListAccessPoliciesResponse represents the ListAccessPoliciesResponse schema from the OpenAPI specification
type ListAccessPoliciesResponse struct {
	Accesspolicysummaries interface{} `json:"accessPolicySummaries,omitempty"`
	Nexttoken interface{} `json:"nextToken,omitempty"`
}

// BatchGetVpcEndpointRequest represents the BatchGetVpcEndpointRequest schema from the OpenAPI specification
type BatchGetVpcEndpointRequest struct {
	Ids interface{} `json:"ids"`
}

// ListSecurityPoliciesResponse represents the ListSecurityPoliciesResponse schema from the OpenAPI specification
type ListSecurityPoliciesResponse struct {
	Securitypolicysummaries interface{} `json:"securityPolicySummaries,omitempty"`
	Nexttoken interface{} `json:"nextToken,omitempty"`
}

// UpdateAccessPolicyResponse represents the UpdateAccessPolicyResponse schema from the OpenAPI specification
type UpdateAccessPolicyResponse struct {
	Accesspolicydetail interface{} `json:"accessPolicyDetail,omitempty"`
}

// ListVpcEndpointsResponse represents the ListVpcEndpointsResponse schema from the OpenAPI specification
type ListVpcEndpointsResponse struct {
	Nexttoken interface{} `json:"nextToken,omitempty"`
	Vpcendpointsummaries interface{} `json:"vpcEndpointSummaries,omitempty"`
}

// ListCollectionsRequest represents the ListCollectionsRequest schema from the OpenAPI specification
type ListCollectionsRequest struct {
	Collectionfilters interface{} `json:"collectionFilters,omitempty"`
	Maxresults interface{} `json:"maxResults,omitempty"`
	Nexttoken interface{} `json:"nextToken,omitempty"`
}

// GetSecurityConfigRequest represents the GetSecurityConfigRequest schema from the OpenAPI specification
type GetSecurityConfigRequest struct {
	Id interface{} `json:"id"`
}

// GetAccessPolicyResponse represents the GetAccessPolicyResponse schema from the OpenAPI specification
type GetAccessPolicyResponse struct {
	Accesspolicydetail interface{} `json:"accessPolicyDetail,omitempty"`
}

// TagResourceRequest represents the TagResourceRequest schema from the OpenAPI specification
type TagResourceRequest struct {
	Tags interface{} `json:"tags"`
	Resourcearn interface{} `json:"resourceArn"`
}

// UntagResourceRequest represents the UntagResourceRequest schema from the OpenAPI specification
type UntagResourceRequest struct {
	Resourcearn interface{} `json:"resourceArn"`
	Tagkeys interface{} `json:"tagKeys"`
}

// BatchGetCollectionRequest represents the BatchGetCollectionRequest schema from the OpenAPI specification
type BatchGetCollectionRequest struct {
	Names interface{} `json:"names,omitempty"`
	Ids interface{} `json:"ids,omitempty"`
}

// SecurityPolicyDetail represents the SecurityPolicyDetail schema from the OpenAPI specification
type SecurityPolicyDetail struct {
	Name interface{} `json:"name,omitempty"`
	Policy interface{} `json:"policy,omitempty"`
	Policyversion interface{} `json:"policyVersion,omitempty"`
	TypeField interface{} `json:"type,omitempty"`
	Createddate interface{} `json:"createdDate,omitempty"`
	Description interface{} `json:"description,omitempty"`
	Lastmodifieddate interface{} `json:"lastModifiedDate,omitempty"`
}

// UpdateVpcEndpointDetail represents the UpdateVpcEndpointDetail schema from the OpenAPI specification
type UpdateVpcEndpointDetail struct {
	Subnetids interface{} `json:"subnetIds,omitempty"`
	Id interface{} `json:"id,omitempty"`
	Lastmodifieddate interface{} `json:"lastModifiedDate,omitempty"`
	Name interface{} `json:"name,omitempty"`
	Securitygroupids interface{} `json:"securityGroupIds,omitempty"`
	Status interface{} `json:"status,omitempty"`
}

// ListSecurityConfigsRequest represents the ListSecurityConfigsRequest schema from the OpenAPI specification
type ListSecurityConfigsRequest struct {
	Maxresults interface{} `json:"maxResults,omitempty"`
	Nexttoken interface{} `json:"nextToken,omitempty"`
	TypeField interface{} `json:"type"`
}

// CollectionSummary represents the CollectionSummary schema from the OpenAPI specification
type CollectionSummary struct {
	Arn interface{} `json:"arn,omitempty"`
	Id interface{} `json:"id,omitempty"`
	Name interface{} `json:"name,omitempty"`
	Status interface{} `json:"status,omitempty"`
}

// DeleteSecurityConfigResponse represents the DeleteSecurityConfigResponse schema from the OpenAPI specification
type DeleteSecurityConfigResponse struct {
}

// GetAccessPolicyRequest represents the GetAccessPolicyRequest schema from the OpenAPI specification
type GetAccessPolicyRequest struct {
	Name interface{} `json:"name"`
	TypeField interface{} `json:"type"`
}

// ListSecurityConfigsResponse represents the ListSecurityConfigsResponse schema from the OpenAPI specification
type ListSecurityConfigsResponse struct {
	Nexttoken interface{} `json:"nextToken,omitempty"`
	Securityconfigsummaries interface{} `json:"securityConfigSummaries,omitempty"`
}

// SecurityConfigStats represents the SecurityConfigStats schema from the OpenAPI specification
type SecurityConfigStats struct {
	Samlconfigcount interface{} `json:"SamlConfigCount,omitempty"`
}

// UpdateSecurityPolicyResponse represents the UpdateSecurityPolicyResponse schema from the OpenAPI specification
type UpdateSecurityPolicyResponse struct {
	Securitypolicydetail interface{} `json:"securityPolicyDetail,omitempty"`
}

// UpdateSecurityConfigRequest represents the UpdateSecurityConfigRequest schema from the OpenAPI specification
type UpdateSecurityConfigRequest struct {
	Configversion interface{} `json:"configVersion"`
	Description interface{} `json:"description,omitempty"`
	Id interface{} `json:"id"`
	Samloptions interface{} `json:"samlOptions,omitempty"`
	Clienttoken interface{} `json:"clientToken,omitempty"`
}

// VpcEndpointSummary represents the VpcEndpointSummary schema from the OpenAPI specification
type VpcEndpointSummary struct {
	Id interface{} `json:"id,omitempty"`
	Name interface{} `json:"name,omitempty"`
	Status interface{} `json:"status,omitempty"`
}

// CreateSecurityConfigRequest represents the CreateSecurityConfigRequest schema from the OpenAPI specification
type CreateSecurityConfigRequest struct {
	Description interface{} `json:"description,omitempty"`
	Name interface{} `json:"name"`
	Samloptions interface{} `json:"samlOptions,omitempty"`
	TypeField interface{} `json:"type"`
	Clienttoken interface{} `json:"clientToken,omitempty"`
}

// GetSecurityPolicyResponse represents the GetSecurityPolicyResponse schema from the OpenAPI specification
type GetSecurityPolicyResponse struct {
	Securitypolicydetail interface{} `json:"securityPolicyDetail,omitempty"`
}

// UpdateAccessPolicyRequest represents the UpdateAccessPolicyRequest schema from the OpenAPI specification
type UpdateAccessPolicyRequest struct {
	Name interface{} `json:"name"`
	Policy interface{} `json:"policy,omitempty"`
	Policyversion interface{} `json:"policyVersion"`
	TypeField interface{} `json:"type"`
	Clienttoken interface{} `json:"clientToken,omitempty"`
	Description interface{} `json:"description,omitempty"`
}

// CreateCollectionDetail represents the CreateCollectionDetail schema from the OpenAPI specification
type CreateCollectionDetail struct {
	TypeField interface{} `json:"type,omitempty"`
	Arn interface{} `json:"arn,omitempty"`
	Description interface{} `json:"description,omitempty"`
	Kmskeyarn interface{} `json:"kmsKeyArn,omitempty"`
	Createddate interface{} `json:"createdDate,omitempty"`
	Id interface{} `json:"id,omitempty"`
	Lastmodifieddate interface{} `json:"lastModifiedDate,omitempty"`
	Name interface{} `json:"name,omitempty"`
	Status interface{} `json:"status,omitempty"`
}

// DeleteCollectionRequest represents the DeleteCollectionRequest schema from the OpenAPI specification
type DeleteCollectionRequest struct {
	Clienttoken interface{} `json:"clientToken,omitempty"`
	Id interface{} `json:"id"`
}

// DeleteVpcEndpointDetail represents the DeleteVpcEndpointDetail schema from the OpenAPI specification
type DeleteVpcEndpointDetail struct {
	Id interface{} `json:"id,omitempty"`
	Name interface{} `json:"name,omitempty"`
	Status interface{} `json:"status,omitempty"`
}

// SecurityPolicyStats represents the SecurityPolicyStats schema from the OpenAPI specification
type SecurityPolicyStats struct {
	Encryptionpolicycount interface{} `json:"EncryptionPolicyCount,omitempty"`
	Networkpolicycount interface{} `json:"NetworkPolicyCount,omitempty"`
}

// AccessPolicyStats represents the AccessPolicyStats schema from the OpenAPI specification
type AccessPolicyStats struct {
	Datapolicycount interface{} `json:"DataPolicyCount,omitempty"`
}

// VpcEndpointErrorDetail represents the VpcEndpointErrorDetail schema from the OpenAPI specification
type VpcEndpointErrorDetail struct {
	Id interface{} `json:"id,omitempty"`
	Errorcode interface{} `json:"errorCode,omitempty"`
	Errormessage interface{} `json:"errorMessage,omitempty"`
}

// ListCollectionsResponse represents the ListCollectionsResponse schema from the OpenAPI specification
type ListCollectionsResponse struct {
	Nexttoken interface{} `json:"nextToken,omitempty"`
	Collectionsummaries interface{} `json:"collectionSummaries,omitempty"`
}

// GetPoliciesStatsRequest represents the GetPoliciesStatsRequest schema from the OpenAPI specification
type GetPoliciesStatsRequest struct {
}

// DeleteSecurityPolicyResponse represents the DeleteSecurityPolicyResponse schema from the OpenAPI specification
type DeleteSecurityPolicyResponse struct {
}

// ListVpcEndpointsRequest represents the ListVpcEndpointsRequest schema from the OpenAPI specification
type ListVpcEndpointsRequest struct {
	Maxresults interface{} `json:"maxResults,omitempty"`
	Nexttoken interface{} `json:"nextToken,omitempty"`
	Vpcendpointfilters interface{} `json:"vpcEndpointFilters,omitempty"`
}

// GetAccountSettingsRequest represents the GetAccountSettingsRequest schema from the OpenAPI specification
type GetAccountSettingsRequest struct {
}

// BatchGetVpcEndpointResponse represents the BatchGetVpcEndpointResponse schema from the OpenAPI specification
type BatchGetVpcEndpointResponse struct {
	Vpcendpointdetails interface{} `json:"vpcEndpointDetails,omitempty"`
	Vpcendpointerrordetails interface{} `json:"vpcEndpointErrorDetails,omitempty"`
}

// ListAccessPoliciesRequest represents the ListAccessPoliciesRequest schema from the OpenAPI specification
type ListAccessPoliciesRequest struct {
	Maxresults interface{} `json:"maxResults,omitempty"`
	Nexttoken interface{} `json:"nextToken,omitempty"`
	Resource interface{} `json:"resource,omitempty"`
	TypeField interface{} `json:"type"`
}

// AccessPolicySummary represents the AccessPolicySummary schema from the OpenAPI specification
type AccessPolicySummary struct {
	TypeField interface{} `json:"type,omitempty"`
	Createddate interface{} `json:"createdDate,omitempty"`
	Description interface{} `json:"description,omitempty"`
	Lastmodifieddate interface{} `json:"lastModifiedDate,omitempty"`
	Name interface{} `json:"name,omitempty"`
	Policyversion interface{} `json:"policyVersion,omitempty"`
}

// UpdateAccountSettingsResponse represents the UpdateAccountSettingsResponse schema from the OpenAPI specification
type UpdateAccountSettingsResponse struct {
	Accountsettingsdetail interface{} `json:"accountSettingsDetail,omitempty"`
}

// GetAccountSettingsResponse represents the GetAccountSettingsResponse schema from the OpenAPI specification
type GetAccountSettingsResponse struct {
	Accountsettingsdetail interface{} `json:"accountSettingsDetail,omitempty"`
}

// VpcEndpointFilters represents the VpcEndpointFilters schema from the OpenAPI specification
type VpcEndpointFilters struct {
	Status interface{} `json:"status,omitempty"`
}

// DeleteCollectionDetail represents the DeleteCollectionDetail schema from the OpenAPI specification
type DeleteCollectionDetail struct {
	Id interface{} `json:"id,omitempty"`
	Name interface{} `json:"name,omitempty"`
	Status interface{} `json:"status,omitempty"`
}

// UpdateSecurityPolicyRequest represents the UpdateSecurityPolicyRequest schema from the OpenAPI specification
type UpdateSecurityPolicyRequest struct {
	Clienttoken interface{} `json:"clientToken,omitempty"`
	Description interface{} `json:"description,omitempty"`
	Name interface{} `json:"name"`
	Policy interface{} `json:"policy,omitempty"`
	Policyversion interface{} `json:"policyVersion"`
	TypeField interface{} `json:"type"`
}

// GetSecurityPolicyRequest represents the GetSecurityPolicyRequest schema from the OpenAPI specification
type GetSecurityPolicyRequest struct {
	Name interface{} `json:"name"`
	TypeField interface{} `json:"type"`
}

// UpdateCollectionResponse represents the UpdateCollectionResponse schema from the OpenAPI specification
type UpdateCollectionResponse struct {
	Updatecollectiondetail interface{} `json:"updateCollectionDetail,omitempty"`
}

// GetSecurityConfigResponse represents the GetSecurityConfigResponse schema from the OpenAPI specification
type GetSecurityConfigResponse struct {
	Securityconfigdetail interface{} `json:"securityConfigDetail,omitempty"`
}

// SamlConfigOptions represents the SamlConfigOptions schema from the OpenAPI specification
type SamlConfigOptions struct {
	Userattribute interface{} `json:"userAttribute,omitempty"`
	Groupattribute interface{} `json:"groupAttribute,omitempty"`
	Metadata interface{} `json:"metadata"`
	Sessiontimeout interface{} `json:"sessionTimeout,omitempty"`
}

// CreateVpcEndpointRequest represents the CreateVpcEndpointRequest schema from the OpenAPI specification
type CreateVpcEndpointRequest struct {
	Name interface{} `json:"name"`
	Securitygroupids interface{} `json:"securityGroupIds,omitempty"`
	Subnetids interface{} `json:"subnetIds"`
	Vpcid interface{} `json:"vpcId"`
	Clienttoken interface{} `json:"clientToken,omitempty"`
}

// TagResourceResponse represents the TagResourceResponse schema from the OpenAPI specification
type TagResourceResponse struct {
}

// GetPoliciesStatsResponse represents the GetPoliciesStatsResponse schema from the OpenAPI specification
type GetPoliciesStatsResponse struct {
	Securitypolicystats interface{} `json:"SecurityPolicyStats,omitempty"`
	Totalpolicycount interface{} `json:"TotalPolicyCount,omitempty"`
	Accesspolicystats interface{} `json:"AccessPolicyStats,omitempty"`
	Securityconfigstats interface{} `json:"SecurityConfigStats,omitempty"`
}

// DeleteAccessPolicyRequest represents the DeleteAccessPolicyRequest schema from the OpenAPI specification
type DeleteAccessPolicyRequest struct {
	Clienttoken interface{} `json:"clientToken,omitempty"`
	Name interface{} `json:"name"`
	TypeField interface{} `json:"type"`
}

// Document represents the Document schema from the OpenAPI specification
type Document struct {
}

// ListTagsForResourceResponse represents the ListTagsForResourceResponse schema from the OpenAPI specification
type ListTagsForResourceResponse struct {
	Tags interface{} `json:"tags,omitempty"`
}

// DeleteVpcEndpointRequest represents the DeleteVpcEndpointRequest schema from the OpenAPI specification
type DeleteVpcEndpointRequest struct {
	Id interface{} `json:"id"`
	Clienttoken interface{} `json:"clientToken,omitempty"`
}

// UpdateVpcEndpointRequest represents the UpdateVpcEndpointRequest schema from the OpenAPI specification
type UpdateVpcEndpointRequest struct {
	Clienttoken interface{} `json:"clientToken,omitempty"`
	Id interface{} `json:"id"`
	Removesecuritygroupids interface{} `json:"removeSecurityGroupIds,omitempty"`
	Removesubnetids interface{} `json:"removeSubnetIds,omitempty"`
	Addsecuritygroupids interface{} `json:"addSecurityGroupIds,omitempty"`
	Addsubnetids interface{} `json:"addSubnetIds,omitempty"`
}

// UpdateCollectionDetail represents the UpdateCollectionDetail schema from the OpenAPI specification
type UpdateCollectionDetail struct {
	TypeField interface{} `json:"type,omitempty"`
	Arn interface{} `json:"arn,omitempty"`
	Createddate interface{} `json:"createdDate,omitempty"`
	Description interface{} `json:"description,omitempty"`
	Id interface{} `json:"id,omitempty"`
	Lastmodifieddate interface{} `json:"lastModifiedDate,omitempty"`
	Name interface{} `json:"name,omitempty"`
	Status interface{} `json:"status,omitempty"`
}

// CollectionDetail represents the CollectionDetail schema from the OpenAPI specification
type CollectionDetail struct {
	Kmskeyarn interface{} `json:"kmsKeyArn,omitempty"`
	Name interface{} `json:"name,omitempty"`
	Createddate interface{} `json:"createdDate,omitempty"`
	Description interface{} `json:"description,omitempty"`
	TypeField interface{} `json:"type,omitempty"`
	Arn interface{} `json:"arn,omitempty"`
	Id interface{} `json:"id,omitempty"`
	Lastmodifieddate interface{} `json:"lastModifiedDate,omitempty"`
	Status interface{} `json:"status,omitempty"`
	Collectionendpoint interface{} `json:"collectionEndpoint,omitempty"`
	Dashboardendpoint interface{} `json:"dashboardEndpoint,omitempty"`
}

// SecurityPolicySummary represents the SecurityPolicySummary schema from the OpenAPI specification
type SecurityPolicySummary struct {
	Lastmodifieddate interface{} `json:"lastModifiedDate,omitempty"`
	Name interface{} `json:"name,omitempty"`
	Policyversion interface{} `json:"policyVersion,omitempty"`
	TypeField interface{} `json:"type,omitempty"`
	Createddate interface{} `json:"createdDate,omitempty"`
	Description interface{} `json:"description,omitempty"`
}

// ListSecurityPoliciesRequest represents the ListSecurityPoliciesRequest schema from the OpenAPI specification
type ListSecurityPoliciesRequest struct {
	TypeField interface{} `json:"type"`
	Maxresults interface{} `json:"maxResults,omitempty"`
	Nexttoken interface{} `json:"nextToken,omitempty"`
	Resource interface{} `json:"resource,omitempty"`
}

// ListTagsForResourceRequest represents the ListTagsForResourceRequest schema from the OpenAPI specification
type ListTagsForResourceRequest struct {
	Resourcearn interface{} `json:"resourceArn"`
}

// CollectionErrorDetail represents the CollectionErrorDetail schema from the OpenAPI specification
type CollectionErrorDetail struct {
	Errorcode interface{} `json:"errorCode,omitempty"`
	Errormessage interface{} `json:"errorMessage,omitempty"`
	Id interface{} `json:"id,omitempty"`
	Name interface{} `json:"name,omitempty"`
}

// CreateSecurityPolicyRequest represents the CreateSecurityPolicyRequest schema from the OpenAPI specification
type CreateSecurityPolicyRequest struct {
	Clienttoken interface{} `json:"clientToken,omitempty"`
	Description interface{} `json:"description,omitempty"`
	Name interface{} `json:"name"`
	Policy interface{} `json:"policy"`
	TypeField interface{} `json:"type"`
}

// UpdateSecurityConfigResponse represents the UpdateSecurityConfigResponse schema from the OpenAPI specification
type UpdateSecurityConfigResponse struct {
	Securityconfigdetail interface{} `json:"securityConfigDetail,omitempty"`
}

// AccessPolicyDetail represents the AccessPolicyDetail schema from the OpenAPI specification
type AccessPolicyDetail struct {
	Policyversion interface{} `json:"policyVersion,omitempty"`
	TypeField interface{} `json:"type,omitempty"`
	Createddate interface{} `json:"createdDate,omitempty"`
	Description interface{} `json:"description,omitempty"`
	Lastmodifieddate interface{} `json:"lastModifiedDate,omitempty"`
	Name interface{} `json:"name,omitempty"`
	Policy interface{} `json:"policy,omitempty"`
}

// UpdateVpcEndpointResponse represents the UpdateVpcEndpointResponse schema from the OpenAPI specification
type UpdateVpcEndpointResponse struct {
	Updatevpcendpointdetail interface{} `json:"UpdateVpcEndpointDetail,omitempty"`
}

// SecurityConfigSummary represents the SecurityConfigSummary schema from the OpenAPI specification
type SecurityConfigSummary struct {
	TypeField interface{} `json:"type,omitempty"`
	Configversion interface{} `json:"configVersion,omitempty"`
	Createddate interface{} `json:"createdDate,omitempty"`
	Description interface{} `json:"description,omitempty"`
	Id interface{} `json:"id,omitempty"`
	Lastmodifieddate interface{} `json:"lastModifiedDate,omitempty"`
}

// DeleteSecurityConfigRequest represents the DeleteSecurityConfigRequest schema from the OpenAPI specification
type DeleteSecurityConfigRequest struct {
	Clienttoken interface{} `json:"clientToken,omitempty"`
	Id interface{} `json:"id"`
}

// CreateCollectionRequest represents the CreateCollectionRequest schema from the OpenAPI specification
type CreateCollectionRequest struct {
	Description interface{} `json:"description,omitempty"`
	Name interface{} `json:"name"`
	Tags interface{} `json:"tags,omitempty"`
	TypeField interface{} `json:"type,omitempty"`
	Clienttoken interface{} `json:"clientToken,omitempty"`
}

// AccountSettingsDetail represents the AccountSettingsDetail schema from the OpenAPI specification
type AccountSettingsDetail struct {
	Capacitylimits CapacityLimits `json:"capacityLimits,omitempty"` // The maximum capacity limits for all OpenSearch Serverless collections, in OpenSearch Compute Units (OCUs). These limits are used to scale your collections based on the current workload. For more information, see <a href="https://docs.aws.amazon.com/opensearch-service/latest/developerguide/serverless-scaling.html">Managing capacity limits for Amazon OpenSearch Serverless</a>.
}

// DeleteAccessPolicyResponse represents the DeleteAccessPolicyResponse schema from the OpenAPI specification
type DeleteAccessPolicyResponse struct {
}

// UpdateAccountSettingsRequest represents the UpdateAccountSettingsRequest schema from the OpenAPI specification
type UpdateAccountSettingsRequest struct {
	Capacitylimits CapacityLimits `json:"capacityLimits,omitempty"` // The maximum capacity limits for all OpenSearch Serverless collections, in OpenSearch Compute Units (OCUs). These limits are used to scale your collections based on the current workload. For more information, see <a href="https://docs.aws.amazon.com/opensearch-service/latest/developerguide/serverless-scaling.html">Managing capacity limits for Amazon OpenSearch Serverless</a>.
}

// SecurityConfigDetail represents the SecurityConfigDetail schema from the OpenAPI specification
type SecurityConfigDetail struct {
	Createddate interface{} `json:"createdDate,omitempty"`
	Description interface{} `json:"description,omitempty"`
	Id interface{} `json:"id,omitempty"`
	Lastmodifieddate interface{} `json:"lastModifiedDate,omitempty"`
	Samloptions interface{} `json:"samlOptions,omitempty"`
	TypeField interface{} `json:"type,omitempty"`
	Configversion interface{} `json:"configVersion,omitempty"`
}

// CreateCollectionResponse represents the CreateCollectionResponse schema from the OpenAPI specification
type CreateCollectionResponse struct {
	Createcollectiondetail interface{} `json:"createCollectionDetail,omitempty"`
}
