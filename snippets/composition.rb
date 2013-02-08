module Sheep
  def sound
    p 'beh'
  end
end

class Herd
  include Sheep
end

Herd.new.sound
