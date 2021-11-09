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

var _ = Describe("Add", func() {
	// Create a new builder server

	Context("adding a file to a working container", func() {
		When("the container does not exist", func() {
			srv := New()

			_, err := srv.Add(context.TODO(), &v1.AddRequest{
				Container: &v1.Container{
					Id: "test",
				},
				Src:  "https://example.com/index.html",
				Dest: "index.html",
			})

			It("should return an error", func() {
				Expect(err).Should(HaveOccurred())
			})
		})

		When("the container exists", func() {
			srv := New()
			iSrv := srv.(*BuilderServer)

			resp, _ := srv.From(context.TODO(), &v1.FromRequest{
				Image: "alpine",
			})

			_, err := srv.Add(context.TODO(), &v1.AddRequest{
				Container: &v1.Container{
					Id: resp.Container.Id,
				},
				Src:  "https://example.com/index.html",
				Dest: "index.html",
			})

			It("should not return an error", func() {
				Expect(err).ShouldNot(HaveOccurred())
			})

			It("should add an ADD line to the working container", func() {
				con, _ := iSrv.store.Get(resp.Container.Id)

				Expect(con.Lines()[1]).To(Equal("ADD https://example.com/index.html index.html"))
			})
		})
	})
})
