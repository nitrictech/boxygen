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
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("New", func() {
	// Create a new builder server

	When("Creating a new builder server", func() {
		srv := New()

		iSrv, ok := srv.(*BuilderServer)

		It("should return a BuilderServer implementation", func() {
			Expect(ok).To(BeTrue())
		})

		It("should contain an empty container state store", func() {
			css, ok := iSrv.store.(*containerStateStoreImpl)

			Expect(ok).To(BeTrue())
			Expect(len(css.store)).To(Equal(0))
		})
	})
})
