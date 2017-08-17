package resources

import (
	"encoding/json"
	"errors"
	"fmt"
)

// AWSWAFRegionalByteMatchSet AWS CloudFormation Resource (AWS::WAFRegional::ByteMatchSet)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wafregional-bytematchset.html
type AWSWAFRegionalByteMatchSet struct {

	// ByteMatchTuples AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wafregional-bytematchset.html#cfn-wafregional-bytematchset-bytematchtuples
	ByteMatchTuples []AWSWAFRegionalByteMatchSet_ByteMatchTuple `json:"ByteMatchTuples,omitempty"`

	// Name AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wafregional-bytematchset.html#cfn-wafregional-bytematchset-name
	Name string `json:"Name,omitempty"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSWAFRegionalByteMatchSet) AWSCloudFormationType() string {
	return "AWS::WAFRegional::ByteMatchSet"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSWAFRegionalByteMatchSet) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// MarshalJSON is a custom JSON marshalling hook that embeds this object into
// an AWS CloudFormation JSON resource's 'Properties' field and adds a 'Type'.
func (r *AWSWAFRegionalByteMatchSet) MarshalJSON() ([]byte, error) {
	type Properties AWSWAFRegionalByteMatchSet
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
func (r *AWSWAFRegionalByteMatchSet) UnmarshalJSON(b []byte) error {
	type Properties AWSWAFRegionalByteMatchSet
	res := &struct {
		Type       string
		Properties *Properties
	}{}
	if err := json.Unmarshal(b, &res); err != nil {
		fmt.Printf("ERROR: %s\n", err)
		return err
	}
	*r = AWSWAFRegionalByteMatchSet(*res.Properties)
	return nil
}

// GetAllAWSWAFRegionalByteMatchSetResources retrieves all AWSWAFRegionalByteMatchSet items from an AWS CloudFormation template
func (t *CloudFormationTemplate) GetAllAWSWAFRegionalByteMatchSetResources() map[string]AWSWAFRegionalByteMatchSet {
	results := map[string]AWSWAFRegionalByteMatchSet{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case AWSWAFRegionalByteMatchSet:
			// We found a strongly typed resource of the correct type; use it
			results[name] = resource
		case map[string]interface{}:
			// We found an untyped resource (likely from JSON) which *might* be
			// the correct type, but we need to check it's 'Type' field
			if resType, ok := resource["Type"]; ok {
				if resType == "AWS::WAFRegional::ByteMatchSet" {
					// The resource is correct, unmarshal it into the results
					if b, err := json.Marshal(resource); err == nil {
						var result AWSWAFRegionalByteMatchSet
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

// GetAWSWAFRegionalByteMatchSetWithName retrieves all AWSWAFRegionalByteMatchSet items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *CloudFormationTemplate) GetAWSWAFRegionalByteMatchSetWithName(name string) (AWSWAFRegionalByteMatchSet, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case AWSWAFRegionalByteMatchSet:
			// We found a strongly typed resource of the correct type; use it
			return resource, nil
		case map[string]interface{}:
			// We found an untyped resource (likely from JSON) which *might* be
			// the correct type, but we need to check it's 'Type' field
			if resType, ok := resource["Type"]; ok {
				if resType == "AWS::WAFRegional::ByteMatchSet" {
					// The resource is correct, unmarshal it into the results
					if b, err := json.Marshal(resource); err == nil {
						var result AWSWAFRegionalByteMatchSet
						if err := json.Unmarshal(b, &result); err == nil {
							return result, nil
						}
					}
				}
			}
		}
	}
	return AWSWAFRegionalByteMatchSet{}, errors.New("resource not found")
}
