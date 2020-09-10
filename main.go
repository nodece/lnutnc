package main

import (
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/url"
	"os"
	"regexp"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/go-resty/resty/v2"
	"github.com/pkg/errors"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
)

var client = resty.New().RemoveProxy().SetTimeout(10 * time.Second)

func Get(u string) (string, error) {
	res, err := client.R().Get(u)
	if err != nil {
		return "", err
	}
	return res.String(), nil
}

func Post(u string, data url.Values) (string, error) {
	res, err := client.R().SetFormDataFromValues(data).SetHeaders(map[string]string{
		"Content-Type": "application/x-www-form-urlencoded",
		"User-Agent":   "Mozilla/5.0 (iPhone; CPU iPhone OS 13_2_3 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/13.0.3 Mobile/15E148 Safari/604.1",
	}).Post(u)
	if err != nil {
		return "", err
	}
	return res.String(), nil
}

func GbkToUtf8(s []byte) ([]byte, error) {
	reader := transform.NewReader(bytes.NewBuffer(s), simplifiedchinese.GBK.NewDecoder())
	return ioutil.ReadAll(reader)
}

func ParsePortalForm(s string) (url.Values, error) {
	data := url.Values{}
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(s))
	if err != nil {
		return nil, err
	}
	doc.Find("form[name=portalForm] input").Each(func(i int, s *goquery.Selection) {
		name, exists := s.Attr("name")
		if !exists {
			return
		}
		value := s.AttrOr("value", "")
		data.Add(name, value)
	})
	return data, nil
}

func DetectConnectAndURL() (bool, string, error) {
	detectURL := "http://baidu.com"
	res, err := client.R().Get(detectURL)
	if err != nil {
		return false, "", err
	}
	if len(res.Body()) != 0 && !bytes.Contains(res.Body(), []byte("portal.do")) {
		return true, "", nil
	}

	re, err := regexp.Compile("(http|https)://[^\"]*")
	if err != nil {
		return false, "", err
	}

	portalURL := re.FindString(string(res.Body()))
	if portalURL == "" {
		return false, "", errors.New("unable to get portal URL")
	}

	return false, portalURL, nil
}

func GoConnect(userid, password string) error {
	connected, portalURL, err := DetectConnectAndURL()
	if err != nil {
		return errors.Wrap(err, "detect network")
	}
	if connected {
		return nil
	}

	portalBody, err := Get(portalURL)
	if err != nil {
		return errors.Wrapf(err, "unable to load the portal HTML %s", portalURL)
	}

	u, _ := url.Parse(portalURL)

	connectURL := fmt.Sprintf("%s://%s/portalAuthAction.do", u.Scheme, u.Host)
	connectData, err := ParsePortalForm(portalBody)
	if err != nil {
		return errors.Wrapf(err, "unable to parse the portal HTML")
	}
	connectData.Set("useridtemp", userid)
	connectData.Set("userid", userid)
	connectData.Set("passwd", password)

	body, err := Post(connectURL, connectData)
	if err != nil {
		return errors.Wrapf(err, "unable to post the connect data %s", connectURL)
	}

	ok, err := GbkToUtf8([]byte(body))
	if err != nil {
		return errors.Wrapf(err, "unable to convert from GKB to UTF8")
	}

	if strings.Contains(string(ok), "/portalDisconnAction.do") {
		return nil
	}
	return errors.New("unable to connect to LNUT network")
}

var (
	u string
	p string
	h bool
	v bool
)

func init() {
	flag.BoolVar(&h, "h", false, "Print the help information")
	flag.BoolVar(&v, "v", false, "Print the version information")
	flag.StringVar(&u, "u", "", "Userid of LNUT network. The environment variable LNUT_USERID has precedence over this flag")
	flag.StringVar(&p, "p", "", "Password of LNUT network. The environment variable LNUT_PASSWORD has precedence over this flag")
}

func Help() {
	bin := os.Args[0]
	fmt.Fprintf(flag.CommandLine.Output(), "Usage of %s:\n", bin)
	flag.PrintDefaults()
	userid := time.Now().Format("2006150405")
	password := "8888"
	fmt.Println("Example:")
	fmt.Printf("%s -u %s -p %s\n", bin, userid, password)
	fmt.Printf("LNUT_USERID=%s LNUT_PASSWORD=%s %s\n", userid, password, bin)
}

func main() {
	flag.Parse()

	if v {
		fmt.Printf("Version: %s\nBuildDate: %s\n", GetVersion(), GetBuildDate())
		return
	}

	if h {
		Help()
		return
	}

	userid := os.Getenv("LNUT_USERID")
	if len(userid) != 0 {
		u = userid
	}

	password := os.Getenv("LNUT_PASSWORD")
	if len(password) != 0 {
		p = password
	}

	if len(u) == 0 || len(p) == 0 {
		Help()
		return
	}

	fmt.Println("Welcome to connect to LNUT network - Zixuan Liu(nodeces@gmail.com)")

	log.Printf("using %s to login to LNUT network\n", u)

	for {
		err := GoConnect(u, p)
		if err != nil {
			log.Println(err)
			time.Sleep(3 * time.Second)
			continue
		}
		log.Println("connected to LNUT network")
		time.Sleep(5 * time.Second)
	}
}
