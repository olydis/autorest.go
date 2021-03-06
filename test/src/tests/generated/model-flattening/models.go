package modelflatteninggroup

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"encoding/json"
	"github.com/Azure/go-autorest/autorest"
)

// ProvisioningStateValues enumerates the values for provisioning state values.
type ProvisioningStateValues string

const (
	// Accepted ...
	Accepted ProvisioningStateValues = "Accepted"
	// Canceled ...
	Canceled ProvisioningStateValues = "canceled"
	// Created ...
	Created ProvisioningStateValues = "Created"
	// Creating ...
	Creating ProvisioningStateValues = "Creating"
	// Deleted ...
	Deleted ProvisioningStateValues = "Deleted"
	// Deleting ...
	Deleting ProvisioningStateValues = "Deleting"
	// Failed ...
	Failed ProvisioningStateValues = "Failed"
	// OK ...
	OK ProvisioningStateValues = "OK"
	// Succeeded ...
	Succeeded ProvisioningStateValues = "Succeeded"
	// Updated ...
	Updated ProvisioningStateValues = "Updated"
	// Updating ...
	Updating ProvisioningStateValues = "Updating"
)

// PossibleProvisioningStateValuesValues returns an array of possible values for the ProvisioningStateValues const type.
func PossibleProvisioningStateValuesValues() []ProvisioningStateValues {
	return []ProvisioningStateValues{Accepted, Canceled, Created, Creating, Deleted, Deleting, Failed, OK, Succeeded, Updated, Updating}
}

// BaseProduct the product documentation.
type BaseProduct struct {
	// ProductID - Unique identifier representing a specific product for a given latitude & longitude. For example, uberX in San Francisco will have a different product_id than uberX in Los Angeles.
	ProductID *string `json:"base_product_id,omitempty"`
	// Description - Description of product.
	Description *string `json:"base_product_description,omitempty"`
}

// Error ...
type Error struct {
	Status  *int32  `json:"status,omitempty"`
	Message *string `json:"message,omitempty"`
	*Error  `json:"parentError,omitempty"`
}

// MarshalJSON is the custom marshaler for Error.
func (e Error) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if e.Status != nil {
		objectMap["status"] = e.Status
	}
	if e.Message != nil {
		objectMap["message"] = e.Message
	}
	if e.Error != nil {
		objectMap["parentError"] = e.Error
	}
	return json.Marshal(objectMap)
}

// UnmarshalJSON is the custom unmarshaler for Error struct.
func (e *Error) UnmarshalJSON(body []byte) error {
	var m map[string]*json.RawMessage
	err := json.Unmarshal(body, &m)
	if err != nil {
		return err
	}
	for k, v := range m {
		switch k {
		case "status":
			if v != nil {
				var status int32
				err = json.Unmarshal(*v, &status)
				if err != nil {
					return err
				}
				e.Status = &status
			}
		case "message":
			if v != nil {
				var message string
				err = json.Unmarshal(*v, &message)
				if err != nil {
					return err
				}
				e.Message = &message
			}
		case "parentError":
			if v != nil {
				var errorVar Error
				err = json.Unmarshal(*v, &errorVar)
				if err != nil {
					return err
				}
				e.Error = &errorVar
			}
		}
	}

	return nil
}

// FlattenedProduct flattened product.
type FlattenedProduct struct {
	*FlattenedProductProperties `json:"properties,omitempty"`
	// ID - Resource Id
	ID *string `json:"id,omitempty"`
	// Type - Resource Type
	Type *string            `json:"type,omitempty"`
	Tags map[string]*string `json:"tags"`
	// Location - Resource Location
	Location *string `json:"location,omitempty"`
	// Name - Resource Name
	Name *string `json:"name,omitempty"`
}

// MarshalJSON is the custom marshaler for FlattenedProduct.
func (fp FlattenedProduct) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if fp.FlattenedProductProperties != nil {
		objectMap["properties"] = fp.FlattenedProductProperties
	}
	if fp.ID != nil {
		objectMap["id"] = fp.ID
	}
	if fp.Type != nil {
		objectMap["type"] = fp.Type
	}
	if fp.Tags != nil {
		objectMap["tags"] = fp.Tags
	}
	if fp.Location != nil {
		objectMap["location"] = fp.Location
	}
	if fp.Name != nil {
		objectMap["name"] = fp.Name
	}
	return json.Marshal(objectMap)
}

// UnmarshalJSON is the custom unmarshaler for FlattenedProduct struct.
func (fp *FlattenedProduct) UnmarshalJSON(body []byte) error {
	var m map[string]*json.RawMessage
	err := json.Unmarshal(body, &m)
	if err != nil {
		return err
	}
	for k, v := range m {
		switch k {
		case "properties":
			if v != nil {
				var flattenedProductProperties FlattenedProductProperties
				err = json.Unmarshal(*v, &flattenedProductProperties)
				if err != nil {
					return err
				}
				fp.FlattenedProductProperties = &flattenedProductProperties
			}
		case "id":
			if v != nil {
				var ID string
				err = json.Unmarshal(*v, &ID)
				if err != nil {
					return err
				}
				fp.ID = &ID
			}
		case "type":
			if v != nil {
				var typeVar string
				err = json.Unmarshal(*v, &typeVar)
				if err != nil {
					return err
				}
				fp.Type = &typeVar
			}
		case "tags":
			if v != nil {
				var tags map[string]*string
				err = json.Unmarshal(*v, &tags)
				if err != nil {
					return err
				}
				fp.Tags = tags
			}
		case "location":
			if v != nil {
				var location string
				err = json.Unmarshal(*v, &location)
				if err != nil {
					return err
				}
				fp.Location = &location
			}
		case "name":
			if v != nil {
				var name string
				err = json.Unmarshal(*v, &name)
				if err != nil {
					return err
				}
				fp.Name = &name
			}
		}
	}

	return nil
}

// FlattenedProductProperties ...
type FlattenedProductProperties struct {
	PName *string `json:"p.name,omitempty"`
	Type  *string `json:"type,omitempty"`
	// ProvisioningStateValues - Possible values include: 'Succeeded', 'Failed', 'Canceled', 'Accepted', 'Creating', 'Created', 'Updating', 'Updated', 'Deleting', 'Deleted', 'OK'
	ProvisioningStateValues ProvisioningStateValues `json:"provisioningStateValues,omitempty"`
	ProvisioningState       *string                 `json:"provisioningState,omitempty"`
}

// GenericURL the Generic URL.
type GenericURL struct {
	// GenericValue - Generic URL value.
	GenericValue *string `json:"generic_value,omitempty"`
}

// ListFlattenedProduct ...
type ListFlattenedProduct struct {
	autorest.Response `json:"-"`
	Value             *[]FlattenedProduct `json:"value,omitempty"`
}

// ListProductWrapper ...
type ListProductWrapper struct {
	autorest.Response `json:"-"`
	Value             *[]ProductWrapper `json:"value,omitempty"`
}

// ProductURL the product URL.
type ProductURL struct {
	// OdataValue - URL value.
	OdataValue *string `json:"@odata.value,omitempty"`
	// GenericValue - Generic URL value.
	GenericValue *string `json:"generic_value,omitempty"`
}

// ProductWrapper the wrapped produc.
type ProductWrapper struct {
	*WrappedProduct `json:"property,omitempty"`
}

// MarshalJSON is the custom marshaler for ProductWrapper.
func (pw ProductWrapper) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if pw.WrappedProduct != nil {
		objectMap["property"] = pw.WrappedProduct
	}
	return json.Marshal(objectMap)
}

// UnmarshalJSON is the custom unmarshaler for ProductWrapper struct.
func (pw *ProductWrapper) UnmarshalJSON(body []byte) error {
	var m map[string]*json.RawMessage
	err := json.Unmarshal(body, &m)
	if err != nil {
		return err
	}
	for k, v := range m {
		switch k {
		case "property":
			if v != nil {
				var wrappedProduct WrappedProduct
				err = json.Unmarshal(*v, &wrappedProduct)
				if err != nil {
					return err
				}
				pw.WrappedProduct = &wrappedProduct
			}
		}
	}

	return nil
}

// Resource ...
type Resource struct {
	// ID - Resource Id
	ID *string `json:"id,omitempty"`
	// Type - Resource Type
	Type *string            `json:"type,omitempty"`
	Tags map[string]*string `json:"tags"`
	// Location - Resource Location
	Location *string `json:"location,omitempty"`
	// Name - Resource Name
	Name *string `json:"name,omitempty"`
}

// MarshalJSON is the custom marshaler for Resource.
func (r Resource) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if r.ID != nil {
		objectMap["id"] = r.ID
	}
	if r.Type != nil {
		objectMap["type"] = r.Type
	}
	if r.Tags != nil {
		objectMap["tags"] = r.Tags
	}
	if r.Location != nil {
		objectMap["location"] = r.Location
	}
	if r.Name != nil {
		objectMap["name"] = r.Name
	}
	return json.Marshal(objectMap)
}

// ResourceCollection ...
type ResourceCollection struct {
	autorest.Response     `json:"-"`
	Productresource       *FlattenedProduct            `json:"productresource,omitempty"`
	Arrayofresources      *[]FlattenedProduct          `json:"arrayofresources,omitempty"`
	Dictionaryofresources map[string]*FlattenedProduct `json:"dictionaryofresources"`
}

// MarshalJSON is the custom marshaler for ResourceCollection.
func (rc ResourceCollection) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if rc.Productresource != nil {
		objectMap["productresource"] = rc.Productresource
	}
	if rc.Arrayofresources != nil {
		objectMap["arrayofresources"] = rc.Arrayofresources
	}
	if rc.Dictionaryofresources != nil {
		objectMap["dictionaryofresources"] = rc.Dictionaryofresources
	}
	return json.Marshal(objectMap)
}

// SetFlattenedProduct ...
type SetFlattenedProduct struct {
	autorest.Response `json:"-"`
	Value             map[string]*FlattenedProduct `json:"value"`
}

// MarshalJSON is the custom marshaler for SetFlattenedProduct.
func (sfp SetFlattenedProduct) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if sfp.Value != nil {
		objectMap["value"] = sfp.Value
	}
	return json.Marshal(objectMap)
}

// SimpleProduct the product documentation.
type SimpleProduct struct {
	autorest.Response        `json:"-"`
	*SimpleProductProperties `json:"details,omitempty"`
	// ProductID - Unique identifier representing a specific product for a given latitude & longitude. For example, uberX in San Francisco will have a different product_id than uberX in Los Angeles.
	ProductID *string `json:"base_product_id,omitempty"`
	// Description - Description of product.
	Description *string `json:"base_product_description,omitempty"`
}

// MarshalJSON is the custom marshaler for SimpleProduct.
func (sp SimpleProduct) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if sp.SimpleProductProperties != nil {
		objectMap["details"] = sp.SimpleProductProperties
	}
	if sp.ProductID != nil {
		objectMap["base_product_id"] = sp.ProductID
	}
	if sp.Description != nil {
		objectMap["base_product_description"] = sp.Description
	}
	return json.Marshal(objectMap)
}

// UnmarshalJSON is the custom unmarshaler for SimpleProduct struct.
func (sp *SimpleProduct) UnmarshalJSON(body []byte) error {
	var m map[string]*json.RawMessage
	err := json.Unmarshal(body, &m)
	if err != nil {
		return err
	}
	for k, v := range m {
		switch k {
		case "details":
			if v != nil {
				var simpleProductProperties SimpleProductProperties
				err = json.Unmarshal(*v, &simpleProductProperties)
				if err != nil {
					return err
				}
				sp.SimpleProductProperties = &simpleProductProperties
			}
		case "base_product_id":
			if v != nil {
				var productID string
				err = json.Unmarshal(*v, &productID)
				if err != nil {
					return err
				}
				sp.ProductID = &productID
			}
		case "base_product_description":
			if v != nil {
				var description string
				err = json.Unmarshal(*v, &description)
				if err != nil {
					return err
				}
				sp.Description = &description
			}
		}
	}

	return nil
}

// SimpleProductProperties the product documentation.
type SimpleProductProperties struct {
	// MaxProductDisplayName - Display name of product.
	MaxProductDisplayName *string `json:"max_product_display_name,omitempty"`
	// Capacity - Capacity of product. For example, 4 people.
	Capacity    *string `json:"max_product_capacity,omitempty"`
	*ProductURL `json:"max_product_image,omitempty"`
}

// MarshalJSON is the custom marshaler for SimpleProductProperties.
func (spp SimpleProductProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if spp.MaxProductDisplayName != nil {
		objectMap["max_product_display_name"] = spp.MaxProductDisplayName
	}
	if spp.Capacity != nil {
		objectMap["max_product_capacity"] = spp.Capacity
	}
	if spp.ProductURL != nil {
		objectMap["max_product_image"] = spp.ProductURL
	}
	return json.Marshal(objectMap)
}

// UnmarshalJSON is the custom unmarshaler for SimpleProductProperties struct.
func (spp *SimpleProductProperties) UnmarshalJSON(body []byte) error {
	var m map[string]*json.RawMessage
	err := json.Unmarshal(body, &m)
	if err != nil {
		return err
	}
	for k, v := range m {
		switch k {
		case "max_product_display_name":
			if v != nil {
				var maxProductDisplayName string
				err = json.Unmarshal(*v, &maxProductDisplayName)
				if err != nil {
					return err
				}
				spp.MaxProductDisplayName = &maxProductDisplayName
			}
		case "max_product_capacity":
			if v != nil {
				var capacity string
				err = json.Unmarshal(*v, &capacity)
				if err != nil {
					return err
				}
				spp.Capacity = &capacity
			}
		case "max_product_image":
			if v != nil {
				var productURL ProductURL
				err = json.Unmarshal(*v, &productURL)
				if err != nil {
					return err
				}
				spp.ProductURL = &productURL
			}
		}
	}

	return nil
}

// WrappedProduct the wrapped produc.
type WrappedProduct struct {
	// Value - the product value
	Value *string `json:"value,omitempty"`
}
