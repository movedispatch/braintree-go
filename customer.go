package braintree

import "encoding/xml"

type Customer struct {
	XMLName    string      `xml:"customer"`
	Id         string      `xml:"id,omitempty"`
	FirstName  string      `xml:"first-name,omitempty"`
	LastName   string      `xml:"last-name,omitempty"`
	Company    string      `xml:"company,omitempty"`
	Email      string      `xml:"email,omitempty"`
	Phone      string      `xml:"phone,omitempty"`
	Fax        string      `xml:"fax,omitempty"`
	Website    string      `xml:"website,omitempty"`
	CreditCard *CreditCard `xml:"credit-card,omitempty`
}

func (this Customer) ToXML() ([]byte, error) {
	xml, err := xml.Marshal(this)
	if err != nil {
		return []byte{}, err
	}
	return xml, nil
}