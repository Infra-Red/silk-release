package integration_test

import (
	"io/ioutil"
	"os"

	"code.cloudfoundry.org/silk/client/config"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gexec"
)

var _ = Describe("error cases", func() {
	var (
		setupConfig    config.Config
		configFilePath string
	)

	BeforeEach(func() {
		stateFile, err := ioutil.TempFile("", "")
		Expect(err).NotTo(HaveOccurred())

		setupConfig = config.Config{
			SubnetRange:    "10.255.0.0/16",
			SubnetMask:     24,
			Database:       testDatabase.DBConfig(),
			UnderlayIP:     "10.244.4.6",
			LocalStateFile: stateFile.Name(),
		}
		configFilePath = writeConfigFile(setupConfig)
	})

	AfterEach(func() {
		os.Remove(setupConfig.LocalStateFile)
		os.Remove(configFilePath)
	})

	Context("when the path to the config is bad", func() {
		It("exits with status 1", func() {
			session := startSetup("some/bad/path")
			Eventually(session, DEFAULT_TIMEOUT).Should(gexec.Exit(1))
			Expect(session.Err.Contents()).To(ContainSubstring("loading config: reading file some/bad/path"))

			session.Interrupt()
		})
	})

	Context("when the contents of the config file cannot be unmarshaled", func() {
		BeforeEach(func() {
			Expect(ioutil.WriteFile(configFilePath, []byte("some-bad-contents"), os.ModePerm)).To(Succeed())
		})

		It("exits with status 1", func() {
			session := startSetup(configFilePath)
			Eventually(session, DEFAULT_TIMEOUT).Should(gexec.Exit(1))
			Expect(session.Err.Contents()).To(ContainSubstring("loading config: unmarshaling contents"))
		})
	})

	Context("when the config has an unsupported type", func() {
		BeforeEach(func() {
			os.Remove(configFilePath)
			setupConfig.Database.Type = "bad-type"
			configFilePath = writeConfigFile(setupConfig)
		})

		It("exits with status 1", func() {
			session := startSetup(configFilePath)
			Eventually(session, DEFAULT_TIMEOUT).Should(gexec.Exit(1))
			Expect(session.Err.Contents()).To(ContainSubstring("creating lease contoller: connecting to database:"))
		})
	})

	Context("when the config has a bad connection string", func() {
		BeforeEach(func() {
			os.Remove(configFilePath)
			setupConfig.Database.ConnectionString = "some-bad-connection-string"
			configFilePath = writeConfigFile(setupConfig)
		})

		It("exits with status 1", func() {
			session := startSetup(configFilePath)
			Eventually(session, "10s").Should(gexec.Exit(1))
			Expect(session.Err.Contents()).To(ContainSubstring("creating lease contoller: connecting to database:"))
		})

		It("deletes the local state file", func() {
			session := startSetup(configFilePath)
			Eventually(session, "10s").Should(gexec.Exit(1))
			_, err := os.Stat(setupConfig.LocalStateFile)
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(ContainSubstring("no such file or directory"))
		})
	})

	XContext("when the lease controller fails to acquire a subnet lease", func() {
		It("exits with status 1", func() {
			// TODO: unpend, figure out how to set up the test so that we can trigger
			// this sort of failure and actually test the behavior in that case
		})
	})

	Context("when the local state file cannot be written", func() {
		BeforeEach(func() {
			os.Remove(configFilePath)
			setupConfig.LocalStateFile = "/some/path/that/does/not/exit"
			configFilePath = writeConfigFile(setupConfig)
		})
		It("exits with status 1", func() {
			session := startSetup(configFilePath)
			Eventually(session, "10s").Should(gexec.Exit(1))
			Expect(session.Err.Contents()).To(ContainSubstring("writing local state file:"))
		})
	})
})
