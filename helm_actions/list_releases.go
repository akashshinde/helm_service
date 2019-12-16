package helm_actions

import (
	"encoding/json"
	"fmt"
	"helm.sh/helm/v3/pkg/action"
	"helm_service/helm_agent"
)

func ListReleases() interface{} {
	cmd := action.NewList(helm_agent.GetActionConfigurations())
	cmd.AllNamespaces = true
	cmd.All = true

	releases, err := cmd.Run()
	if err != nil {
		panic(err)
	}
	txt, _ := json.Marshal(releases)
	fmt.Println(string(txt))
	return releases
}
