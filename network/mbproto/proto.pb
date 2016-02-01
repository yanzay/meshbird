syntax = "proto3";

package mbproto;

message RawMessage {
	bytes IV = 1;
	bytes Nonce = 2;
	bytes Payload = 3;
}

message TransmitMessage {
	bytes payload = 1;
}

message HandshakeMessage {
	bytes  networkKey = 1;
	uint32 privateIP = 2;
}