// Code generated by github.com/whyrusleeping/cbor-gen. DO NOT EDIT.

package hello

import (
	"fmt"
	"io"
	"math"
	"sort"

	abi "github.com/filecoin-project/go-state-types/abi"
	cid "github.com/ipfs/go-cid"
	cbg "github.com/whyrusleeping/cbor-gen"
	xerrors "golang.org/x/xerrors"
)

var _ = xerrors.Errorf
var _ = cid.Undef
var _ = math.E
var _ = sort.Sort

var lengthBufHelloMessage = []byte{132}

func (t *HelloMessage) MarshalCBOR(w io.Writer) error {
	if t == nil {
		_, err := w.Write(cbg.CborNull)
		return err
	}

	cw := cbg.NewCborWriter(w)

	if _, err := cw.Write(lengthBufHelloMessage); err != nil {
		return err
	}

	// t.HeaviestTipSet ([]cid.Cid) (slice)
	if len(t.HeaviestTipSet) > cbg.MaxLength {
		return xerrors.Errorf("Slice value in field t.HeaviestTipSet was too long")
	}

	if err := cw.WriteMajorTypeHeader(cbg.MajArray, uint64(len(t.HeaviestTipSet))); err != nil {
		return err
	}
	for _, v := range t.HeaviestTipSet {
		if err := cbg.WriteCid(w, v); err != nil {
			return xerrors.Errorf("failed writing cid field t.HeaviestTipSet: %w", err)
		}
	}

	// t.HeaviestTipSetHeight (abi.ChainEpoch) (int64)
	if t.HeaviestTipSetHeight >= 0 {
		if err := cw.WriteMajorTypeHeader(cbg.MajUnsignedInt, uint64(t.HeaviestTipSetHeight)); err != nil {
			return err
		}
	} else {
		if err := cw.WriteMajorTypeHeader(cbg.MajNegativeInt, uint64(-t.HeaviestTipSetHeight-1)); err != nil {
			return err
		}
	}

	// t.HeaviestTipSetWeight (big.Int) (struct)
	if err := t.HeaviestTipSetWeight.MarshalCBOR(cw); err != nil {
		return err
	}

	// t.GenesisHash (cid.Cid) (struct)

	if err := cbg.WriteCid(cw, t.GenesisHash); err != nil {
		return xerrors.Errorf("failed to write cid field t.GenesisHash: %w", err)
	}

	return nil
}

func (t *HelloMessage) UnmarshalCBOR(r io.Reader) (err error) {
	*t = HelloMessage{}

	cr := cbg.NewCborReader(r)

	maj, extra, err := cr.ReadHeader()
	if err != nil {
		return err
	}
	defer func() {
		if err == io.EOF {
			err = io.ErrUnexpectedEOF
		}
	}()

	if maj != cbg.MajArray {
		return fmt.Errorf("cbor input should be of type array")
	}

	if extra != 4 {
		return fmt.Errorf("cbor input had wrong number of fields")
	}

	// t.HeaviestTipSet ([]cid.Cid) (slice)

	maj, extra, err = cr.ReadHeader()
	if err != nil {
		return err
	}

	if extra > cbg.MaxLength {
		return fmt.Errorf("t.HeaviestTipSet: array too large (%d)", extra)
	}

	if maj != cbg.MajArray {
		return fmt.Errorf("expected cbor array")
	}

	if extra > 0 {
		t.HeaviestTipSet = make([]cid.Cid, extra)
	}

	for i := 0; i < int(extra); i++ {

		c, err := cbg.ReadCid(cr)
		if err != nil {
			return xerrors.Errorf("reading cid field t.HeaviestTipSet failed: %w", err)
		}
		t.HeaviestTipSet[i] = c
	}

	// t.HeaviestTipSetHeight (abi.ChainEpoch) (int64)
	{
		maj, extra, err := cr.ReadHeader()
		var extraI int64
		if err != nil {
			return err
		}
		switch maj {
		case cbg.MajUnsignedInt:
			extraI = int64(extra)
			if extraI < 0 {
				return fmt.Errorf("int64 positive overflow")
			}
		case cbg.MajNegativeInt:
			extraI = int64(extra)
			if extraI < 0 {
				return fmt.Errorf("int64 negative oveflow")
			}
			extraI = -1 - extraI
		default:
			return fmt.Errorf("wrong type for int64 field: %d", maj)
		}

		t.HeaviestTipSetHeight = abi.ChainEpoch(extraI)
	}
	// t.HeaviestTipSetWeight (big.Int) (struct)

	{

		if err := t.HeaviestTipSetWeight.UnmarshalCBOR(cr); err != nil {
			return xerrors.Errorf("unmarshaling t.HeaviestTipSetWeight: %w", err)
		}

	}
	// t.GenesisHash (cid.Cid) (struct)

	{

		c, err := cbg.ReadCid(cr)
		if err != nil {
			return xerrors.Errorf("failed to read cid field t.GenesisHash: %w", err)
		}

		t.GenesisHash = c

	}
	return nil
}

var lengthBufLatencyMessage = []byte{130}

func (t *LatencyMessage) MarshalCBOR(w io.Writer) error {
	if t == nil {
		_, err := w.Write(cbg.CborNull)
		return err
	}

	cw := cbg.NewCborWriter(w)

	if _, err := cw.Write(lengthBufLatencyMessage); err != nil {
		return err
	}

	// t.TArrival (int64) (int64)
	if t.TArrival >= 0 {
		if err := cw.WriteMajorTypeHeader(cbg.MajUnsignedInt, uint64(t.TArrival)); err != nil {
			return err
		}
	} else {
		if err := cw.WriteMajorTypeHeader(cbg.MajNegativeInt, uint64(-t.TArrival-1)); err != nil {
			return err
		}
	}

	// t.TSent (int64) (int64)
	if t.TSent >= 0 {
		if err := cw.WriteMajorTypeHeader(cbg.MajUnsignedInt, uint64(t.TSent)); err != nil {
			return err
		}
	} else {
		if err := cw.WriteMajorTypeHeader(cbg.MajNegativeInt, uint64(-t.TSent-1)); err != nil {
			return err
		}
	}
	return nil
}

func (t *LatencyMessage) UnmarshalCBOR(r io.Reader) (err error) {
	*t = LatencyMessage{}

	cr := cbg.NewCborReader(r)

	maj, extra, err := cr.ReadHeader()
	if err != nil {
		return err
	}
	defer func() {
		if err == io.EOF {
			err = io.ErrUnexpectedEOF
		}
	}()

	if maj != cbg.MajArray {
		return fmt.Errorf("cbor input should be of type array")
	}

	if extra != 2 {
		return fmt.Errorf("cbor input had wrong number of fields")
	}

	// t.TArrival (int64) (int64)
	{
		maj, extra, err := cr.ReadHeader()
		var extraI int64
		if err != nil {
			return err
		}
		switch maj {
		case cbg.MajUnsignedInt:
			extraI = int64(extra)
			if extraI < 0 {
				return fmt.Errorf("int64 positive overflow")
			}
		case cbg.MajNegativeInt:
			extraI = int64(extra)
			if extraI < 0 {
				return fmt.Errorf("int64 negative oveflow")
			}
			extraI = -1 - extraI
		default:
			return fmt.Errorf("wrong type for int64 field: %d", maj)
		}

		t.TArrival = int64(extraI)
	}
	// t.TSent (int64) (int64)
	{
		maj, extra, err := cr.ReadHeader()
		var extraI int64
		if err != nil {
			return err
		}
		switch maj {
		case cbg.MajUnsignedInt:
			extraI = int64(extra)
			if extraI < 0 {
				return fmt.Errorf("int64 positive overflow")
			}
		case cbg.MajNegativeInt:
			extraI = int64(extra)
			if extraI < 0 {
				return fmt.Errorf("int64 negative oveflow")
			}
			extraI = -1 - extraI
		default:
			return fmt.Errorf("wrong type for int64 field: %d", maj)
		}

		t.TSent = int64(extraI)
	}
	return nil
}
