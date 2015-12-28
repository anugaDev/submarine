# == Schema Information
#
# Table name: room_members
#
#  id         :integer          not null, primary key
#  user_id    :integer
#  room_id    :integer
#  created_at :datetime         not null
#  updated_at :datetime         not null
#
# Indexes
#
#  index_room_members_on_room_id  (room_id)
#  index_room_members_on_user_id  (user_id) UNIQUE
#

class RoomMember < ActiveRecord::Base
  belongs_to :user
  belongs_to :room, :counter_cache => true
end