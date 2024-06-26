package gotenberg

import "fmt"

const (
	formFieldURL               string = "url"
	remoteURLBaseHTTPHeaderKey string = "Gotenberg-Remoteurl-"
)

// URLRequest facilitates remote URL conversion
// with the Gotenberg API.
type URLRequest struct {
	*chromeRequest
}

// NewURLRequest create URLRequest.
func NewURLRequest(url string) *URLRequest {
	req := &URLRequest{newChromeRequest()}
	req.values[formFieldURL] = url
	return req
}

func (req *URLRequest) postURL() string {
	return "/forms/chromium/convert/url"
}

func (req *URLRequest) screenshotURL() string {
	return "/forms/chromium/screenshot/url"
}

// AddRemoteURLHTTPHeader add a remote URL custom HTTP header.
func (req *URLRequest) AddRemoteURLHTTPHeader(key, value string) {
	key = fmt.Sprintf("%s%s", remoteURLBaseHTTPHeaderKey, key)
	req.httpHeaders[key] = value
}

func (req *URLRequest) formFiles() map[string]Document {
	files := make(map[string]Document)
	if req.header != nil {
		files["header.html"] = req.header
	}
	if req.footer != nil {
		files["footer.html"] = req.footer
	}
	return files
}

// Compile-time checks to ensure type implements desired interfaces.
var (
	_ = Request(new(URLRequest))
)
