package hackedu_test

import (
	. "github.com/hackedu/backend/hackedu"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"appengine/aetest"
	"appengine/datastore"
)

var _ = Describe("Application", func() {

	var (
		mockApplication Application
	)

	BeforeEach(func() {
		mockApplication = Application{
			HighSchool:         "Some Datacenter",
			InterestingProject: "hackEDU",
			SystemHacked:       "I hacked my datacenter using x, y, and z.",
			Passion:            "I'm passionate about Go. I really like Go.",
			Story:              "I'm just a little server in a big world.",
			Why:                "They don't have CS classes in my server farm.",
		}
	})

	Describe("Creation", func() {
		Context("With valid information", func() {

			var (
				validApplication   Application
				createdApplication Application
				creationError      error
			)

			BeforeEach(func() {
				validApplication = mockApplication

				c, err := aetest.NewContext(nil)
				if err != nil {
					Fail(err.Error())
				}
				defer c.Close()

				var key *datastore.Key
				key, creationError = ConstructApplication(c, &validApplication)
				if err != nil {
					Fail(err.Error())
				}

				err = datastore.Get(c, key, &createdApplication)
				if err != nil {
					Fail(err.Error())
				}
			})

			It("should create a new application in the database", func() {
				Expect(createdApplication).ToNot(BeNil())
			})

			It("should set the CreatedAt variable", func() {
				Expect(createdApplication.CreatedAt).ToNot(BeNil())
			})

			PIt("should set the User's datastore key", func() {
			})

			It("should not return an error", func() {
				Expect(creationError).To(BeNil())
			})
		})

		Context("With invalid information", func() {

			var (
				invalidApplication Application
				creationError      error
				key                *datastore.Key
			)

			register := func() {
				c, err := aetest.NewContext(nil)
				if err != nil {
					Fail(err.Error())
				}
				defer c.Close()

				key, creationError = ConstructApplication(c, &invalidApplication)
			}

			BeforeEach(func() {
				invalidApplication = mockApplication
			})

			Context("High school doesn't exist", func() {

				BeforeEach(func() {
					invalidApplication.HighSchool = ""
					register()
				})

				It("should not create a new application in the database", func() {
					Expect(key).To(BeNil())
				})

				It("should return an error", func() {
					Expect(creationError.Error()).To(ContainSubstring("high school"))
				})
			})

			Context("Interesting project doesn't exist", func() {

				BeforeEach(func() {
					invalidApplication.InterestingProject = ""
					register()
				})

				It("should not create a new application in the database", func() {
					Expect(key).To(BeNil())
				})

				It("should return an error", func() {
					Expect(creationError.Error()).To(ContainSubstring("project"))
				})
			})

			Context("System hacked doesn't exist", func() {

				BeforeEach(func() {
					invalidApplication.SystemHacked = ""
					register()
				})

				It("should not create a new application in the database", func() {
					Expect(key).To(BeNil())
				})

				It("should return an error", func() {
					Expect(creationError.Error()).To(ContainSubstring("system hacked"))
				})
			})

			Context("Passion doesn't exist", func() {

				BeforeEach(func() {
					invalidApplication.Passion = ""
					register()
				})

				It("should not create a new application in the database", func() {
					Expect(key).To(BeNil())
				})

				It("should return an error", func() {
					Expect(creationError.Error()).To(ContainSubstring("passion"))
				})
			})

			Context("Story doesn't exist", func() {

				BeforeEach(func() {
					invalidApplication.Story = ""
					register()
				})

				It("should not create a new application in the database", func() {
					Expect(key).To(BeNil())
				})

				It("should return an error", func() {
					Expect(creationError.Error()).To(ContainSubstring("story"))
				})
			})

			Context("Why doesn't exist", func() {

				BeforeEach(func() {
					invalidApplication.Why = ""
					register()
				})

				It("should not create a new application in the database", func() {
					Expect(key).To(BeNil())
				})

				It("should return an error", func() {
					Expect(creationError.Error()).To(ContainSubstring("why"))
				})
			})
		})
	})
})
