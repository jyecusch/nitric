// Copyright 2021 Nitric Pty Ltd.
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

package worker

import (
	"context"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Worker", func() {
	Context("UnimplementedWorker", func() {
		uiWrkr := &UnimplementedWorker{}

		When("calling HandlesTrigger", func() {
			It("should return false", func() {
				Expect(uiWrkr.HandlesTrigger(nil)).To(BeFalse())
			})
		})

		When("calling HandleTrigger", func() {
			It("should return an error", func() {
				_, err := uiWrkr.HandleTrigger(context.TODO(), nil)
				Expect(err).Should(HaveOccurred())
			})
		})
	})
})
