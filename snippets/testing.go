func TestKeyEncoding(t *testing.T) {
	key := 123
	if decodeKey(encodeKey(key)) != key {
		t.Error("Encoding and then decoding doesn't yield the same value")
	}
}
