package manifests_test

import (
	"testing"

	lokiv1 "github.com/grafana/loki/operator/apis/loki/v1"
	"github.com/grafana/loki/operator/internal/manifests"
	"github.com/stretchr/testify/require"
)

func TestNewQueryFrontendDeployment_SelectorMatchesLabels(t *testing.T) {
	ss := manifests.NewQueryFrontendDeployment(manifests.Options{
		Name:      "abcd",
		Namespace: "efgh",
		Stack: lokiv1.LokiStackSpec{
			Template: &lokiv1.LokiTemplateSpec{
				QueryFrontend: &lokiv1.LokiComponentSpec{
					Replicas: 1,
				},
			},
		},
	})
	l := ss.Spec.Template.GetObjectMeta().GetLabels()
	for key, value := range ss.Spec.Selector.MatchLabels {
		require.Contains(t, l, key)
		require.Equal(t, l[key], value)
	}
}

func TestNewQueryFrontendDeployment_HasTemplateConfigHashAnnotation(t *testing.T) {
	ss := manifests.NewQueryFrontendDeployment(manifests.Options{
		Name:       "abcd",
		Namespace:  "efgh",
		ConfigSHA1: "deadbeef",
		Stack: lokiv1.LokiStackSpec{
			Template: &lokiv1.LokiTemplateSpec{
				QueryFrontend: &lokiv1.LokiComponentSpec{
					Replicas: 1,
				},
			},
		},
	})
	expected := "loki.grafana.com/config-hash"
	annotations := ss.Spec.Template.Annotations
	require.Contains(t, annotations, expected)
	require.Equal(t, annotations[expected], "deadbeef")
}
