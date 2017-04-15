package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"strings"
	"testing"

	"golang.org/x/net/html"
)

// // not thread safe
// func captureStdout(f func()) string {
//   old := os.Stdout
//   r, w, _ := os.Pipe()
//   os.Stdout = w
//
//   f()
//
//   w.Close()
//   os.Stdout = old
//
//   var buf bytes.Buffer
//   io.Copy(&buf, r)
//   trimmed := bytes.Trim(buf.Bytes(), "\x00")
//   trimmed = bytes.TrimSpace(trimmed)
//   return string(trimmed)
// }

func compareHTML(n1, n2 *html.Node) (bool, error) {
	var b1, b2 bytes.Buffer

	if err := html.Render(&b1, n1); err != nil {
		return false, err
	}
	if err := html.Render(&b2, n2); err != nil {
		return false, err
	}

	b1s := strings.TrimSpace(b1.String())
	b2s := strings.TrimSpace(b2.String())

	// fmt.Println(b1s == b2s)
	// fmt.Println(string(b1s))
	// fmt.Println(string(b2s))

	return (b1s == b2s), nil
}

func compareHTMLStrings(s1 string, s2 string) (bool, error) {
	var b1, b2 bytes.Buffer

	// parse string as HTML then convert back to ...
	// ... detect bad formatting & remove whitespace
	parsed1, err := html.Parse(strings.NewReader(s1))
	if err != nil {
		return false, err
	}
	if err := html.Render(&b1, parsed1); err != nil {
		return false, err
	}

	parsed2, err := html.Parse(strings.NewReader(s2))
	if err != nil {
		return false, err
	}
	if err := html.Render(&b2, parsed2); err != nil {
		return false, err
	}

	// fmt.Printf("string:\n\n%s", b2.String())
	// fmt.Printf("\n\nnode:\n\n%s", b1.S1tring())

	return strings.Contains(b2.String(), b1.String()), nil
}

var url string = "http://localhost:3000"
var fileName string = "testdata/index2.html"

func zTestOutlineParse(t *testing.T) {
	// resp, err := http.Get(url)
	// if err != nil {
	//   t.Error(err)
	// }
	// defer resp.Body.Close()
	//
	// doc, err := html.Parse(resp.Body)
	// if err != nil {
	//   t.Error(err)
	// }

	s, err := ioutil.ReadFile(fileName)
	if err != nil {
		t.Error(err)
	}

	doc, err := html.Parse(bytes.NewReader(s))
	if err != nil {
		t.Error(err)
	}

	//!+call
	out := captureStdout(func() {
		forEachNode(doc, startElement, endElement)
	})

	expected := `<html><head/><body>a b</body></html>`
	doc2, _ := html.Parse(strings.NewReader(expected))
	match, err := compareHTML(doc, doc2)
	if err != nil {
		t.Error(err)
	}

	if !match {
		fmt.Println(s)
		fmt.Println(out)

		t.Errorf("\nExpected (len %d):\n\n-%s-\n\nGot (len %d):\n\n-%s-", len(s), s, len(out), out)
	}

	// if s != out {
	//   fmt.Println(strings.TrimSpace(s))
	//   fmt.Println(strings.TrimSpace(out))
	//
	//   fmt.Println([]byte(strings.TrimSpace(s)))
	//   fmt.Println([]byte(strings.TrimSpace(out)))
	//   t.Errorf("\nExpected (len %d):\n\n-%s-\n\nGot (len %d):\n\n-%s-", len(s), s, len(out), out)
	// }

	// out2 := strings.TrimSpace(out)
	// expected2 := strings.TrimSpace(expected)
	//
	// parsedActual, _ := html.Parse(strings.NewReader(out2))
	// parsedExpected, _ := html.Parse(strings.NewReader(expected2))
	//
	// _, matched := compareHTML(parsedActual, parsedExpected)
	//
	// if !matched {
	//   t.Errorf("\nExpected (len %d):\n\n-%s-\n\nGot (len %d):\n\n-%s-", len(expected2), expected2, len(out2), out2)
	// }
}

//
// func zsTest_htmlEqualsString(t *testing.T) {
//   htmlString := `<html><head/><body>a b</body></html>`
//   // htmlString2 := `<html><head/><body>a b c</body></html>`
//   htmlString3 := `<html> <head/>
//
//   <body>a b</body></html>
//
//
//   `
//
//   htmlNode, err := html.Parse(strings.NewReader(htmlString3))
//   if err != nil {
//     t.Error(err)
//   }
//
//   expected := htmlString
//   actual := htmlNode
//
//   match, err := htmlEqualsString(expected, actual)
//   if err != nil {
//     t.Error(err)
//   }
//
//   if !match {
//     t.Fail()
//   }
// }

func TestOutline(t *testing.T) {
	htmlString := `<html><head/><body><div id=thediv><span class="foo">hello world</span></div></body></html>`
	actual, err := TidyHTMLString(htmlString)

	if err != nil {
		t.Error(err)
	}

	expectedRaw := `
<html>
  <head/>
  <body>
    <div id='thediv'>
      <span class='foo'>
        hello world
      </span>
    </div>
  </body>
</html>
`
	expected := strings.TrimSpace(expectedRaw)

	if expected != strings.TrimSpace(actual) {
		t.Errorf("Expected:\n%v\n\nGot:\n%v", []byte(expected), []byte(actual))
	}
}
