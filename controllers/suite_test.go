/*
Copyright 2023.

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

package controllers

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"k8s.io/client-go/kubernetes/scheme"

	cro "github.com/RHsyseng/cluster-relocation-operator/api/v1beta1"
	relocationv1alpha1 "github.com/carbonin/cluster-relocation-service/api/v1alpha1"
	bmh_v1alpha1 "github.com/metal3-io/baremetal-operator/apis/metal3.io/v1alpha1"
	//+kubebuilder:scaffold:imports
)

func TestAPIs(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Controller Suite")
}

var _ = BeforeSuite(func() {
	Expect(cro.AddToScheme(scheme.Scheme)).To(Succeed())
	Expect(relocationv1alpha1.AddToScheme(scheme.Scheme)).To(Succeed())
	Expect(bmh_v1alpha1.AddToScheme(scheme.Scheme)).To(Succeed())
})
