package portsstream

// Stream returns the stream channel.
func (s *Stream) Stream() chan Result {
	return s.stream
}
