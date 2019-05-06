// Copyright 2019 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package messages

import (
	"github.com/wmnsk/go-gtp/v2/ies"
)

// ContextRequest is a ContextRequest Header and its IEs above.
type ContextRequest struct {
	*Header
	IMSI                               *ies.IE
	GUTI                               *ies.IE
	RAI                                *ies.IE
	PTMSI                              *ies.IE
	PTMSISignature                     *ies.IE
	CompleteTAURequestMessage          *ies.IE
	AddressAndTEIDForCPlane            *ies.IE
	UDPSourcePortNumber                *ies.IE
	RATType                            *ies.IE
	Indication                         *ies.IE
	HopCounter                         *ies.IE
	TargetPLMNID                       *ies.IE
	MMESGSNLDN                         *ies.IE
	SGSNNodeName                       *ies.IE
	MMENodeName                        *ies.IE
	SGSNNumber                         *ies.IE
	SGSNIdentifier                     *ies.IE
	MMEIdentifier                      *ies.IE
	CIoTOptimizationsSupportIndication *ies.IE
	PrivateExtension                   *ies.IE
	AdditionalIEs                      []*ies.IE
}

// NewContextRequest creates a new ContextRequest.
func NewContextRequest(teid, seq uint32, ie ...*ies.IE) *ContextRequest {
	c := &ContextRequest{
		Header: NewHeader(
			NewHeaderFlags(2, 0, 1),
			MsgTypeContextRequest, teid, seq, nil,
		),
	}

	for _, i := range ie {
		if i == nil {
			continue
		}
		switch i.Type {
		case ies.IMSI:
			c.IMSI = i
		case ies.GUTI:
			c.GUTI = i
		case ies.UserLocationInformation:
			c.RAI = i
		case ies.PacketTMSI:
			c.PTMSI = i
		case ies.PTMSISignature:
			c.PTMSISignature = i
		case ies.CompleteRequestMessage:
			c.CompleteTAURequestMessage = i
		case ies.FullyQualifiedTEID:
			c.AddressAndTEIDForCPlane = i
		case ies.PortNumber:
			c.UDPSourcePortNumber = i
		case ies.RATType:
			c.RATType = i
		case ies.Indication:
			c.Indication = i
		case ies.HopCounter:
			c.HopCounter = i
		case ies.ServingNetwork:
			c.TargetPLMNID = i
		case ies.LocalDistinguishedName:
			c.MMESGSNLDN = i
		case ies.FullyQualifiedDomainName:
			switch i.Instance() {
			case 0:
				c.SGSNNodeName = i
			case 1:
				c.MMENodeName = i
			default:
				c.AdditionalIEs = append(c.AdditionalIEs, i)
			}
		case ies.NodeNumber:
			c.SGSNNumber = i
		case ies.NodeIdentifier:
			switch i.Instance() {
			case 0:
				c.SGSNIdentifier = i
			case 1:
				c.MMEIdentifier = i
			default:
				c.AdditionalIEs = append(c.AdditionalIEs, i)
			}
		case ies.CIoTOptimizationsSupportIndication:
			c.CIoTOptimizationsSupportIndication = i
		case ies.PrivateExtension:
			c.PrivateExtension = i
		default:
			c.AdditionalIEs = append(c.AdditionalIEs, i)
		}
	}

	c.SetLength()
	return c
}

// Serialize serializes ContextRequest into bytes.
func (c *ContextRequest) Serialize() ([]byte, error) {
	b := make([]byte, c.Len())
	if err := c.SerializeTo(b); err != nil {
		return nil, err
	}
	return b, nil
}

// SerializeTo serializes ContextRequest into bytes.
func (c *ContextRequest) SerializeTo(b []byte) error {
	if c.Header.Payload != nil {
		c.Header.Payload = nil
	}
	c.Header.Payload = make([]byte, c.Len()-c.Header.Len())

	offset := 0
	if ie := c.IMSI; ie != nil {
		if err := ie.SerializeTo(c.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.Len()
	}
	if ie := c.GUTI; ie != nil {
		if err := ie.SerializeTo(c.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.Len()
	}
	if ie := c.RAI; ie != nil {
		if err := ie.SerializeTo(c.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.Len()
	}
	if ie := c.PTMSI; ie != nil {
		if err := ie.SerializeTo(c.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.Len()
	}
	if ie := c.PTMSISignature; ie != nil {
		if err := ie.SerializeTo(c.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.Len()
	}
	if ie := c.CompleteTAURequestMessage; ie != nil {
		if err := ie.SerializeTo(c.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.Len()
	}
	if ie := c.AddressAndTEIDForCPlane; ie != nil {
		if err := ie.SerializeTo(c.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.Len()
	}
	if ie := c.UDPSourcePortNumber; ie != nil {
		if err := ie.SerializeTo(c.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.Len()
	}
	if ie := c.RATType; ie != nil {
		if err := ie.SerializeTo(c.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.Len()
	}
	if ie := c.Indication; ie != nil {
		if err := ie.SerializeTo(c.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.Len()
	}
	if ie := c.HopCounter; ie != nil {
		if err := ie.SerializeTo(c.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.Len()
	}
	if ie := c.TargetPLMNID; ie != nil {
		if err := ie.SerializeTo(c.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.Len()
	}
	if ie := c.MMESGSNLDN; ie != nil {
		if err := ie.SerializeTo(c.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.Len()
	}
	if ie := c.SGSNNodeName; ie != nil {
		if err := ie.SerializeTo(c.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.Len()
	}
	if ie := c.MMENodeName; ie != nil {
		if err := ie.SerializeTo(c.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.Len()
	}
	if ie := c.SGSNNumber; ie != nil {
		if err := ie.SerializeTo(c.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.Len()
	}
	if ie := c.SGSNIdentifier; ie != nil {
		if err := ie.SerializeTo(c.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.Len()
	}
	if ie := c.MMEIdentifier; ie != nil {
		if err := ie.SerializeTo(c.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.Len()
	}
	if ie := c.CIoTOptimizationsSupportIndication; ie != nil {
		if err := ie.SerializeTo(c.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.Len()
	}
	if ie := c.PrivateExtension; ie != nil {
		if err := ie.SerializeTo(c.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.Len()
	}

	for _, ie := range c.AdditionalIEs {
		if ie == nil {
			continue
		}
		if err := ie.SerializeTo(c.Header.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.Len()
	}

	c.Header.SetLength()
	return c.Header.SerializeTo(b)
}

// DecodeContextRequest decodes given bytes as ContextRequest.
func DecodeContextRequest(b []byte) (*ContextRequest, error) {
	c := &ContextRequest{}
	if err := c.DecodeFromBytes(b); err != nil {
		return nil, err
	}
	return c, nil
}

// DecodeFromBytes decodes given bytes as ContextRequest.
func (c *ContextRequest) DecodeFromBytes(b []byte) error {
	var err error
	c.Header, err = DecodeHeader(b)
	if err != nil {
		return err
	}
	if len(c.Header.Payload) < 2 {
		return nil
	}

	decodedIEs, err := ies.DecodeMultiIEs(c.Header.Payload)
	if err != nil {
		return err
	}
	for _, i := range decodedIEs {
		if i == nil {
			continue
		}
		switch i.Type {
		case ies.IMSI:
			c.IMSI = i
		case ies.GUTI:
			c.GUTI = i
		case ies.UserLocationInformation:
			c.RAI = i
		case ies.PacketTMSI:
			c.PTMSI = i
		case ies.PTMSISignature:
			c.PTMSISignature = i
		case ies.CompleteRequestMessage:
			c.CompleteTAURequestMessage = i
		case ies.FullyQualifiedTEID:
			c.AddressAndTEIDForCPlane = i
		case ies.PortNumber:
			c.UDPSourcePortNumber = i
		case ies.RATType:
			c.RATType = i
		case ies.Indication:
			c.Indication = i
		case ies.HopCounter:
			c.HopCounter = i
		case ies.ServingNetwork:
			c.TargetPLMNID = i
		case ies.LocalDistinguishedName:
			c.MMESGSNLDN = i
		case ies.FullyQualifiedDomainName:
			switch i.Instance() {
			case 0:
				c.SGSNNodeName = i
			case 1:
				c.MMENodeName = i
			default:
				c.AdditionalIEs = append(c.AdditionalIEs, i)
			}
		case ies.NodeNumber:
			c.SGSNNumber = i
		case ies.NodeIdentifier:
			switch i.Instance() {
			case 0:
				c.SGSNIdentifier = i
			case 1:
				c.MMEIdentifier = i
			default:
				c.AdditionalIEs = append(c.AdditionalIEs, i)
			}
		case ies.CIoTOptimizationsSupportIndication:
			c.CIoTOptimizationsSupportIndication = i
		case ies.PrivateExtension:
			c.PrivateExtension = i
		default:
			c.AdditionalIEs = append(c.AdditionalIEs, i)
		}
	}

	return nil
}

// Len returns the actual length in int.
func (c *ContextRequest) Len() int {
	l := c.Header.Len() - len(c.Header.Payload)
	if ie := c.IMSI; ie != nil {
		l += ie.Len()
	}
	if ie := c.GUTI; ie != nil {
		l += ie.Len()
	}
	if ie := c.RAI; ie != nil {
		l += ie.Len()
	}
	if ie := c.PTMSI; ie != nil {
		l += ie.Len()
	}
	if ie := c.PTMSISignature; ie != nil {
		l += ie.Len()
	}
	if ie := c.CompleteTAURequestMessage; ie != nil {
		l += ie.Len()
	}
	if ie := c.AddressAndTEIDForCPlane; ie != nil {
		l += ie.Len()
	}
	if ie := c.UDPSourcePortNumber; ie != nil {
		l += ie.Len()
	}
	if ie := c.RATType; ie != nil {
		l += ie.Len()
	}
	if ie := c.Indication; ie != nil {
		l += ie.Len()
	}
	if ie := c.HopCounter; ie != nil {
		l += ie.Len()
	}
	if ie := c.TargetPLMNID; ie != nil {
		l += ie.Len()
	}
	if ie := c.MMESGSNLDN; ie != nil {
		l += ie.Len()
	}
	if ie := c.SGSNNodeName; ie != nil {
		l += ie.Len()
	}
	if ie := c.MMENodeName; ie != nil {
		l += ie.Len()
	}
	if ie := c.SGSNNumber; ie != nil {
		l += ie.Len()
	}
	if ie := c.SGSNIdentifier; ie != nil {
		l += ie.Len()
	}
	if ie := c.MMEIdentifier; ie != nil {
		l += ie.Len()
	}
	if ie := c.CIoTOptimizationsSupportIndication; ie != nil {
		l += ie.Len()
	}
	if ie := c.PrivateExtension; ie != nil {
		l += ie.Len()
	}

	for _, ie := range c.AdditionalIEs {
		if ie == nil {
			continue
		}
		l += ie.Len()
	}
	return l
}

// SetLength sets the length in Length field.
func (c *ContextRequest) SetLength() {
	c.Header.Length = uint16(c.Len() - 4)
}

// MessageTypeName returns the name of protocol.
func (c *ContextRequest) MessageTypeName() string {
	return "Context Request"
}

// TEID returns the TEID in uint32.
func (c *ContextRequest) TEID() uint32 {
	return c.Header.teid()
}