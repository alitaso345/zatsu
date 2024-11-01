require 'discordrb'
require 'dotenv/load'

bot = Discordrb::Bot.new(token: ENV['DISCORD_TOKEN'])
bot.message(context: 'Ping!') do |event|
  event.respond('Pong!')
end

bot.run

__END__
require 'sinatra'

get '/' do
  'hello'
end