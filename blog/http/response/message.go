package response

// Message constant.
const (
	// Success
	MessageSuccess         = "Success!"
	MessageCreateSuccess   = "Create Success!"
	MessageUpdateSuccess   = "Update Success!"
	MessageDeleteSuccess   = "Delete Success!"
	MessageLoginSuccees    = "Logout successfully!"
	MessageRegisterSuccees = "Register successfully!"

	// Error
	MessageRegisterError = "Cannot create user!"

	// Validate
	MessageNotExistsValidate           = "Entity not exists!"
	MessageEmailExistsValidate         = "This email exists! Please pick another."
	MessageIncorrectCredentailValidate = "Username or password incorrect!"
)
