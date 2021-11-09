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
	"crypto/sha256"
	"encoding/hex"
	"fmt"

	v1 "github.com/nitrictech/boxygen/pkg/proto/builder/v1"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("From", func() {
	// Create a new builder server
	srv := New()
	iSrv := srv.(*BuilderServer)

	Context("creating a new working container", func() {
		When("putting to the container state store succeeds", func() {
			resp, err := srv.From(context.TODO(), &v1.FromRequest{
				Image: "alpine",
			})

			It("should not return an error", func() {
				Expect(err).ShouldNot(HaveOccurred())
			})

			containerId := resp.Container.Id

			It("should return a container id that is a hash of the parent image", func() {
				// SHA256
				h := sha256.New()
				h.Write([]byte("alpine"))
				sum := h.Sum(nil)

				id := hex.EncodeToString(sum)
				Expect(containerId).To(Equal(id))
			})

			It("should store the generated id in the container state store", func() {
				Expect(iSrv.store.Has(containerId)).To(BeTrue())
			})

			container, _ := iSrv.store.Get(containerId)

			It("should add a FROM line to the stored container", func() {
				Expect(container.Lines()[0]).To(Equal(fmt.Sprintf("FROM alpine as layer-%s", containerId)))
			})
		})

		PWhen("putting to the container state store fails", func() {
			// TODO: Add failure test case
		})
	})
})
