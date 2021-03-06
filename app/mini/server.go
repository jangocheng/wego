package mini

import (
	"bytes"
	"encoding/xml"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/godcong/wego/core"
	"github.com/godcong/wego/core/message"
	"github.com/godcong/wego/crypt"
	"github.com/godcong/wego/log"
)

/*Server Server */
type Server struct {
	CryptResponse   bool
	message         *core.Message
	mType           string
	bizMsg          *crypt.BizMsg
	defaultCallback []core.MessageCallback
	callback        map[message.MsgType][]core.MessageCallback
}

/*RegisterCallback 注册回调 */
func (s *Server) RegisterCallback(sc core.MessageCallback, types ...message.MsgType) {
	size := len(types)
	if size == 0 {
		s.defaultCallback = append(s.defaultCallback, sc)
		return
	}
	for _, t := range types {
		if callback, b := s.callback[t]; b {
			s.callback[t] = append(callback, sc)
		} else {
			s.callback[t] = []core.MessageCallback{sc}
		}
	}
}

// ServeHTTP ...
func (s *Server) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	var bodyBytes []byte
	var rltXML []byte
	var err error
	if req.Body != nil {
		bodyBytes, _ = ioutil.ReadAll(req.Body)
	}
	// Restore the io.ReadCloser to its original state
	req.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))

	if len(bodyBytes) == 0 {
		w.WriteHeader(http.StatusOK)
	}
	query, err := url.ParseQuery(req.URL.RawQuery)
	if err != nil {
		log.Error(err)
		w.WriteHeader(http.StatusOK)
	}
	encryptType := query.Get("encrypt_type")
	ts := query.Get("timestamp")
	nonce := query.Get("nonce")
	msgSignature := query.Get("msg_signature")

	if encryptType == "aes" {
		log.Debug(ts, nonce, msgSignature, string(bodyBytes))
		bodyBytes, err = s.bizMsg.Decrypt(string(bodyBytes), msgSignature, ts, nonce)
		if err != nil {
			log.Error(err)
			w.WriteHeader(http.StatusOK)
		}
	}

	message := new(core.Message)
	log.Debug(string(bodyBytes))
	err = xml.Unmarshal(bodyBytes, message)
	if err != nil {
		log.Error(err)
		return
	}
	result := s.CallbackFunc(message)
	w.WriteHeader(http.StatusOK)

	rltXML, err = result.ToXML()
	if err != nil {
		log.Error(err)
		return
	}

	//if encryptType == "aes" {
	//	tmpStr, err := s.bizMsg.RSAEncrypt(string(rltXML), ts, nonce)
	//	if err != nil {
	//		log.Error(err)
	//		return
	//	}
	//	rltXML = []byte(tmpStr)
	//}
	if s.mType == "xml" {
		header := w.Header()
		if val := header["Content-Type"]; len(val) == 0 {
			header["Content-Type"] = []string{"application/xml; charset=utf-8"}
		}
	} else {
		header := w.Header()
		if val := header["Content-Type"]; len(val) == 0 {
			header["Content-Type"] = []string{"application/json; charset=utf-8"}
		}
	}
	log.Debug(string(rltXML))
	w.Write(rltXML)
	return
}

/*CallbackFunc CallbackFunc */
func (s *Server) CallbackFunc(msg *core.Message) message.Messager {
	var result message.Messager
	for _, v := range s.defaultCallback {
		if r := v(msg); r != nil {
			result = r
		}
	}

	if v0, b := s.callback[msg.GetType()]; b {
		for _, v := range v0 {
			if r := v(msg); r != nil {
				result = r
			}
		}
	}
	return result
}

func newServer(token, key, id string) *Server {
	return &Server{
		mType:           "xml",
		bizMsg:          crypt.NewBizMsg(token, key, id),
		message:         nil,
		defaultCallback: []core.MessageCallback{},
		callback:        map[message.MsgType][]core.MessageCallback{},
	}
}

/*NewServer NewServer */
func NewServer(config *core.Config) *Server {
	log.Debug(config.Get("accessToken"), config.Get("aes_key"), config.Get("app_id"))
	return newServer(config.GetString("accessToken"), config.GetString("aes_key"), config.GetString("app_id"))
}
