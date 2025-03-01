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

package queue_test

import (
	"context"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/nitrictech/nitric/core/pkg/plugins/queue"
)

var _ = Describe("Unimplemented Queue Plugin Tests", func() {
	uiqp := &queue.UnimplementedQueuePlugin{}

	Context("Send", func() {
		When("Calling Send on UnimplementedQueuePlugin", func() {
			err := uiqp.Send(context.TODO(), "test", queue.NitricTask{})

			It("should return an unimplemented error", func() {
				Expect(err).Should(HaveOccurred())
				Expect(err.Error()).To(ContainSubstring("UNIMPLEMENTED"))
			})
		})
	})

	Context("Complete", func() {
		When("Calling Complete on UnimplementedQueuePlugin", func() {
			err := uiqp.Complete(context.TODO(), "test", "test")

			It("should return an unimplemented error", func() {
				Expect(err).Should(HaveOccurred())
				Expect(err.Error()).To(ContainSubstring("UNIMPLEMENTED"))
			})
		})
	})

	Context("Receive", func() {
		When("Calling Receive on UnimplementedQueuePlugin", func() {
			_, err := uiqp.Receive(context.TODO(), queue.ReceiveOptions{})

			It("should return an unimplemented error", func() {
				Expect(err).Should(HaveOccurred())
				Expect(err.Error()).To(ContainSubstring("UNIMPLEMENTED"))
			})
		})
	})

	Context("SendBatch", func() {
		When("Calling SendBatch on UnimplementedQueuePlugin", func() {
			_, err := uiqp.SendBatch(context.TODO(), "test", nil)

			It("should return an unimplemented error", func() {
				Expect(err).Should(HaveOccurred())
				Expect(err.Error()).To(ContainSubstring("UNIMPLEMENTED"))
			})
		})
	})
})
