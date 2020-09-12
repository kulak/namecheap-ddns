package main

/*
<?xml version="1.0"?>
<interface-response>
	<Command>SETDNSHOST</Command>
	<Language>eng</Language>
	<ErrCount>1</ErrCount>
	<errors>
		<Err1>Domain name not found</Err1>
	</errors>
	<ResponseCount>1</ResponseCount>
	<responses>
		<response>
			<ResponseNumber>316153</ResponseNumber>
			<ResponseString>Validation error; not found; domain name(s)</ResponseString>
		</response>
	</responses>
	<Done>true</Done>
	<debug><![CDATA[]]></debug>
</interface-response>
*/

// NamecheapResponse holds namecheap response message.
// Note: Errors are not handled dynamically and only 1st error is properly parsed.
type NamecheapResponse struct {
	Command            string      `xml:"Command"`
	Language           string      `xml:"Language"`
	ErrCount           int         `xml:"ErrCount"`
	Errors             NRErrors    `xml:"errors"`
	ResponseCount      int         `xml:"ResponseCount"`
	ResponsesContainer NRResponses `xml:"responses"`
	Done               bool        `xml:"Done"`
	Debug              []byte      `xml:"debug"`
}

type NRErrors struct {
	Err1 string `xml:"Err1"`
}

type NRResponses struct {
	Responses []NRResponse `xml:"response"`
}

type NRResponse struct {
	ResponseNumber int    `xml:"ResponseNumber"`
	ResponseString string `xml:"ResponseString"`
}
