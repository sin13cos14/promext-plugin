package config

const Mapping = `{
    "settings": {
        "number_of_shards": 6,
        "number_of_replicas": 1
    },
    "mappings": {
        "metric": {
            "_all": {
                "enabled": false
            },
			"dynamic_templates":[
				{ "dubbo_values_as_double": {
					  "match": "memory*", 
					  "mapping":{"type": "double"}
				}},
				{ "http_values_as_double": {
					  "match": "cpu*", 
					  "mapping":{"type": "double"}
				}},
				{ "http_values_as_double": {
					  "match": "disk*", 
					  "mapping":{"type": "double"}
				}},
				{ "string_as_key": {
					  "match": "*", 
					  "match_mapping_type": "string",
					  "mapping":{"type": "keyword"}
				}}
			],
            "properties": {
                "ip": {
                    "type": "ip"
                },
                "project": {
                    "type": "keyword"
                },
				"ctime":{
					"type":"date",
					"format": "epoch_second"
				}
            }
        }
    }
}`