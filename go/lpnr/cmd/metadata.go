package cmd

import (
  "encoding/json"
  "encoding/xml"
  "fmt"
  "github.com/calebcase/pretty"
  "github.com/golang/protobuf/jsonpb"
  "github.com/golang/protobuf/proto"
  "github.com/spf13/cobra"
  "io/ioutil"
  "os"
  md_pb "github.com/calebcase/libphonenumber/go/pb"
  md_xml "github.com/calebcase/libphonenumber/go/xml"
)

func check(e error) {
  if e != nil {
    panic(e)
  }
}

// metadata xml to...

var xmlToJsonCmd = &cobra.Command{
  Use: "json",
  Short: "Convert xml format to json.",
  Run: func (cmd *cobra.Command, args []string) {
    var path = args[0]
    xml_bytes, err := ioutil.ReadFile(path)
    check(err)

    xml_tree := &md_xml.PhoneNumberMetadata{}
    err = xml.Unmarshal(xml_bytes, xml_tree)
    check(err)

    json_bytes, err := json.MarshalIndent(xml_tree.Territories, "", "  ")
    check(err)

    fmt.Println(string(json_bytes))
  },
}

var xmlToPbCmd = &cobra.Command{
  Use: "pb",
  Short: "Convert xml format to pb.",
  Run: func (cmd *cobra.Command, args []string) {
    var path = args[0]
    xml_bytes, err := ioutil.ReadFile(path)
    check(err)

    xml_tree := &md_xml.PhoneNumberMetadata{}
    err = xml.Unmarshal(xml_bytes, xml_tree)
    check(err)

    json_bytes, err := json.Marshal(xml_tree.Territories)
    check(err)

    pb_tree := &md_pb.PhoneMetadataCollection{}
    err = jsonpb.UnmarshalString(string(json_bytes), pb_tree)
    check(err)

    pb_bytes, err := proto.Marshal(pb_tree)
    check(err)

    os.Stdout.Write(pb_bytes)
  },
}

var xmlToXmlCmd = &cobra.Command{
  Use: "xml",
  Short: "Convert xml format to xml.",
  Run: func (cmd *cobra.Command, args []string) {
    var path = args[0]
    data, err := ioutil.ReadFile(path)
    check(err)

    xml_tree := &md_xml.PhoneNumberMetadata{}
    err = xml.Unmarshal(data, xml_tree)
    check(err)

    xml_bytes, err := xml.Marshal(xml_tree)
    check(err)

    fmt.Println(string(xml_bytes))
  },
}

var xmlToCmd = &cobra.Command{
  Use: "to",
  Short: "Convert xml format to...",
}

var xmlCmd = &cobra.Command{
  Use:   "xml",
  Short: "Manage xml format.",
}

// metadata json to...

var jsonToJsonCmd = &cobra.Command{
  Use: "json",
  Short: "Convert json format to json.",
  Run: func (cmd *cobra.Command, args []string) {
    var path = args[0]
    data, err := ioutil.ReadFile(path)
    check(err)

    pb_tree := &md_pb.PhoneMetadataCollection{}
    err = jsonpb.UnmarshalString(string(data), pb_tree)
    check(err)

    json_bytes, err := json.MarshalIndent(pb_tree, "", "  ")
    check(err)

    fmt.Println(string(json_bytes))
  },
}

var jsonToPbCmd = &cobra.Command{
  Use: "pb",
  Short: "Convert json format to pb.",
  Run: func (cmd *cobra.Command, args []string) {
    var path = args[0]
    json_bytes, err := ioutil.ReadFile(path)
    check(err)

    pb_tree := &md_pb.PhoneMetadataCollection{}
    err = jsonpb.UnmarshalString(string(json_bytes), pb_tree)
    check(err)

    pb_bytes, err := proto.Marshal(pb_tree)
    check(err)

    os.Stdout.Write(pb_bytes)
  },
}

var jsonToXmlCmd = &cobra.Command{
  Use: "xml",
  Short: "Convert json format to xml.",
  Run: func (cmd *cobra.Command, args []string) {
    var path = args[0]
    json_bytes, err := ioutil.ReadFile(path)
    check(err)

    xml_tree := &md_xml.PhoneNumberMetadata{
      Territories: &md_xml.Territories{},
    }
    err = json.Unmarshal(json_bytes, xml_tree.Territories)
    check(err)

    xml_bytes, err := xml.Marshal(xml_tree)
    check(err)

    fmt.Println(string(xml_bytes))
  },
}

var jsonToCmd = &cobra.Command{
  Use: "to",
  Short: "Convert json format to...",
}

var jsonCmd = &cobra.Command{
  Use:   "json",
  Short: "Manage json format.",
}

// metadata pb to...

var pbToJsonCmd = &cobra.Command{
  Use: "json",
  Short: "Convert pb format to json.",
  Run: func (cmd *cobra.Command, args []string) {
    var path = args[0]
    pb_bytes, err := ioutil.ReadFile(path)
    check(err)

    pb_tree := &md_pb.PhoneMetadataCollection{}
    err = proto.Unmarshal(pb_bytes, pb_tree)
    check(err)

    json_bytes, err := json.MarshalIndent(pb_tree, "", "  ")
    check(err)

    fmt.Println(string(json_bytes))
  },
}

var pbToPbCmd = &cobra.Command{
  Use: "pb",
  Short: "Convert pb format to pb.",
  Run: func (cmd *cobra.Command, args []string) {
    var path = args[0]
    data, err := ioutil.ReadFile(path)
    check(err)

    pb_tree := &md_pb.PhoneMetadataCollection{}
    err = proto.Unmarshal(data, pb_tree)
    check(err)

    pb_bytes, err := proto.Marshal(pb_tree)
    check(err)

    os.Stdout.Write(pb_bytes)
  },
}

var pbToXmlCmd = &cobra.Command{
  Use: "xml",
  Short: "Convert pb format to xml.",
  Run: func (cmd *cobra.Command, args []string) {
    var path = args[0]
    pb_bytes, err := ioutil.ReadFile(path)
    check(err)

    pb_tree := &md_pb.PhoneMetadataCollection{}
    err = proto.Unmarshal(pb_bytes, pb_tree)
    check(err)

    json_bytes, err := json.MarshalIndent(pb_tree, "", "  ")
    check(err)

    xml_tree := &md_xml.PhoneNumberMetadata{
      Territories: &md_xml.Territories{},
    }
    err = json.Unmarshal(json_bytes, xml_tree.Territories)
    check(err)

    xml_bytes, err := xml.Marshal(xml_tree)
    check(err)

    fmt.Println(string(xml_bytes))
  },
}

var pbToGoCmd = &cobra.Command{
  Use: "go",
  Short: "Convert pb format to go.",
  Run: func (cmd *cobra.Command, args []string) {
    var path = args[0]
    pb_bytes, err := ioutil.ReadFile(path)
    check(err)

    pb_tree := &md_pb.PhoneMetadataCollection{}
    err = proto.Unmarshal(pb_bytes, pb_tree)
    check(err)

    pretty.Println(pb_tree)
  },
}

var pbToCmd = &cobra.Command{
  Use: "to",
  Short: "Convert pb format to...",
}

var pbCmd = &cobra.Command{
  Use:   "pb",
  Short: "Manage pb format.",
  Run: func (cmd *cobra.Command, args []string) {
    fmt.Println("Hi there!")
  },
}

var metadataCmd = &cobra.Command{
  Use:   "metadata",
  Short: "Managing phone metadata data.",
}

func init() {
  // metadata xml to...
  xmlToCmd.AddCommand(xmlToJsonCmd)
  xmlToCmd.AddCommand(xmlToPbCmd)
  xmlToCmd.AddCommand(xmlToXmlCmd)
  xmlCmd.AddCommand(xmlToCmd)
  metadataCmd.AddCommand(xmlCmd)

  // metadata json to...
  jsonToCmd.AddCommand(jsonToJsonCmd)
  jsonToCmd.AddCommand(jsonToPbCmd)
  jsonToCmd.AddCommand(jsonToXmlCmd)
  jsonCmd.AddCommand(jsonToCmd)
  metadataCmd.AddCommand(jsonCmd)

  // metadata pb to...
  pbToCmd.AddCommand(pbToJsonCmd)
  pbToCmd.AddCommand(pbToPbCmd)
  pbToCmd.AddCommand(pbToXmlCmd)
  pbToCmd.AddCommand(pbToGoCmd)
  pbCmd.AddCommand(pbToCmd)
  metadataCmd.AddCommand(pbCmd)

  // metadata...
  RootCmd.AddCommand(metadataCmd)
}
