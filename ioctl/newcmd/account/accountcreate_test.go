// Copyright (c) 2019 IoTeX Foundation
// This is an alpha (internal) release and is not suitable for production. This source code is provided 'as is' and no
// warranties are given as to title or non-infringement, merchantability or fitness for purpose and, to the extent
// permitted by law, all liability for your use of the code is disclaimed. This source code is governed by Apache
// License 2.0 that can be found in the LICENSE file.

package account

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"

	"github.com/MeshBoxFoundation/mesh-core/ioctl/config"
	"github.com/MeshBoxFoundation/mesh-core/ioctl/output"
	"github.com/MeshBoxFoundation/mesh-core/ioctl/util"
	"github.com/MeshBoxFoundation/mesh-core/test/mock/mock_ioctlclient"
)

func TestNewAccountCreate(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	client := mock_ioctlclient.NewMockClient(ctrl)
	client.EXPECT().SelectTranslation(gomock.Any()).Return("mockTranslationString",
		config.English).Times(12)
	client.EXPECT().PrintInfo(gomock.Any()).Do(func(info string) {
		output.PrintResult(info)
	}).Times(2)

	t.Run("CryptoSm2 is false", func(t *testing.T) {
		client.EXPECT().HasCryptoSm2().Return(false).Times(1)
		cmd := NewAccountCreate(client)
		_, err := util.ExecuteCmd(cmd)
		require.NoError(t, err)
	})

	t.Run("CryptoSm2 is true", func(t *testing.T) {
		client.EXPECT().HasCryptoSm2().Return(true).Times(1)
		cmd := NewAccountCreate(client)
		_, err := util.ExecuteCmd(cmd)
		require.NoError(t, err)
	})
}
