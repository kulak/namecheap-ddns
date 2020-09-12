package main

import (
	"encoding/xml"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestXmlResponse(t *testing.T) {
	const responseStr = `
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
`
	var nr NamecheapResponse
	err := xml.Unmarshal([]byte(responseStr), &nr)
	require.NoError(t, err)
	t.Log(nr)
	require.Equal(t, "SETDNSHOST", nr.Command)
	require.Equal(t, 1, nr.ErrCount)
	require.Equal(t, "Domain name not found", nr.Errors.Err1)
	require.Equal(t, 1, nr.ResponseCount)
	require.Equal(t, 1, len(nr.ResponsesContainer.Responses))
	require.Equal(t, "Validation error; not found; domain name(s)", nr.ResponsesContainer.Responses[0].ResponseString)
	require.Equal(t, true, nr.Done)
	require.Equal(t, 0, len(nr.Debug))
}
