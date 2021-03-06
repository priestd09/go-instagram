package request

import (
	"net/http"

	"github.com/hieven/go-instagram/src/constants"
	sessionMocks "github.com/hieven/go-instagram/src/utils/session/mocks"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/parnurzeal/gorequest"
)

var _ = Describe("common", func() {
	Describe("#withDefaultHeader", func() {
		var (
			mockSessionManager *sessionMocks.SessionManager
			req                *gorequest.SuperAgent

			result *gorequest.SuperAgent
		)

		BeforeEach(func() {
			mockSessionManager = &sessionMocks.SessionManager{}
			mockSessionManager.On("GetCookies").Return([]*http.Cookie{})

			req = gorequest.New()
		})

		JustBeforeEach(func() {
			result = withDefaultHeader(mockSessionManager, req)
		})

		Context("when success", func() {
			It("should return", func() {
				Expect(result).To(Equal(req))
				Expect(result.Header["Connection"]).To(Equal("close"))
				Expect(result.Header["Accept"]).To(Equal("*/*"))
				Expect(result.Header["X-IG-Connection-Type"]).To(Equal("WIFI"))
				Expect(result.Header["X-IG-Capabilities"]).To(Equal("3QI="))
				Expect(result.Header["Accept-Language"]).To(Equal("en-US"))
				Expect(result.Header["Host"]).To(Equal(constants.Hostname))
				Expect(result.Header["User-Agent"]).To(Equal("Instagram " + constants.AppVersion + " Android (21/5.1.1; 401dpi; 1080x1920; Oppo; A31u; A31u; en_US)"))
			})
		})
	})
})
