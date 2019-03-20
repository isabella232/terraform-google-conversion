// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    AUTO GENERATED CODE     ***
//
// ----------------------------------------------------------------------------
//
//     This file is automatically generated by Magic Modules and manual
//     changes will be clobbered when the file is regenerated.
//
//     Please read more about how to change this file in
//     .github/CONTRIBUTING.md.
//
// ----------------------------------------------------------------------------

package google

import (
	"fmt"
	"reflect"
	"regexp"

	"github.com/hashicorp/terraform/helper/schema"
)

// compareTpuNodeSchedulingConfig diff suppresses for the default
// scheduling, i.e. if preemptible is false, the API may either return no
// schedulingConfig or an empty schedulingConfig.
func compareTpuNodeSchedulingConfig(k, old, new string, d *schema.ResourceData) bool {
	if k == "scheduling_config.0.preemptible" {
		return old == "" && new == "false"
	}
	if k == "scheduling_config.#" {
		o, n := d.GetChange("scheduling_config.0.preemptible")
		return o.(bool) == n.(bool)
	}
	return false
}

func validateHttpHeaders() schema.SchemaValidateFunc {
	return func(i interface{}, k string) (s []string, es []error) {
		headers := i.(map[string]interface{})
		if _, ok := headers["Content-Length"]; ok {
			es = append(es, fmt.Errorf("Cannot set the Content-Length header on %s", k))
			return
		}
		for key := range headers {
			match, _ := regexp.MatchString("(X-Google-|X-AppEngine-).*", key)
			if match {
				es = append(es, fmt.Errorf("Cannot set the %s header on %s", key, k))
				return
			}
		}

		return
	}
}

func GetTpuNodeCaiObject(d TerraformResourceData, config *Config) (Asset, error) {
	name, err := replaceVars(d, config, "//tpu.googleapis.com/projects/{{project}}/locations/{{zone}}/nodes/{{name}}")
	if err != nil {
		return Asset{}, err
	}
	if obj, err := GetTpuNodeApiObject(d, config); err == nil {
		return Asset{
			Name: name,
			Type: "google.tpu.Node",
			Resource: &AssetResource{
				Version:              "v1",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/tpu/v1/rest",
				DiscoveryName:        "Node",
				Data:                 obj,
			},
		}, nil
	} else {
		return Asset{}, err
	}
}

func GetTpuNodeApiObject(d TerraformResourceData, config *Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	nameProp, err := expandTpuNodeName(d.Get("name"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("name"); !isEmptyValue(reflect.ValueOf(nameProp)) && (ok || !reflect.DeepEqual(v, nameProp)) {
		obj["name"] = nameProp
	}
	descriptionProp, err := expandTpuNodeDescription(d.Get("description"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("description"); !isEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	acceleratorTypeProp, err := expandTpuNodeAcceleratorType(d.Get("accelerator_type"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("accelerator_type"); !isEmptyValue(reflect.ValueOf(acceleratorTypeProp)) && (ok || !reflect.DeepEqual(v, acceleratorTypeProp)) {
		obj["acceleratorType"] = acceleratorTypeProp
	}
	tensorflowVersionProp, err := expandTpuNodeTensorflowVersion(d.Get("tensorflow_version"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("tensorflow_version"); !isEmptyValue(reflect.ValueOf(tensorflowVersionProp)) && (ok || !reflect.DeepEqual(v, tensorflowVersionProp)) {
		obj["tensorflowVersion"] = tensorflowVersionProp
	}
	networkProp, err := expandTpuNodeNetwork(d.Get("network"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("network"); !isEmptyValue(reflect.ValueOf(networkProp)) && (ok || !reflect.DeepEqual(v, networkProp)) {
		obj["network"] = networkProp
	}
	cidrBlockProp, err := expandTpuNodeCidrBlock(d.Get("cidr_block"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("cidr_block"); !isEmptyValue(reflect.ValueOf(cidrBlockProp)) && (ok || !reflect.DeepEqual(v, cidrBlockProp)) {
		obj["cidrBlock"] = cidrBlockProp
	}
	schedulingConfigProp, err := expandTpuNodeSchedulingConfig(d.Get("scheduling_config"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("scheduling_config"); !isEmptyValue(reflect.ValueOf(schedulingConfigProp)) && (ok || !reflect.DeepEqual(v, schedulingConfigProp)) {
		obj["schedulingConfig"] = schedulingConfigProp
	}
	labelsProp, err := expandTpuNodeLabels(d.Get("labels"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("labels"); !isEmptyValue(reflect.ValueOf(labelsProp)) && (ok || !reflect.DeepEqual(v, labelsProp)) {
		obj["labels"] = labelsProp
	}

	return obj, nil
}

func expandTpuNodeName(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandTpuNodeDescription(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandTpuNodeAcceleratorType(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandTpuNodeTensorflowVersion(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandTpuNodeNetwork(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandTpuNodeCidrBlock(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandTpuNodeSchedulingConfig(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedPreemptible, err := expandTpuNodeSchedulingConfigPreemptible(original["preemptible"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedPreemptible); val.IsValid() && !isEmptyValue(val) {
		transformed["preemptible"] = transformedPreemptible
	}

	return transformed, nil
}

func expandTpuNodeSchedulingConfigPreemptible(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandTpuNodeLabels(v interface{}, d TerraformResourceData, config *Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}
