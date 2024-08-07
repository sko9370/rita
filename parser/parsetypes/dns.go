package parsetypes

import (
	"github.com/activecm/rita-legacy/config"
	"github.com/globalsign/mgo/bson"
)

// DNS provides a data structure for entries in the zeek DNS log
type DNS struct {
	// ID contains the id set by mongodb
	ID bson.ObjectId `bson:"_id,omitempty"`
	// TimeStamp of this connection
	TimeStamp int64 `bson:"ts" bro:"ts" brotype:"time" json:"-"`
	// TimeStampGeneric is used when reading from json files
	TimeStampGeneric interface{} `bson:"-" json:"ts"`
	// UID is the Unique Id for this connection (generated by Bro)
	UID string `bson:"uid" bro:"uid" brotype:"string" json:"uid"`
	// Source is the source address for this connection
	Source string `bson:"id_orig_h" bro:"id.orig_h" brotype:"addr" json:"id.orig_h"`
	// SourcePort is the source port of this connection
	SourcePort int `bson:"id_orig_p" bro:"id.orig_p" brotype:"port" json:"id.orig_p"`
	// Destination is the destination of the connection
	Destination string `bson:"id_resp_h" bro:"id.resp_h" brotype:"addr" json:"id.resp_h"`
	// DestinationPort is the port at the destination host
	DestinationPort int `bson:"id_resp_p" bro:"id.resp_p" brotype:"port" json:"id.resp_p"`
	// Proto is the string protocol identifier for this connection
	Proto string `bson:"proto" bro:"proto" brotype:"enum" json:"proto"`
	// TransID contains a 16 bit identifier assigned by the program that generated
	// the query
	TransID int64 `bson:"trans_id" bro:"trans_id" brotype:"count" json:"trans_id"`
	// RTT contains the round trip time of this request / response
	RTT float64 `bson:"rtt" bro:"rtt" brotype:"interval" json:"rtt"`
	// Query contains the query string
	Query string `bson:"query" bro:"query" brotype:"string" json:"query"`
	// QClass contains a the qclass of the query
	QClass int64 `bson:"qclass" bro:"qclass" brotype:"count" json:"qclass"`
	// QClassName contains a descriptive name for the query
	QClassName string `bson:"qclass_name" bro:"qclass_name" brotype:"string" json:"qclass_name"`
	// QType contains the value of the query type
	QType int64 `bson:"qtype" bro:"qtype" brotype:"count" json:"qtype"`
	// QTypeName provides a descriptive name for the query
	QTypeName string `bson:"qtype_name" bro:"qtype_name" brotype:"string" json:"qtype_name"`
	// RCode contains the response code value from the DNS messages
	RCode int64 `bson:"rcode" bro:"rcode" brotype:"count" json:"rcode"`
	// RCodeName provides a descriptive name for RCode
	RCodeName string `bson:"rcode_name" bro:"rcode_name" brotype:"string" json:"rcode_name"`
	// AA represents the state of the authoritive answer bit of the resp messages
	AA bool `bson:"AA" bro:"AA" brotype:"bool" json:"AA"`
	// TC represents the truncation bit of the message
	TC bool `bson:"TC" bro:"TC" brotype:"bool" json:"TC"`
	// RD represens the recursion desired bit of the message
	RD bool `bson:"RD" bro:"RD" brotype:"bool" json:"RD"`
	// RA represents the recursion available bit of the message
	RA bool `bson:"RA" bro:"RA" brotype:"bool" json:"RA"`
	// Z represents the state of a reseverd field that should be zero in qll queries
	Z int64 `bson:"Z" bro:"Z" brotype:"count" json:"Z"`
	// Answers contains the set of resource descriptions in the query answer
	Answers []string `bson:"answers" bro:"answers" brotype:"vector[string]" json:"answers"`
	// TTLs contains a vector of interval type time to live values
	TTLs []float64 `bson:"TTLs" bro:"TTLs" brotype:"vector[interval]" json:"TTLs"`
	// Rejected indicates if this query was rejected or not
	Rejected bool `bson:"rejected" bro:"rejected" brotype:"bool" json:"rejected"`
	// AgentHostname names which sensor recorded this event. Only set when combining logs from multiple sensors.
	AgentHostname string `bson:"agent_hostname" bro:"agent_hostname" brotype:"string" json:"agent_hostname"`
	// AgentUUID identifies which sensor recorded this event. Only set when combining logs from multiple sensors.
	AgentUUID string `bson:"agent_uuid" bro:"agent_uuid" brotype:"string" json:"agent_uuid"`
}

// TargetCollection returns the mongo collection this entry should be inserted
func (line *DNS) TargetCollection(config *config.StructureTableCfg) string {
	return config.DNSTable
}

// ConvertFromJSON performs any extra conversions necessary when reading from JSON
func (line *DNS) ConvertFromJSON() {
	line.TimeStamp = convertTimestamp(line.TimeStampGeneric)
}
