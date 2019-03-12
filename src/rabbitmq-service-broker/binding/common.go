package binding

func (b Builder) firstHostname() string {
	return b.Hostnames[0]
}

func (b Builder) amqpScheme() string {
	if b.TLS {
		return "amqps"
	}
	return "amqp"
}

func ternary(expression bool, ifTrue, ifFalse string) string {
	if expression {
		return ifTrue
	}
	return ifFalse
}
