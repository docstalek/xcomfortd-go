package xc

const (
	MGW_PT_TX     = 0xB1
	MGW_PT_CONFIG = 0xB2
	MGW_PT_RX     = 0xC1
	MGW_PT_STATUS = 0xC3
	MGW_PT_FW     = 0xD1
)

/* Events that can be sent to datapoints with MGW_PT_TX.  Not all events
   are valid for all devices. */

const (
	TX_EVENT_SWITCH         = 0x0A
	TX_EVENT_DIM            = 0x0D // Dimmer, Analogue
	TX_EVENT_JALO           = 0x0E // Jalousie
	TX_EVENT_PUSHBUTTON     = 0x50 // All types
	TX_EVENT_REQUEST        = 0x0B // All mains powered
	TX_EVENT_BASIC_MODE     = 0x80
	TX_EVENT_INT16_1POINT   = 0x11
	TX_EVENT_FLOAT          = 0x1A // Room Manager, Analogue input
	TX_EVENT_RM_TIME        = 0x2A // Room Manager, HRV
	TX_EVENT_RM_DATE        = 0x2B // Room Manager, HRV
	TX_EVENT_RC_DATA        = 0x2C
	TX_EVENT_UINT32         = 0x30 // Room Manager
	TX_EVENT_UINT32_1POINT  = 0x31
	TX_EVENT_UINT32_2POINT  = 0x32
	TX_EVENT_UINT32_3POINT  = 0x33
	TX_EVENT_UINT16         = 0x40
	TX_EVENT_UINT16_1POINT  = 0x41
	TX_EVENT_UINT16_2POINT  = 0x42
	TX_EVENT_UINT16_3POINT  = 0x43
	TX_EVENT_DIMPLEX_CONFIG = 0x44
	TX_EVENT_DIMPLEX_TEMP   = 0x45
	TX_EVENT_HRV_IN         = 0x46
)

const (
	// TX_EVENT_SWITCH
	TX_EVENTDATA_OFF = 0x00
	TX_EVENTDATA_ON  = 0x01

	// TX_EVENT_DIM
	TX_EVENTDATA_STOP     = 0x00
	TX_EVENTDATA_DARKER   = 0x04
	TX_EVENTDATA_BRIGHTER = 0x0F
	TX_EVENTDATA_PERCENT  = 0x40

	// TX_EVENT_JALO
	TX_EVENTDATA_CLOSE      = 0x00
	TX_EVENTDATA_OPEN       = 0x01
	TX_EVENTDATA_JSTOP      = 0x02
	TX_EVENTDATA_STEP_CLOSE = 0x10
	TX_EVENTDATA_STEP_OPEN  = 0x11

	// TX_EVENT_PUSHBUTTON
	TX_EVENTDATA_UP            = 0x50
	TX_EVENTDATA_DOWN          = 0x51
	TX_EVENTDATA_UP_PRESSED    = 0x54
	TX_EVENTDATA_UP_RELEASED   = 0x55
	TX_EVENTDATA_DOWN_PRESSED  = 0x56
	TX_EVENTDATA_DOWN_RELEASED = 0x57

	// TX_EVENT_REQUEST
	TX_EVENTDATA_DUMMY = 0x00

	// TX_EVENT_BASIC_MODE
	TX_EVENTDATA_LEARNMODE_OFF   = 0x00
	TX_EVENTDATA_LEARNMODE_ON    = 0x01
	TX_EVENTDATA_ASSIGN_ACTUATOR = 0x10
	TX_EVENTDATA_REMOVE_ACTUATOR = 0x20
	TX_EVENTDATA_REMOVE_SENSORS  = 0x30
)

/* Events that can be received via MGW_PT_RX.  These are events that are
   known by MRF, they may not all be applicable to datapoints. */

const (
	RX_EVENT_ON            = 0x50
	RX_EVENT_OFF           = 0x51
	RX_EVENT_SWITCH_ON     = 0x52
	RX_EVENT_SWITCH_OFF    = 0x53
	RX_EVENT_UP_PRESSED    = 0x54
	RX_EVENT_UP_RELEASED   = 0x55
	RX_EVENT_DOWN_PRESSED  = 0x56
	RX_EVENT_DOWN_RELEASED = 0x57
	RX_EVENT_FORCED        = 0x5A
	RX_EVENT_SINGLE_ON     = 0x5B
	RX_EVENT_VALUE         = 0x62
	RX_EVENT_TOO_COLD      = 0x63
	RX_EVENT_TOO_WARM      = 0x64
	RX_EVENT_STATUS        = 0x70
	RX_EVENT_STATUS_EXT    = 0x73
	RX_EVENT_BASIC_MODE    = 0x80
)

// Data types for RX events

const (
	RX_DATA_TYPE_NO_DATA       = 0x00
	RX_DATA_TYPE_PERCENT       = 0x01
	RX_DATA_TYPE_UINT8         = 0x02
	RX_DATA_TYPE_INT16_1POINT  = 0x03
	RX_DATA_TYPE_FLOAT         = 0x04
	RX_DATA_TYPE_UINT16        = 0x0D
	RX_DATA_TYPE_UINT16_1POINT = 0x21
	RX_DATA_TYPE_UINT16_2POINT = 0x22
	RX_DATA_TYPE_UINT16_3POINT = 0x23
	RX_DATA_TYPE_UINT32        = 0x0E
	RX_DATA_TYPE_UINT32_1POINT = 0x0F
	RX_DATA_TYPE_UINT32_2POINT = 0x10
	RX_DATA_TYPE_UINT32_3POINT = 0x11
	RX_DATA_TYPE_RC_DATA       = 0x17
	RX_DATA_TYPE_RM_TIME       = 0x1E
	RX_DATA_TYPE_RM_DATE       = 0x1F
	RX_DATA_TYPE_ROSETTA       = 0x35
	RX_DATA_TYPE_HRV_OUT       = 0x37
	RX_DATA_TYPE_SERIAL_NUMBER = 0x39
)

// INFO_SHORT values

const (
	RX_IS_OFF    = 0x00
	RX_IS_ON     = 0x01
	RX_IS_OFF_NG = 0x02
	RX_IS_ON_NG  = 0x03

	RX_IS_STOP  = 0x00
	RX_IS_OPEN  = 0x01
	RX_IS_CLOSE = 0x02
)

/* Config commands that can be sent to the stick itself. These are sent with
   the MGW_PT_CONFIG message and responses are received via the MGW_PT_STATUS
   message. */

const (
	CONF_CONNEX          = 0x02
	CONF_RS232_BAUD      = 0x03
	CONF_SEND_OK_MRF     = 0x04
	CONF_RS232_FLOW      = 0x05
	CONF_RS232_CRC       = 0x06
	CONF_TIMEACCOUNT     = 0x0A
	CONF_COUNTER_RX      = 0x0B
	CONF_COUNTER_TX      = 0x0C
	CONF_SERIAL          = 0x0E
	CONF_LED             = 0x0F
	CONF_LED_DIM         = 0x1A
	CONF_RELEASE         = 0x1B
	CONF_SEND_CLASS      = 0x1D
	CONF_SEND_RFSEQNO    = 0x1E
	CONF_BACK_TO_FACTORY = 0x1F
)

const (
	// CONF_CONNEX
	CF_DATA_AUTO  = 0x01
	CF_DATA_USB   = 0x02
	CF_DATA_RS232 = 0x03

	CF_DATA_STATUS = 0x00

	// CONF_RS232_BAUD
	CF_DATA_BD1200  = 0x01
	CF_DATA_BD2400  = 0x02
	CF_DATA_BD4800  = 0x03
	CF_DATA_BD9600  = 0x04
	CF_DATA_BD14400 = 0x05
	CF_DATA_BD19200 = 0x06
	CF_DATA_BD38400 = 0x07
	CF_DATA_BD56700 = 0x08

	// CONF_RS232_CRC, CONF_SEND_RFSEQNO, etc.
	CF_DATA_CLEAR = 0x0F
	CF_DATA_SET   = 0x01

	// CONF_COUNTER_RX, TX, SERIAL
	CF_DATA_GET = 0x00

	// CONF_LED
	CF_DATA_LED_STANDARD  = 0x01
	CF_DATA_REVERSE_GREEN = 0x02
	CF_DATA_LED_OFF       = 0x03

	// CONF_RELEASE
	CF_DATA_GET_REVISION = 0x10

	// CONF_BACK_TO_FACTORY
	CF_DATA_BTF_GW  = 0x0F
	CF_DATA_BTF_MRF = 0xF0
	CF_DATA_BTF_ALL = 0xFF
)

/* Status message received from the stick itself, via the MGW_PT_STATUS
   message. */

const (
	STATUS_TYPE_CONNEX       = 0x02
	STATUS_TYPE_RS232_BAUD   = 0x03
	STATUS_TYPE_RS232_FLOW   = 0x05
	STATUS_TYPE_RS232_CRC    = 0x06
	STATUS_TYPE_ERROR        = 0x09
	STATUS_TYPE_TIMEACCOUNT  = 0x0A
	STATUS_TYPE_SEND_OK_MRF  = 0x0D
	STATUS_TYPE_SERIAL       = 0x0E
	STATUS_TYPE_LED          = 0x0F
	STATUS_TYPE_LED_DIM      = 0x1A
	STATUS_TYPE_RELEASE      = 0x1B
	STATUS_TYPE_OK           = 0x1C
	STATUS_TYPE_SEND_CLASS   = 0x1D
	STATUS_TYPE_SEND_RFSEQNO = 0x1E
)

const (
	// STATUS_TYPE_ERROR
	STATUS_GENERAL     = 0x00
	STATUS_UNKNOWN     = 0x01
	STATUS_DP_OOR      = 0x02
	STATUS_BUSY_MRF    = 0x03
	STATUS_BUSY_MRF_RX = 0x04
	STATUS_TX_MSG_LOST = 0x05
	STATUS_NO_ACK      = 0x06

	// STATUS_TYPE_TIMEACCOUNT
	STATUS_DATA    = 0x00
	STATUS_IS_0    = 0x01
	STATUS_LESS_10 = 0x02
	STATUS_MORE_15 = 0x03

	// STATUS_TYPE_RELEASE
	STATUS_REVISION = 0x10

	// STATUS_TYPE_OK
	STATUS_OK_MRF    = 0x04
	STATUS_OK_CONFIG = 0x05
	STATUS_OK_BTFACT = 0xCE
)

const (
	STATUS_DATA_OKMRF_NOINFO     = 0x00
	STATUS_DATA_OKMRF_ACK_DIRECT = 0x10
	STATUS_DATA_OKMRF_ACK_ROUTED = 0x20
	STATUS_DATA_OKMRF_ACK        = 0x30
	STATUS_DATA_OKMRF_ACK_BM     = 0x40
	STATUS_DATA_OKMRF_DPREMOVED  = 0x80
)
