// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/terraform-providers/terraform-provider-oci/internal/acctest"
	"github.com/terraform-providers/terraform-provider-oci/internal/resourcediscovery"
	"github.com/terraform-providers/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	DeployOkeBlueGreenTrafficShiftStageRequiredOnlyResource = DeployOkeBlueGreenTrafficShiftStageResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_devops_deploy_stage", "test_deploy_stage", acctest.Required, acctest.Create, deployOkeBlueGreenTrafficShiftStageRepresentation)

	DeployOkeBlueGreenTrafficShiftStageResourceConfig = DeployOkeBlueGreenTrafficShiftStageResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_devops_deploy_stage", "test_deploy_stage", acctest.Optional, acctest.Update, deployOkeBlueGreenTrafficShiftStageRepresentation)

	deployOkeBlueGreenTrafficShiftStageSingularDataSourceRepresentation = map[string]interface{}{
		"deploy_stage_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_devops_deploy_stage.test_deploy_stage.id}`},
	}

	deployOkeBlueGreenTrafficShiftStageRepresentation = acctest.GetUpdatedRepresentationCopy("deploy_stage_type", acctest.Representation{RepType: acctest.Required, Create: `OKE_BLUE_GREEN_TRAFFIC_SHIFT`},
		acctest.RepresentationCopyWithNewProperties(acctest.RepresentationCopyWithRemovedProperties(deployStageRepresentation, []string{"wait_criteria", "deploy_stage_predecessor_collection"}), map[string]interface{}{
			"oke_blue_green_deploy_stage_id":      acctest.Representation{RepType: acctest.Required, Create: `${oci_devops_deploy_stage.test_oke_blue_green_deploy_stage.id}`},
			"deploy_stage_predecessor_collection": acctest.RepresentationGroup{RepType: acctest.Required, Group: okeBlueGreenTrafficShiftDeployStageDeployStagePredecessorCollectionRepresentation},
		}))

	okeBlueGreenTrafficShiftDeployStageDeployStagePredecessorCollectionRepresentation = map[string]interface{}{
		"items": acctest.RepresentationGroup{RepType: acctest.Required, Group: okeBlueGreenTrafficShiftDeployStageDeployStagePredecessorCollectionItemsRepresentation},
	}

	okeBlueGreenTrafficShiftDeployStageDeployStagePredecessorCollectionItemsRepresentation = map[string]interface{}{
		"id": acctest.Representation{RepType: acctest.Required, Create: `${oci_devops_deploy_stage.test_oke_blue_green_deploy_stage.id}`},
	}

	DeployOkeBlueGreenTrafficShiftStageResourceDependencies = acctest.GenerateResourceFromRepresentationMap("oci_devops_deploy_artifact", "test_deploy_inline_artifact", acctest.Required, acctest.Create, deployArtifactRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_devops_deploy_environment", "test_deploy_kubernetes_environment", acctest.Required, acctest.Create, deployEnvironmentRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_devops_deploy_pipeline", "test_deploy_pipeline", acctest.Required, acctest.Create, deployPipelineRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_devops_project", "test_project", acctest.Required, acctest.Create, devopsProjectRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_devops_deploy_stage", "test_oke_blue_green_deploy_stage", acctest.Required, acctest.Create, deployOkeBlueGreenStageRepresentation) +
		DefinedTagsDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_ons_notification_topic", "test_notification_topic", acctest.Required, acctest.Create, notificationTopicRepresentation)
)

// issue-routing-tag: devops/default
func TestDevopsDeployStageResource_okeBlueGreenTrafficShift(t *testing.T) {
	httpreplay.SetScenario("TestDevopsDeployStageResource_okeBlueGreenTrafficShift")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_devops_deploy_stage.test_deploy_stage"
	datasourceName := "data.oci_devops_deploy_stages.test_deploy_stages"
	singularDatasourceName := "data.oci_devops_deploy_stage.test_deploy_stage"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "Create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+DeployOkeBlueGreenTrafficShiftStageResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_devops_deploy_stage", "test_deploy_stage", acctest.Optional, acctest.Create, deployOkeBlueGreenTrafficShiftStageRepresentation), "devops", "deployStage", t)

	acctest.ResourceTest(t, testAccCheckDevopsDeployStageDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + DeployOkeBlueGreenTrafficShiftStageResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_devops_deploy_stage", "test_deploy_stage", acctest.Required, acctest.Create, deployOkeBlueGreenTrafficShiftStageRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "deploy_pipeline_id"),
				resource.TestCheckResourceAttr(resourceName, "deploy_stage_predecessor_collection.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "deploy_stage_predecessor_collection.0.items.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "deploy_stage_predecessor_collection.0.items.0.id"),
				resource.TestCheckResourceAttr(resourceName, "deploy_stage_type", "OKE_BLUE_GREEN_TRAFFIC_SHIFT"),
				resource.TestCheckResourceAttrSet(resourceName, "oke_blue_green_deploy_stage_id"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + DeployOkeBlueGreenTrafficShiftStageResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + DeployOkeBlueGreenTrafficShiftStageResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_devops_deploy_stage", "test_deploy_stage", acctest.Optional, acctest.Create, deployOkeBlueGreenTrafficShiftStageRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "compartment_id"),
				resource.TestCheckResourceAttrSet(resourceName, "deploy_pipeline_id"),
				resource.TestCheckResourceAttr(resourceName, "deploy_stage_predecessor_collection.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "deploy_stage_predecessor_collection.0.items.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "deploy_stage_predecessor_collection.0.items.0.id"),
				resource.TestCheckResourceAttr(resourceName, "deploy_stage_type", "OKE_BLUE_GREEN_TRAFFIC_SHIFT"),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "3"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "project_id"),
				resource.TestCheckResourceAttrSet(resourceName, "oke_blue_green_deploy_stage_id"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					if isEnableExportCompartment, _ := strconv.ParseBool(utils.GetEnvSettingWithDefault("enable_export_compartment", "true")); isEnableExportCompartment {
						if errExport := resourcediscovery.TestExportCompartmentWithResourceName(&resId, &compartmentId, resourceName); errExport != nil {
							return errExport
						}
					}
					return err
				},
			),
		},

		// verify updates to updatable parameters
		{
			Config: config + compartmentIdVariableStr + DeployOkeBlueGreenTrafficShiftStageResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_devops_deploy_stage", "test_deploy_stage", acctest.Optional, acctest.Update, deployOkeBlueGreenTrafficShiftStageRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "compartment_id"),
				resource.TestCheckResourceAttrSet(resourceName, "deploy_pipeline_id"),
				resource.TestCheckResourceAttr(resourceName, "deploy_stage_predecessor_collection.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "deploy_stage_predecessor_collection.0.items.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "deploy_stage_predecessor_collection.0.items.0.id"),
				resource.TestCheckResourceAttr(resourceName, "description", "description2"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "3"),
				resource.TestCheckResourceAttrSet(resourceName, "project_id"),
				resource.TestCheckResourceAttrSet(resourceName, "oke_blue_green_deploy_stage_id"),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("Resource recreated when it was supposed to be updated.")
					}
					return err
				},
			),
		},
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_devops_deploy_stages", "test_deploy_stages", acctest.Optional, acctest.Update, deployStageDataSourceRepresentation) +
				compartmentIdVariableStr + DeployOkeBlueGreenTrafficShiftStageResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_devops_deploy_stage", "test_deploy_stage", acctest.Optional, acctest.Update, deployOkeBlueGreenTrafficShiftStageRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(datasourceName, "deploy_pipeline_id"),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttrSet(datasourceName, "id"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_devops_deploy_stage", "test_deploy_stage", acctest.Required, acctest.Create, deployOkeBlueGreenTrafficShiftStageSingularDataSourceRepresentation) +
				compartmentIdVariableStr + DeployOkeBlueGreenTrafficShiftStageResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "deploy_stage_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "compartment_id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "deploy_stage_predecessor_collection.#", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "deploy_stage_predecessor_collection.0.items.0.id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "deploy_stage_type", "OKE_BLUE_GREEN_TRAFFIC_SHIFT"),
				resource.TestCheckResourceAttr(singularDatasourceName, "description", "description2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "defined_tags.%", "3"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "project_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "oke_blue_green_deploy_stage_id"),
			),
		},
		// verify resource import
		{
			Config:                  config + DeployStageRequiredOnlyResource,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}