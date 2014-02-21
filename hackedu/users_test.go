package hackedu_test

import (
	. "github.com/hackedu/backend/hackedu"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"appengine/aetest"
	"appengine/datastore"
)

var _ = Describe("Users", func() {

	var mockUser User

	BeforeEach(func() {
		mockUser = User{
			FirstName:      "foo",
			LastName:       "bar",
			Email:          "foo@bar.com",
			Password:       "foobarfoobar",
			PasswordVerify: "foobarfoobar",
		}
	})

	Describe("Registration", func() {
		Context("With valid information", func() {

			var validUser User

			BeforeEach(func() {
				validUser = mockUser
			})

			Context("and a valid application", func() {

				var (
					registeredUser User
					registerError  error
				)

				BeforeEach(func() {
					c, err := aetest.NewContext(nil)
					if err != nil {
						Fail(err.Error())
					}
					defer c.Close()

					var key *datastore.Key
					key, registerError = RegisterUser(c, &validUser)
					if registerError != nil {
						Fail(registerError.Error())
					}

					if err = datastore.Get(c, key, &registeredUser); err != nil {
						Fail(err.Error())
					}
				})

				It("should create a new user in the database", func() {
					Expect(registeredUser.FirstName).ToNot(BeNil())
				})

				It("should set the CreatedAt variable", func() {
					Expect(registeredUser.CreatedAt).ToNot(BeNil())
				})

				It("should clear the Password and PasswordVerify fields", func() {
					Expect(registeredUser.Password).To(Equal(""))
					Expect(registeredUser.PasswordVerify).To(Equal(""))
				})

				PIt("should set the application's datastore key", func() {
				})

				PIt("should email me with the user's info on registration", func() {
				})

				It("should not return an error", func() {
					Expect(registerError).To(BeNil())
				})
			})

			PContext("and an invalid application", func() {
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

		Context("With invalid information", func() {

			var (
				invalidUser   User
				registerError error
				key           *datastore.Key
			)

			register := func() {
				c, err := aetest.NewContext(nil)
				if err != nil {
					Fail(err.Error())
				}
				defer c.Close()

				key, registerError = RegisterUser(c, &invalidUser)
			}

			BeforeEach(func() {
				invalidUser = mockUser
			})

			Context("First name doesn't exist", func() {

				BeforeEach(func() {
					invalidUser.FirstName = ""
					register()
				})

				It("should not create a new user in the database", func() {
					Expect(key).To(BeNil())
				})

				It("should return an error", func() {
					Expect(registerError.Error()).To(ContainSubstring("first name"))
				})
			})

			Context("Last name doesn't exist", func() {

				BeforeEach(func() {
					invalidUser.LastName = ""
					register()
				})

				It("should not create a new user in the database", func() {
					Expect(key).To(BeNil())
				})

				It("should return an error", func() {
					Expect(registerError.Error()).To(ContainSubstring("last name"))
				})
			})

			Context("Short password", func() {

				BeforeEach(func() {
					invalidUser.Password = "foo"
					register()
				})

				It("should not create a new user in the database", func() {
					Expect(key).To(BeNil())
				})

				It("should return an error", func() {
					Expect(registerError.Error()).To(ContainSubstring("at least"))
				})
			})

			Context("Password doesn't match password verify", func() {

				BeforeEach(func() {
					invalidUser.Password = "foobarfoobar"
					invalidUser.PasswordVerify = "barfoobarfoo"
					register()
				})

				It("should not create a new user in the database", func() {
					Expect(key).To(BeNil())
				})

				It("should return an error", func() {
					Expect(registerError.Error()).To(ContainSubstring("match"))
				})
			})

			Context("Invalid email", func() {

				BeforeEach(func() {
					invalidUser.Email = "not a valid email"
					register()
				})

				It("should not create a new user in the database", func() {
					Expect(key).To(BeNil())
				})

				It("should return an error", func() {
					Expect(registerError.Error()).To(ContainSubstring("valid email"))
				})
			})

			PContext("Email already taken", func() {
				BeforeEach(func() {
					// Make copies of invalidUser because the object is modified in-place
					first := invalidUser
					second := invalidUser

					invalidUser = first
					register()
					invalidUser = second
					register()
				})

				PIt("should not create a new user in the database", func() {
					Expect(key).To(BeNil())
				})

				PIt("should return an error", func() {
					Expect(registerError.Error()).To(ContainSubstring("already taken"))
				})
			})
		})
	})
})
