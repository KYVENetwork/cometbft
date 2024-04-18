package test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/KYVENetwork/cometbft/v37/types"
)

func TestMakeHeader(t *testing.T) {
	header := MakeHeader(t, &types.Header{})
	require.NotNil(t, header)

	require.NoError(t, header.ValidateBasic())
}
