		
parameter: {
	name:   string
	image:  string
	memory: string
	cpu:    string
	port:   string
	url:    string
}

name: parameter.name
type: "Webservice"
properties: {
	image:           parameter.image
	imagePullPolicy: "IfNotPresent"
	memory:          parameter.memory
	cpu:             parameter.cpu
	ports: [
		{
		port: parameter.port
	},
]
env: [
	{
		name: "value"
	},
]
traits: [
	{
		type: "expose"
		properties: {
			port: parameter.port
			type: "NodePort"
		}
	},
	{
		type: "scaler"
		properties: {
			replicas: 1
		}
	},
]
}