package cmd

import (
	"bufio"
	"bytes"
	"encoding/gob"
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"regexp"
	"strconv"
	"text/template"
)

// geocoding txt to...

var goTemplate = `// Autogenerated, DO NOT EDIT
package {{.Package}}

import (
	"fmt"
	"os"
	"strconv"
)

var valueStore []{{.ValueType}} = {{.Values | printf "%#v"}}

type child map[{{.KeyType}}]trie

type trie struct {
	valueIndex int
	child child
}

var data trie = {{.Data}}

func (t *trie) lookup(keys []{{.KeyType}}) {{.ValueType}} {
	var final *trie = t

	for _, k := range keys {
		c, found := final.child[k]
		if found {
			final = &c
		}
	}

	return valueStore[final.valueIndex]
}

func main() {
	var needle []int
	for _, v := range os.Args {
		i, _ := strconv.Atoi(v)
		needle = append(needle, i)
	}
	fmt.Println("Found:", data.lookup(needle))
}
`

var valueStore []string
var valueToIndex map[string]int = make(map[string]int)

type child map[int]trie

type trie struct {
	ValueIndex int
	Child      child
}

func (c child) String() string {
	var buf bytes.Buffer
	buf.WriteString("child{")

	size := len(c)
	count := 0
	for k, v := range c {
		count += 1
		if count == size {
			buf.WriteString(fmt.Sprint(k, ":", v))
		} else {
			buf.WriteString(fmt.Sprint(k, ":", v, ","))
		}
	}

	buf.WriteString("}")
	return buf.String()
}

func (t trie) String() string {
	return fmt.Sprint("trie{", t.ValueIndex, ",", t.Child, "}")
}

func (t *trie) insert(keys []int, value string) {
	var final *trie = t

	for _, k := range keys {
		if ft, found := final.Child[k]; !found {
			storedIndex, foundIndex := valueToIndex[value]
			if !foundIndex {
				valueStore = append(valueStore, value)
				storedIndex = len(valueStore) - 1
				valueToIndex[value] = storedIndex
			}

			var nt trie
			nt = trie{storedIndex, make(child)}
			final.Child[k] = nt
			final = &nt
		} else {
			final = &ft
		}
	}
}

var data trie

func load(path string) {
	f, err := os.Open(path)
	check(err)

	valueStore = append(valueStore, "NA")
	data = trie{0, make(child)}
	s := bufio.NewScanner(f)
	re := regexp.MustCompile(`([0-9]+)[|](.+)`)
	for s.Scan() {
		matches := re.FindStringSubmatch(s.Text())
		if matches != nil {
			prefixes := []int{}
			for _, a := range matches[1] {
				prefix, err := strconv.Atoi(string(a))
				check(err)

				prefixes = append(prefixes, prefix)
			}

			data.insert(prefixes, matches[2])
		}
	}
}

var geocodingTxtToGoCmd = &cobra.Command{
	Use:   "go",
	Short: "Convert txt format to go.",
	Run: func(cmd *cobra.Command, args []string) {
		var path = args[0]
		load(path)

		t := template.Must(template.New("package").Parse(goTemplate))
		t.Execute(os.Stdout, struct {
			Package   string
			KeyType   string
			ValueType string
			Data      trie
			Values    []string
		}{
			"main",
			"int",
			"string",
			data,
			valueStore,
		})
	},
}

var geocodingTxtToGobCmd = &cobra.Command{
	Use:   "gob",
	Short: "Convert txt format to gob.",
	Run: func(cmd *cobra.Command, args []string) {
		var path = args[0]
		load(path)

		encoder := gob.NewEncoder(os.Stdout)
		err := encoder.Encode(struct {
			ValueIndex []string
			Data       trie
		}{
			valueStore,
			data,
		})
		check(err)
	},
}

var geocodingTxtToCmd = &cobra.Command{
	Use:   "to",
	Short: "Convert txt format to...",
}

var geocodingTxtCmd = &cobra.Command{
	Use:   "txt",
	Short: "Manage txt format.",
}

var geocodingCmd = &cobra.Command{
	Use:   "geocoding",
	Short: "Managing geocoding data.",
}

func init() {
	// geocoding txt to...
	geocodingTxtToCmd.AddCommand(geocodingTxtToGoCmd)
	geocodingTxtToCmd.AddCommand(geocodingTxtToGobCmd)
	geocodingTxtCmd.AddCommand(geocodingTxtToCmd)
	geocodingCmd.AddCommand(geocodingTxtCmd)

	// metadata...
	RootCmd.AddCommand(geocodingCmd)
}
