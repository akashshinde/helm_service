package helm_actions

import (
	"helm.sh/helm/v3/pkg/action"
	"helm_service/helm_agent"
)

func ListReleases() (interface{}, error) {
	cmd := action.NewList(helm_agent.GetActionConfigurations())
	cmd.AllNamespaces = true
	cmd.All = true

	releases, err := cmd.Run()
	if err != nil {
		return nil, err
	}
	return releases, nil
}
