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

var _ = Describe("Run", func() {
	Context("running a command on a working container", func() {
		When("the container does not already exist", func() {
			srv := New()
			resp, err := srv.Run(context.TODO(), &v1.RunRequest{
				Container: &v1.Container{
					Id: "test",
				},
			})

			It("should return an error", func() {
				Expect(err).Should(HaveOccurred())
			})

			It("should return a nil response", func() {
				Expect(resp).To(BeNil())
			})
		})

		When("the container exists", func() {
			// TODO: Add failure test case
			srv := New()
			iSrv := srv.(*BuilderServer)
			store := iSrv.store.(*containerStateStoreImpl)
			// TODO: Mock container state store
			resp, err := srv.From(context.TODO(), &v1.FromRequest{
				Image: "alpine",
			})

			_, err = srv.Run(context.TODO(), &v1.RunRequest{
				Container: &v1.Container{
					Id: resp.Container.Id,
				},
				Command: []string{"apk update -y"},
			})

			It("should not return an error", func() {
				Expect(err).ShouldNot(HaveOccurred())
			})

			It("should add a RUN line to the container state", func() {
				Expect(store.store[resp.Container.Id].Lines()[1]).To(Equal("RUN apk update -y"))
			})
		})
	})
})
