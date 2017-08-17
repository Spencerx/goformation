package resources

import (
	"encoding/json"
	"errors"
	"fmt"
)

// AWSEC2VPNGatewayRoutePropagation AWS CloudFormation Resource (AWS::EC2::VPNGatewayRoutePropagation)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-vpn-gatewayrouteprop.html
type AWSEC2VPNGatewayRoutePropagation struct {

	// RouteTableIds AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-vpn-gatewayrouteprop.html#cfn-ec2-vpngatewayrouteprop-routetableids
	RouteTableIds []string `json:"RouteTableIds,omitempty"`

	// VpnGatewayId AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-vpn-gatewayrouteprop.html#cfn-ec2-vpngatewayrouteprop-vpngatewayid
	VpnGatewayId string `json:"VpnGatewayId,omitempty"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSEC2VPNGatewayRoutePropagation) AWSCloudFormationType() string {
	return "AWS::EC2::VPNGatewayRoutePropagation"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSEC2VPNGatewayRoutePropagation) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// MarshalJSON is a custom JSON marshalling hook that embeds this object into
// an AWS CloudFormation JSON resource's 'Properties' field and adds a 'Type'.
func (r *AWSEC2VPNGatewayRoutePropagation) MarshalJSON() ([]byte, error) {
	type Properties AWSEC2VPNGatewayRoutePropagation
	return json.Marshal(&struct {
		Type       string
		Properties Properties
	}{
		Type:       r.AWSCloudFormationType(),
		Properties: (Properties)(*r),
	})
}

// UnmarshalJSON is a custom JSON unmarshalling hook that strips the outer
// AWS CloudFormation resource object, and just keeps the 'Properties' field.
func (r *AWSEC2VPNGatewayRoutePropagation) UnmarshalJSON(b []byte) error {
	type Properties AWSEC2VPNGatewayRoutePropagation
	res := &struct {
		Type       string
		Properties *Properties
	}{}
	if err := json.Unmarshal(b, &res); err != nil {
		fmt.Printf("ERROR: %s\n", err)
		return err
	}
	*r = AWSEC2VPNGatewayRoutePropagation(*res.Properties)
	return nil
}

// GetAllAWSEC2VPNGatewayRoutePropagationResources retrieves all AWSEC2VPNGatewayRoutePropagation items from an AWS CloudFormation template
func (t *CloudFormationTemplate) GetAllAWSEC2VPNGatewayRoutePropagationResources() map[string]AWSEC2VPNGatewayRoutePropagation {
	results := map[string]AWSEC2VPNGatewayRoutePropagation{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case AWSEC2VPNGatewayRoutePropagation:
			// We found a strongly typed resource of the correct type; use it
			results[name] = resource
		case map[string]interface{}:
			// We found an untyped resource (likely from JSON) which *might* be
			// the correct type, but we need to check it's 'Type' field
			if resType, ok := resource["Type"]; ok {
				if resType == "AWS::EC2::VPNGatewayRoutePropagation" {
					// The resource is correct, unmarshal it into the results
					if b, err := json.Marshal(resource); err == nil {
						var result AWSEC2VPNGatewayRoutePropagation
						if err := json.Unmarshal(b, &result); err == nil {
							results[name] = result
						}
					}
				}
			}
		}
	}
	return results
}

// GetAWSEC2VPNGatewayRoutePropagationWithName retrieves all AWSEC2VPNGatewayRoutePropagation items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *CloudFormationTemplate) GetAWSEC2VPNGatewayRoutePropagationWithName(name string) (AWSEC2VPNGatewayRoutePropagation, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case AWSEC2VPNGatewayRoutePropagation:
			// We found a strongly typed resource of the correct type; use it
			return resource, nil
		case map[string]interface{}:
			// We found an untyped resource (likely from JSON) which *might* be
			// the correct type, but we need to check it's 'Type' field
			if resType, ok := resource["Type"]; ok {
				if resType == "AWS::EC2::VPNGatewayRoutePropagation" {
					// The resource is correct, unmarshal it into the results
					if b, err := json.Marshal(resource); err == nil {
						var result AWSEC2VPNGatewayRoutePropagation
						if err := json.Unmarshal(b, &result); err == nil {
							return result, nil
						}
					}
				}
			}
		}
	}
	return AWSEC2VPNGatewayRoutePropagation{}, errors.New("resource not found")
}
