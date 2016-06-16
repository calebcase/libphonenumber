package xml

import (
  "strings"
)

func check(e error) {
  if e != nil {
    panic(e)
  }
}

type Pattern struct {
  Value string
}

func (p *Pattern) UnmarshalText(text []byte) (err error) {
  p.Value = string(text[0:len(text)])
  return nil
}

func (p *Pattern) MarshalText() (json []byte, err error) {
  // Patterns my have newlines and leading whitespace. We remove this
  // whitespace when rendering.

  // Split the string on newlines.
  vs := strings.Split(p.Value, "\n")

  // Create a temporary array for storing the stripped lines.
  vsm := make([]string, len(vs))

  // For each line in the pattern:
  for i, v := range vs {
    // Remove leading spaces and tabs.
    vsm[i] = strings.TrimLeft(v, " \t")
  }

  // Join the resulting array without newlines.
  json = []byte(strings.Join(vsm, ""))
  err = nil

  return
}

type NumberFormat struct {
  Pattern                              *string  `xml:"pattern,omitempty"                              json:"pattern,omitempty"`
  Format                               *string  `xml:"format,omitempty"                               json:"format,omitempty"`
  LeadingDigitsPattern                 []string `xml:"leadingDigitsPattern,omitempty"                 json:"leading_digits_pattern,omitempty"`
  NationalPrefixFormattingRule         *string  `xml:"nationalPrefixFormattingRule,omitempty"         json:"national_prefix_formatting_rule,omitempty"`
  NationalPrefixOptionalWhenFormatting *bool    `xml:"nationalPrefixOptionalWhenFormatting,omitempty" json:"national_prefix_optional_when_formatting,omitempty"`
  DomesticCarrierCodeFormattingRule    *string  `xml:"domesticCarrierCodeFormattingRule,omitempty"    json:"domestic_carrier_code_formatting_rule,omitempty"`
}

type PhoneNumberDesc struct {
  NationalNumberPattern *Pattern `xml:"nationalNumberPattern,omitempty" json:"national_number_pattern,omitempty"`
  PossibleNumberPattern *Pattern `xml:"possibleNumberPattern,omitempty" json:"possible_number_pattern,omitempty"`
  ExampleNumber         *string  `xml:"exampleNumber,omitempty"         json:"example_number,omitempty"`
}

type Territory struct {
  // Attributes
  Id                            *string  `xml:"id,attr,omitempty"                            json:"id,omitempty"`
  CountryCode                   *int32   `xml:"countryCode,attr,omitempty"                   json:"country_code,omitempty"`
  InternationalPrefix           *string  `xml:"internationalPrefix,attr,omitempty"           json:"international_prefix,omitempty"`
  PreferredInternationalPrefix  *string  `xml:"preferredInternationalPrefix,attr,omitempty"  json:"preferred_international_prefix,omitempty"`
  NationalPrefix                *string  `xml:"nationalPrefix,attr,omitempty"                json:"national_prefix,omitempty"`
  PreferredExtnPrefix           *string  `xml:"preferredExtnPrefix,attr,omitempty"           json:"preferred_extn_prefix,omitempty"`
  NationalPrefixForParsing      *Pattern `xml:"nationalPrefixForParsing,attr,omitempty"      json:"national_prefix_for_parsing,omitempty"`
  NationalPrefixTransformRule   *string  `xml:"nationalPrefixTransformRule,attr,omitempty"   json:"national_prefix_transform_rule,omitempty"`
  SameMobileAndFixedLinePattern *bool    `xml:"sameMobileAndFixedLinePattern,attr,omitempty" json:"same_mobile_and_fixed_line_pattern,omitempty"`
  MainCountryForCode            *bool    `xml:"mainCountryForCode,attr,omitempty"            json:"main_country_for_code,omitempty"`
  LeadingDigits                 *string  `xml:"leadingDigits,attr,omitempty"                 json:"leading_digits,omitempty"`
  LeadingZeroPossible           *bool    `xml:"leadingZeroPossible,attr,omitempty"           json:"leading_zero_possible,omitempty"`
  MobileNumberPortableRegion    *bool    `xml:"mobileNumberPortableRegion,attr,omitempty"    json:"mobile_number_portable_region,omitempty"`

  // Phone Number Description Children
  GeneralDesc             *PhoneNumberDesc `xml:"generalDesc,omitempty"             json:"general_desc,omitempty"`
  FixedLine               *PhoneNumberDesc `xml:"fixedLine,omitempty"               json:"fixed_line,omitempty"`
  Mobile                  *PhoneNumberDesc `xml:"mobile,omitempty"                  json:"mobile,omitempty"`
  TollFree                *PhoneNumberDesc `xml:"tollFree,omitempty"                json:"toll_free,omitempty"`
  PremiumRate             *PhoneNumberDesc `xml:"premiumRate,omitempty"             json:"premium_rate,omitempty"`
  SharedCost              *PhoneNumberDesc `xml:"sharedCost,omitempty"              json:"shared_cost,omitempty"`
  PersonalNumber          *PhoneNumberDesc `xml:"personalNumber,omitempty"          json:"personal_number,omitempty"`
  Voip                    *PhoneNumberDesc `xml:"voip,omitempty"                    json:"voip,omitempty"`
  Pager                   *PhoneNumberDesc `xml:"pager,omitempty"                   json:"pager,omitempty"`
  Uan                     *PhoneNumberDesc `xml:"uan,omitempty"                     json:"uan,omitempty"`
  Emergency               *PhoneNumberDesc `xml:"emergency,omitempty"               json:"emergency,omitempty"`
  Voicemail               *PhoneNumberDesc `xml:"voicemail,omitempty"               json:"voicemail,omitempty"`
  ShortCode               *PhoneNumberDesc `xml:"shortCode,omitempty"               json:"short_code,omitempty"`
  StandardRate            *PhoneNumberDesc `xml:"standardRate,omitempty"            json:"standard_rate,omitempty"`
  CarrierSpecific         *PhoneNumberDesc `xml:"carrierSpecific,omitempty"         json:"carrier_specific,omitempty"`
  NoInternationalDialling *PhoneNumberDesc `xml:"noInternationalDialling,omitempty" json:"no_international_dialling,omitempty"`

  // Number Format Children
  NumberFormat     []*NumberFormat `xml:"numberFormat,omitempty"     json:"number_format,omitempty"`
  IntlNumberFormat []*NumberFormat `xml:"intlNumberFormat,omitempty" json:"intl_number_format,omitempty"`
}

type Territories struct {
  Territory []*Territory `xml:"territory,omitempty" json:"metadata,omitempty"`
}

type PhoneNumberMetadata struct {
  XMLName string `xml:"phoneNumberMetadata"`
  Territories *Territories `xml:"territories,omitempty" json:"phone_number_metadata,omitempty"`
}
