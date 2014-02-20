package hackedu_test

import (
	. "github.com/hackedu/backend/hackedu"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = PDescribe("Users", func() {

	PDescribe("Registration", func() {
		PContext("With valid information", func() {
			PContext("and a valid application", func() {
				PIt("should create a new user in the database", func() {
				})

				PIt("should set the CreatedAt variable", func() {
				})

				PIt("should clear the Password and PasswordVerify fields", func() {
				})

				PIt("should not return an error", func() {
				})
			})

			PContext("and an invalid applicatino", func() {
				PIt("should not create a new user in the database", func() {
				})

				PIt("should return an error", func() {
				})
			})

			PContext("but no application", func() {
				PIt("should not create a new user in the database", func() {
				})

				PIt("should return an error", func() {
				})
			})
		})

		PContext("With invalid information", func() {
			PIt("should not create a new user in the database", func() {
			})

			PIt("should return an error", func() {
			})
		})
	})
})
