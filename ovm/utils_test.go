package ovm

import (
	"reflect"
	"testing"

	"github.com/devans10/go-ovm-helper/ovmhelper"
)

func TestFlattenID(t *testing.T) {
	cases := []struct {
		id       *ovmhelper.ID
		expected map[string]interface{}
	}{
		{
			id: &ovmhelper.ID{
				Name:  "testexample",
				Value: "0004fb0000130000dfc5261750e0df78",
				Type:  "com.oracle.ovm.mgr.ws.model.VmDiskMapping",
				URI:   "https://localhost:7002//ovm/core/wsapi/rest/VmDiskMapping/0004fb0000130000dfc5261750e0df78",
			},
			expected: map[string]interface{}{
				"name":  "testexample",
				"value": "0004fb0000130000dfc5261750e0df78",
				"type":  "com.oracle.ovm.mgr.ws.model.VmDiskMapping",
				"uri":   "https://localhost:7002//ovm/core/wsapi/rest/VmDiskMapping/0004fb0000130000dfc5261750e0df78",
			},
		},
	}

	for _, c := range cases {
		out := flattenID(c.id)
		if !reflect.DeepEqual(out, c.expected) {
			t.Fatalf("Error matching output and expected: %#v vs %#v", out, c.expected)
		}
	}
}

func TestFlattenIDs(t *testing.T) {
	cases := []struct {
		list     []*ovmhelper.ID
		expected []map[string]interface{}
	}{
		{
			list: []*ovmhelper.ID{
				{
					Name:  "testexample",
					Value: "0004fb0000130000dfc5261750e0df78",
					Type:  "com.oracle.ovm.mgr.ws.model.VmDiskMapping",
					URI:   "https://localhost:7002//ovm/core/wsapi/rest/VmDiskMapping/0004fb0000130000dfc5261750e0df78",
				},
			},
			expected: []map[string]interface{}{
				{
					"name":  "testexample",
					"value": "0004fb0000130000dfc5261750e0df78",
					"type":  "com.oracle.ovm.mgr.ws.model.VmDiskMapping",
					"uri":   "https://localhost:7002//ovm/core/wsapi/rest/VmDiskMapping/0004fb0000130000dfc5261750e0df78",
				},
			},
		},
	}

	for _, c := range cases {
		out := flattenIds(c.list)
		if !reflect.DeepEqual(out, c.expected) {
			t.Fatalf("Error matching output and expected: %#v vs %#v", out, c.expected)
		}
	}
}
