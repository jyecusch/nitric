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

package events_test

import (
	"context"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/nitrictech/nitric/core/pkg/plugins/events"
)

var _ = Describe("Unimplemented Events Plugin Tests", func() {
	uiep := &events.UnimplementedEventsPlugin{}

	Context("Publish", func() {
		When("Calling Send on UnimplementedEventsPlugin", func() {
			err := uiep.Publish(context.TODO(), "test", 0, &events.NitricEvent{})

			It("should return an unimplemented error", func() {
				Expect(err).Should(HaveOccurred())
				Expect(err.Error()).To(ContainSubstring("UNIMPLEMENTED"))
			})
		})
	})

	Context("ListTopics", func() {
		When("Calling ListTopics on UnimplementedQueuePlugin", func() {
			_, err := uiep.ListTopics(context.TODO())

			It("should return an unimplemented error", func() {
				Expect(err).Should(HaveOccurred())
				Expect(err.Error()).To(ContainSubstring("UNIMPLEMENTED"))
			})
		})
	})
})
