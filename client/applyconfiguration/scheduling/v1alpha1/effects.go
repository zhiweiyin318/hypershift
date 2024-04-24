/*


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

package v1alpha1

import (
	resource "k8s.io/apimachinery/pkg/api/resource"
)

// EffectsApplyConfiguration represents an declarative configuration of the Effects type for use
// with apply.
type EffectsApplyConfiguration struct {
	KASGoMemLimit                 *resource.Quantity                  `json:"kasGoMemLimit,omitempty"`
	ControlPlanePriorityClassName *string                             `json:"controlPlanePriorityClassName,omitempty"`
	EtcdPriorityClassName         *string                             `json:"etcdPriorityClassName,omitempty"`
	APICriticalPriorityClassName  *string                             `json:"APICriticalPriorityClassName,omitempty"`
	ResourceRequests              []ResourceRequestApplyConfiguration `json:"resourceRequests,omitempty"`
}

// EffectsApplyConfiguration constructs an declarative configuration of the Effects type for use with
// apply.
func Effects() *EffectsApplyConfiguration {
	return &EffectsApplyConfiguration{}
}

// WithKASGoMemLimit sets the KASGoMemLimit field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the KASGoMemLimit field is set to the value of the last call.
func (b *EffectsApplyConfiguration) WithKASGoMemLimit(value resource.Quantity) *EffectsApplyConfiguration {
	b.KASGoMemLimit = &value
	return b
}

// WithControlPlanePriorityClassName sets the ControlPlanePriorityClassName field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ControlPlanePriorityClassName field is set to the value of the last call.
func (b *EffectsApplyConfiguration) WithControlPlanePriorityClassName(value string) *EffectsApplyConfiguration {
	b.ControlPlanePriorityClassName = &value
	return b
}

// WithEtcdPriorityClassName sets the EtcdPriorityClassName field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the EtcdPriorityClassName field is set to the value of the last call.
func (b *EffectsApplyConfiguration) WithEtcdPriorityClassName(value string) *EffectsApplyConfiguration {
	b.EtcdPriorityClassName = &value
	return b
}

// WithAPICriticalPriorityClassName sets the APICriticalPriorityClassName field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the APICriticalPriorityClassName field is set to the value of the last call.
func (b *EffectsApplyConfiguration) WithAPICriticalPriorityClassName(value string) *EffectsApplyConfiguration {
	b.APICriticalPriorityClassName = &value
	return b
}

// WithResourceRequests adds the given value to the ResourceRequests field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the ResourceRequests field.
func (b *EffectsApplyConfiguration) WithResourceRequests(values ...*ResourceRequestApplyConfiguration) *EffectsApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithResourceRequests")
		}
		b.ResourceRequests = append(b.ResourceRequests, *values[i])
	}
	return b
}
