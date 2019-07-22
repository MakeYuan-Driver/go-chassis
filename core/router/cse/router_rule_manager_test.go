package cse_test

import (
	"testing"

	"github.com/go-chassis/go-chassis/core/common"
	"github.com/go-chassis/go-chassis/core/config"
	"github.com/go-chassis/go-chassis/core/config/model"
	"github.com/go-chassis/go-chassis/core/lager"
)

// the service key in router config file
const (
	svcNone               = "svcNone"
	svcRoute              = "svcRoute"
	svcDarkLaunch         = "svcDarkLaunch"
	svcRouteAndDarkLaunch = "svcRouteAndDarkLaunch"
)

func genSvcRouteRule() []*model.RouteRule {
	r := []*model.RouteRule{
		{
			Precedence: 0,
			Routes: []*model.RouteTag{
				{
					Tags: map[string]string{
						common.BuildinTagVersion: "0.0.1",
						common.BuildinTagApp:     svcRoute,
					},
					Weight: 20,
				},
			},
		},
	}
	return r
}

func genSvcDarkLaunchRule() string {
	return `{"policyType":"RATE","ruleItems":[{"groupName":"s0"},{"groupName":"s1"}]}`
}

func genSvcRouteAndDarkLaunchRule() ([]*model.RouteRule, string) {
	r := []*model.RouteRule{
		{
			Precedence: 1,
			Routes: []*model.RouteTag{
				{
					Tags: map[string]string{
						common.BuildinTagVersion: "0.0.1",
						common.BuildinTagApp:     svcRouteAndDarkLaunch,
					},
					Weight: 20,
				},
			},
		},
	}
	return r, `{"policyType":"RATE","ruleItems":[{"groupName":"s0"},{"groupName":"s1"},{"groupName":"s2"}]}`
}

func genDarkLaunchRuleAfterAdd() string {
	return `{"policyType":"RATE","ruleItems":[{"groupName":"s0"},{"groupName":"s1"},{"groupName":"s2"},{"groupName":"s3"}]}`
}

func genDarkLaunchRuleAfterUpdate() string {
	return `{"policyType":"RATE","ruleItems":[{"groupName":"s0"},{"groupName":"s1"},{"groupName":"s2"},{"groupName":"s3"},{"groupName":"s4"}]}`
}

func init() {
	lager.Initialize("", "DEBUG", "", "size", true, 1, 10, 7)
}

func initRouterDefinition() {
	config.RouterDefinition = &model.RouterConfig{Router: model.Router{}}
}

func TestRouter_Init(t *testing.T) {

}
