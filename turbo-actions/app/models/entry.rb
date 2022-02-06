class Entry < ApplicationRecord
  validates :title, :body, presence: true
end
