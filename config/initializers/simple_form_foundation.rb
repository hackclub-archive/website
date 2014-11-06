SimpleForm.setup do |config|
  config.wrappers :foundation, class: :input, hint_class: :field_with_hint,
    error_class: :error do |b|
    b.use :html5
    b.use :placeholder
    b.optional :maxlength
    b.optional :pattern
    b.optional :min_max
    b.optional :readonly
    b.use :label_input
    b.use :error, wrap_with: { tag: :small, class: :error }
    b.use :hint,  wrap_with: { tag: :span, class: :hint }
  end

  # CSS class for buttons
  config.button_class = 'button'

  # CSS class to add for error notification helper.
  config.error_notification_class = 'error'

  # The default wrapper to be used by the FormBuilder.
  config.default_wrapper = :foundation

  # Render checkboxes inline
  config.boolean_style = :nested

  # Disable browser validations
  config.browser_validations = false
end
