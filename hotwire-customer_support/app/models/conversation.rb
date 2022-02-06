class Conversation < ApplicationRecord
  belongs_to :contact
  has_many :posts

  broadcasts_to ->(c) { "conversations" }, inserts_by: :prepend, target: "conversations"
end
