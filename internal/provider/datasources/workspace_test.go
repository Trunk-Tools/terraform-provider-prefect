package datasources_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/prefecthq/terraform-provider-prefect/internal/testutils"
)

func fixtureAccWorkspaceByHandle(workspace, handle string) string {
	return fmt.Sprintf(`
%s

data "prefect_workspace" "testdata" {
	handle = "%s"
	depends_on = [prefect_workspace.test]
}
`, workspace, handle)
}

func fixtureAccWorkspaceByID(workspace string) string {
	return fmt.Sprintf(`
%s

data "prefect_workspace" "testdata" {
	id = prefect_workspace.test.id
}
`, workspace)
}

//nolint:paralleltest // we use the resource.ParallelTest helper instead
func TestAccDatasource_workspace(t *testing.T) {
	// Workspaces are not supportedin OSS.
	testutils.SkipTestsIfOSS(t)

	ephemeralWorkspace := testutils.NewEphemeralWorkspace()
	dataSourceName := "data.prefect_workspace.testdata"

	resource.ParallelTest(t, resource.TestCase{
		ProtoV6ProviderFactories: testutils.TestAccProtoV6ProviderFactories,
		PreCheck:                 func() { testutils.AccTestPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: fixtureAccWorkspaceByHandle(ephemeralWorkspace.Resource, ephemeralWorkspace.Name),
				ConfigStateChecks: []statecheck.StateCheck{
					testutils.ExpectKnownValue(dataSourceName, "handle", ephemeralWorkspace.Name),
					testutils.ExpectKnownValueNotNull(dataSourceName, "created"),
					testutils.ExpectKnownValueNotNull(dataSourceName, "updated"),
					testutils.ExpectKnownValueNotNull(dataSourceName, "name"),
				},
			},
			{
				Config: fixtureAccWorkspaceByID(ephemeralWorkspace.Resource),
				ConfigStateChecks: []statecheck.StateCheck{
					testutils.ExpectKnownValue(dataSourceName, "handle", ephemeralWorkspace.Name),
					testutils.ExpectKnownValueNotNull(dataSourceName, "created"),
					testutils.ExpectKnownValueNotNull(dataSourceName, "updated"),
					testutils.ExpectKnownValueNotNull(dataSourceName, "name"),
				},
			},
		},
	})
}
