# Place all the behaviors and hooks related to the matching controller here.
# All this logic will automatically be available in application.js.
# You can use CoffeeScript in this file: http://coffeescript.org/

DOLLAR_COST_PER_STUDENT = 3;
IMPACT_DURATION_RECURRING = 'every month'
IMPACT_DURATION_ONE_TIME = 'for a month'

@['donations#new'] = (data) ->
  impactCount = $('#donation-impact-count');
  impactDuration = $('#donation-impact-duration');

  # Adjust displayed impact count based on donation amount
  $('input[name="amount"]').change ->
    amount = parseInt this.value
    count = Math.floor(amount / DOLLAR_COST_PER_STUDENT)
    impactCount.text constructImpactCount(count)

  # Adjust displayed impact duration based on whether the donation is recurring
  $('input:radio[name="recurring"]').change ->
    switch this.value
      when 'true'
        impactDuration.text IMPACT_DURATION_RECURRING
      when 'false'
        impactDuration.text IMPACT_DURATION_ONE_TIME

constructImpactCount = (count) =>
  if count == 1
    count + ' student'
  else
    count + ' students'
