package plugin

import (
	"bytes"
	"crypto/sha256"
	"math/rand"
	"sort"
	"testing"

	"github.com/lazyledger/lazyledger-core/types"
	"github.com/lazyledger/nmt"
)

func TestDataSquareRowOrColumnRawInputParserCidEqNmtRoot(t *testing.T) {
	const (
		namespaceSize = types.NamespaceSize
		shareSize = types.ShareSize
	)

	tests := []struct {
		name     string
		leafData [][]byte
	}{
		{"16 leaves", generateRandNamespacedRawData(16, namespaceSize, shareSize)},
		// TODO add at least a row of an extended data square (incl. parity bytes) as a test-vector too
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			n := nmt.New(sha256.New())
			buf := bytes.NewBuffer(make([]byte, 0))
			for _, share := range tt.leafData {
				_, err := buf.Write(share)
				if err != nil {
					t.Errorf("buf.Write() unexpected error = %v", err)
					return
				}
				err = n.Push(share[:namespaceSize], share[namespaceSize:])
				if err != nil {
					t.Errorf("nmt.Push() unexpected error = %v", err)
					return
				}
			}
			gotNodes, err := DataSquareRowOrColumnRawInputParser(buf, 0, 0)
			if err != nil {
				t.Errorf("DataSquareRowOrColumnRawInputParser() unexpected error = %v", err)
				return
			}
			lastNodeCid := gotNodes[len(gotNodes)-1].Cid()
			multiHashOverhead := 2
			lastNodeHash := lastNodeCid.Hash()
			if got, want := lastNodeHash[multiHashOverhead:], n.Root().Bytes(); !bytes.Equal(got, want) {
				t.Errorf("hashes don't match\ngot: %v\nwant: %v", got, want)
			}
			firstNodeCid := gotNodes[0].Cid()
			if gotHash, wantHash := firstNodeCid.Hash(), hashLeaf(tt.leafData[0]); !bytes.Equal(gotHash[multiHashOverhead:], wantHash) {
				t.Errorf("first node's hash does not match the Cid\ngot: %v\nwant: %v", gotHash[multiHashOverhead:], wantHash)
			}
			nodePrefixOffset := 1 // leaf / inner node prefix is one byte
			firstNodeData := gotNodes[0].RawData()
			if gotData, wantData := firstNodeData[nodePrefixOffset:], tt.leafData[0]; !bytes.Equal(gotData, wantData) {
				t.Errorf("first node's data does not match the leaf's data\ngot: %v\nwant: %v", gotData, wantData)
			}
		})
	}
}

// this snippet of the nmt internals is copied here:
func hashLeaf(data []byte) []byte {
	h := sha256.New()
	nID := data[:types.NamespaceSize]
	toCommittToDataWithoutNID := data[types.NamespaceSize:]

	res := append(append(make([]byte, 0), nID...), nID...)
	data = append([]byte{nmt.LeafPrefix}, toCommittToDataWithoutNID...)
	h.Write(data)
	return h.Sum(res)
}

// TODO add a dag put using a IPFS daemon and see if the returned leaf data on dag get matches

func generateRandNamespacedRawData(total int, nidSize int, leafSize int) [][]byte {
	data := make([][]byte, total)
	for i := 0; i < total; i++ {
		nid := make([]byte, nidSize)
		rand.Read(nid)
		data[i] = nid
	}
	sortByteArrays(data)
	for i := 0; i < total; i++ {
		d := make([]byte, leafSize)
		rand.Read(d)
		data[i] = append(data[i], d...)
	}

	return data
}

func sortByteArrays(src [][]byte) {
	sort.Slice(src, func(i, j int) bool { return bytes.Compare(src[i], src[j]) < 0 })
}
