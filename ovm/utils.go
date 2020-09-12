package ovm

import "github.com/devans10/go-ovm-helper/ovmHelper"

func flattenIds(list []*ovmHelper.Id) []map[string]interface{} {
	result := make([]map[string]interface{}, 0, len(list))
	for _, i := range list {
		l := map[string]interface{}{
			"name":  i.Name,
			"value": i.Value,
			"type":  i.Type,
			"uri":   i.Uri,
		}
		result = append(result, l)
	}
	return result
}

func flattenID(id *ovmHelper.Id) map[string]interface{} {

	result := map[string]interface{}{
		"name":  id.Name,
		"value": id.Value,
		"type":  id.Type,
		"uri":   id.Uri,
	}

	return result
}
