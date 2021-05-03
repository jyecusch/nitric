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

package pubsub_service_test

import (
	mocks "github.com/nitric-dev/membrane/mocks/pubsub"
	pubsub_plugin "github.com/nitric-dev/membrane/plugins/eventing/pubsub"
	"github.com/nitric-dev/membrane/sdk"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Pubsub Plugin", func() {
	When("Listing Available Topics", func() {
		When("There are no topics available", func() {
			pubsubClient := mocks.NewMockPubsubClient(mocks.MockPubsubOptions{
				Topics: []string{},
			})
			pubsubPlugin, _ := pubsub_plugin.NewWithClient(pubsubClient)

			It("Should return an empty list of topics", func() {
				topics, err := pubsubPlugin.ListTopics()
				Expect(err).To(BeNil())
				Expect(topics).To(BeEmpty())
			})
		})

		When("There are topics available", func() {
			pubsubClient := mocks.NewMockPubsubClient(mocks.MockPubsubOptions{
				Topics: []string{"Test"},
			})
			pubsubPlugin, _ := pubsub_plugin.NewWithClient(pubsubClient)

			It("Should return all available topics", func() {
				topics, err := pubsubPlugin.ListTopics()
				Expect(err).To(BeNil())
				Expect(topics).To(ContainElement("Test"))
			})
		})
	})

	When("Publishing Messages", func() {
		event := &sdk.NitricEvent{
			ID:          "Test",
			PayloadType: "Test",
			Payload: map[string]interface{}{
				"Test": "Test",
			},
		}

		When("To a topic that does not exist", func() {
			pubsubClient := mocks.NewMockPubsubClient(mocks.MockPubsubOptions{
				Topics: []string{},
			})
			pubsubPlugin, _ := pubsub_plugin.NewWithClient(pubsubClient)

			It("should return an error", func() {
				err := pubsubPlugin.Publish("Test", event)
				Expect(err).Should(HaveOccurred())
			})
		})

		When("To a topic that does exist", func() {
			pubsubClient := mocks.NewMockPubsubClient(mocks.MockPubsubOptions{
				Topics: []string{"Test"},
			})
			pubsubPlugin, _ := pubsub_plugin.NewWithClient(pubsubClient)

			It("should successfully publish the message", func() {
				err := pubsubPlugin.Publish("Test", event)
				Expect(err).ShouldNot(HaveOccurred())
				Expect(pubsubClient.PublishedMessages["Test"]).To(HaveLen(1))
			})
		})
	})
})
