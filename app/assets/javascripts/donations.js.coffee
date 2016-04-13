# Place all the behaviors and hooks related to the matching controller here.
# All this logic will automatically be available in application.js.
# You can use CoffeeScript in this file: http://coffeescript.org/

DOLLAR_COST_PER_STUDENT = 3;
IMPACT_DURATION_RECURRING = 'every month'
IMPACT_DURATION_ONE_TIME = 'for a month'

@['donations#new'] = (data) ->
  impactCount = $('#donation-impact-count')
  impactDuration = $('#donation-impact-duration')

  contributeBtn = $('#contribute')
  amountInput = $('input[name="amount"]')
  recurringInput = $('input[name="recurring"]')

  amount = amountInput.filter(':checked').val()
  recurring = recurringInput.filter(':checked').val()
  stripeHandler = StripeCheckout.configure(
    key: 'REPLACE_ME'
    image: 'TODO',
    locale: 'auto',
    token: (token) ->
      # TODO?
  )

  # Adjust displayed impact count based on donation amount
  amountInput.change ->
    amount = parseInt this.value
    count = Math.floor(amount / DOLLAR_COST_PER_STUDENT)
    impactCount.text constructImpactCount(count)

  # Adjust displayed impact duration based on whether the donation is recurring
  recurringInput.change ->
    # Convert the selected value to a boolean and store it in recurring
    recurring = this.value == 'true'

    if recurring
      impactDuration.text IMPACT_DURATION_RECURRING
    else
      impactDuration.text IMPACT_DURATION_ONE_TIME

  # Open Stripe Checkout when the contribute button is clicked
  contributeBtn.on 'click', (e) ->
    stripeHandler.open
      name: 'Hack Club',
      description: constructStripeCheckoutDescription recurring
      amount: amount * 100 # Gotta convert it from dollars to pennies
    e.preventDefault()

  # Close Stripe Checkout on page navigation
  $(window).on 'popstate', ->
    stripeHandler.close()

constructImpactCount = (count) =>
  if count == 1
    count + ' student'
  else
    count + ' students'

constructStripeCheckoutDescription = (recurring) =>
  suffix = 'Hack Club contribution'

  if recurring
    'Monthly ' + suffix
  else
    'One time ' + suffix
