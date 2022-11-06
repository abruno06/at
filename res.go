package at

import "strings"

// Response represents a numerical response from AT command
type Response struct {
	// ID is usually used to detect an Responseion from a number in string.
	ID int
	// Description contains a human-readable description of an Responseion.
	Description string
}

// StringResponse represents a string Responseion from AT command
type StringResponse struct {
	// ID is usually used to detect an Responseion from a substring in string.
	ID string
	// Description contains a human-readable description of an Responseion.
	Description string
}

// UnknownResponse represents an Responseion that was parsed incorrectly or was not parsed at all.
var UnknownResponse = Response{ID: -1, Description: "-"}

// UnknownStringResponse represents a string Responseion that was parsed incorrectly or was not parsed at all.
var UnknownStringResponse = StringResponse{ID: "nil", Description: "Unknown"}

// KillCmd is an artificial AT command that may be successfully sent to device in order
// to emulate the response from it. In other words, if a connection with device stalled and
// no bytes could be read, then this command is used to read something and then close the connection.

const KillCmd = "AT_KILL"

// NoopCmd is like a ping command that signals that the device is responsive.
const NoopCmd = "AT"

type (
	responseMap     map[int]Response
	stringResponses []StringResponse
)

func (r responseMap) Decode(id int) Response {
	if res, ok := r[id]; ok {
		return res
	}
	return UnknownResponse
}

func (s stringResponses) Decode(str string) StringResponse {
	for _, v := range s {
		if strings.HasPrefix(str, v.ID) {
			return v
		}
	}
	return UnknownStringResponse
}

// ModemState represents the device state including cellular Response ,
// signal quality, current operator name, service status.
type DevicModemState struct {
	ServiceState   Response
	ServiceDomain  Response
	RoamingState   Response
	SystemMode     Response
	SystemSubmode  Response
	SimState       Response
	ModelName      string
	OperatorName   string
	IMEI           string
	SignalStrength int
}

/*
// NewDeviceState returns a clean state with unknown Responseions.
func NewDeviceState() *DeviceState {
	return &DeviceState{
		ServiceState:  UnknownResponse,
		ServiceDomain: UnknownResponse,
		RoamingState:  UnknownResponse,
		SystemMode:    UnknownResponse,
		SystemSubmode: UnknownResponse,
		SimState:      UnknownResponse,
	}
}

var sim = ResponseMap{
	0:   Response{0, "Invalid USIM card or pin code locked"},
	1:   Response{1, "Valid USIM card"},
	2:   Response{2, "USIM is invalid for cellular service"},
	3:   Response{3, "USIM is invalid for packet service"},
	4:   Response{4, "USIM is not valid for cellular nor packet services"},
	255: Response{255, "USIM card is not exist"},
}

// SimStates represent the possible data card states.
var SimStates = struct {
	Resolve func(int) Response

	Invalid     Response
	Valid       Response
	InvalidCS   Response
	InvalidPS   Response
	InvalidCSPS Response
	NoCard      Response
}{
	func(id int) Response { return sim.Resolve(id) },

	sim[0], sim[1], sim[2], sim[3], sim[4], sim[255],
}

var service = ResponseMap{
	0: Response{0, "No service"},
	1: Response{1, "Restricted service"},
	2: Response{2, "Valid service"},
	3: Response{3, "Restricted regional service"},
	4: Response{4, "Power-saving and deep sleep state"},
}

// ServiceStates represent the possible service states.
var ServiceStates = struct {
	Resolve func(int) Response

	None               Response
	Restricted         Response
	Valid              Response
	RestrictedRegional Response
	PowerSaving        Response
}{
	func(id int) Response { return service.Resolve(id) },

	service[0], service[1], service[2], service[3], service[4],
}

var domain = ResponseMap{
	0: Response{0, "No service"},
	1: Response{1, "Cellular service only"},
	2: Response{2, "Packet service only"},
	3: Response{3, "Packet and Cellular services"},
	4: Response{4, "Searching"},
}

// ServiceDomains represent the possible service domains.
var ServiceDomains = struct {
	Resolve func(int) Response

	None               Response
	Restricted         Response
	Valid              Response
	RestrictedRegional Response
	PowerSaving        Response
}{
	func(id int) Response { return domain.Resolve(id) },

	domain[0], domain[1],
	domain[2], domain[3], domain[4],
}

var roaming = ResponseMap{
	0: Response{0, "Non roaming"},
	1: Response{1, "Roaming"},
}

// RoamingStates represent the state of roaming.
var RoamingStates = struct {
	Resolve func(int) Response

	NotRoaming Response
	Roaming    Response
}{
	func(id int) Response { return roaming.Resolve(id) },

	roaming[0], roaming[1],
}

var mode = ResponseMap{
	0:  Response{0, "No service"},
	1:  Response{1, "AMPS"},
	2:  Response{2, "CDMA"},
	3:  Response{3, "GSM/GPRS"},
	4:  Response{4, "HDR"},
	5:  Response{5, "WCDMA"},
	6:  Response{6, "GPS"},
	7:  Response{7, "GSM/WCDMA"},
	8:  Response{8, "CDMA/HDR HYBRID"},
	15: Response{15, "TD-SCDMA"},
}

// SystemModes represent the possible system operating modes.
var SystemModes = struct {
	Resolve func(int) Response

	NoService Response
	AMPS      Response
	CDMA      Response
	GsmGprs   Response
	HDR       Response
	WCDMA     Response
	GPS       Response
	GsmWcdma  Response
	CdmaHdr   Response
	SCDMA     Response
}{
	func(id int) Response { return mode.Resolve(id) },

	mode[0], mode[1], mode[2], mode[3], mode[4],
	mode[5], mode[6], mode[7], mode[8], mode[15],
}

var submode = ResponseMap{
	0:  Response{0, "No service"},
	1:  Response{1, "GSM"},
	2:  Response{2, "GPRS"},
	3:  Response{3, "EDGE"},
	4:  Response{4, "WCDMA"},
	5:  Response{5, "HSDPA"},
	6:  Response{6, "HSUPA"},
	7:  Response{7, "HSDPA and HSUPA"},
	8:  Response{8, "TD-SCDMA"},
	9:  Response{9, "HSPA+"},
	17: Response{17, "HSPA+(64QAM)"},
	18: Response{18, "HSPA+(MIMO)"},
}

// SystemSubmodes represent the possible system operating submodes.
var SystemSubmodes = struct {
	Resolve func(int) Response

	NoService  Response
	GSM        Response
	GPRS       Response
	EDGE       Response
	WCDMA      Response
	HSDPA      Response
	HSUPA      Response
	HsdpaHsupa Response
	SCDMA      Response
	HspaPlus   Response
	Hspa64QAM  Response
	HspaMIMO   Response
}{
	func(id int) Response { return submode.Resolve(id) },

	submode[0], submode[1], submode[2], submode[3],
	submode[4], submode[5], submode[6], submode[7],
	submode[8], submode[9], submode[17], submode[18],
}

var result = stringResponses{
	{"AT", "Noop"},
	{"OK", "Success"},
	{"CONNECT", "Connect"},
	{"RING", "Ringing"},
	{"NO CARRIER", "No carrier"},
	{"ERROR", "Error"},
	{"NO DIALTONE", "No dialtone"},
	{"BUSY", "Busy"},
	{"NO ANSWER", "No answer"},
	{"+CME ERROR:", "CME Error"},
	{"+CMS ERROR:", "CMS Error"},
	{"COMMAND NOT SUPPORT", "Command is not supported"},
	{"TOO MANY PARAMETERS", "Too many parameters"},
	{"AT_KILL", "Timeout"},
}

// FinalResults represent the possible replies from a modem.
var FinalResults = struct {
	Resolve func(string) StringResponse

	Noop              StringResponse
	Ok                StringResponse
	Connect           StringResponse
	Ring              StringResponse
	NoCarrier         StringResponse
	Error             StringResponse
	NoDialtone        StringResponse
	Busy              StringResponse
	NoAnswer          StringResponse
	CmeError          StringResponse
	CmsError          StringResponse
	NotSupported      StringResponse
	TooManyParameters StringResponse
	Timeout           StringResponse
}{
	func(str string) StringResponse { return result.Resolve(str) },

	result[0], result[1], result[2], result[3],
	result[4], result[5], result[6], result[7],
	result[8], result[9], result[10], result[11],
	result[12], result[13],
}

var resultReporting = ResponseMap{
	0: Response{0, "Disabled"},
	1: Response{1, "Enabled"},
	2: Response{2, "Exit"},
}

// UssdResultReporting represents the available Responseions of USSD reporting.
var UssdResultReporting = struct {
	Resolve func(int) Response

	Disable Response
	Enable  Response
	Exit    Response
}{
	func(id int) Response { return resultReporting.Resolve(id) },

	resultReporting[0],
	resultReporting[1],
	resultReporting[2],
}

var reports = stringResponses{
	{"+CUSD:", "USSD reply"},
	{"+CMTI:", "Incoming SMS"},
	{"^RSSI:", "Signal strength"},
	{"^BOOT:", "Boot handshake"},
	{"^MODE:", "System mode"},
	{"^SRVST:", "Service state"},
	{"^SIMST:", "Sim state"},
	{"^STIN:", "STIN"},
	{"+CLIP:", "Incoming Caller ID"},
}

// Reports represent the possible state reports from a modem.
var Reports = struct {
	Resolve func(string) StringResponse

	Ussd           StringResponse
	Message        StringResponse
	SignalStrength StringResponse
	BootHandshake  StringResponse
	Mode           StringResponse
	ServiceState   StringResponse
	SimState       StringResponse
	Stin           StringResponse
	CallerID       StringResponse
}{
	func(str string) StringResponse { return reports.Resolve(str) },

	reports[0], reports[1], reports[2], reports[3],
	reports[4], reports[5], reports[6], reports[7], reports[8],
}

var mem = stringResponses{
	{"ME", "NV RAM"},
	{"MT", "ME-associated storage"},
	{"SM", "Sim message storage"},
	{"SR", "State report storage"},
}

// MemoryTypes represent the available Responseions of message storage.
var MemoryTypes = struct {
	Resolve func(string) StringResponse

	NvRAM       StringResponse
	Associated  StringResponse
	Sim         StringResponse
	StateReport StringResponse
}{
	func(str string) StringResponse { return mem.Resolve(str) },

	mem[0], mem[1], mem[2], mem[3],
}

var delResponses = ResponseMap{
	0: Response{0, "Delete message by index"},
	1: Response{1, "Delete all read messages except MO"},
	2: Response{2, "Delete all read messages except unsent MO"},
	3: Response{3, "Delete all except unread"},
	4: Response{4, "Delete all messages"},
}

// DeleteResponseions represent the available Responseions of message deletion masks.
var DeleteResponseions = struct {
	Resolve func(int) Response

	Index            Response
	AllReadNotMO     Response
	AllReadNotUnsent Response
	AllNotUnread     Response
	All              Response
}{
	func(id int) Response { return resultReporting.Resolve(id) },

	delResponses[0], delResponses[1], delResponses[2], delResponses[3], delResponses[4],
}

var msgFlags = ResponseMap{
	0: Response{0, "Unread"},
	1: Response{1, "Read"},
	2: Response{2, "Unsent"},
	3: Response{3, "Sent"},
	4: Response{4, "Any"},
}

// MessageFlags represent the available states of messages in memory.
var MessageFlags = struct {
	Resolve func(int) Response

	Unread Response
	Read   Response
	Unsent Response
	Sent   Response
	Any    Response
}{
	func(id int) Response { return resultReporting.Resolve(id) },

	msgFlags[0], msgFlags[1], msgFlags[2], msgFlags[3], msgFlags[4],
}

var callerIDType = ResponseMap{
	129: Response{129, "Network Specific Caller ID"},
	145: Response{145, "International Caller ID"},
}

// CallerIDTypes represent the possible caller id types.
var CallerIDTypes = struct {
	Resolve func(int) Response

	NetworkSpecific Response
	International   Response
}{
	func(id int) Response { return callerIDType.Resolve(id) },

	callerIDType[129], callerIDType[145],
}

var callerIDValidity = ResponseMap{
	0: Response{0, "Valid"},
	1: Response{1, "Rejected by calling party"},
	2: Response{2, "Denied by network"},
}

// CallerIDValidityStates represent the possible caller id validity states.
var CallerIDValidityStates = struct {
	Resolve func(int) Response

	Valid    Response
	Rejected Response
	Denied   Response
}{
	func(id int) Response { return callerIDValidity.Resolve(id) },

	callerIDValidity[0], callerIDValidity[1], callerIDValidity[2],
}


*/
