# This file was generated by typhen-api

module TyphenApi::Model::Submarine::Battle
  class Speed
    include Virtus.model(:strict => true)

    attribute :max, Number, :required => true
    attribute :accelerated_at, Integer, :required => true
    attribute :duration, Integer, :required => true
  end
end
