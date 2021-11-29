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

	"github.com/golang/mock/gomock"
	mock_v1 "github.com/nitrictech/boxygen/mocks/proto"
	v1 "github.com/nitrictech/boxygen/pkg/proto/builder/v1"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Run", func() {
	Context("running a command on a working container", func() {
		When("the container does not already exist", func() {
			srv := New()
			err := srv.Run(&v1.RunRequest{
				Container: &v1.Container{
					Id: "test",
				},
			}, nil)

			It("should return an error", func() {
				Expect(err).Should(HaveOccurred())
			})
		})

		When("the container exists", func() {
			It("should append a RUN line to the working container state", func() {
				srv := New()
				iSrv := srv.(*BuilderServer)
				store := iSrv.store
				ctrl := gomock.NewController(GinkgoT())
				mockStr := mock_v1.NewMockBuilder_RunServer(ctrl)
				// TODO: Mock container state store
				resp, _ := srv.From(context.TODO(), &v1.FromRequest{
					Image: "alpine",
				})

				err := srv.Run(&v1.RunRequest{
					Container: &v1.Container{
						Id: resp.Container.Id,
					},
					Command: []string{"apk update -y"},
				}, mockStr)

				By("should not return an error")
				Expect(err).ShouldNot(HaveOccurred())

				By("appening RUN to the container state")
				con, _ := store.Get(resp.Container.Id)
				Expect(con.Lines()[1]).To(Equal("RUN apk update -y"))

			})
		})
	})
})
