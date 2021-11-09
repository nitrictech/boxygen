// Copyright 2021 Nitric Technologies Pty Ltd.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package docker_server

import (
	"context"

	v1 "github.com/nitrictech/boxygen/pkg/proto/builder/v1"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Config", func() {
	// Create a new builder server

	Context("When the container does not exist", func() {
		srv := New()

		_, err := srv.Config(context.TODO(), &v1.ConfigRequest{
			Container: &v1.Container{
				Id: "test",
			},
		})

		It("should return an error", func() {
			Expect(err).Should(HaveOccurred())
		})
	})

	Context("When the container exists", func() {
		Context("Entrypoint", func() {

		})

		Context("Cmd", func() {

		})

		Context("Env", func() {

		})

		Context("Ports", func() {

		})

		Context("Volume", func() {

		})
	})
})
