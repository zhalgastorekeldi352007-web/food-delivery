package infra

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRenderTemplate(t *testing.T) {
	body, err := RenderTemplate(OrderConfirmationTemplate, map[string]string{"OrderID": "abc123"})
	require.NoError(t, err)
	require.Contains(t, body, "abc123")
}
