package helm_actions

import (
	"helm.sh/helm/v3/pkg/action"
	"helm_service/helm_agent"
)

func InstallChart() interface{} {
	cmd := action.NewInstall(helm_agent.GetActionConfigurations())
	cmd.Namespace = "akash-helm-server"
	cmd.ReleaseName = "install-test"


}
