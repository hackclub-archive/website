# Place all the behaviors and hooks related to the matching controller here.
# All this logic will automatically be available in application.js.
# You can use CoffeeScript in this file: http://coffeescript.org/

PENNY_COST_PER_STUDENT = 300;
IMPACT_DURATION_RECURRING = 'every month'
IMPACT_DURATION_ONE_TIME = 'for a month'

@['donations#new'] = (data) ->
  # If this page is being loaded after a successful donation, trigger the
  # confirmation modal
  if data != undefined && data.donation_successful == true
    $('#confirmation-modal').foundation('reveal', 'open')

  # Load Stripe Checkout before doing anything else
  $.getScript 'https://checkout.stripe.com/checkout.js', ->
    donationLogic data

donationLogic = (data) ->
  contributionAmount = $('#contribution-amount')
  impactCount = $('#donation-impact-count')
  impactDuration = $('#donation-impact-duration')

  form = $('#donation-form')
  contributeBtn = $('#contribute')
  customAmountDiv = $('#custom-amount')
  amountInput = $('input[name="amount"]')
  customAmountInput = $('input[name="custom_amount"]')
  recurringInput = $('input[name="recurring"]')

  amount = amountInput.filter(':checked').val()
  recurring = recurringInput.filter(':checked').val()
  stripeHandler = StripeCheckout.configure(
    key: data.stripe_publishable_key,
    image: data.logo_url,
    locale: 'auto'
    token: (token) ->
      # Add the Stripe token and email to the form, then submit it
      tokenInput = $('<input type="hidden" name="stripe_token" />').val(token.id)
      emailInput = $('<input type="hidden" name="stripe_email" />').val(token.email)

      form.append(tokenInput).append(emailInput).submit()
  )

  # Adjust amount and impact count based on donation amount
  amountInput.change ->
    if this.value == 'custom'
      customAmountDiv.removeClass 'hide'

      amount = Math.round(parseFloat(customAmountInput.val()) * 100) || 0
      updateImpactCount(amount, impactCount)
      updateContributionAmount(amount, contributionAmount)
    else
      customAmountDiv.addClass 'hide'

      amount = parseInt this.value
      updateImpactCount(amount, impactCount)
      updateContributionAmount(amount, contributionAmount)

  # Adjust amount and impact count for custom donation amount
  customAmountInput.on 'input', ->
    # If the custom box is checked, then register this input
    if amountInput.filter(':checked').val() == 'custom'
      givenAmount = parseFloat this.value

      # Don't allow negative numbers or anything that's not a valid number
      if this.value < 0 || givenAmount == NaN
        this.value = ''
        return

      # Convert to pennies or set to 0 if empty
      amount = Math.round(givenAmount * 100) || 0
      updateImpactCount(amount, impactCount)
      updateContributionAmount(amount, contributionAmount)

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
      amount: amount
    e.preventDefault()

  # Close Stripe Checkout on page navigation
  $(window).on 'popstate', ->
    stripeHandler.close()

  constructImpactCount = (count) =>
    if count == 1
      count + ' student'
    else
      count + ' students'

  # Formats a given amount in pennies for display. Ex. turns 1500 into $15
  #
  # More examples:
  #
  # 1212 -> $12.12
  # 300 -> $3
  # 4300 -> $43
  # 4201 -> $42.01
  formatAmount = (amount) =>
    floatingAmount = amount * 0.01

    if floatingAmount % 1 != 0
      "$#{floatingAmount.toFixed(2)}"
    else
      "$#{floatingAmount.toFixed(0)}"

  updateImpactCount = (amount, impactCount) =>
    count = Math.floor(amount / PENNY_COST_PER_STUDENT)
    impactCount.text constructImpactCount(count)

  updateContributionAmount = (amount, contributionAmount) =>
    contributionAmount.text formatAmount amount

  constructStripeCheckoutDescription = (recurring) =>
    suffix = 'Hack Club contribution'

    if recurring
      'Monthly ' + suffix
    else
      'One time ' + suffix
