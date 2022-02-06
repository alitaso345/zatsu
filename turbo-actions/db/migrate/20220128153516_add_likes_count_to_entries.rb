class AddLikesCountToEntries < ActiveRecord::Migration[7.0]
  def change
    add_column :entries, :likes_count, :integer
  end
end
