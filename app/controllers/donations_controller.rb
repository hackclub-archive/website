class DonationsController < ApplicationController
  def new
  end

  def create
    @amount = 0

    if params[:amount] == 'custom'
      @amount = (params[:custom_amount].to_f * 100).to_i # Gotta convert it to pennies
    else
      @amount = params[:amount].to_i
    end

    # Gotta convert it from a string to a boolean
    recurring = (params[:recurring] == 'true')

    # Create the Stripe customer
    customer = Stripe::Customer.create(
      email: params[:stripe_email],
      source: params[:stripe_token]
    )

    if recurring
      plan = plan_for(@amount)
      customer.subscriptions.create(plan: plan.id)
    else
      charge = Stripe::Charge.create(
        customer: customer.id,
        amount: @amount,
        description: 'Hack Club donation',
        currency: 'usd'
      )
    end

  rescue Stripe::CardError => e
    flash[:alert] = e.message
    redirect_to new_donation_path
  end

  private

  # Monthly plans are named `donation_m_<amount>`. Examples: `donation_m_2500`,
  # `donation_m_500`, and `donation_m_1500` for the $25/m, $5/m, and $15/m
  # plans, respectively.
  #
  # If there isn't already a plan for the given amount, we create it.
  def plan_for(amount)
    plan_id = "donation_m_#{amount}"
    plan_name = "#{format_amount amount} Monthly Contribution"

    begin
      Stripe::Plan.retrieve(plan_id)
    rescue Stripe::InvalidRequestError => e
      # Raise the exception again if the error is anything but a 404. If it's a
      # 404, then we should go ahead and create a new plan.
      unless e.http_status == 404
        raise e
      end

      Stripe::Plan.create(
        amount: amount,
        interval: 'month',
        name: plan_name,
        currency: 'usd',
        id: plan_id,
      )
    end
  end

  # Formats amount for display. Example: converts 1523 to $15.23.
  def format_amount(amount)
    ActionController::Base.helpers.number_to_currency(amount.to_f / 100)
  end
end
