it "makes encoding and decoding symmetric" do
  key = 123
  decode_key(encode_key(key)).should == key
end
