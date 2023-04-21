parameter: {
	name:       string
	url:        string
	components: [...string]
}

split: "---"

apiVersion: "core.oam.dev/v1beta1"
kind:       "Application"
metadata: name: parameter.name
spec: {
	components: 
		parameter.components
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
