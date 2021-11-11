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

var _ = Describe("Copy", func() {
	// Create a new builder server

	Context("Copying files to a working container", func() {
		When("specified container does not exist", func() {
			srv := New()

			err := srv.Copy(&v1.CopyRequest{
				Container: &v1.Container{
					Id: "test",
				},
				Source: "test.txt",
				Dest:   "test.txt",
			}, nil)

			It("should return an error", func() {
				Expect(err).Should(HaveOccurred())
			})
		})

		When("copying from the workspace", func() {
			It("should successfully append a COPY line", func() {
				srv := New()
				iSrv := srv.(*BuilderServer)
				ctrl := gomock.NewController(GinkgoT())
				mockStr := mock_v1.NewMockBuilder_CopyServer(ctrl)

				resp, _ := srv.From(context.TODO(), &v1.FromRequest{
					Image: "alpine",
				})

				By("logging out the COPY append")
				mockStr.EXPECT().Send(gomock.Any())

				err := srv.Copy(&v1.CopyRequest{
					Container: &v1.Container{
						Id: resp.Container.Id,
					},
					Source: "test.txt",
					Dest:   "test.txt",
				}, mockStr)

				By("not returning an error")
				Expect(err).ShouldNot(HaveOccurred())

				By("adding a COPY line to the working container")
				c, _ := iSrv.store.Get(resp.Container.Id)
				Expect(c.Lines()[1]).To(Equal("COPY test.txt test.txt"))

				ctrl.Finish()
			})
		})

		When("copying from a previous stage", func() {
			When("the stage exists", func() {
				It("should successfully append the copy line", func() {
					srv := New()
					iSrv := srv.(*BuilderServer)
					ctrl := gomock.NewController(GinkgoT())
					mockStr := mock_v1.NewMockBuilder_CopyServer(ctrl)

					resp1, _ := srv.From(context.TODO(), &v1.FromRequest{
						Image: "golang:alpine",
					})

					resp2, _ := srv.From(context.TODO(), &v1.FromRequest{
						Image: "alpine",
					})

					By("logging out the COPY append")
					mockStr.EXPECT().Send(gomock.Any())

					err := srv.Copy(&v1.CopyRequest{
						Container: &v1.Container{
							Id: resp2.Container.Id,
						},
						From:   resp1.Container.Id,
						Source: "test.txt",
						Dest:   "test.txt",
					}, mockStr)

					By("not returning an error")
					Expect(err).ShouldNot(HaveOccurred())

					By("appending a COPY line to the container state")
					c, _ := iSrv.store.Get(resp2.Container.Id)
					Expect(c.Lines()[1]).To(Equal(fmt.Sprintf("COPY --from=layer-%s test.txt test.txt", resp1.Container.Id)))

					ctrl.Finish()
				})
			})

			When("the stage does not exist", func() {
				It("should return an error", func() {
					srv := New()

					resp, _ := srv.From(context.TODO(), &v1.FromRequest{
						Image: "alpine",
					})

					err := srv.Copy(&v1.CopyRequest{
						Container: &v1.Container{
							Id: resp.Container.Id,
						},
						From:   "test",
						Source: "test.txt",
						Dest:   "test.txt",
					}, nil)

					By("returning an error")
					Expect(err).Should(HaveOccurred())
				})
			})
		})
	})
})
