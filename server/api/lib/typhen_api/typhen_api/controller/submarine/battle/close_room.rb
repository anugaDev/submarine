# This file was generated by typhen-api

module TyphenApi::Controller::Submarine::Battle
  module CloseRoom
    extend ActiveSupport::Concern

    class RequestType
      include Virtus.model(:strict => true)

      attribute :room_id, Integer, :required => true
    end

    ResponseType = nil
    ErrorType = TyphenApi::Model::Submarine::Error

    def no_authentication_required?
      true
    end
  end
end
