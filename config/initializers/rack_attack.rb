class Rack::Attack
  spammer_string = (Figaro.env.REFERRERS_TO_BLOCK || '')
  spammers = spammer_string.split('|').map { |spammer| Regexp.new(spammer) }

  blacklist('block referral spam') do |request|
    spammers.find { |spammer| request.referer =~ spammer }
  end
end
