class ApplicationSpamService
  WORD_CUTOFF = 5

  def is_spam?(application)
    if word_count(application.interesting_project) <= WORD_CUTOFF &&
       word_count(application.system_hacked) <= WORD_CUTOFF
      return true
    end

    false
  end

  private

  def word_count(phrase)
    phrase.split(' ').length
  end
end
