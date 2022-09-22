/*
Copyright 2021 The Crossplane Authors.

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

// Code generated by terrajet. DO NOT EDIT.

// Package apis contains Kubernetes API for the provider.
package apis

import (
	"k8s.io/apimachinery/pkg/runtime"

	v1alpha1 "github.com/crossplane-contrib/provider-jet-gcp/apis/accessapproval/v1alpha1"
	v1alpha1accesscontextmanager "github.com/crossplane-contrib/provider-jet-gcp/apis/accesscontextmanager/v1alpha1"
	v1alpha1activedirectory "github.com/crossplane-contrib/provider-jet-gcp/apis/activedirectory/v1alpha1"
	v1alpha1apigee "github.com/crossplane-contrib/provider-jet-gcp/apis/apigee/v1alpha1"
	v1alpha1apikeys "github.com/crossplane-contrib/provider-jet-gcp/apis/apikeys/v1alpha1"
	v1alpha1appengine "github.com/crossplane-contrib/provider-jet-gcp/apis/appengine/v1alpha1"
	v1alpha1assuredworkloads "github.com/crossplane-contrib/provider-jet-gcp/apis/assuredworkloads/v1alpha1"
	v1alpha1bigquery "github.com/crossplane-contrib/provider-jet-gcp/apis/bigquery/v1alpha1"
	v1alpha1bigtable "github.com/crossplane-contrib/provider-jet-gcp/apis/bigtable/v1alpha1"
	v1alpha1billing "github.com/crossplane-contrib/provider-jet-gcp/apis/billing/v1alpha1"
	v1alpha1binaryauthorization "github.com/crossplane-contrib/provider-jet-gcp/apis/binaryauthorization/v1alpha1"
	v1alpha1cloudasset "github.com/crossplane-contrib/provider-jet-gcp/apis/cloudasset/v1alpha1"
	v1alpha1cloudbuild "github.com/crossplane-contrib/provider-jet-gcp/apis/cloudbuild/v1alpha1"
	v1alpha1cloudfunctions "github.com/crossplane-contrib/provider-jet-gcp/apis/cloudfunctions/v1alpha1"
	v1alpha1cloudidentity "github.com/crossplane-contrib/provider-jet-gcp/apis/cloudidentity/v1alpha1"
	v1alpha1cloudiot "github.com/crossplane-contrib/provider-jet-gcp/apis/cloudiot/v1alpha1"
	v1alpha1cloudplatform "github.com/crossplane-contrib/provider-jet-gcp/apis/cloudplatform/v1alpha1"
	v1alpha2 "github.com/crossplane-contrib/provider-jet-gcp/apis/cloudplatform/v1alpha2"
	v1alpha1cloudrun "github.com/crossplane-contrib/provider-jet-gcp/apis/cloudrun/v1alpha1"
	v1alpha1cloudscheduler "github.com/crossplane-contrib/provider-jet-gcp/apis/cloudscheduler/v1alpha1"
	v1alpha1cloudtasks "github.com/crossplane-contrib/provider-jet-gcp/apis/cloudtasks/v1alpha1"
	v1alpha1composer "github.com/crossplane-contrib/provider-jet-gcp/apis/composer/v1alpha1"
	v1alpha1compute "github.com/crossplane-contrib/provider-jet-gcp/apis/compute/v1alpha1"
	v1alpha2compute "github.com/crossplane-contrib/provider-jet-gcp/apis/compute/v1alpha2"
	v1alpha1container "github.com/crossplane-contrib/provider-jet-gcp/apis/container/v1alpha1"
	v1alpha2container "github.com/crossplane-contrib/provider-jet-gcp/apis/container/v1alpha2"
	v1alpha1containeranalysis "github.com/crossplane-contrib/provider-jet-gcp/apis/containeranalysis/v1alpha1"
	v1alpha1datacatalog "github.com/crossplane-contrib/provider-jet-gcp/apis/datacatalog/v1alpha1"
	v1alpha1dataflow "github.com/crossplane-contrib/provider-jet-gcp/apis/dataflow/v1alpha1"
	v1alpha1datafusion "github.com/crossplane-contrib/provider-jet-gcp/apis/datafusion/v1alpha1"
	v1alpha1datalossprevention "github.com/crossplane-contrib/provider-jet-gcp/apis/datalossprevention/v1alpha1"
	v1alpha1dataproc "github.com/crossplane-contrib/provider-jet-gcp/apis/dataproc/v1alpha1"
	v1alpha1datastore "github.com/crossplane-contrib/provider-jet-gcp/apis/datastore/v1alpha1"
	v1alpha1deploymentmanager "github.com/crossplane-contrib/provider-jet-gcp/apis/deploymentmanager/v1alpha1"
	v1alpha1dialogflow "github.com/crossplane-contrib/provider-jet-gcp/apis/dialogflow/v1alpha1"
	v1alpha1dialogflowcx "github.com/crossplane-contrib/provider-jet-gcp/apis/dialogflowcx/v1alpha1"
	v1alpha1dns "github.com/crossplane-contrib/provider-jet-gcp/apis/dns/v1alpha1"
	v1alpha1endpoints "github.com/crossplane-contrib/provider-jet-gcp/apis/endpoints/v1alpha1"
	v1alpha1essentialcontacts "github.com/crossplane-contrib/provider-jet-gcp/apis/essentialcontacts/v1alpha1"
	v1alpha1eventarc "github.com/crossplane-contrib/provider-jet-gcp/apis/eventarc/v1alpha1"
	v1alpha1filestore "github.com/crossplane-contrib/provider-jet-gcp/apis/filestore/v1alpha1"
	v1alpha1firebaserules "github.com/crossplane-contrib/provider-jet-gcp/apis/firebaserules/v1alpha1"
	v1alpha1firestore "github.com/crossplane-contrib/provider-jet-gcp/apis/firestore/v1alpha1"
	v1alpha1gameservices "github.com/crossplane-contrib/provider-jet-gcp/apis/gameservices/v1alpha1"
	v1alpha1gkehub "github.com/crossplane-contrib/provider-jet-gcp/apis/gkehub/v1alpha1"
	v1alpha1healthcare "github.com/crossplane-contrib/provider-jet-gcp/apis/healthcare/v1alpha1"
	v1alpha1iap "github.com/crossplane-contrib/provider-jet-gcp/apis/iap/v1alpha1"
	v1alpha1identityplatform "github.com/crossplane-contrib/provider-jet-gcp/apis/identityplatform/v1alpha1"
	v1alpha1kms "github.com/crossplane-contrib/provider-jet-gcp/apis/kms/v1alpha1"
	v1alpha1logging "github.com/crossplane-contrib/provider-jet-gcp/apis/logging/v1alpha1"
	v1alpha1memcache "github.com/crossplane-contrib/provider-jet-gcp/apis/memcache/v1alpha1"
	v1alpha1mlengine "github.com/crossplane-contrib/provider-jet-gcp/apis/mlengine/v1alpha1"
	v1alpha1monitoring "github.com/crossplane-contrib/provider-jet-gcp/apis/monitoring/v1alpha1"
	v1alpha2monitoring "github.com/crossplane-contrib/provider-jet-gcp/apis/monitoring/v1alpha2"
	v1alpha1network "github.com/crossplane-contrib/provider-jet-gcp/apis/network/v1alpha1"
	v1alpha1networkmanagement "github.com/crossplane-contrib/provider-jet-gcp/apis/networkmanagement/v1alpha1"
	v1alpha1networkservices "github.com/crossplane-contrib/provider-jet-gcp/apis/networkservices/v1alpha1"
	v1alpha1notebooks "github.com/crossplane-contrib/provider-jet-gcp/apis/notebooks/v1alpha1"
	v1alpha1orgpolicy "github.com/crossplane-contrib/provider-jet-gcp/apis/orgpolicy/v1alpha1"
	v1alpha1osconfig "github.com/crossplane-contrib/provider-jet-gcp/apis/osconfig/v1alpha1"
	v1alpha1oslogin "github.com/crossplane-contrib/provider-jet-gcp/apis/oslogin/v1alpha1"
	v1alpha1privateca "github.com/crossplane-contrib/provider-jet-gcp/apis/privateca/v1alpha1"
	v1alpha1pubsub "github.com/crossplane-contrib/provider-jet-gcp/apis/pubsub/v1alpha1"
	v1alpha1recaptcha "github.com/crossplane-contrib/provider-jet-gcp/apis/recaptcha/v1alpha1"
	v1alpha2redis "github.com/crossplane-contrib/provider-jet-gcp/apis/redis/v1alpha2"
	v1alpha1resourcemanager "github.com/crossplane-contrib/provider-jet-gcp/apis/resourcemanager/v1alpha1"
	v1alpha1scc "github.com/crossplane-contrib/provider-jet-gcp/apis/scc/v1alpha1"
	v1alpha1secretmanager "github.com/crossplane-contrib/provider-jet-gcp/apis/secretmanager/v1alpha1"
	v1alpha1servicenetworking "github.com/crossplane-contrib/provider-jet-gcp/apis/servicenetworking/v1alpha1"
	v1alpha1sourcerepo "github.com/crossplane-contrib/provider-jet-gcp/apis/sourcerepo/v1alpha1"
	v1alpha1spanner "github.com/crossplane-contrib/provider-jet-gcp/apis/spanner/v1alpha1"
	v1alpha2sql "github.com/crossplane-contrib/provider-jet-gcp/apis/sql/v1alpha2"
	v1alpha1storage "github.com/crossplane-contrib/provider-jet-gcp/apis/storage/v1alpha1"
	v1alpha2storage "github.com/crossplane-contrib/provider-jet-gcp/apis/storage/v1alpha2"
	v1alpha1storagetransfer "github.com/crossplane-contrib/provider-jet-gcp/apis/storagetransfer/v1alpha1"
	v1alpha1tags "github.com/crossplane-contrib/provider-jet-gcp/apis/tags/v1alpha1"
	v1alpha1tpu "github.com/crossplane-contrib/provider-jet-gcp/apis/tpu/v1alpha1"
	v1alpha1apis "github.com/crossplane-contrib/provider-jet-gcp/apis/v1alpha1"
	v1alpha1vertexai "github.com/crossplane-contrib/provider-jet-gcp/apis/vertexai/v1alpha1"
	v1alpha1vpcaccess "github.com/crossplane-contrib/provider-jet-gcp/apis/vpcaccess/v1alpha1"
	v1alpha1workflows "github.com/crossplane-contrib/provider-jet-gcp/apis/workflows/v1alpha1"
)

func init() {
	// Register the types with the Scheme so the components can map objects to GroupVersionKinds and back
	AddToSchemes = append(AddToSchemes,
		v1alpha1.SchemeBuilder.AddToScheme,
		v1alpha1accesscontextmanager.SchemeBuilder.AddToScheme,
		v1alpha1activedirectory.SchemeBuilder.AddToScheme,
		v1alpha1apigee.SchemeBuilder.AddToScheme,
		v1alpha1apikeys.SchemeBuilder.AddToScheme,
		v1alpha1appengine.SchemeBuilder.AddToScheme,
		v1alpha1assuredworkloads.SchemeBuilder.AddToScheme,
		v1alpha1bigquery.SchemeBuilder.AddToScheme,
		v1alpha1bigtable.SchemeBuilder.AddToScheme,
		v1alpha1billing.SchemeBuilder.AddToScheme,
		v1alpha1binaryauthorization.SchemeBuilder.AddToScheme,
		v1alpha1cloudasset.SchemeBuilder.AddToScheme,
		v1alpha1cloudbuild.SchemeBuilder.AddToScheme,
		v1alpha1cloudfunctions.SchemeBuilder.AddToScheme,
		v1alpha1cloudidentity.SchemeBuilder.AddToScheme,
		v1alpha1cloudiot.SchemeBuilder.AddToScheme,
		v1alpha1cloudplatform.SchemeBuilder.AddToScheme,
		v1alpha2.SchemeBuilder.AddToScheme,
		v1alpha1cloudrun.SchemeBuilder.AddToScheme,
		v1alpha1cloudscheduler.SchemeBuilder.AddToScheme,
		v1alpha1cloudtasks.SchemeBuilder.AddToScheme,
		v1alpha1composer.SchemeBuilder.AddToScheme,
		v1alpha1compute.SchemeBuilder.AddToScheme,
		v1alpha2compute.SchemeBuilder.AddToScheme,
		v1alpha1container.SchemeBuilder.AddToScheme,
		v1alpha2container.SchemeBuilder.AddToScheme,
		v1alpha1containeranalysis.SchemeBuilder.AddToScheme,
		v1alpha1datacatalog.SchemeBuilder.AddToScheme,
		v1alpha1dataflow.SchemeBuilder.AddToScheme,
		v1alpha1datafusion.SchemeBuilder.AddToScheme,
		v1alpha1datalossprevention.SchemeBuilder.AddToScheme,
		v1alpha1dataproc.SchemeBuilder.AddToScheme,
		v1alpha1datastore.SchemeBuilder.AddToScheme,
		v1alpha1deploymentmanager.SchemeBuilder.AddToScheme,
		v1alpha1dialogflow.SchemeBuilder.AddToScheme,
		v1alpha1dialogflowcx.SchemeBuilder.AddToScheme,
		v1alpha1dns.SchemeBuilder.AddToScheme,
		v1alpha1endpoints.SchemeBuilder.AddToScheme,
		v1alpha1essentialcontacts.SchemeBuilder.AddToScheme,
		v1alpha1eventarc.SchemeBuilder.AddToScheme,
		v1alpha1filestore.SchemeBuilder.AddToScheme,
		v1alpha1firebaserules.SchemeBuilder.AddToScheme,
		v1alpha1firestore.SchemeBuilder.AddToScheme,
		v1alpha1gameservices.SchemeBuilder.AddToScheme,
		v1alpha1gkehub.SchemeBuilder.AddToScheme,
		v1alpha1healthcare.SchemeBuilder.AddToScheme,
		v1alpha1iap.SchemeBuilder.AddToScheme,
		v1alpha1identityplatform.SchemeBuilder.AddToScheme,
		v1alpha1kms.SchemeBuilder.AddToScheme,
		v1alpha1logging.SchemeBuilder.AddToScheme,
		v1alpha1memcache.SchemeBuilder.AddToScheme,
		v1alpha1mlengine.SchemeBuilder.AddToScheme,
		v1alpha1monitoring.SchemeBuilder.AddToScheme,
		v1alpha2monitoring.SchemeBuilder.AddToScheme,
		v1alpha1network.SchemeBuilder.AddToScheme,
		v1alpha1networkmanagement.SchemeBuilder.AddToScheme,
		v1alpha1networkservices.SchemeBuilder.AddToScheme,
		v1alpha1notebooks.SchemeBuilder.AddToScheme,
		v1alpha1orgpolicy.SchemeBuilder.AddToScheme,
		v1alpha1osconfig.SchemeBuilder.AddToScheme,
		v1alpha1oslogin.SchemeBuilder.AddToScheme,
		v1alpha1privateca.SchemeBuilder.AddToScheme,
		v1alpha1pubsub.SchemeBuilder.AddToScheme,
		v1alpha1recaptcha.SchemeBuilder.AddToScheme,
		v1alpha2redis.SchemeBuilder.AddToScheme,
		v1alpha1resourcemanager.SchemeBuilder.AddToScheme,
		v1alpha1scc.SchemeBuilder.AddToScheme,
		v1alpha1secretmanager.SchemeBuilder.AddToScheme,
		v1alpha1servicenetworking.SchemeBuilder.AddToScheme,
		v1alpha1sourcerepo.SchemeBuilder.AddToScheme,
		v1alpha1spanner.SchemeBuilder.AddToScheme,
		v1alpha2sql.SchemeBuilder.AddToScheme,
		v1alpha1storage.SchemeBuilder.AddToScheme,
		v1alpha2storage.SchemeBuilder.AddToScheme,
		v1alpha1storagetransfer.SchemeBuilder.AddToScheme,
		v1alpha1tags.SchemeBuilder.AddToScheme,
		v1alpha1tpu.SchemeBuilder.AddToScheme,
		v1alpha1apis.SchemeBuilder.AddToScheme,
		v1alpha1vertexai.SchemeBuilder.AddToScheme,
		v1alpha1vpcaccess.SchemeBuilder.AddToScheme,
		v1alpha1workflows.SchemeBuilder.AddToScheme,
	)
}

// AddToSchemes may be used to add all resources defined in the project to a Scheme
var AddToSchemes runtime.SchemeBuilder

// AddToScheme adds all Resources to the Scheme
func AddToScheme(s *runtime.Scheme) error {
	return AddToSchemes.AddToScheme(s)
}
