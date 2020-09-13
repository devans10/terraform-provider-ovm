package ovm

import "github.com/devans10/go-ovm-helper/ovmhelper"

func flattenIds(list []*ovmhelper.ID) []map[string]interface{} {
	result := make([]map[string]interface{}, 0, len(list))
	for _, i := range list {
		l := map[string]interface{}{
			"id":    i.Value,
			"name":  i.Name,
			"value": i.Value,
			"type":  i.Type,
			"uri":   i.URI,
		}
		result = append(result, l)
	}
	return result
}

func flattenID(id *ovmhelper.ID) map[string]interface{} {

	result := map[string]interface{}{
		"id":    id.Value,
		"name":  id.Name,
		"value": id.Value,
		"type":  id.Type,
		"uri":   id.URI,
	}

	return result
}
