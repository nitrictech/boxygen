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

	v1 "github.com/nitrictech/boxygen/pkg/proto/builder/v1"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Copy", func() {
	// Create a new builder server

	Context("Copying files to a working container", func() {
		When("specified container does not exist", func() {
			srv := New()

			r, err := srv.Copy(context.TODO(), &v1.CopyRequest{
				Container: &v1.Container{
					Id: "test",
				},
				Source: "test.txt",
				Dest:   "test.txt",
			})

			It("should return an error", func() {
				Expect(err).Should(HaveOccurred())
			})

			It("should return a nil response", func() {
				Expect(r).To(BeNil())
			})
		})

		When("copying from the workspace", func() {
			srv := New()
			iSrv := srv.(*BuilderServer)

			resp, _ := srv.From(context.TODO(), &v1.FromRequest{
				Image: "alpine",
			})

			_, err := srv.Copy(context.TODO(), &v1.CopyRequest{
				Container: &v1.Container{
					Id: resp.Container.Id,
				},
				Source: "test.txt",
				Dest:   "test.txt",
			})

			It("should return an error", func() {
				Expect(err).ShouldNot(HaveOccurred())
			})

			It("should add a COPY line to the working container", func() {
				c, _ := iSrv.store.Get(resp.Container.Id)
				Expect(c.Lines()[1]).To(Equal("COPY /workspace/test.txt test.txt"))
			})
		})

		When("copying from a previous stage", func() {
			When("the stage exists", func() {
				srv := New()
				iSrv := srv.(*BuilderServer)

				resp1, _ := srv.From(context.TODO(), &v1.FromRequest{
					Image: "golang:alpine",
				})

				resp2, _ := srv.From(context.TODO(), &v1.FromRequest{
					Image: "alpine",
				})

				_, err := srv.Copy(context.TODO(), &v1.CopyRequest{
					Container: &v1.Container{
						Id: resp2.Container.Id,
					},
					From:   resp1.Container.Id,
					Source: "test.txt",
					Dest:   "test.txt",
				})

				It("should not return an error", func() {
					Expect(err).ShouldNot(HaveOccurred())
				})

				It("should add a COPY line to the working container", func() {
					c, _ := iSrv.store.Get(resp2.Container.Id)
					Expect(c.Lines()[1]).To(Equal(fmt.Sprintf("COPY --from layer-%s test.txt test.txt", resp1.Container.Id)))
				})
			})

			When("the stage does not exist", func() {
				srv := New()

				resp, _ := srv.From(context.TODO(), &v1.FromRequest{
					Image: "alpine",
				})

				_, err := srv.Copy(context.TODO(), &v1.CopyRequest{
					Container: &v1.Container{
						Id: resp.Container.Id,
					},
					From:   "test",
					Source: "test.txt",
					Dest:   "test.txt",
				})

				It("should return an error", func() {
					Expect(err).Should(HaveOccurred())
				})
			})

		})
	})
})
