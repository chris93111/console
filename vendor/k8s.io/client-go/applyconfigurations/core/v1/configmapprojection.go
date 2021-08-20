/*
Copyright The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1

// ConfigMapProjectionApplyConfiguration represents an declarative configuration of the ConfigMapProjection type for use
// with apply.
type ConfigMapProjectionApplyConfiguration struct {
	LocalObjectReferenceApplyConfiguration `json:",inline"`
	Items                                  []KeyToPathApplyConfiguration `json:"items,omitempty"`
	Optional                               *bool                         `json:"optional,omitempty"`
}

// ConfigMapProjectionApplyConfiguration constructs an declarative configuration of the ConfigMapProjection type for use with
// apply.
func ConfigMapProjection() *ConfigMapProjectionApplyConfiguration {
	return &ConfigMapProjectionApplyConfiguration{}
}

// WithName sets the Name field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Name field is set to the value of the last call.
func (b *ConfigMapProjectionApplyConfiguration) WithName(value string) *ConfigMapProjectionApplyConfiguration {
	b.Name = &value
	return b
}

// WithItems adds the given value to the Items field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Items field.
func (b *ConfigMapProjectionApplyConfiguration) WithItems(values ...*KeyToPathApplyConfiguration) *ConfigMapProjectionApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithItems")
		}
		b.Items = append(b.Items, *values[i])
	}
	return b
}

// WithOptional sets the Optional field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Optional field is set to the value of the last call.
func (b *ConfigMapProjectionApplyConfiguration) WithOptional(value bool) *ConfigMapProjectionApplyConfiguration {
	b.Optional = &value
	return b
}