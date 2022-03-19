/*
Copyright 2022.

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

package resume

import (
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"sigs.k8s.io/controller-runtime/pkg/client"

	resumesv1alpha1 "github.com/jefedavis/resume-operator/apis/resumes/v1alpha1"
)

// CreateConfigMapResumeConfig creates the resume-config ConfigMap resource.
func CreateConfigMapResumeConfig(
	parent *resumesv1alpha1.Profile,
) ([]client.Object, error) {

	resourceObjs := []client.Object{}
	var resourceObj = &unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "v1",
			"kind":       "ConfigMap",
			"metadata": map[string]interface{}{
				"name": "resume-config",
				"labels": map[string]interface{}{
					//controlled by field: profile.firstName
					//controlled by field: profile.lastName
					"resume.jefedavis.dev/candidate": "" + parent.Spec.Profile.FirstName + "" + parent.Spec.Profile.LastName + "",
				},
			},
			"data": map[string]interface{}{
				//controlled by field: baseURL
				//controlled by field: pageTitle
				//controlled by field: pageCount
				"config.toml": `languageCode = "en-us"
defaultContentLanguage = "en"
enableRobotsTXT = true
enableEmoji = true

theme = "resume"
disableKinds = ["page", "section", "taxonomy", "term", "RSS", "sitemap"]

baseURL = "https://` + parent.Spec.BaseURL + `/"
title = "` + parent.Spec.PageTitle + `"
#googleAnalytics = ""

[params]
enableMetaTags = true
colorLight = "#fff"
colorDark = "#666"
colorPageBackground = "#ddd"
colorPrimary = "#4C7535" #LightGreen
colorSecondary = "#68B3C2" #LightTeal
colorHeader = "#3E762A" #DarkGreen
colorHeader2 = "#33779D" #DarkTeal
colorIconPrimary = "#fff"
colorIconBackground = "#96B986"
colorRightColumnBackground = "#f5f5f5"
colorRightColumnHeadingText = "#4C7535"
colorRightColumnBodyText = "#666"
colorRightColumnIconPrimary = "#fff"
colorRightColumnIconBackground = "#96B986"
pages = ` + parent.Spec.PageCount + ``,
			},
		},
	}

	resourceObj.SetNamespace(parent.Namespace)

	resourceObjs = append(resourceObjs, resourceObj)

	return resourceObjs, nil
}