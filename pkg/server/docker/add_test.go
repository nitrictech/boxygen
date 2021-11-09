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
	"fmt"

	"github.com/golang/mock/gomock"
	mock_v1 "github.com/nitrictech/boxygen/mocks/proto"
	v1 "github.com/nitrictech/boxygen/pkg/proto/builder/v1"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Add", func() {
	Context("adding a file to a working container", func() {
		When("the container does not exist", func() {
			srv := New()

			err := srv.Add(&v1.AddRequest{
				Container: &v1.Container{
					Id: "test",
				},
				Src:  "https://example.com/index.html",
				Dest: "index.html",
			}, nil)

			It("should return an error", func() {
				Expect(err).Should(HaveOccurred())
			})
		})

		When("the container exists", func() {

			It("should append ADD to the containers state", func() {
				srv := New()
				iSrv := srv.(*BuilderServer)

				ctrl := gomock.NewController(GinkgoT())
				mockStr := mock_v1.NewMockBuilder_AddServer(ctrl)

				resp, _ := srv.From(context.TODO(), &v1.FromRequest{
					Image: "alpine",
				})

				By("Logging out line append")
				mockStr.EXPECT().Send(&v1.OutputResponse{
					Log: []string{
						fmt.Sprintf("Append [ADD https://example.com/index.html index.html] to container %s", resp.Container.Id),
					},
				})

				err := srv.Add(&v1.AddRequest{
					Container: &v1.Container{
						Id: resp.Container.Id,
					},
					Src:  "https://example.com/index.html",
					Dest: "index.html",
				}, mockStr)

				By("Not returning an error")
				Expect(err).ShouldNot(HaveOccurred())

				By("Updating the container state store")
				con, _ := iSrv.store.Get(resp.Container.Id)
				Expect(con.Lines()[1]).To(Equal("ADD https://example.com/index.html index.html"))

				ctrl.Finish()
			})

		})
	})
})
