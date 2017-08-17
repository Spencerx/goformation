package resources

import (
	"encoding/json"
	"errors"
	"fmt"
)

// AWSECRRepository AWS CloudFormation Resource (AWS::ECR::Repository)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecr-repository.html
type AWSECRRepository struct {

	// RepositoryName AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecr-repository.html#cfn-ecr-repository-repositoryname
	RepositoryName string `json:"RepositoryName,omitempty"`

	// RepositoryPolicyText AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecr-repository.html#cfn-ecr-repository-repositorypolicytext
	RepositoryPolicyText interface{} `json:"RepositoryPolicyText,omitempty"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSECRRepository) AWSCloudFormationType() string {
	return "AWS::ECR::Repository"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSECRRepository) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// MarshalJSON is a custom JSON marshalling hook that embeds this object into
// an AWS CloudFormation JSON resource's 'Properties' field and adds a 'Type'.
func (r *AWSECRRepository) MarshalJSON() ([]byte, error) {
	type Properties AWSECRRepository
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
func (r *AWSECRRepository) UnmarshalJSON(b []byte) error {
	type Properties AWSECRRepository
	res := &struct {
		Type       string
		Properties *Properties
	}{}
	if err := json.Unmarshal(b, &res); err != nil {
		fmt.Printf("ERROR: %s\n", err)
		return err
	}
	*r = AWSECRRepository(*res.Properties)
	return nil
}

// GetAllAWSECRRepositoryResources retrieves all AWSECRRepository items from an AWS CloudFormation template
func (t *CloudFormationTemplate) GetAllAWSECRRepositoryResources() map[string]AWSECRRepository {
	results := map[string]AWSECRRepository{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case AWSECRRepository:
			// We found a strongly typed resource of the correct type; use it
			results[name] = resource
		case map[string]interface{}:
			// We found an untyped resource (likely from JSON) which *might* be
			// the correct type, but we need to check it's 'Type' field
			if resType, ok := resource["Type"]; ok {
				if resType == "AWS::ECR::Repository" {
					// The resource is correct, unmarshal it into the results
					if b, err := json.Marshal(resource); err == nil {
						var result AWSECRRepository
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

// GetAWSECRRepositoryWithName retrieves all AWSECRRepository items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *CloudFormationTemplate) GetAWSECRRepositoryWithName(name string) (AWSECRRepository, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case AWSECRRepository:
			// We found a strongly typed resource of the correct type; use it
			return resource, nil
		case map[string]interface{}:
			// We found an untyped resource (likely from JSON) which *might* be
			// the correct type, but we need to check it's 'Type' field
			if resType, ok := resource["Type"]; ok {
				if resType == "AWS::ECR::Repository" {
					// The resource is correct, unmarshal it into the results
					if b, err := json.Marshal(resource); err == nil {
						var result AWSECRRepository
						if err := json.Unmarshal(b, &result); err == nil {
							return result, nil
						}
					}
				}
			}
		}
	}
	return AWSECRRepository{}, errors.New("resource not found")
}
