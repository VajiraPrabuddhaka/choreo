/*
 * Copyright (c) 2025, WSO2 Inc. (http://www.wso2.org) All Rights Reserved.
 *
 * WSO2 Inc. licenses this file to you under the Apache License,
 * Version 2.0 (the "License"); you may not use this file except
 * in compliance with the License.
 * You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied. See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */

package kubernetes

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	choreov1 "github.com/choreo-idp/choreo/api/v1"
	"github.com/choreo-idp/choreo/internal/dataplane"
	"github.com/choreo-idp/choreo/internal/labels"
)

func TestDeploymentIntegrationKubernetes(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Deployment Integration Kubernetes Suite")
}

// Create a new DeploymentContext for testing. Each test should create a new context
// and set the required fields for the test.
func newTestDeploymentContext() *dataplane.DeploymentContext {
	deployCtx := &dataplane.DeploymentContext{}

	deployCtx.Project = &choreov1.Project{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "my-project",
			Namespace: "test-organization",
			Labels: map[string]string{
				labels.LabelKeyOrganizationName: "test-organization",
				labels.LabelKeyName:             "my-project",
			},
		},
	}
	deployCtx.Environment = &choreov1.Environment{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "development",
			Namespace: "test-namespace",
			Labels: map[string]string{
				labels.LabelKeyOrganizationName: "test-organization",
				labels.LabelKeyName:             "development",
			},
		},
	}
	deployCtx.Component = &choreov1.Component{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "my-component",
			Namespace: "test-namespace",
			Labels: map[string]string{
				labels.LabelKeyOrganizationName: "test-organization",
				labels.LabelKeyProjectName:      "my-project",
				labels.LabelKeyName:             "my-component",
			},
		},
	}
	deployCtx.DeploymentTrack = &choreov1.DeploymentTrack{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "my-main-track",
			Namespace: "test-namespace",
			Labels: map[string]string{
				labels.LabelKeyOrganizationName: "test-organization",
				labels.LabelKeyProjectName:      "my-project",
				labels.LabelKeyComponentName:    "my-component",
				labels.LabelKeyName:             "my-main-track",
			},
		},
	}
	deployCtx.DeployableArtifact = &choreov1.DeployableArtifact{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "my-artifact",
			Namespace: "test-namespace",
			Labels: map[string]string{
				labels.LabelKeyOrganizationName:    "test-organization",
				labels.LabelKeyProjectName:         "my-project",
				labels.LabelKeyComponentName:       "my-component",
				labels.LabelKeyDeploymentTrackName: "my-main-track",
				labels.LabelKeyName:                "my-artifact",
			},
		},
	}

	deployCtx.Deployment = &choreov1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "my-deployment",
			Namespace: "test-namespace",
			Labels: map[string]string{
				labels.LabelKeyOrganizationName:    "test-organization",
				labels.LabelKeyProjectName:         "my-project",
				labels.LabelKeyEnvironmentName:     "my-environment",
				labels.LabelKeyComponentName:       "my-component",
				labels.LabelKeyDeploymentTrackName: "my-main-track",
				labels.LabelKeyName:                "my-deployment",
			},
		},
	}

	deployCtx.ContainerImage = "my-image:latest"

	return deployCtx
}
