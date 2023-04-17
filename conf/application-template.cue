parameter: {
	name:   string
	image:  string
	memory: string
	cpu:    string
	port:   string
	url:    string
}

split: "---"

apiVersion: "core.oam.dev/v1beta1"
kind:       "Application"
metadata: name: parameter.name
spec: {
	components: [{
		// name: parameter.name
		// type: "Webservice"
		// properties: {
		// 	image:           parameter.image
		// 	imagePullPolicy: "IfNotPresent"
		// 	memory:          parameter.memory
		// 	cpu:             parameter.cpu
		// 	ports: [
		// 		{
		// 			port: parameter.port
		// 		},
		// 	]
		// 	env: [
		// 		{
		// 			name: "value"
		// 		},
		// 	]
		// 	traits: [
		// 		{
		// 			type: "expose"
		// 			properties: {
		// 				port: parameter.port
		// 				type: "NodePort"
		// 			}
		// 		},
		// 		{
		// 			type: "scaler"
		// 			properties: {
		// 				replicas: 1
		// 			}
		// 		},
		// 	]
		// }

	}]
	policies: [
		{
			name: "target-default"
			type: "topology"
			properties: {
				clusters: ["local"]
				namespace: "default"
			}
		},
	]
	workflow: [
		{
			steps: [
				{
					name:    "deploy2default"
					type:    "deploy"
					timeout: "120s"
					properties: {
						policies: ["target-default"]
						auto: true
					}
				},
				{
					name:    "reportfaild"
					type:    "request"
					timeout: "120s"
					if:      "status.deploy2default.phase == \"failed\""
					properties: {
						url:    parameter.url
						method: "POST"
						body: {"appId": "demo-app", "status": "failed"}

					}

				},
			]
		},

	]
}
