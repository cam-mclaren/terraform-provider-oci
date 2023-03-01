// Copyright (c) 2017, 2023, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package management_agent

import "github.com/oracle/terraform-provider-oci/internal/tfresource"

func RegisterDatasource() {
	tfresource.RegisterDatasource("oci_management_agent_management_agent", ManagementAgentManagementAgentDataSource())
	tfresource.RegisterDatasource("oci_management_agent_management_agent_available_histories", ManagementAgentManagementAgentAvailableHistoriesDataSource())
	tfresource.RegisterDatasource("oci_management_agent_management_agent_count", ManagementAgentManagementAgentCountDataSource())
	tfresource.RegisterDatasource("oci_management_agent_management_agent_get_auto_upgradable_config", ManagementAgentManagementAgentGetAutoUpgradableConfigDataSource())
	tfresource.RegisterDatasource("oci_management_agent_management_agent_images", ManagementAgentManagementAgentImagesDataSource())
	tfresource.RegisterDatasource("oci_management_agent_management_agent_install_key", ManagementAgentManagementAgentInstallKeyDataSource())
	tfresource.RegisterDatasource("oci_management_agent_management_agent_install_keys", ManagementAgentManagementAgentInstallKeysDataSource())
	tfresource.RegisterDatasource("oci_management_agent_management_agent_plugin_count", ManagementAgentManagementAgentPluginCountDataSource())
	tfresource.RegisterDatasource("oci_management_agent_management_agent_plugins", ManagementAgentManagementAgentPluginsDataSource())
	tfresource.RegisterDatasource("oci_management_agent_management_agents", ManagementAgentManagementAgentsDataSource())
}